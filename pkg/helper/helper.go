package helper

import (
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
)

type Helper struct {
	Minio *MinioConfig
}

func NewHelper(config *viper.Viper, client *minio.Client) *Helper {
	return &Helper{
		Minio: NewMinioHelper(config, client),
	}
}
