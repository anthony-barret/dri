package main

import (
	"flag"
	"log"
	"sync"

	dri "github.com/anthony-barret/dri/download-reddit-images"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.yaml", "The configuration file")
	flag.Parse()
	config, err := dri.ParseConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	var images []string
	images = make([]string, 0)
	for _, sub := range config.SubReddits {
		imageLinks, err := dri.GetImagesLinksFromSubreddit(config, sub)
		if err != nil {
			log.Println("ERROR:", err)
		}
		images = append(images, imageLinks...)
	}
	var wg sync.WaitGroup
	for _, img := range images {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := dri.DownloadImage(img, config.Config.Directory)
			if err != nil {
				log.Println("ERROR:", err)
			}
		}()
	}
	wg.Wait()
}
