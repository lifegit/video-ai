package app

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lifegit/go-gulu/v2/nice/file"
	"github.com/lifegit/go-gulu/v2/pkg/viperine"
	"github.com/lifegit/video/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

var Global GlobalConf

type GlobalConf struct {
	Workspace models.Workspace `toml:"workspace"`
	Aliyun    models.Aliyun    `toml:"aliyun"`
	DB        models.Db        `toml:"db"`
}

const DEV = "dev"

func (g *GlobalConf) isDev() bool {
	return g.getEnv() == DEV
}
func (g *GlobalConf) getEnv() (res string) {
	if res = os.Getenv("GO_ENV"); res == "" {
		res = DEV
	}

	return res
}
func SetUpConf() {
	basePath := recursionPath("conf")
	v, err := viperine.LocalConfToViper([]string{
		path.Join(basePath, "base.toml"),
		//path.Join(basePath, Global.getEnv(), "conf.toml"),
	}, &Global, func(event fsnotify.Event, viper *viper.Viper) {
		if event.Op != fsnotify.Remove {
			_ = viper.Unmarshal(&Global)
		}
	})

	if err != nil {
		logrus.WithError(err).Fatal(err, v)
	}
}

func recursionPath(dirName string) (dirPath string) {
	var dir string
	for i := 0; i < 10; i++ {
		dirPath = path.Join(dir, dirName)
		dir = path.Join(dir, "../")

		if file.IsDir(dirPath) {
			return
		}
	}

	return
}
