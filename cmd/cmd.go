package cmd

import (
	"github.com/lifegit/video/api"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{}

func init() {
	RootCmd.AddCommand(transCmd)
	RootCmd.AddCommand(warehousingCmd)
	RootCmd.AddCommand(tranCmd)
	RootCmd.AddCommand(ossPutCmd)
	RootCmd.AddCommand(aiPutCmd)
	RootCmd.AddCommand(aiResultCmd)
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

var transCmd = &cobra.Command{
	Use:     "trans",
	Short:   "warehousing + tran + oss + ai-put",
	Example: "video trans",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return api.AIResult()
	},
}
