package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	utils "github.com/th-release/utils/go"
)

func createFileIfNotExist(filePath string, initialContent interface{}) error {
	_, err := os.Open(filePath)
	if err != nil {
		err := os.MkdirAll(utils.GetFilePath(""), os.ModePerm)
		if err != nil {
			return fmt.Errorf("setup error: %v", err)
		}

		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("setup error: %v", err)
		}
		defer file.Close()

		if initialContent != nil {
			// If initial content is provided, write it to the file
			data, err := json.Marshal(initialContent)
			if err != nil {
				return fmt.Errorf("error marshaling initial content: %v", err)
			}

			_, err = file.Write(data)
			if err != nil {
				return fmt.Errorf("error writing initial content to file: %v", err)
			}
		}
	}

	return nil
}

func main() {
	c_err := createFileIfNotExist(utils.GetFilePath(""), nil)
	if c_err != nil {
		log.Fatalf("error: %v", c_err)
	}

	app := fiber.New()
	SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error Starting Server: %v", err)
	}
}
