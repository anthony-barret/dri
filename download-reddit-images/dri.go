package downloadredditimages

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

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
