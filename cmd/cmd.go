package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/lifegit/video/api"
	"github.com/spf13/cobra"
	"time"
)

var RootCmd = &cobra.Command{
	Short: "将视频(mp4)的语音转为文本",
	Long:  " 执行步骤: \r\n 1.warehousing \r\n 2. tran \r\n 3. oss \r\n 4. ai-put \r\n 5. ai-result \r\n ‼️ 1. 步骤1-4，可以直接用 video trans。 \r\n ‼️ 2.执行完第四步先等待一会再执行第五步，识别是需要时间的。",
}

func init() {
	RootCmd.AddCommand(transCmd)
	transCmd.Flags().IntVarP(&Trans.Wait, "wait", "w", 60, "发送到AI后等待多少秒取回结果")
	RootCmd.AddCommand(warehousingCmd)
	RootCmd.AddCommand(tranCmd)
	RootCmd.AddCommand(ossPutCmd)
	RootCmd.AddCommand(aiPutCmd)
	RootCmd.AddCommand(aiResultCmd)
}

var Trans struct {
	Wait int
}
var transCmd = &cobra.Command{
	Use:     "trans",
	Short:   "warehousing + tran + oss + ai-put + ai-result",
	Example: "video trans",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = api.Warehousing()
		if err != nil {
			return
		}
		err = api.Tran()
		if err != nil {
			return
		}
		err = api.OssPut()
		if err != nil {
			return
		}
		err = api.AIPut()
		if err != nil {
			return
		}
		for i := range make([]int, Trans.Wait) {
			fmt.Println(fmt.Sprintf("还需要等待%s秒", color.YellowString("%d", Trans.Wait-i)))
			time.Sleep(time.Second)
		}
		err = api.AIResult()

		return
	},
}

var warehousingCmd = &cobra.Command{
	Use:     "warehousing",
	Short:   "将mp4文件入库",
	Example: "video warehousing",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return api.Warehousing()
	},
}

var tranCmd = &cobra.Command{
	Use:     "tran",
	Short:   "转换mp4文件到mp3",
	Example: "video tran",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return api.Tran()
	},
}

var ossPutCmd = &cobra.Command{
	Use:     "oss",
	Short:   "将mp3上传到oss",
	Example: "video oss",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return api.OssPut()
	},
}

var aiPutCmd = &cobra.Command{
	Use:     "ai-put",
	Short:   "将mp3(oss)上传到ai语音识别",
	Example: "video ai-put",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return api.AIPut()
	},
}

var aiResultCmd = &cobra.Command{
	Use:     "ai-result",
	Short:   "获取ai语音识别结果",
	Example: "video ai-result",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return api.AIResult()
	},
}
