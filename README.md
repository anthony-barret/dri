# DownloadRedditImages

## About this project

The goal of this project is to provide a simple program to download the images
from multiple subreddits.

## How to build

In order to build this project you have to have go installed on your machine, then run:
```sh
go build
```

You can simply execute the program with:
```sh
./dri
```

Command line option:
* `--config`: the configuration file to use, by default it will use the file
named `config.yaml` in the current directory.

## Configuration file example

```yaml
config:
  directory: img
  limit: 10
  over_18: no

subreddits:
  - memes
  - programminghumour
```

The configuration is in the YAML format and is composed of 2 sections.

First, the section `config` is composed of the parameter `directory` which is
the name of the directory to save the images downloaded during the program
execution, `limit` the maximum number of images to download by subreddits,
`over_18` allows to filter NSFW contents.

The section `subreddits` is simply a list of subreddits that you want to
download the images from.
