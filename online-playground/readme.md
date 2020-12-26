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

- All Playground: [repl.it](https://repl.it/) | [Paiza.io](https://paiza.io/en) | [CodePad:Java/Python/C++14](http://codepad.yenhsuan.xyz/) ...
- C/CPP Playground: [CPlayground](https://cplayground.com/) | [C++Shell](http://cpp.sh/) | [Lab Stack](https://code.labstack.com/cpp) | [TutorialsPoint](https://www.tutorialspoint.com/compile_cpp_online.php) ...
- Golang Playground: [official](https://play.golang.org/) | [go-plus-play](https://play.goplus.org/) | [go-plus-repl](https://repl.goplus.org/) | [study-golang](https://play.studygolang.com/)
- PHP Playground: [PHP Sandbox](https://sandbox.onlinephpfunctions.com/) | [teh-playground](https://www.tehplayground.com/) | [LaravelPlayground](https://laravelplayground.com/#/) | [php-fiddle](http://phpfiddle.org/) | [Fast-Site-PHP](https://www.fastsitephp.com/en/playground) ...
- Web Front(HTML/CSS/Javascript): [codepen](https://codepen.io/) | [jsfiddle](https://jsfiddle.net/) | [PlayCode](https://playcode.io/empty/) ...
- [Typescript Playground](https://www.typescriptlang.org/play/)
- Python Playground: [Jupyter](https://jupyter.org/) | [trinket](https://trinket.io/python) | [Kata Coda](https://www.katacoda.com/courses/python/playground) ...
- Java Playground: [Lab Stack](https://code.labstack.com/java) | [Study Tonight](https://www.studytonight.com/code/playground/java/) ...
- SQL Playground: [db-fiddle](https://www.db-fiddle.com/) | [SQL Fiddle](http://sqlfiddle.com/) | [sql pad](https://sqlpad.io/playground/) ...
- [SwiftPlayground](http://online.swiftplayground.run/)
- RubyPlayground: [official](https://try.ruby-lang.org/playground/) | [Kata Coda](https://www.katacoda.com/courses/ruby/playground) ...
- RustPlayground: [official](https://play.rust-lang.org/) | [TutorialsPoint](https://www.tutorialspoint.com/compile_rust_online.php) ...
- ...

### Ref

- [awesome-playgrounds](https://github.com/ymyzk/awesome-playgrounds)
- [systray](https://github.com/getlantern/systray)
- [2goarray](https://github.com/cratonica/2goarray)