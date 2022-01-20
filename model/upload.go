package model

import (
	"VideoWeb/config"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

func UpLoadFile(file multipart.File, filesize int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: config.Bucket,
	}
	mac := qbox.NewMac(config.AccessKey, config.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(
		context.Background(),
		&ret,
		upToken,
		file,
		filesize,
		&putExtra,
	)
	if err != nil {
		return "", err
	}
	url := config.QiniuServer + ret.Key
	return url, nil
}
