package main

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
)

var binary string
var (
	//go:embed resources
	resources          embed.FS
	applicationSupport = os.Getenv("HOME") + "/Library/" + binary
	chmod777           = os.FileMode(0777)
)

func main() {
	_, err := os.Open(applicationSupport + "/resources/app.jar")
	if err != nil {
		os.MkdirAll(applicationSupport+"/resources", chmod777)
		traverseDir("resources")
	}
	cmd := exec.Command(applicationSupport+"/resources/jre/Contents/Home/bin/java", "-jar", applicationSupport+"/resources/app.jar")
	cmd.Run()
}

func traverseDir(dir string) {
	entries, _ := resources.ReadDir(dir)
	for _, entry := range entries {
		info, _ := entry.Info()
		if !info.IsDir() {
			writeFile(dir, entry)
		} else {
			writeDirectory(dir, entry)
		}
	}
}

func writeFile(currentDir string, entry fs.DirEntry) {
	res := currentDir + "/" + entry.Name()
	//println("File: " + res)
	f, _ := resources.Open(res)
	info, _ := entry.Info()
	size := info.Size()
	bytes := make([]byte, size)
	f.Read(bytes)
	ioutil.WriteFile(applicationSupport+"/"+res, bytes, chmod777)
}

func writeDirectory(currentDir string, entry fs.DirEntry) {
	res := currentDir + "/" + entry.Name()
	//println("Dir: " + res)
	os.MkdirAll(applicationSupport+"/"+res, chmod777)
	traverseDir(res)
}
