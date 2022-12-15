package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("sa-east-1"),
		},
	)
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
	s3Bucket = "go-expert-hlvl-bucket-exemplo"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 100)
	errorFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case fileName := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				uploadFile(fileName, uploadControl, errorFileUpload)
			}
		}
	}()
	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error reading directory: %s\n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(fileName string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", fileName)
	log.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)
	f, err := os.Open(completeFileName)
	if err != nil {
		log.Printf("Error opening file %s\n", completeFileName)
		<-uploadControl // esvazia o canal
		errorFileUpload <- fileName
		return
	}
	defer f.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   f,
	})
	if err != nil {
		log.Printf("Error uploading file %s - %s\n", completeFileName, err)
		<-uploadControl
		errorFileUpload <- fileName
		return
	}
	log.Printf("File %s uploaded successfully\n", completeFileName)
	<-uploadControl
}
