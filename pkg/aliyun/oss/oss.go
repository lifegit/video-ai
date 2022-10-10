package oss

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"path"
)

type Oss struct {
	Client *oss.Bucket
}

func (o *Oss) OssPutAll(file string) (url string, err error) {
	v := path.Base(file)
	url = fmt.Sprintf("https://%s.%s/%s", o.Client.BucketName, o.Client.Client.Config.Endpoint, v)
	err = o.Client.PutObjectFromFile(v, file)

	return
}

func NewOss(endpoint, accessKeyID, accessKeySecret, bucketName string) (res *Oss, err error) {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, errors.New("oss-new 初始化失败")
	}
	bucket, err := client.Bucket(bucketName)

	return &Oss{Client: bucket}, err
}
