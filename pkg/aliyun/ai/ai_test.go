package ai_test

import (
	"fmt"
	"github.com/lifegit/video/pkg/aliyun/ai"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	a, err := ai.New("", "", "", "")
	if err != nil {
		return
	}
	taskId, err := a.Send("https://gw.alipayobjects.com/os/bmw-prod/0574ee2e-f494-45a5-820f-63aee583045a.wav")
	fmt.Println(taskId, err)

	for true {
		<-time.After(time.Second * 10)
		res, err := a.GetResult(taskId)
		fmt.Println(res, err)
		if err == nil {
			break
		}
	}
}
