package downloadredditimages

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func GetImagesLinksFromSubreddit(config Config, subRedditLink string) ([]string, error) {
	response, err := http.Get(subRedditLink)
	if err != nil {
		return []string{}, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []string{}, err
	}
	var redditJsonPage RedditResponse
	err = json.Unmarshal(body, &redditJsonPage)
	if err != nil {
		return []string{}, err
	}
	return redditJsonPage.GetLinks(config), nil
}

func DownloadImage(imageLink string, directory string) error {
	fileName := path.Join(directory, path.Base(imageLink))
	// check if file already exists
	_, err := os.Stat(fileName)
	if !os.IsNotExist(err) {
		log.Println("The file", path.Base(imageLink), "has already been downloaded")
		return nil
	} else {
		log.Println("Downloading file", path.Base(imageLink))
	}
	response, err := http.Get(imageLink)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(body)
	if err != nil {
		return err
	}
	return nil
}
