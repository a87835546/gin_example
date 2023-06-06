package service

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Read(filepath string) []byte {
	f, err := os.Open(filepath)
	if err != nil {
		log.Println("read file fail", err)
		return nil
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("read to fd fail", err)
		return nil
	}

	return fd
}

func UploadFileToS3(s *session.Session, raw []byte, filename string) (string, error) {

	tempFileName := "pic/" + filename

	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String("oner-dev"), // bucket名称
		Key:         aws.String(tempFileName),
		Body:        bytes.NewReader(raw),
		ContentType: aws.String(http.DetectContentType(raw)),
	})
	if err != nil {
		log.Println("PUT err", err)
		return "", err
	}
	fileUrl := "https://oner-dev.s3.ap-southeast-1.amazonaws.com/" + tempFileName

	return fileUrl, err
}

func main() {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-east-1"), // 替换自己账户的region
		Credentials: credentials.NewStaticCredentials(
			"AccessKeyID",
			"SecretAccessKey",
			"SessionToken"), // Sessiontoken是进程相关，应该是连接中可以返回 （可为空）
	})
	if err != nil {
		log.Println("aws  failed", err)
	}
	for i := 0; i < 10000; i++ {
		f1 := fmt.Sprintf("green_%d.gif", i)
		res := Read("./green/" + f1)
		fileName, err1 := UploadFileToS3(s, res, f1)
		if err1 != nil {
			//log.Println("Upload failed ", err1)
		} else {
			log.Println("upload success ", fileName)
		}
	}
}
