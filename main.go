package main

import (
	"flag"
	"sync"

	dri "github.com/anthony-barret/dri/download-reddit-images"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.yaml", "The configuration file")
	flag.Parse()
	config := dri.ParseConfig(configFile)
	var images []string
	images = make([]string, 0)
	for _, sub := range config.SubReddits {
		imageLinks := dri.GetImagesLinksFromSubreddit(config, sub)
		images = append(images, imageLinks...)
	}
	var wg sync.WaitGroup
	for _, img := range images {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dri.DownloadImage(img, config.Config.Directory)
		}()
	}
	wg.Wait()
}
