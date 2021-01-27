package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	return ioutil.WriteFile(p.Title, p.Body, 0666)
}
func (p *Page) load(title string) (err error) {
	p.Title = title
	p.Body, err = ioutil.ReadFile(p.Title)
	return err
}
func main() {
	page := Page{
		"test.md",
		[]byte("# Page\n## Section1\nThis is section1."),
	}
	page.save()

	var new_page Page
	new_page.load(page.Title)
	fmt.Println(string(new_page.Body))
}
