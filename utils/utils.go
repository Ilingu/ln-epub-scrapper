package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// It will cut an Array in n slices, e.g: if n = 4 you would get an array containing 4 arrays.
// Note that it it approximative: if "len(arr) / n" doesn't result in an integer you would get n+1 slices
func CutArray[T any](arr []T, n int) [][]T {
	ratio := len(arr) / n
	newArr := [][]T{}

	var remains []T
	for i := range arr {
		if i != 0 && i%ratio == 0 {
			newArr = append(newArr, arr[i-ratio:i])
			remains = arr[i:]
		}
	}
	newArr = append(newArr, remains) // add remainder
	return newArr
}

func GetImageCover(cover string) (string, bool) {
	resp, err := http.Get(cover)
	if err != nil {
		log.Println(err)
		return "", false
	}
	defer resp.Body.Close()

	file, err := os.CreateTemp("", "epub-scrapper-cover-*.png")
	if err != nil {
		log.Println(err)
		return "", false
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Println(err)
		return "", false
	}

	return file.Name(), true
}

func CleanHTML(html string) string {
	html = strings.ReplaceAll(html, "”", "'")
	html = strings.ReplaceAll(html, "&nbsp;", "")
	html = strings.ReplaceAll(html, "<br>", "<br />")
	return html
}
