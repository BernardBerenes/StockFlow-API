package config

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func NewMinio(config *viper.Viper) *minio.Client {
	minioClient, err := minio.New(fmt.Sprintf("%s:%d", config.GetString("MINIO_HOST"), config.GetInt("MINIO_PORT")), &minio.Options{
		Creds:  credentials.NewStaticV4(config.GetString("MINIO_ACCESS_KEY_ID"), config.GetString("MINIO_SECRET_ACCESS_KEY"), ""),
		Secure: false,
	})
	if err != nil {
		panic(fmt.Errorf("fatal error connecting minio: %w", err))
	}

	var exists bool

	exists, err = minioClient.BucketExists(ctx, config.GetString("MINIO_BUCKET"))
	if err != nil {
		panic(fmt.Errorf("fatal error check existing bucket: %w", err))
	}

	if !exists {
		err = minioClient.MakeBucket(ctx, config.GetString("MINIO_BUCKET"), minio.MakeBucketOptions{})
		if err != nil {
			panic(fmt.Errorf("fatal error make bucket: %w", err))
		}
	}

	return minioClient
}
