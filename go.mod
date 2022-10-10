module github.com/lifegit/video

go 1.16

require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1799
	github.com/aliyun/aliyun-oss-go-sdk v2.2.5+incompatible
	github.com/fatih/color v1.7.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/lifegit/go-gulu/v2 v2.1.8
	github.com/mitchellh/mapstructure v1.4.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	gorm.io/driver/sqlite v1.3.1
	gorm.io/gorm v1.23.2
)

replace gorm.io/datatypes v1.0.6 => github.com/lifegit/datatypes v1.0.7
