package file_store

import (
	minio "github.com/minio/minio-go"
)

func NewClient() Repo {
	endpoint := "play.minio.io:9000"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := false

	conn, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		panic(err)
	}

	bucketName := "unicampus_admission_student"
	location := "us-east-1"

	if err = conn.MakeBucket(bucketName, location); err != nil {
		if exists, err := conn.BucketExists(bucketName); err != nil || !exists {
			panic(err)
		}
	}

	return Repo{
		conn: conn,
	}
}

type Repo struct {
	conn *minio.Client
}
