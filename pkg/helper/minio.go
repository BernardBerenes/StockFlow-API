package helper

import (
	"context"
	"fmt"
	"mime/multipart"
	"path"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
)

type MinioConfig struct {
	Config      *viper.Viper
	MinioClient *minio.Client
}

func NewMinioHelper(config *viper.Viper, client *minio.Client) *MinioConfig {
	return &MinioConfig{
		Config:      config,
		MinioClient: client,
	}
}

var ctx = context.Background()

func (h *MinioConfig) Delete(fileName string) error {
	return h.MinioClient.RemoveObject(ctx, h.Config.GetString("MINIO_BUCKET"), fileName, minio.RemoveObjectOptions{})
}

func (h *MinioConfig) Insert(fileRequest *multipart.FileHeader, fileName string) (string, error) {
	file, err := fileRequest.Open()
	if err != nil {
		return "", err
	}

	objectName := fmt.Sprintf("%s%s", fileName, path.Ext(fileRequest.Filename))

	_, err = h.MinioClient.PutObject(ctx, h.Config.GetString("MINIO_BUCKET"), objectName, file, fileRequest.Size, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})

	return objectName, err
}
