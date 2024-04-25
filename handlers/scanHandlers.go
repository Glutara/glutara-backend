package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"cloud.google.com/go/vertexai/genai"
	"github.com/gin-gonic/gin"
)

func ScanFood(c *gin.Context) {
	var requestBody map[string]interface{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON request body"})
		return
	}

	imageURL, ok := requestBody["image_url"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image URL"})
		return
	}

	projectID := "upheld-acumen-420202"
	location := "us-central1"
	modelName := "gemini-1.0-pro-vision"
	temperature := 0.4

	// construct this multimodal prompt:
	newImage, err := partFromImageURL(imageURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
		return
	}

	// create a multimodal (multipart) prompt
	prompt := []genai.Part{
		genai.Text("Describe this picture."),
		newImage,
	}

	// generate the response
	response, err := generateMultimodalContent(prompt, projectID, location, modelName, float32(temperature))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to generate response: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
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

// partFromImageURL create a multimodal prompt part from an image URL
func partFromImageURL(image string) (genai.Part, error) {
	var img genai.Blob

	imageURL, err := url.Parse(image)
	if err != nil {
		return img, err
	}
	res, err := http.Get(image)
	if err != nil || res.StatusCode != 200 {
		return img, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return img, fmt.Errorf("unable to read from http: %v", err)
	}

	position := strings.LastIndex(imageURL.Path, ".")
	if position == -1 {
		return img, fmt.Errorf("couldn't find a period to indicate a file extension")
	}
	ext := imageURL.Path[position+1:]

	img = genai.ImageData(ext, data)
	return img, nil
}
