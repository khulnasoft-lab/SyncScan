package main

import (
	"context"

	"github.com/khulnasoft-lab/syncscan/pkg/config"
	"github.com/khulnasoft-lab/syncscan/pkg/extractor"
	"github.com/khulnasoft-lab/syncscan/pkg/scanner"
)

func run() {
	containerID := "9612e7b41bc5"

	cfg, err := config.ParseConfig("integ-tests/config.yaml")
	if err != nil {
		println(err.Error())
		return
	}
	extract, err := extractor.NewContainerExtractor(config.Config2Filter(cfg), "", containerID)
	if err != nil {
		println(err.Error())
		return
	}
	defer extract.Close()

	scanner.ApplyScan(context.Background(), extract, func(f extractor.ExtractedFile) {
		println(f.Filename)
	})
}
