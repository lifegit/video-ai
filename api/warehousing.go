package api

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/lifegit/go-gulu/v2/nice/file"
	"github.com/lifegit/video/app"
	"github.com/lifegit/video/models"
	"strings"
)

// Warehousing is put mp4 video to db
func Warehousing() (err error) {
	videos, err := file.GetAllFile(app.Global.Workspace.Mp4, []string{".mp4"})
	if err != nil {
		return
	}
	for _, video := range videos {
		err = app.DB.Create(&models.Videos{Mp4: video}).Error
		if err != nil {
			if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
				fmt.Println(color.YellowString("[WARN] UNIQUE constraint failed: %s", video))
			} else {
				return err
			}
		}
	}
	return nil
}
