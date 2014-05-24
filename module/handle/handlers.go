package handle

import (
	"PushKids/module/database"
	"PushKids/module/resize"
	"PushKids/module/utility"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Username string
	Email    string
	Password string
	Date     time.Time
}

type Fish struct {
	Id       int
	Type     string
	Username string
	Weight   float64
	Length   float64
	Location string
	Date     string
	Lure     string
	Info     string
	Picture  string
}

var db = database.SetupDB()

func Handler(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" { // Check if the request is for the root
		http.NotFound(rw, req)
		return
	}

	temp, err := template.ParseFiles("template/index.html")
	utility.ShitAppend(err)

	fmt.Println("[*]Handling Request from : " + req.RemoteAddr)
	var _, er = req.Cookie("fishme")
	if er != nil {
		signin, err := utility.LoadPage("article/signin")
		utility.ShitAppend(err)
		SignButton := template.HTML(string(signin.Body))
		temp.Execute(rw, SignButton)
	} else {
		home, err := utility.LoadPage("article/home")
		utility.ShitAppend(err)
		HomeButton := template.HTML(string(home.Body))
		temp.Execute(rw, HomeButton)
	}
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	fishes := make([]Fish, 20)

	temp, err := template.ParseFiles("template/home.html")
	utility.ShitAppend(err)
	fmt.Println("[*]Handling Request from : " + req.RemoteAddr + " At [/home]")

	var cookie, er = req.Cookie("fishme")
	if er != nil {
		fmt.Println("[*]" + req.RemoteAddr + " not able to connect")
		http.Redirect(rw, req, "/", http.StatusFound)
	} else {
		cookieVal := cookie.Value
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM fish WHERE username = $1", cookieVal).Scan(&count)
		utility.ShitAppend(err)
		for loop := 0; loop <= count; loop++ {
			//fishes := append(fishes, make(Fish{}))
			err := db.QueryRow("SELECT * FROM fish WHERE username = $1 LIMIT 1 OFFSET $2", cookieVal, loop).Scan(
				&fishes[loop].Id,
				&fishes[loop].Type,
				&fishes[loop].Username,
				&fishes[loop].Weight,
				&fishes[loop].Length,
				&fishes[loop].Location,
				&fishes[loop].Date,
				&fishes[loop].Lure,
				&fishes[loop].Info,
				&fishes[loop].Picture)

			if fishes[loop].Picture == "" {
				fishes[loop].Picture = "img/fish/" + fishes[loop].Type + ".jpg"
			}

			utility.ShitAppend(err)
			fmt.Println("[*] Fish => " + fishes[loop].Type + " loaded")
		}
		fishes = fishes[0:count]
		fmt.Println("[*] Cookie value for " + req.RemoteAddr + " is " + cookie.Value)
		temp.Execute(rw, fishes)
	}

}

func SignIn(rw http.ResponseWriter, req *http.Request) {

	username := req.FormValue("username")
	password := req.FormValue("password")
	user := User{}
	err := db.QueryRow("SELECT username,password FROM users WHERE username = $1 AND password = $2", username, password).Scan(&user.Username, &user.Password)
	if err != nil {
		http.Redirect(rw, req, "/", http.StatusFound)
	}
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "fishme", Value: user.Username, Path: "/", Expires: expire, MaxAge: 86400}
	http.SetCookie(rw, &cookie)
	http.Redirect(rw, req, "/home", http.StatusFound)
}

func LogOut(rw http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: "fishme", Path: "/", MaxAge: -1}
	http.SetCookie(rw, &cookie)
	http.Redirect(rw, req, "/", http.StatusFound)
}

func AddUser(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		user := User{
			Username: req.FormValue("username"),
			Email:    req.FormValue("email"),
			Password: req.FormValue("password"),
			Date:     time.Now()}

		var userCheck, mailCheck string

		err := db.QueryRow("SELECT username, email FROM users WHERE username = $1 AND email= $2", user.Username, user.Email).Scan(&userCheck, &mailCheck)
		utility.ShitAppend(err)
		if userCheck != "" || mailCheck != "" {
			http.Redirect(rw, req, "/", http.StatusFound)

		} else {
			rows, erro := db.Query("INSERT INTO users(username, email, password, date) VALUES($1, $2, $3, $4)",
				user.Username,
				user.Email,
				user.Password,
				user.Date)

			utility.ShitAppend(erro)
			defer rows.Close()
			utility.Mkdir("img/users/" + user.Username)
			timeNow := time.Now().Format(time.RFC822)
			_, er := db.Query("INSERT INTO fish(type, username, weight, length, location, date, lure, info, picture) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
				"Goldfish",
				user.Username,
				0.2,
				1.2,
				"Toilette",
				timeNow,
				"Net",
				"Just the Beginning !",
				"")
			utility.ShitAppend(er)

			http.Redirect(rw, req, "/", http.StatusFound)
		}
	}
}

func AddFish(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var cookie, er = req.Cookie("fishme")
		utility.ShitAppend(er)
		species := req.FormValue("type")
		var path string
		// Picture Upload
		file, handler, err := req.FormFile("picture")
		utility.ShitAppend(err)
		if err != nil {
			fmt.Println("[*] NO PICTURE UPLOADED")
			path = "img/fish/" + species + ".jpg"
		} else {
			data, err := ioutil.ReadAll(file)
			utility.ShitAppend(err)
			path = "img/users/" + cookie.Value + "/" + handler.Filename
			err = ioutil.WriteFile(path, data, 0777)
			utility.ShitAppend(err)
			resize.ResizeMe(path)
		}
		// End

		poid, err := strconv.ParseFloat(req.FormValue("weigth"), 32)
		utility.ShitAppend(err)

		long, err := strconv.ParseFloat(req.FormValue("length"), 32)
		utility.ShitAppend(err)

		timeNow := time.Now().Format(time.RFC822)
		fmt.Println(timeNow)

		fish := Fish{
			Type:     species,
			Username: cookie.Value,
			Weight:   poid,
			Length:   long,
			Location: req.FormValue("location"),
			Date:     timeNow,
			Lure:     req.FormValue("lure"),
			Info:     req.FormValue("info"),
			Picture:  path}

		rows, err := db.Query("INSERT INTO fish(type, username, weight, length, location, date, lure, info, picture) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
			fish.Type,
			fish.Username,
			fish.Weight,
			fish.Length,
			fish.Location,
			fish.Date,
			fish.Lure,
			fish.Info,
			fish.Picture)
		utility.ShitAppend(err)
		defer rows.Close()
		http.Redirect(rw, req, "/home", http.StatusFound)
	}
}
