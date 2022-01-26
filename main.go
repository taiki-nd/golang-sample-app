package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "title", Body: []byte("this is a sample page")}
	p1.Save() // title.txtで"this is a sample page"と書いてあるファイルを作成する

	p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))
}
