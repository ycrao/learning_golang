# 在线演练场（online-playground）

[English](readme.md)

>   一个托盘应用，可快速链接到某种或多种语言的在线 `playground` （演练场）地址。

### 截图

![](snapshot.png)

### 编译与运行

- 编辑 `playground.json` 配置文件，可以增删一些 `playgrounds`。 
- 对于 `Windows` 已有编译好的名为 `playground.zip` 的 [压缩档](https://github.com/ycrao/learning_golang/releases) ，解压到某一目录，双击 `playground.exe` 即可运行。
- `Linux` 与 `MacOS` 操作系统，需要您自行编译。

应用图标等相关问题请参考：[systray](https://github.com/getlantern/systray) 文档。

```bash
# 在  Windows 下编译
# Windows Powershell
$env:GOOS='windows'; $env:GOARCH='amd64'; go build -ldflags "-s -w -H=windowsgui" -o playground.exe; $env:GOGS=''; $env:GOARCH='';
# Linux 下编译
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o playground
# MacOS 下编译，需要自行打包成 .app 应用，参考 https://github.com/getlantern/systray#macos
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o playground
```

### 链接

查看 [playground.json](playground.json) 文件，可能存在一些死链。

### 参考

- [awesome-playgrounds](https://github.com/ymyzk/awesome-playgrounds)
- [systray](https://github.com/getlantern/systray)
- [2goarray](https://github.com/cratonica/2goarray)