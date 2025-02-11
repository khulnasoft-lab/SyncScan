package main

import (
	"bytes"
	"context"
	"os/exec"

	"github.com/khulnasoft-lab/syncscan/pkg/config"
	"github.com/khulnasoft-lab/syncscan/pkg/extractor"
	"github.com/khulnasoft-lab/syncscan/pkg/scanner"
)

func runTarExtraction() {
	imgTar := "/tmp/xmrig/image.tar"
	imageName := "metal3d/xmrig:latest"
	cmd := exec.Command("skopeo", "copy", "--insecure-policy", "--src-tls-verify=false",
		"docker://"+imageName, "docker-archive:"+imgTar)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	errorOnRun := cmd.Run()
	if errorOnRun != nil {
		println("Error: ", errorOnRun.Error())
		println("stderr: ", stderr.String())
		return
	}

	cfg, err := config.ParseConfig("integ-tests/config.yaml")
	if err != nil {
		println(err.Error())
		return
	}

	extract, err := extractor.NewTarExtractor(config.Config2Filter(cfg), "", imgTar)
	if err != nil {
		println(err.Error())
		return
	}
	defer extract.Close()

	scanner.ApplyScan(context.Background(), extract, func(f extractor.ExtractedFile) {
		println(f.Filename)
	})
}
