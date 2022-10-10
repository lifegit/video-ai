package models

type Aliyun struct {
	User User `toml:"user"`
	Ai   Ai   `toml:"ai"`
	Oss  Oss  `toml:"oss"`
}
type User struct {
	AccessKeyID     string `toml:"accessKeyID"`
	AccessKeySecret string `toml:"accessKeySecret"`
}
type Ai struct {
	RegionID string `toml:"regionId"`
	Appkey   string `toml:"appkey"`
}
type Oss struct {
	Endpoint   string `toml:"endpoint"`
	BucketName string `toml:"bucketName"`
}
