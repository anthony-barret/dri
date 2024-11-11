package downloadredditimages

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func GetImagesLinksFromSubreddit(subRedditLink string) []string {
	response, err := http.Get(subRedditLink)
	if err != nil {
		log.Fatalln("Error when requesting", subRedditLink)
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

func DownloadImage(imageLink string, directory string) {
	fileName := path.Join(directory, path.Base(imageLink))
	// check if file already exists
	_, err := os.Stat(fileName)
	if !os.IsNotExist(err) {
		log.Println("The file", path.Base(imageLink), "has already been downloaded")
		return
	} else {
		log.Println("Downloading file", path.Base(imageLink))
	}
	response, err := http.Get(imageLink)
	if err != nil {
		log.Fatalln("Error when requesting", imageLink)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
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
