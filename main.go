package main

import (
	"io/ioutil"

	"github.com/gofiber/fiber"
	"github.com/otiai10/gosseract"
)

type Response struct {
	Text           string `json:"text"`
	TranslatedText string `json:"translatedText"`
}

/* func handleTranslate(text string) (*Response, error){
	translateClient := translate.NewClient("APP_ID", "API_KEY", "SECRET_KEY")

    // Set the text to be translated and the target language
    text := "Hello, world!"
    target := "fr"

    // Call the API to translate the text
    result, err := client.Translate(text, target, "auto")
    if err != nil {
        fmt.Println(err)
        return
    }

} */

func handleGet(c *fiber.Ctx) {
	// Get the file data from the form-data
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(err)
		c.Status(fiber.StatusInternalServerError)
	}

	// Read the file data
	fileData, err := file.Open()
	if err != nil {
		c.JSON(err)
		c.Status(fiber.StatusInternalServerError)
	}

	defer fileData.Close()
	data, err := ioutil.ReadAll(fileData)
	if err != nil {
		c.JSON(err)
		c.Status(fiber.StatusInternalServerError)
	}

	// Use gosseract to extract text from the file data
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImageFromBytes(data)
	text, err := client.Text()
	if err != nil {
		c.JSON(err)
		c.Status(fiber.StatusInternalServerError)
	}

	// Create the response data
	response := Response{
		Text: text,
	}

	// Write the response data to the response writer
	if err := c.JSON(response); err != nil {
		c.JSON(err)
		c.Status(fiber.StatusInternalServerError)
	}

	switch err {
	case nil:
		c.JSON(response)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
}

func main() {
	app := fiber.New()

	app.Get("/getText", handleGet)

	app.Listen(8080)
}
