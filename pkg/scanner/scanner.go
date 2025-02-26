package scanner

import (
	"context"
	"io"

	"github.com/khulnasoft-lab/syncscan/pkg/extractor"
	"github.com/khulnasoft-lab/syncscan/pkg/log"
)

func ApplyScan(ctx context.Context,
	extract extractor.FileExtractor,
	scan func(extractor.ExtractedFile)) {

	var (
		err  error
		file extractor.ExtractedFile
	)
	for err != io.EOF {
		select {
		case <-ctx.Done():
			return
		default:
			file, err = extract.NextFile()
		}
		if err != nil {
			if err != io.EOF {
				log.ErrLogger(err)
			}
			continue
		}
		scan(file)
		file.Close()
	}
	if err == io.EOF {
		log.InfoLogger("scan completed")
	}
}
