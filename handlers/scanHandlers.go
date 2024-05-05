package handlers

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
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

	projectID := os.Getenv("PROJECT_ID")
	location := os.Getenv("LOCATION")
	modelName := os.Getenv("MODEL_NAME")
	temperature := 0.4

	// create a multimodal (multipart) prompt
	prompt := []genai.Part{
		genai.Text(os.Getenv("PROMPT_IMG")),
		genai.FileData{
			MIMEType: "image/jpeg",
			FileURI:  gcsURI,
		},
	}

	// generate the response
	responseAny, err := generateMultimodalContent(prompt, projectID, location, modelName, float32(temperature))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to generate response: %v", err)})
		return
	}

	// Convert responseAny to string
	response := fmt.Sprintf("%v", responseAny)

	// Define regex patterns
	patterns := map[string]string{
		"foodname": `'foodname':\s*'([^']+)`,
		"calories": `'calories':\s*(\d+)`,
		"carbs":    `'carbs':\s*(\d+)`,
		"protein":  `'protein':\s*(\d+)`,
		"fat":      `'fat':\s*(\d+)`,
		"fiber":    `'fiber':\s*(\d+)`,
		"glucose":  `'glucose':\s*(\d+)`,
	}

	// Parse JSON fields using regex
	foodname := findField(patterns["foodname"], response)
	calories := findField(patterns["calories"], response)
	carbs := findField(patterns["carbs"], response)
	protein := findField(patterns["protein"], response)
	fat := findField(patterns["fat"], response)
	fiber := findField(patterns["fiber"], response)
	glucose := findField(patterns["glucose"], response)

	// Create a Gin H map for JSON response
	parsedResponse := gin.H{
		"foodname": foodname,
		"calories": calories,
		"carbs":    carbs,
		"protein":  protein,
		"fat":      fat,
		"fiber":    fiber,
		"glucose":  glucose,
	}

	c.JSON(http.StatusOK, parsedResponse)
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

// Function to find a field using a regex pattern
func findField(pattern string, text string) string {
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(text)
	if len(match) < 2 {
		return ""
	}
	return match[1]
}
