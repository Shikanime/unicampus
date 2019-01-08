package services

import (
	"github.com/minio/minio-go"
)

type MinioService struct {
	driver *minio.Client
}

func NewMinioService() *minio.Client {
	endpoint := "play.minio.io:9000"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := false

	conn, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		panic(err)
	}

	return conn
}

func (s *MinioService) Driver() *minio.Client {
	return s.driver
}

func (s *MinioService) Close() {
}
