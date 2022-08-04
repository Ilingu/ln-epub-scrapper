package main

import (
	"fmt"
	"path/filepath"

	"github.com/bmaupin/go-epub"
)

type ChaptersList []*Chapter

func (chapters ChaptersList) ToEpub(config ePubConfig) (string, bool) {
	// Create a new EPUB
	e := epub.NewEpub("Nurturing the Hero to Avoid Death")

	// Metadata
	e.SetTitle(config.title)
	e.SetAuthor(config.author)
	e.SetDescription(config.synopsis)
	e.SetLang("english")

	// CoverPath, ok := utils.GetImageCover(config.coverUrl)
	e.SetCover(config.coverUrl, "")

	for _, ch := range chapters {
		if ch.failed {
			continue
		}

		sectionBody := fmt.Sprintf(`<h1>%s</h1>%s`, ch.title, ch.textContent)
		e.AddSection(sectionBody, ch.title, "", "")
	}

	outputFilePath := filepath.Join(config.output, fmt.Sprintf("%s.epub", config.title))
	err := e.Write(outputFilePath)
	if err != nil {
		return "", false
	}

	return outputFilePath, true
}
