# cmd-bass-player

>   基于 [bass.dll](http://us2.un4seen.com/files/bass24.zip) 和 [Golang syscall](https://github.com/golang/go/wiki/WindowsDLLs) 实现的命令行版播放器。

### 安装使用说明

`Golang` 实现源码请参考本人 [gist](https://gist.github.com/ycrao/e7d1df181f870091b4a6d298d6ea2770#file-bass_play-go) 文件。

>   目前支持命令行播放 `mp3` 音频文件，仅支持 `Windows` 64位系统。

解压 `release.zip` 到特定目录下（请保证 `bass.dll` 和 `play.exe` 位于同一目录下），然后将该目录加入系统环境变量 `PATH` 中。
后续即可在 `cmd/powershell` 终端执行以下命令，播放 `mp3` 文件。

```bash
# play {mp3 file path}
# 为了稳妥路径可识别性，mp3文件名请使用英文ASCII字符串。
play sample.mp3 
play hello.mp3
```