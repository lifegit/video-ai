package api

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/lifegit/video/app"
	"github.com/lifegit/video/models"
	"github.com/lifegit/video/pkg/aliyun/oss"
)

func OssPut() (err error) {
	var videos []models.Videos
	err = app.DB.
		Clause(fire.WhereCompare("mp3", "", fire.CompareNeq)).
		Clause(fire.WhereCompare("oss", "", fire.CompareEq)).
		CrudAll(models.Videos{}, &videos)
	if err != nil {
		return
	}
	aliOss, err := oss.NewOss(app.Global.Aliyun.Oss.Endpoint, app.Global.Aliyun.User.AccessKeyID, app.Global.Aliyun.User.AccessKeySecret, app.Global.Aliyun.Oss.BucketName)
	if err != nil {
		return
	}
	for _, video := range videos {
		url, err := aliOss.OssPutAll(video.Mp3)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] %s. error: %s",
				color.RedString("oss-put fail"),
				video.Mp3,
				color.RedString(err.Error()),
			))
			continue
		}

		app.DB.CrudUpdate(models.Videos{Model: models.Model{ID: video.ID}}, models.Videos{Oss: url})
		fmt.Println(fmt.Sprintf("[%s] %s",
			color.GreenString("oss-put successfully"),
			url,
		))
	}

	return
}
