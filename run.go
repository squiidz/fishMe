package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Page struct {
	Title string
	Body  []byte
}

type User struct {
	Username string
	Email    string
	Password string
	Date     time.Time
}

type Fish struct {
	Type     string
	Username string
	Weight   int
	Length   int
	Location string
	Date     time.Time
	Lure     string
	Info     string
}

/*
	func (p *Page) save(path string) error {
		filename := path + "/" + p.Title + ".pk"
		return ioutil.WriteFile(filename, p.Body, 0600)
	}
*/

func loadPage(title string) (*Page, error) {
	filename := title + ".pk"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func SetupDB() *sql.DB {
	DB, err := sql.Open("postgres", "host=fishme.io user=squiidz password=Bazooka4me dbname=fishme sslmode=disable")
	shitAppend(err)
	return DB
}

func shitAppend(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/home", HomeHandler)
	/*
		http.HandleFunc("/home", homeHandler)
		http.HandleFunc("/admin", adminHandler)
		http.HandleFunc("/Sign", SignUp)
		http.HandleFunc("/addFish", FishForm)
	*/

	http.HandleFunc("/add", AddUser)
	http.HandleFunc("/fish", AddFish)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.ListenAndServe(":8080", nil)
}

func Handler(rw http.ResponseWriter, req *http.Request) {
	file, err := loadPage("article/push")
	shitAppend(err)

	temp, err := template.ParseFiles("template/index.html")
	shitAppend(err)

	fmt.Println("[*]Handling Request from : " + req.RemoteAddr)
	temp.Execute(rw, file)
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	file, err := loadPage("article/push")
	shitAppend(err)

	temp, err := template.ParseFiles("template/home.html")
	shitAppend(err)

	fmt.Println("[*]Handling Request from : " + req.RemoteAddr)
	temp.Execute(rw, file)
}

/*
	func adminHandler(rw http.ResponseWriter, req *http.Request) {
		temp, _ := template.ParseFiles("template/admin.html")
		temp.Execute(rw, nil)
		if req.Method == "POST" {
			page := &Page{Title: req.FormValue("title"), Body: []byte(req.FormValue("content"))}
			page.save("article")
		}
	}


	func homeHandler(rw http.ResponseWriter, req *http.Request) {
		newPage := &Page{"Your at Home", []byte("This is home !!")}
		temp, _ := template.ParseFiles("template/home.html")
		temp.Execute(rw, newPage)
	}
*/

func AddUser(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		db := SetupDB()
		user := User{
			Username: req.FormValue("username"),
			Email:    req.FormValue("email"),
			Password: req.FormValue("password"),
			Date:     time.Now()}

		rows, err := db.Query("INSERT INTO users VALUES($1, $2, $3, $4, $5)", 3, user.Username, user.Email, user.Password, user.Date)
		shitAppend(err)
		defer rows.Close()
		http.Redirect(rw, req, "/", http.StatusFound)
	}
}

func AddFish(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		db := SetupDB()

		poid, err := strconv.Atoi(req.FormValue("weigth"))
		shitAppend(err)
		long, err := strconv.Atoi(req.FormValue("length"))
		shitAppend(err)

		fish := Fish{
			Type:     req.FormValue("type"),
			Username: req.FormValue("username"),
			Weight:   poid,
			Length:   long,
			Location: req.FormValue("location"),
			Date:     time.Now(),
			Lure:     req.FormValue("lure"),
			Info:     req.FormValue("info")}

		rows, err := db.Query("INSERT INTO fish VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
			8,
			fish.Type,
			fish.Username,
			fish.Weight,
			fish.Length,
			fish.Location,
			fish.Date,
			fish.Lure,
			fish.Info)
		shitAppend(err)
		defer rows.Close()
		http.Redirect(rw, req, "/", http.StatusFound)
	}
}

/*
	func FishForm(rw http.ResponseWriter, req *http.Request) {
		temp, _ := template.ParseFiles("template/fish.html")
		temp.Execute(rw, nil)
	}

	func SignUp(rw http.ResponseWriter, req *http.Request) {
		temp, _ := template.ParseFiles("template/add.html")
		temp.Execute(rw, nil)
	}
*/
