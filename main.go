package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"fmt"
	"log"
	"bytes"
)

func main() {
	path := os.Args[0]
	args := os.Args[1:]
	fileName := filepath.Base(path)
	ext := filepath.Ext(path)
	path = path[:strings.LastIndex(path, fileName)]
	if ext != "" {
		fileName = fileName[:strings.LastIndex(fileName, ext)]
	}

	logFile, err := os.OpenFile(fileName + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer logFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(logFile)

	var stdout, stderr bytes.Buffer

	log.Println(fileName + ext + " " + strings.Join(args, " "))
	cmd := exec.Command(filepath.Join(path, "original", fileName + ext), args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	fmt.Print(stdout.String())
}
