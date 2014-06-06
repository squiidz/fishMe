package handle

import (
	"PushKids/module/utility"
	"bytes"
	"html/template"
	"log"
	"net/http"
)

func Handler(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" { // Check if the request is for the root
		http.NotFound(rw, req)
		return
	}

	var _, er = req.Cookie("fishme")
	log.Println("[*]Handling Request from : " + req.RemoteAddr)

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
		log.Println("[*] " + req.RemoteAddr + " Redirected to index")
	} else {
		temp, err := template.ParseFiles("template/profil.html")
		utility.ShitAppend(err)

		temp = template.Must(temp.Parse(ParseNavbarFile("article/home"))) // Add The content of the home.pk to the current template

		temp.Execute(rw, nil)
	}
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	fishes := make([]Fish, 20)

	log.Println("[*]Handling Request from : " + req.RemoteAddr + " At [/home]")

	var cookie, er = req.Cookie("fishme")
	if er != nil {
		log.Println("[*]" + req.RemoteAddr + " not able to connect")
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
			log.Println("[*] Fish => " + fishes[loop].Type + " loaded")
		}
		fishes = fishes[0:count]
		log.Println("[*] Cookie value for " + req.RemoteAddr + " is " + cookie.Value)

		temp, err := template.ParseFiles("template/home.html") // Parse the home.html file
		utility.ShitAppend(err)

		userId := "{{define \"userId\"}}" + cookie.Value + "{{end}}" // Create a template on the fly to get the username with the cookie value

		temp = template.Must(temp.Parse(ParseNavbarFile("article/home"))) // Add The content of the home.pk to the current template
		temp = template.Must(temp.Parse(userId))                          // Same as above but for userId

		temp.Execute(rw, ParseFishFile(fishes)) // Execute the template and push it to the ResponseWrite

	}

}

// Parse the fish.tmpl, execute it to a buffer, transform the buffer to a HTML string and return it !
func ParseFishFile(fishes []Fish) template.HTML {
	var buff bytes.Buffer

	fishTemp := template.Must(template.New("fishes").ParseFiles("template/fish.tmpl"))
	fishTemp.ExecuteTemplate(&buff, "fishes", fishes)
	fishData := template.HTML(buff.String())

	return fishData
}

// End of the ParseFishFile Function !
func ParseNavbarFile(file string) string {
	home, err := utility.LoadPage(file) // Load the content of the home.pk file (Navbar)
	utility.ShitAppend(err)
	return string(home.Body)
}
