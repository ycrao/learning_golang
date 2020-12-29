package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	getExePath()
	getAbsPath()
	getWorkingDirPath()
}

func getExePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(ex)
	fmt.Println("exePath:", exePath)
	return exePath
}

func getAbsPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	fmt.Println("absPath:", dir)
	return dir
}

func getWorkingDirPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("workingDirPath:", dir)
	return dir
}
