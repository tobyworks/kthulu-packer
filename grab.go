package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	resources := "resources"
	exec.Command("rm", "-rf", resources).Run()
	exec.Command("mkdir", resources).Run()
	os := runtime.GOOS
	arch := runtime.GOARCH
	jre := ""
	println("Running os: " + os)
	println("Arch : " + arch)
	if os == "darwin" {
		if arch == "arm64" {
			jre = "https://github.com/adoptium/temurin11-binaries/releases/download/jdk-11.0.15%2B10/OpenJDK11U-jre_aarch64_mac_hotspot_11.0.15_10.tar.gz"
		} else {
			jre = "https://github.com/adoptium/temurin11-binaries/releases/download/jdk-11.0.15%2B10/OpenJDK11U-jre_x64_mac_hotspot_11.0.15_10.tar.gz"
		}
	}
	println("Download: " + jre)
	DownloadFile(resources+"/jre.tar.gz", jre)
	Unpack(resources+"/jre.tar.gz", resources)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func Unpack(input string, resFolder string) {
	version := "jdk-11.0.15+10-jre"
	exec.Command("tar", "-xvf", input, "-C", resFolder).Run()
	exec.Command("rm", input).Run()
	exec.Command("mv", resFolder+"/"+version, resFolder+"/jre").Run()
}
