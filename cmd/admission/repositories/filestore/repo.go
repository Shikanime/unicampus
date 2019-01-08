package filestore

import (
	minio "github.com/minio/minio-go"
)

func NewRepository(conn *minio.Client) Repo {
	bucketName := "unicampus_api_admission_v1alpha1_student"
	location := "us-east-1"

	if err := conn.MakeBucket(bucketName, location); err != nil {
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
