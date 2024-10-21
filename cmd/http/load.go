package main

import (
	"os"
	"regexp"
	"sort"
	"strings"
)

// Post struct contains necessary data for a post
type Post struct {
	FileName string
	Title    string
	Date     string
	Tags     []string
	Image    string
}

// Posts stuct contains a collection of type Post
type Posts struct {
	Contents []*Post
}

// Read all found files and load them into a stuct
func (app *application) aggregate(location string) (p *Posts, err error) {
	var posts []*Post

	files, err := os.ReadDir(location)
	if err != nil {
		return nil, err
	}

	// Loop over every file in the directory and read the contents.
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".tmpl.html") {
			newPost, err := app.readFile(location + "/" + file.Name())
			if err != nil {
				return nil, err
			}

			posts = append(posts, newPost)
		}
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date > posts[j].Date
	})

	return &Posts{Contents: posts}, nil
}

func (app *application) readFile(location string) (p *Post, err error) {
	fileContent, err := os.ReadFile(location)
	if err != nil {
		return nil, err
	}

	var post *Post = new(Post)

	fileName := strings.TrimSuffix(strings.Split(location, "/")[2], ".tmpl.html")

	// title
	title := strings.ReplaceAll(fileName, "_", " ")

	// date
	datePattern := regexp.MustCompile(`{{define "uploaded-on"}}(\d{4}-\d{2}-\d{2}){{end}}`)
	dateMatching := datePattern.FindStringSubmatch(string(fileContent))

	var date string
	if len(dateMatching) > 1 {
		date = dateMatching[1]
	} else {
		date = ""
	}

	// tags
	tagsPattern := regexp.MustCompile(`{{define "keywords"}}([\w\s]+){{end}}`)
	matchingTags := tagsPattern.FindStringSubmatch(string(fileContent))

	var tags []string
	if len(matchingTags) > 1 {
		tags = strings.Fields(matchingTags[1])
	} else {
		tags = []string{}
	}

	// thumbnail image
	imagePattern := regexp.MustCompile(`{{define "thumbnail"}}<img src="(.+)" class="mainImage"( alt="(.+)")* />{{end}}`)
	matchingImage := imagePattern.FindStringSubmatch(string(fileContent))

	var image string
	if len(matchingImage) > 1 {
		image = strings.Trim(matchingImage[0], "{{define \"thumbnail\"}}")
	} else {
		image = ""
	}

	post.FileName = fileName
	post.Title = title
	post.Date = date
	post.Tags = tags
	post.Image = image

	return post, nil

}
