package main

import (
	"fmt"
	"sync"

	dri "github.com/anthony-barret/dri/download-reddit-images"
)

func main() {
	fmt.Println("DRI")
	subReddits := dri.ParseConfig("./conf.txt")
	var images []string
	images = make([]string, 0)
	for _, sub := range subReddits {
		imageLinks := dri.GetImagesLinksFromSubreddit(sub)
		fmt.Println(imageLinks)
		images = append(images, imageLinks...)
	}
	var wg sync.WaitGroup
	for _, img := range images {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dri.DownloadImage(img)
		}()
	}
	wg.Wait()
}
