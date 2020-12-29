# go-path-test

### 差异/Diff

```bash
# Go RUN 源码运行/just run in source
$go run main.go
exePath: /tmp/go-build854301005/b001/exe
absPath: /tmp/go-build854301005/b001/exe
workingDirPath: /path-to-project/go-path-test/
# 编译后运行二进制文件/build and run its binary file
$go build -o main main.go
$./main
exePath: /path-to-project/go-path-test
absPath: /path-to-project/go-path-test
workingDirPath: /path-to-project/go-path-test
# 在其它目录下运行二进制文件/run its binary file in another directory
$cd ..
$./go-path-test/main
exePath: /path-to-project/go-path-test
absPath: /path-to-project/go-path-test
workingDirPath: /path-to-project
```