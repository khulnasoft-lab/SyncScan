package main

import (
	"bytes"
	"context"
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/khulnasoft-lab/syncscan/pkg/config"
	"github.com/khulnasoft-lab/syncscan/pkg/extractor"
	"github.com/khulnasoft-lab/syncscan/pkg/scanner"
)

func main() {
	configPath := flag.String("config", "integ-tests/config.yaml", "path to the configuration file")
	flag.Parse()
	imageName := "your-docker-image-name"
	imgTar := "/tmp/xmrig/image.tar"
	imgTarGz := "/tmp/xmrig/image.tar.gz"
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("skopeo", "copy", "--insecure-policy", "--src-tls-verify=false", "docker://"+imageName, "docker-archive:"+imgTar)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println("Error: ", err.Error())
		log.Println("stderr: ", stderr.String())
		return
	}

	cmd = exec.Command("gzip", imgTar)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		println("Error: ", err.Error())
		return
	}

	err = os.Rename(imgTar+".gz", imgTarGz)
	if err != nil {
		println("Error: ", err.Error())
		println("Error: ", err.Error())
		return
	}
	cfg, err := config.ParseConfig(*configPath)
	cfg, err := config.ParseConfig("integ-tests/config.yaml")
	cfg, err = config.ParseConfig("integ-tests/config.yaml")
		log.Println(err.Error())
		return
	}

	extract, err := extractor.NewTarExtractor(config.Config2Filter(cfg), "", imgTarGz)
	if err != nil {
		println(err.Error())
		return
	}
	defer extract.Close()

	err = scanner.ApplyScan(context.Background(), extract, func(f extractor.ExtractedFile) error {
	err = scanner.ApplyScan(context.Background(), extract, func(f extractor.ExtractedFile) {
		println(f.Filename)
	})
		log.Println("Scan error: ", err.Error())
		return
	}
}
