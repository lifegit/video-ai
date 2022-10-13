package ffmpeg

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/nice/file"
	"os/exec"
	"path"
	"strings"
)

type ffmpeg struct {
	File string
}

// mp4 => mp3
func (f *ffmpeg) Tran(video, mp3Dir string) (filename string, err error) {
	err = file.IsNotExistMkDir(mp3Dir)
	if err != nil {
		return
	}

	name := strings.TrimSuffix(path.Base(video), path.Ext(video))
	filename = path.Join(mp3Dir, name+".mp3")
	// ffmpeg -i a.mp4 -y -ar 16000 a.mp3
	// https://zhuanlan.zhihu.com/p/67878761
	// -i 转换
	// -ar 设置 码率
	// -y 文件存时覆盖
	cmd := exec.Command(f.File, "-i", video, "-y", "-ar", "16000", filename)
	_, err = cmd.Output()
	//res, err := cmd.Output()
	//v := cmd.String()
	//fmt.Println(string(res), v)

	return filename, err
}

func New(ffmpegFile string) (f *ffmpeg, err error) {
	if !file.IsExist(ffmpegFile) {
		return nil, fmt.Errorf("ffmpeg file is not exist in %s", ffmpegFile)
	}

	return &ffmpeg{
		File: ffmpegFile,
	}, nil
}
