package handle

import (
	"bytes"
	"html/template"
	"time"

	"github.com/squiidz/fishMe/module/database"
	"github.com/squiidz/fishMe/module/utility"
)

type User struct {
	Id       int
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

type ById []Fish

func (f ById) Len() int           { return len(f) }
func (f ById) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ById) Less(i, j int) bool { return f[i].Id > f[j].Id }

var db = database.SetupDB()

// Parse the fish.tmpl, execute it to a buffer, transform the buffer to a HTML string and return it !
func ParseFishFile(fishes []Fish) template.HTML {
	var buff bytes.Buffer

	fishTemp := template.Must(template.New("fishes").ParseFiles("template/fish.tmpl"))
	fishTemp.ExecuteTemplate(&buff, "fishes", fishes)
	fishData := template.HTML(buff.String())

	return fishData
}

func ParseSecureFishFile(fishes []Fish) template.HTML {
	var buff bytes.Buffer

	fishTemp := template.Must(template.New("fishes").ParseFiles("template/secureFish.tmpl"))
	fishTemp.ExecuteTemplate(&buff, "secureFishes", fishes)
	fishData := template.HTML(buff.String())

	return fishData
}

// End of the ParseFishFile Function !
func ParseNavbarFile(file string) string {
	home, err := utility.LoadTemplate(file) // Load the content of the home.pk file (Navbar)
	utility.ShitAppend(err)
	return string(home.Body)
}

// Search in DB
func FindUser(name string) *User {
	result := User{}
	err := db.QueryRow("SELECT * FROM users WHERE username = $1", name).Scan(
		&result.Id,
		&result.Username,
		&result.Email,
		&result.Password,
		&result.Date,
	)
	utility.ShitAppend(err)

	return &result
}

func FindFish(name string) []Fish {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM fish WHERE type = $1", name).Scan(&count)
	utility.ShitAppend(err)

	fishes := make([]Fish, count)

	for loop := 0; loop <= count-1; loop++ {
		//fishes := append(fishes, Fish{})
		err := db.QueryRow("SELECT * FROM fish WHERE type = $1", name).Scan(
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

	return fishes
}
