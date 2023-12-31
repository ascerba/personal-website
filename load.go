package main

import (
	"os"
	"regexp"
	"sort"
	"strings"
)

type Post struct {
	FileName string
	Title    string
	Date     string
	Tags     []string
	Image    string
}

type Posts struct {
	Contents []*Post
}

func (p Post) containsTag(filterTag string) bool {
	if filterTag == "" {
		return true
	}

	for _, tag := range p.Tags {
		if filterTag == tag {
			return true
		}
	}

	return false
}

func (app *application) loadPosts(location string, postCount int) (p *Posts, err error) {
	if postCount == 0 || postCount < -1 {
		return nil, os.ErrInvalid
	}

	var posts []*Post

	files, err := os.ReadDir(location)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".tmpl.html") {
			newPost, err := app.readFile(location + "/" + file.Name())
			if err != nil {
				return nil, err
			}
			/*
				// filtering by tag
				if !newPost.containsTag(filterTag) {
					continue
				} */

			posts = append(posts, newPost)
		}
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date > posts[j].Date
	})

	if postCount == -1 {
		return &Posts{Contents: posts}, nil
	} else if postCount < len(posts) {
		return &Posts{Contents: posts[:postCount]}, nil
	} else {
		return &Posts{Contents: posts}, nil
	}
}

func (app *application) readFile(location string) (p *Post, err error) {
	fileContent, err := os.ReadFile(location)
	if err != nil {
		return nil, err
	}

	var tmp *Post = new(Post)

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
	tagsMatching := tagsPattern.FindStringSubmatch(string(fileContent))

	var tags []string
	if len(tagsMatching) > 1 {
		tags = strings.Fields(tagsMatching[1])
	} else {
		tags = []string{}
	}

	// thumbnail image
	imagePattern := regexp.MustCompile(`<img src="(.+)" class="mainImage"( alt="(.+)")* />`)
	imageMatching := imagePattern.FindStringSubmatch(string(fileContent))

	var image string
	if len(imageMatching) > 1 {
		image = imageMatching[0]
	} else {
		image = ""
	}

	tmp.FileName = fileName
	tmp.Title = title
	tmp.Date = date
	tmp.Tags = tags
	tmp.Image = image

	return tmp, nil

}
