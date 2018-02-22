package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"fmt"
	"log"
)

func main() {
	path := os.Args[0]
	args := strings.Join(os.Args[1:], " ")
	fileName := filepath.Base(path)
	ext := filepath.Ext(path)
	if ext != "" {
		fileName = fileName[:strings.LastIndex(fileName, ext)]
	}

	logFile, err := os.OpenFile(fileName + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer logFile.Close()
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)

	out, err := exec.Command("./" + fileName + "_orig" + ext, args).Output()
	log.Println(fileName + ext + " " + args)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
