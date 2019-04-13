package myminio

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	minio "github.com/minio/minio-go"

	"start/config"
)

var minioClient *minio.Client

func Init() {
	config := config.GetConfig()
	var errMinioClient error
	minioClient, errMinioClient = minio.New(
		config.GetString("minio.host"),
		config.GetString("minio.access_key"),
		config.GetString("minio.secret_key"),
		config.GetBool("minio.use_ssl"),
	)
	if errMinioClient != nil {
		log.Fatalln(errMinioClient)
	}
	policy := `{
		"Version": "2012-10-17",
		"Statement": [
			{"Action": ["s3:GetObject"],
			"Effect": "Allow",
			"Principal": {"AWS": ["*"]},
			"Resource": ["arn:aws:s3:::test/*"]
			,"Sid": "Public"}
			]
		}`
	errPolicy := minioClient.SetBucketPolicy(config.GetString("minio.bucket"), policy)
	if errPolicy != nil {
		log.Fatalln(errPolicy)
	}
}

func UploadFile(file *multipart.File, bucket, filename string) (string, error) {
	var buff bytes.Buffer
	io.Copy(&buff, *file)
	bFile := buff.Bytes()
	lengthFile := buff.Len()
	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	// contentTypeFile := http.DetectContentType(bFile)

	lenUpload, errUpload := minioClient.PutObject(
		bucket,
		filename, bytes.NewReader(bFile),
		int64(lengthFile),
		minio.PutObjectOptions{ContentType: "application/octet-stream", UserMetadata: userMetaData})
	if errUpload != nil {
		log.Println("ERROR_UPLOAD:", errUpload)
		return "", errUpload
	}
	fmt.Println("Successfully uploaded bytes: ", lenUpload)
	return filename, nil
}
