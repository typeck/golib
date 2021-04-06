package aws

import (

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)


//s3manager.UploadInput{
//		Bucket: aws.String(bucket),
//		Key:    aws.String(filePath),
//		Body:   bytes.NewReader(file),
//	}
//size: buffer size
func UploadS3(region string,in *s3manager.UploadInput, size int)error{
	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		// Define a strategy that will buffer 25 MiB in memory
		u.BufferProvider = s3manager.NewBufferedReadSeekerWriteToPool(size)
	})

	_, err := uploader.Upload(in)
	if err != nil {
		return err
	}
	return  nil
}
