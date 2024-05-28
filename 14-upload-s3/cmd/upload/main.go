package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	wg sync.WaitGroup
	dirPath = "../tmp"
) 
func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	s3Client = s3.New(sess)
}

func main(){
	// esse channel suporta ate 100 valores (structs) dentro dele
	// controla a quantidade de upload prto s3
	ch := make( chan struct{}, 100)

	// esse channel vai controlar quem falhou para fazer a retentativa de upload do arquivo.
	// todo erro que ocorrer, jogar o nome do arquivo nesse canal e se cair algo aqui retentar o upload
	ch2 := make(chan string, 10)

	
	// readFiles(dirPath)
	readFilesAndUploadToS3( ch, ch2)
}
func readFilesAndUploadToS3(uploadControl chan struct{}, errorFileUpload chan string){
	dir, err := os.Open(dirPath)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	go retryErrorFileUpload(uploadControl, errorFileUpload)

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}

		wg.Add(1)
		// pq colocando uma struct vazia: pq ele vai encher e vai fazer em ate 100 simultaneos e nao tentando fazer tudo de maneira concorrente.
		uploadControl <- struct{}{}
	
		go uploadToS3( files[0].Name(),  "s3-sqs-ts-bucket-dev-321123", uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadToS3( fileName, bucketName string, uploadControl <-chan struct{}, errorFileUpload chan<- string)  {
	defer wg.Done()

	content, err := readFilesFromDir(fileName)
	//  soh ocorre aqui quando nao eh possivel ler o arquivo de dentro da pasta
	if err != nil {
		<-uploadControl // esvazia o canal se der erro para outro tentar
		errorFileUpload <- fileName // carrega o nome do arquivo para retentativa
		return 
	}

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(content),
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		fmt.Printf("Error uploading file %s\n", fileName)
		<-uploadControl // esvazia o canal 
		errorFileUpload <- fileName // carrega o nome do arquivo para retentativa
		return
	}

	fmt.Printf("File %s uploaded successfully\n", fileName)
	<-uploadControl // esvazia o canal 
	return 
}

func retryErrorFileUpload(uploadControl chan struct{}, errorFileUpload chan string){
	for{
		select {
		case fileName := <- errorFileUpload:
			wg.Add(1)
			// para retentar, deve sinalizar 
			uploadControl <- struct{}{}
		
			go uploadToS3( fileName,  "s3-sqs-ts-bucket-dev-321123", uploadControl, errorFileUpload)
		}
	}
}

func readFilesFromDir(fileName string) ([]byte, error){
	filePath := filepath.Join(dirPath, fileName)
	content, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Printf("Error opening file %s\n", filePath)
		return nil, err
	}
	fmt.Printf("File: %s\tContent: %s\n", filePath, content)
	return content, nil
}

//  using range -> one upload at a time
// func readFiles(dirPath string){
//     // Open the directory
// 	dirFiles, err := os.ReadDir(dirPath)
    
// 	if err != nil {
//         log.Fatal(err)
// 	}

//     // Iterate through the files and read their content
//     for _, fileInfo := range dirFiles {
//         if !fileInfo.IsDir() {
//             filePath := filepath.Join(dirPath, fileInfo.Name())
//             content, err := os.ReadFile(filePath)

//             if err != nil {
//                 log.Fatal(err)
//             }
//             fmt.Printf("File: %s\nContent:\n%s\n", filePath, content)
//         }
//     }
// }

