package minioconnect

import (
	"awesomeProject/dou-yin/common/viperconfigread"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v6"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
)

func MinioConnect() (*minio.Client, error) {
	minioConfig, err := viperconfigread.ConfigReadToMinio()
	if err != nil {
		logx.Errorf("MinioConnect error:%v", err)
		return nil, err
	}
	minioClient, err := minio.New(minioConfig.Endpoint, minioConfig.AccessKeyID,
		minioConfig.SecretAccessKey, minioConfig.UseSSL)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	logx.Infof("%v", minioClient) // minioClient初使化成功
	return minioClient, nil
}

func MinioMakeBucket() error {
	// 创建一个叫mymusic的存储桶。
	// bucketName := "test-minio"
	// location := "yzx_bucket" // 存储桶被创建的region

	// err := minioClient.MakeBucket(bucketName, location)
	// if err != nil {
	// 	// 检查存储桶是否已经存在。
	// 	exists, err := minioClient.BucketExists(bucketName)
	// 	if err == nil && exists {
	// 		logx.Debug("We already own %s\n", bucketName)
	// 	} else {
	// 		logx.Error(err)
	// 		return err
	// 	}
	// }
	// logx.Debug("Successfully created %s\n", bucketName)
	return nil
}

func MinioFileUploader(minioClient *minio.Client, bucketName string, objectPre string, filePath string) (string, error) {
	newfilePath := filepath.Base(filePath)
	newfilePath = uuid.New().String() + "-" + newfilePath
	logx.Infof("%s", newfilePath)

	objectName := objectPre + newfilePath // 要上传的文件的名字
	logx.Infof("MinioFileUploader, objectName:%s, newfilePath:%s", objectName, newfilePath)
	contentType := ""
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		logx.Error(err)
		return "", err
	}
	logx.Infof("Successfully uploaded %s of size %d\n", objectName, n)
	str := bucketName + "_" + objectName
	strbytes := []byte(str)
	bucket_filepath := "minio_" + base64.StdEncoding.EncodeToString(strbytes)
	return bucket_filepath, nil
}

func MinioFileUploader_byte(minioClient *minio.Client, bucketName string, objectPre string, filePath string, videoIO io.Reader, objectSize int64) (string, error) {

	newfilePath := filepath.Base(filePath)
	newfilePath = uuid.New().String() + "-" + newfilePath
	logx.Infof("%s", newfilePath)

	objectName := objectPre + newfilePath // 要上传的文件的名字
	logx.Infof("MinioFileUploader, objectName:%s, newfilePath:%s", objectName, newfilePath)
	contentType := "video/mp4" // 类型

	//n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType}

	n, err := minioClient.PutObject(bucketName, objectName, videoIO, objectSize, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		logx.Error(err)
		return "", err
	}
	logx.Infof("Successfully uploaded %s of size %d\n", objectName, n)
	str := bucketName + "_" + objectName
	strbytes := []byte(str)
	bucket_filepath := "minio_" + base64.StdEncoding.EncodeToString(strbytes)
	return bucket_filepath, nil
}

type MinioKeyVal struct {
	SourceType string
	Bucket     string
	Key        string
}

const (
	sourcetype = "minio"
	separator  = "_"
)

func GetPlayUrl(playUrl string) (string, error) {
	if playUrl == "" {
		logx.Infof("[pkg]BaseInterface [func]GetPlayUrl [msg]playUrl is nil")
		return "", nil
	}
	decodeKey, err := DecodeFileKey(playUrl)
	if err != nil {
		logx.Errorf("[pkg]BaseInterface [func]GetPlayUrl [msg]decode base64 error:%v", err)
		return "", err
	}

	ConfigReadToMinio, err := viperconfigread.ConfigReadToMinio()
	if err != nil {
		logx.Errorf("[pkg]BaseInterface [func]GetPlayUrl [msg]SqlConnect error:%v", err)
		return "获取域名失败，->", err
	}
	minioUrl := "http://" + ConfigReadToMinio.Endpoint
	resPlayUrl := fmt.Sprintf("%s/%s/%s", minioUrl, decodeKey.Bucket, decodeKey.Key)
	return resPlayUrl, nil
}

func DecodeFileKey(key string) (*MinioKeyVal, error) {
	keyval := &MinioKeyVal{}
	if !strings.Contains(key, separator) {
		return nil, errors.New("invalid filekey fail")
	}
	keyparts := strings.Split(key, separator)
	if len(keyparts) != 2 {
		return nil, errors.New("cant Split")
	}
	keyval.SourceType = keyparts[0]
	keyData, err := base64.StdEncoding.DecodeString(keyparts[1])
	if err != nil {
		logx.Errorf("decode base64 error:", err.Error())
		return nil, err
	}

	decodeString := string(keyData)
	index := strings.Index(decodeString, separator)
	if index <= 0 {
		return nil, errors.New("cant find separator")
	}

	keyval.Bucket = decodeString[:index]
	keyval.Key = decodeString[index+len(separator):]
	return keyval, nil
}
