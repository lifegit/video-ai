package api

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/lifegit/video/app"
	"github.com/lifegit/video/models"
	"github.com/lifegit/video/pkg/aliyun/ai"
)

func AIPut() (err error) {
	var videos []models.Videos
	err = app.DB.
		Clause(fire.WhereCompare("oss", "", fire.CompareNeq)).
		Clause(fire.WhereCompare("ai_task_id", "", fire.CompareEq)).
		CrudAll(models.Videos{}, &videos)
	if err != nil {
		return
	}
	aliAI, err := ai.New(app.Global.Aliyun.Ai.RegionID, app.Global.Aliyun.User.AccessKeyID, app.Global.Aliyun.User.AccessKeySecret, app.Global.Aliyun.Ai.Appkey)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	for _, video := range videos {
		taskID, err := aliAI.Send(video.Oss)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] %s. %s.error: %s",
				color.RedString("ai-put fail"),
				video.AITaskID,
				video.Oss,
				color.RedString(err.Error()),
			))
			continue
		}

		app.DB.CrudUpdate(models.Videos{Model: models.Model{ID: video.ID}}, models.Videos{AITaskID: taskID})
		fmt.Println(fmt.Sprintf("[%s] %s. %s",
			color.GreenString("ai-put successfully"),
			taskID,
			video.Oss,
		))
	}

	return
}
