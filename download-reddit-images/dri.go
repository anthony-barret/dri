package downloadredditimages

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type RedditResponse struct {
	Data struct {
		Children []struct {
			Data struct {
				URL string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func (r RedditResponse) GetLinks() []string {
	var links []string
	for _, child := range r.Data.Children {
		if strings.HasSuffix(child.Data.URL, ".jpg") || strings.HasSuffix(child.Data.URL, ".png") {
			links = append(links, string(child.Data.URL))
		}
	}
	return links
}

func ParseConfig(configFile string) []string {
	// Read config file
	fmt.Println(configFile)
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalln("Cannot read the configuration file")
	}
	lines := strings.Split(string(data), "\n")
	// Extract subreddits
	links := make([]string, 0)
	for _, line := range lines {
		if trimmedLine := strings.TrimSpace(line); trimmedLine != "" {
			link := "https://www.reddit.com/r/" + trimmedLine + "/new.json?limit=10"
			links = append(links, link)
		}

	}
	// Return list of subreddits
	return links
}

func GetImagesLinksFromSubreddit(subredditLink string) []string {
	response, err := http.Get(subredditLink)
	if err != nil {
		log.Fatalln("Error when requesting", subredditLink)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var redditJsonPage RedditResponse
	err = json.Unmarshal(body, &redditJsonPage)
	if err != nil {
		log.Fatalln(err)
	}
	return redditJsonPage.GetLinks()
}

func DownloadImage(imageLink string) {
	response, err := http.Get(imageLink)
	if err != nil {
		log.Fatalln("Error when requesting", imageLink)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fileName := "img/" + path.Base(imageLink)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("Error creating file", err)
		return
	}
	defer file.Close()
	_, err = file.Write(body)
	if err != nil {
		log.Fatalln("Error writing in file", err)
		return
	}
}
