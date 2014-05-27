package handle

import (
	"PushKids/module/resize"
	"PushKids/module/utility"
	"io/ioutil"
	"net/http"
	//"os"
	"strconv"
	"time"
)

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

func DeleteFish(rw http.ResponseWriter, req *http.Request) {

	picture := req.FormValue("picture")
	id, err := strconv.Atoi(req.FormValue("id"))
	utility.ShitAppend(err)

	fish := Fish{Id: id, Picture: picture}

	_, err = db.Query("DELETE FROM fish WHERE id = $1", fish.Id)
	utility.ShitAppend(err)

	//err = os.Remove(fish.Picture)
	//utility.ShitAppend(err)

	http.Redirect(rw, req, "/home", http.StatusFound)
}
