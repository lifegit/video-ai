package models

import "time"

type Model struct {
	ID        uint `gorm:"primarykey" form:"id" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Videos struct {
	Model

	Mp4                string `gorm:"not null;comment:mp4视频位置;unique" json:"mp4"`
	Mp3                string `gorm:"not null;comment:mp3语音位置" json:"mp3"`
	Oss                string `gorm:"not null;comment:mp3语音远程oss位置" json:"oss"`
	AITaskID           string `gorm:"not null;comment:ai识别id" json:"ai_task_id"`
	AITaskResult       string `gorm:"not null;comment:ai识别结果" json:"ai_task_result"`
	AITaskResultSimple string `gorm:"not null;comment:ai识别简单结果" json:"ai_task_result_simple"`
}
