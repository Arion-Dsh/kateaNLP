package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
)

// GetCurrentFilePath ...
func GetCurrentFilePath(filename string) string {
	_, filePath, _, _ := runtime.Caller(1)
	filePath = path.Join(path.Dir(filePath), filename)
	return filePath
}

// DownloadFile  ...
func DownloadFile(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// StrList fmt list to string
func StrList(list interface{}) string {
	return fmt.Sprintf("%q", list)
}

// GetSentences split docs to []string like: ["sentence1","sentence2",...]
func GetSentences(docs string, cutAll bool) []string {
	sentences := []string{}

	linePoint := regexp.MustCompile(`[\r\n]`)
	delimiter := regexp.MustCompile(`[，。？！；]`)
	for _, line := range linePoint.Split(docs, -1) {
		s := strings.Trim(line, " ")
		if s == "" {
			continue
		}
		if cutAll {
			for _, i := range delimiter.Split(s, -1) {
				if i == "" {
					continue
				}
				sentences = append(sentences, i)
			}
		} else {
			sentences = append(sentences, s)
		}
	}
	return sentences
}

// FliterStopWord fliter the stopword
func FliterStopWord(words []string) []string {
	stopwordsFile := GetCurrentFilePath("stopwords.txt")
	f, err := ioutil.ReadFile(stopwordsFile)
	if err != nil {
		panic("load stropwords.txt error")
	}
	stopwords := map[string]string{}
	for _, word := range strings.Split(string(f), "\n") {
		s := strings.Trim(word, " ")
		stopwords[s] = s
	}
	ret := []string{}
	for _, word := range words {
		w := strings.Trim(word, " ")
		if _, in := stopwords[w]; !in {
			ret = append(ret, w)
		}
	}

	return ret
}
