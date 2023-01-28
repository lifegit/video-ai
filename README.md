将视频(mp4)的语音转为文本

> 使用之前，你需要先设置 `conf/base.sample.toml` 的配置！同时把 `.sample`去掉。

### 流程
1. 建立文件目录到 DB
2. 通过 [ffmpeg](https://www.ffmpeg.org/download.html) 将mp4视频 转换到 mp3 语音
3. 将 mp3 上传到 [阿里云oss](https://www.aliyun.com/product/oss)
4. 将 阿里云oss 的语音上传到[阿里云语音识别](https://ai.aliyun.com/nls/filetrans)


### 使用方法
1. `video warehousing`
2. `video tran`
3. `video oss`
4. `video ai-put`
5. `video ai-result`
   
> 1. video 指的是本ci程序。
> 3. 执行完第四步请耐心等待，识别是需要时间的。
> 4. 你也可使用 `video trans` 命令将上述五步合并运行，通过 `-w `控制第四步的等待时间（默认300秒）。
> 5. 命令是允许重复执行的。
