package handlers

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"cloud.google.com/go/vertexai/genai"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

func ScanFood(c *gin.Context) {
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to get image from request: %v", err)})
		return
	}
	defer file.Close()

	// Upload image to GCS
	gcsURI, err := uploadImageToGCS(c, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to upload image to GCS: %v", err)})
		return
	}

	projectID := "upheld-acumen-420202"
	location := "us-central1"
	modelName := "gemini-1.0-pro-vision"
	temperature := 0.4

	// create a multimodal (multipart) prompt
	prompt := []genai.Part{
		genai.Text("For the given image, return a JSON object that has the fields foodname, calories, carbohydrate, protein, fat, fiber, and glucose with the value being the approximate value (integer) in gram (or cal for calories). Just output the JSON object without explanation or description. example output: { 'foodname': 'spaghetti', 'calories': 300, 'carbs': 20, 'protein': 10, 'fat': 13, 'fiber': 2, 'glucose': 2 }."),
		genai.FileData{
			MIMEType: "image/jpeg",
			FileURI:  gcsURI, // Use the URI obtained from GCS
		},
	}

	// generate the response
	response, err := generateMultimodalContent(prompt, projectID, location, modelName, float32(temperature))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to generate response: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}

// Upload image to Google Cloud Storage and return the URI
func uploadImageToGCS(c *gin.Context, file multipart.File) (string, error) {
	ctx := appengine.NewContext(c.Request)

	// Create a Google Cloud Storage client
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("application_default_credentials.json"))
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Get a bucket handle
	bucket := client.Bucket("glutara-scan")

	// Generate object name with current date and time
    currentTime := time.Now().Format("2006-01-02_15:04:05") // Format as "YYYY-MM-DD_HH:MM:SS"
    objectName := "images/" + currentTime + ".jpg"

	// Create a writer for uploading the file
	wc := bucket.Object(objectName).NewWriter(ctx)
	defer wc.Close()

	// Copy the file data to the GCS object
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	// Return the URI of the uploaded file
	return "gs://glutara-scan/" + objectName, nil
}

// generateMultimodalContent provide a generated response using multimodal input
func generateMultimodalContent(parts []genai.Part, projectID, location, modelName string, temperature float32) (any, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)
	model.SetTemperature(temperature)

	res, err := model.GenerateContent(ctx, parts...)
	if err != nil {
		return "", fmt.Errorf("unable to generate contents: %v", err)
	}

	// Check if there are any parts in the response
	if len(res.Candidates) == 0 || len(res.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("generated response contains no content")
	}

	// Return the first part of the content
	return res.Candidates[0].Content.Parts[0], nil
}
