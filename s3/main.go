package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"text/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Data struct to hold dynamic data
type Data struct {
	Name    string
	Age     int
	Country string
	// Add more fields as needed
}

func generatePDFFromHTML(htmlContent string) ([]byte, error) {
	// Initialize a new PDF document
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(htmlContent))))
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	err = pdfg.WriteFile("./test.pdf")

	if err != nil {
		return nil, err
	}

	return pdfg.Buffer().Bytes(), nil
}

func readHTMLTemplateFromS3(bucket, key string) (string, error) {
	sess := session.Must(session.NewSession())
	s3Client := s3.New(sess)

	obj, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}

	defer obj.Body.Close()
	htmlBytes, err := io.ReadAll(obj.Body)
	if err != nil {
		return "", err
	}

	return string(htmlBytes), nil
}

func uploadToS3(fileName string, data []byte, bucketName string) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc := s3.New(sess)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(data),
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return err
	}

	return nil
}

type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func handleRequest(ctx context.Context) (Response, error) {
	// Fetch dynamic data
	data := Data{
		Name:    "John Doe",
		Age:     30,
		Country: "USA",
	}

	// Read HTML template file from S3
	htmlTemplate, err := readHTMLTemplateFromS3("s3-sqs-ts-bucket-dev-321123", "templates/template-golang.html")
	if err != nil {
		return Response{StatusCode: 402}, err
	}

	fmt.Println(htmlTemplate)

	// Parse HTML template
	tmpl, err := template.New("htmlTemplate").Parse(htmlTemplate)
	if err != nil {
		return Response{StatusCode: 404}, err
	}

	// Populate template with dynamic data
	var populatedHTML bytes.Buffer
	err = tmpl.Execute(&populatedHTML, data)
	if err != nil {
		return Response{StatusCode: 404}, err
	}

	// Generate PDF from populated HTML
	pdfBytes, err := generatePDFFromHTML(populatedHTML.String())
	if err != nil {
		return Response{StatusCode: 400}, err
	}

	// Store the generated PDF back to S3
	err = uploadToS3("templates/generated_pdf.pdf", pdfBytes, "s3-sqs-ts-bucket-dev-321123")
	if err != nil {
		return Response{StatusCode: 400}, err
	}

		resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            "done",
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "world-handler",
		},
	}

	return resp, nil
}



func main() {
	lambda.Start(handleRequest)
}