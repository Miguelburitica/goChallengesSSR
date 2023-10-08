package controllers

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
)

// Initialize aws session
var sess, _ = session.NewSession(&aws.Config{
	Region: aws.String("us-east-2")},
)
// Initialize s3 service client
var svc = s3.New(sess)

func UploadFileController(c *fiber.Ctx) error {
	// Parse the multipart form:
	form, err := c.MultipartForm()
	if err != nil {
			return err
	}

	// Get all files from "documents" key:
	files := form.File["documents"]
	// print the files
	log.Println(files)

	// Loop through files:
	for _, file := range files {
			// Source
			src, err := file.Open()
			if err != nil {
					return err
			}
			defer src.Close()

			// Upload the file to the s3 given bucket
			_, err = svc.PutObject(&s3.PutObjectInput{
					Bucket: aws.String("loaded-files"),
					ContentType: aws.String(file.Header.Get("Content-Type")),
					Key:    aws.String(file.Filename),
					Body:   src,
			})
			if err != nil {
					return err
			}
	}

	// Return success message
	return c.SendString("Files uploaded successfully!")
}

func UploadFilesViewController(c *fiber.Ctx) error {
	return c.Render("upload-file", fiber.Map{})
}