将视频(mp4)的语音转为文本

> 使用之前，你需要先设置 `conf/base.toml.sample` 的配置！同时把 `.sample`去掉。

### 流程
1. 建立文件目录到DB
2. 将mp4视频 转换到 mp3 语音
3. 将mp3上传到阿里云oss
4. 将阿里云oss的语音上传到阿里云语音识别


### 使用方法
1. `video warehousing`
2. `video tran`
3. `video oss`
4. `video ai-put`
5. `video ai-result`

> 1. video 指的是本ci程序。
> 2. 步骤1-4，可以直接用 `video trans`
> 3. 执行完第四步请耐心等待，识别是需要时间的。
> 4. 命令是允许重复执行的。
