package api

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/lifegit/video/app"
	"github.com/lifegit/video/models"
	"github.com/lifegit/video/pkg/ffmpeg"
)

// Tran is mp4 tran to mp3
func Tran() (err error) {
	fmg, err := ffmpeg.New(app.Global.Workspace.Ffmpeg)
	if err != nil {
		return
	}

	var videos []models.Videos
	err = app.DB.CrudAll(models.Videos{}, &videos)
	if err != nil {
		return
	}
	for _, video := range videos {
		mp3, err := fmg.Tran(video.Mp4, app.Global.Workspace.Mp3)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] %s. error: %s",
				color.RedString("tran fail"),
				video.Mp4,
				color.RedString(err.Error()),
			))
			continue
		}

		app.DB.CrudUpdate(models.Videos{Model: models.Model{ID: video.ID}}, models.Videos{Mp3: mp3})
		fmt.Println(fmt.Sprintf("[%s] %s", color.GreenString("tran successfully"), mp3))
	}

	return
}
