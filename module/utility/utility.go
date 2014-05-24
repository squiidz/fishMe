package utility

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func ShitAppend(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".pk"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func SaveFile(data string) {

}
