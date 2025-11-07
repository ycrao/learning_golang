# online-playground

[简体中文](readme_zh-CN.md)

>   A tray/taskbar app, one-click then navigate to some online programing playground sites.

### Snapshot

![](snapshot.png)

### Compile and Run

- Modify `playground.json` file to `add` or `remove` some playgrounds.
- For `Windows`, you can download [release zip file](https://github.com/ycrao/learning_golang/releases) named `playground.zip` , unzip to some directory, just click `playground.exe` to run.
- For Linux and MacOS, you need compile and build yourself.

Some issues about app-icon or compiling problem please ref to [systray](https://github.com/getlantern/systray) documentation.

```bash
# Compile For Windows
# in Windows Powershell
$env:GOOS='windows'; $env:GOARCH='amd64'; go build -ldflags "-s -w -H=windowsgui" -o playground.exe; $env:GOGS=''; $env:GOARCH='';
# Compile For Linux
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o playground
# Compile For MacOS, need package to `.app` by yourself, ref to https://github.com/getlantern/systray#macos
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o playground
```

### Links

See [playground.json](playground.json), may be have some dead links in it.

### Ref

- [awesome-playgrounds](https://github.com/ymyzk/awesome-playgrounds)
- [systray](https://github.com/getlantern/systray)
- [2goarray](https://github.com/cratonica/2goarray)