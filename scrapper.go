package main

import (
	"fmt"
	"go-epub-scrapper/utils"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/go-rod/rod"
)

type Chapter struct {
	id          int
	url         string
	title       string
	textContent string

	failed bool
}

type ePubConfig struct {
	title    string
	author   string
	coverUrl string
	synopsis string

	output string
}

func handleFailure() {
	if r := recover(); r != nil {
		appFailure = true
		fmt.Println("RECOVER", r)
	}
}

func Scrap(InfoUrl string) (ChaptersList, ePubConfig, bool) {
	defer handleFailure()
	browser := NewBrowser()
	defer browser.MustClose()

	mainPage := browser.MustPage(InfoUrl)
	mainPage.MustWaitLoad()

	title := mainPage.MustElement("article header.post-header > h1.post-title").MustText()
	author := mainPage.MustElement("article div.post-content > p:nth-child(6)").MustText()
	cover := mainPage.MustElement("article div.post-content > p:nth-child(6) img").MustAttribute("src")
	if cover == nil {
		mokecover := ""
		cover = &mokecover
	}
	synopsis := mainPage.MustElement("article div.post-content > p:nth-child(23)").MustText()

	chaptersList := mainPage.MustElements("article div.post-content > p:nth-child(28) a")
	if len(chaptersList) <= 0 {
		return nil, ePubConfig{}, false
	}

	chapters := []*Chapter{}
	for i, chapter := range chaptersList {
		chapterStruct := Chapter{
			id:    i,
			title: chapter.MustText(),
		}

		chLink := chapter.MustAttribute("href")
		if chLink == nil || len(strings.TrimSpace(*chLink)) <= 0 || !strings.HasPrefix(*chLink, "https://perpetualdaydreams.com/novel/") {
			chapterStruct.failed = true
		} else {
			chapterStruct.url = *chLink
		}

		chapters = append(chapters, &chapterStruct)
	}

	// Split the chapters list in 5 because otherwise the browser crash (due to 100+ tabs open at once...)
	for i, chGroup := range utils.CutArray(chapters, 2) {
		log.Printf("--- Part #%d ---", i+1)

		wg.Add(len(chGroup))
		for _, ch := range chGroup {
			go ch.ExtractChInfo(browser)
		}
		wg.Wait()

		if len(chapters) >= 100 {
			time.Sleep(time.Second) // to not surchage the site
		}
	}

	return chapters, ePubConfig{title: title, author: author, coverUrl: *cover, synopsis: synopsis}, true
}

func NewBrowser() *rod.Browser {
	return rod.New().MustConnect().MustIncognito().NoDefaultDevice().Timeout(4 * time.Minute)
}

var wg sync.WaitGroup

func (ch *Chapter) ExtractChInfo(browser *rod.Browser) {
	defer wg.Done()
	if ch.failed {
		return
	}

	subPage := browser.MustPage(ch.url).Timeout(1 * time.Minute)
	defer subPage.Close() // not important if page don't close
	subPage.MustWaitLoad()

	if subPage.MustInfo().URL != ch.url {
		ch.failed = true
		log.Printf("Chapter #%d failed ❌", ch.id)
		return
	}

	chContentParagraph := subPage.MustElements("article div.post-content > p")
	var chTextContent string

	for _, paragraph := range chContentParagraph {
		chTextContent += utils.CleanHTML(paragraph.MustHTML())
	}
	ch.textContent = chTextContent
	log.Printf("Chapter #%d Succeed ✅", ch.id)
}
