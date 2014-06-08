package utility

import (
	"fmt"
	"io/ioutil"
	"os"
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

func LoadTemplate(title string) (*Page, error) {
	filename := title + ".tmpl"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func Mkdir(name string) {
	err := os.Mkdir(name, 0777)
	if err != nil {
		fmt.Println("[*] ERROR WHEN CREATING DIRECTORY : " + name)
	} else {
		fmt.Println("[+] DIRECTORY CREATED => " + name)
	}
}

func SaveFile(data string) {

}
