package handle

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/squiidz/fishMe/module/utility"
)

func Handler(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" { // Check if the request is for the root
		http.NotFound(rw, req)
		return
	}
	var _, er = req.Cookie("fishme")

	if er != nil {
		temp, err := template.ParseFiles("template/index.html")
		signin, err := utility.LoadPage("article/signin")
		utility.ShitAppend(err)
		SignButton := template.HTML(string(signin.Body))
		temp.Execute(rw, SignButton)
	} else {
		http.Redirect(rw, req, "/home", http.StatusFound)
	}
}

func ProfilHandler(rw http.ResponseWriter, req *http.Request) {
	var _, er = req.Cookie("fishme")
	if er != nil {
		http.Redirect(rw, req, "/", http.StatusFound)
	} else {
		temp, err := template.ParseFiles("template/profil.html")
		utility.ShitAppend(err)

		temp = template.Must(temp.Parse(ParseNavbarFile("template/navbar"))) // Add The content of the navbar.tmpl to the current template
		temp = template.Must(temp.ParseFiles("template/add_fish.tmpl"))

		temp.Execute(rw, nil)
	}
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	fishes := make([]Fish, 20)

	var cookie, er = req.Cookie("fishme")
	if er != nil {
		http.Redirect(rw, req, "/", http.StatusFound)
	} else {
		cookieVal := cookie.Value
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM fish WHERE username = $1", cookieVal).Scan(&count)
		utility.ShitAppend(err)
		for loop := 0; loop <= count-1; loop++ {
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
		}

		fishes = fishes[0:count]
		sort.Sort(ById(fishes))

		temp, err := template.ParseFiles("template/home.html") // Parse the home.html file
		utility.ShitAppend(err)

		userId := "{{define \"userId\"}}" + cookie.Value + "{{end}}" // Create a template on the fly to get the username with the cookie value

		temp = template.Must(temp.Parse(ParseNavbarFile("template/navbar"))) // Add The content of the navbar.tmpl to the current template
		temp = template.Must(temp.Parse(userId))                             // Same as above but for userId
		temp = template.Must(temp.ParseFiles("template/add_fish.tmpl"))      // Add Fish modal
		temp = template.Must(temp.ParseFiles("template/map.tmpl"))           // Map template

		temp.Execute(rw, ParseFishFile(fishes)) // Execute the template and push it to the ResponseWrite

	}

}

func FindHandler(rw http.ResponseWriter, req *http.Request) {
	var _, er = req.Cookie("fishme")

	if er != nil {
		http.Redirect(rw, req, "/", http.StatusFound)
	} else {
		term := req.FormValue("find")
		if term != "" {
			term = strings.ToLower(term)
			term = strings.Replace(term, term[:1], strings.ToUpper(term[:1]), 1)
			log.Println(term)

			temp, err := template.ParseFiles("template/find.html")
			utility.ShitAppend(err)

			temp = template.Must(temp.Parse(ParseNavbarFile("template/navbar"))) // Add The content of the navbar.tmpl to the current template
			temp = template.Must(temp.ParseFiles("template/add_fish.tmpl"))
			temp = template.Must(temp.ParseFiles("template/map.tmpl")) // Map template

			fishFind := FindFish(term)

			temp.Execute(rw, ParseSecureFishFile(fishFind))
		} else {
			http.Redirect(rw, req, "/home", http.StatusFound)
		}

	}
}
