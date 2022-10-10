package api

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/lifegit/video/app"
	"github.com/lifegit/video/models"
	"github.com/lifegit/video/pkg/aliyun/ai"
)

func AIResult() (err error) {
	var videos []models.Videos
	err = app.DB.
		Clause(fire.WhereCompare("ai_task_id", "", fire.CompareNeq)).
		Clause(fire.WhereCompare("ai_task_result", "", fire.CompareEq)).
		CrudAll(models.Videos{}, &videos)
	if err != nil {
		return
	}
	aliAI, err := ai.New(app.Global.Aliyun.Ai.RegionID, app.Global.Aliyun.User.AccessKeyID, app.Global.Aliyun.User.AccessKeySecret, app.Global.Aliyun.Ai.Appkey)
	if err != nil {
		return
	}
	for _, video := range videos {
		task, err := aliAI.GetResult(video.AITaskID)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] %s. oss: %s .error: %s",
				color.RedString("ai-result fail"),
				color.YellowString(video.AITaskID),
				video.Oss,
				color.RedString(err.Error()),
			))
			continue
		}

		result, _ := json.Marshal(task)
		var resultSimple string
		for _, t := range task {
			resultSimple = resultSimple + t.Text
		}

		app.DB.CrudUpdate(models.Videos{Model: models.Model{ID: video.ID}}, models.Videos{AITaskResult: string(result), AITaskResultSimple: resultSimple})
		fmt.Println(fmt.Sprintf("[%s] %s. %s  %s",
			color.GreenString("ai-result successfully"),
			video.AITaskID,
			color.GreenString(resultSimple),
			video.Oss,
		))
	}

	return
}
