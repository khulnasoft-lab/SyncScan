package main

import (
	"context"

	"github.com/khulnasoft-lab/syncscan/pkg/config"
	"github.com/khulnasoft-lab/syncscan/pkg/extractor"
	"github.com/khulnasoft-lab/syncscan/pkg/scanner"
)

func mainImage() {
	imageID := "e0c9858e10ed"

	cfg, err := config.ParseConfig("integ-tests/config.yaml")
	if err != nil {
		println(err.Error())
		return
	}

	extract, err := extractor.NewImageExtractor(config.Config2Filter(cfg), "", imageID)
	if err != nil {
		println(err.Error())
		return
	}
	defer extract.Close()

	scanner.ApplyScan(context.Background(), extract, func(f extractor.ExtractedFile) {
		println(f.Filename)
	})
}
