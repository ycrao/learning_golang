# play-audio-in-cmd

>   在命令行下播放音频文件，可尝试下面跨平台软件：http://sox.sourceforge.net/

### 各平台安装办法

```bash
# Ubuntu
sudo apt-get install sox
sudo apt-get install libsox-fmt-mp3
# macOS with brew
brew install lame
brew install sox
# Window *存在不兼容性见后面说明* https://sourceforge.net/projects/sox/files/sox/14.4.2/sox-14.4.2-win32.zip/download 解压到某个目录，然后将该目录加入系统环境变量PATH中
copy sox.exe play.exe
#play mp3
play -q -t mp3 hello.mp3
```

### 兼容性问题

>   https://stackoverflow.com/questions/3537155/sox-fail-util-unable-to-load-mad-decoder-library-libmad-function-mad-stream

官方最新 `14.4.2` 版 `sox` 在 `Windows 10` 下可能出现 `Sorry, there is no default audio device configured` 错误，这里使用他的[旧版]()。另外为了支持 `mp3` 格式，需要下载[依赖的 `dll`](http://www.videohelp.com/download/sox-14.4.0-libmad-libmp3lame.zip) 。本目录下，也打包了一份[副本](sox-14.4.0-libmad-libmp3lame.zip)。