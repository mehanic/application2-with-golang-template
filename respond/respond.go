package respond

import (
	"bytes"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"server-application3/function"
	"server-application3/models"
	"strings"
	"time"
)

// var templates = template.Must(template.ParseGlob("templates/*.html"))

func InternalServerError(conn net.Conn) {
	body := "500 Internal Server Error"
	fmt.Fprintf(conn, "HTTP/1.1 500 Internal Server Error\r\n"+
		"Content-Type: text/plain\r\n"+
		"Content-Length: %d\r\n\r\n%s", len(body), body)
}

func BadRequest(conn net.Conn) {
	body := "400 Bad Request"
	fmt.Fprintf(conn, "HTTP/1.1 400 Bad Request\r\n"+
		"Content-Type: text/plain\r\n"+
		"Content-Length: %d\r\n\r\n%s", len(body), body)
}

var fm = template.FuncMap{
	"uc":       strings.ToUpper,
	"ft":       function.FirstThree,
	"fdateMDY": function.MonthDayYear,
	"fdbl":     function.Double,
	"fsq":      function.Square,
	"fsqrt":    function.SqRoot,
	"sub": func(a, b int) int { // 👈 додано
		return a - b
	},
}

var templates = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.html"))

func Respond(w http.ResponseWriter, r *http.Request) {
	// func Respond(conn net.Conn, method, path string, keepAlive bool, form map[string]string) {
	// statusLine := "HTTP/1.1 200 OK\r\n"
	var tmpl string
	var data interface{}

	switch r.URL.Path { // comented switch path
	case "/":
		tmpl = "index.html"
		thinkers := []models.SecondSaga{
			{Name2: "Confucius", Motto2: "Everything has beauty, but not everyone sees it."},
			{Name2: "Socrates", Motto2: "The unexamined life is not worth living."},
			{Name2: "Lao Tzu", Motto2: "A journey of a thousand miles begins with a single step."},
		}

		vehicles := []models.SecondCar{
			{Manufacturer: "Audi", Model: "A8", Doors: 4},
			{Manufacturer: "Mercedes-Benz", Model: "GLE", Doors: 5},
		}
		data = models.NewData{
			PageData: models.PageData{
				Title:   "Home",
				Heading: "Welcome to the Home Page",
				Message: "This is served with Go templates.",
			},

			Thinkers: thinkers,
			Vehicles: vehicles,
			SomeDate: time.Now(),
			Number:   8,
			FloatNum: 49.0,
		}
		if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
			http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
			return
		}

	case "/about":
		tmpl = "about.html"

		wisdomList := []models.Sagen{
			{Name1: "Buddha", Motto1: "The belief of no beliefs"},
			{Name1: "Gandhi", Motto1: "Be the change"},
			{Name1: "Martin Luther King", Motto1: "Hatred never ceases with hatred but with love alone is healed."},
		}

		transportList := []models.Car{
			{Manufacturer: "Tesla", Model: "Model S", Doors: 4},
			{Manufacturer: "BMW", Model: "X5", Doors: 5},
		}
		data = models.AboutData{
			Information: models.Information{
				PageData: models.PageData{
					Title:   "About",
					Heading: "About Page",
					Message: "This is a custom server in Go.",
				},
				Items: map[string]string{
					"India":    "Gandhi",
					"America":  "MLK",
					"Meditate": "Buddha",
					"Love":     "Jesus",
					"Prophet":  "Muhammad",
				},
				Sage: models.Sage{
					Name:  "the result with which are very famous",
					Motto: "we try to do infrustucture",
				},
			},
			Wisdom:    wisdomList,
			Transport: transportList,
		}
	case "/contact":
		tmpl = "contact.html"
		// data = map[string]interface{}{
		// "Title":   "Contact",
		// "Heading": "Contact Page",
		// "Message": "Email: example@example.com",
		// }
		xs := []string{"zero", "one", "two", "three", "four", "five"}
		users := []models.User{
			{Name: "Buddha", Motto: "The belief of no beliefs", Admin: false},
			{Name: "Gandhi", Motto: "Be the change", Admin: true},
			{Name: "", Motto: "Nobody", Admin: true},
		}

		g1 := struct {
			Score1 int
			Score2 int
		}{
			7,
			9,
		}
		// data = models.PageData{
		// Title:   "Contact",
		// Heading: "Contact Page",
		// Message: "Email: example@example.com",
		// }
		// anonim function
		data = struct {
			models.PageData
			Words  []string
			Lname  string
			Users  []models.User
			Scores struct {
				Score1 int
				Score2 int
			}
		}{
			PageData: models.PageData{
				Title:   "Contact",
				Heading: "Contact Page",
				Message: "Email: example@example.com",
			},
			Words:  xs,
			Lname:  "McLeod",
			Users:  users,
			Scores: g1,
		}

	case "/information":
		tmpl = "education.html"
		// data = models.PageData{
		// Title:   "Home",
		// Heading: "Welcome to the Home Page",
		// Message: "This is served with Go templates.",
		// }
		years := []models.Year{
			{
				AcaYear: "2020-2021",
				Fall: models.Semester{
					Term: "Fall",
					Courses: []models.Course{
						{Number: "CSCI-40", Name: "Introduction to Programming in Go", Units: "4"},
						{Number: "CSCI-130", Name: "Introduction to Web Programming with Go", Units: "4"},
						{Number: "CSCI-140", Name: "Mobile Apps Using Go", Units: "4"},
					},
				},
				Spring: models.Semester{
					Term: "Spring",
					Courses: []models.Course{
						{Number: "CSCI-50", Name: "Advanced Go", Units: "5"},
						{Number: "CSCI-190", Name: "Advanced Web Programming with Go", Units: "5"},
						{Number: "CSCI-191", Name: "Advanced Mobile Apps With Go", Units: "5"},
					},
				},
			},
			{
				AcaYear: "2021-2022",
				Fall: models.Semester{
					Term: "Fall",
					Courses: []models.Course{
						{Number: "CSCI-40", Name: "Introduction to Programming in Go", Units: "4"},
						{Number: "CSCI-130", Name: "Introduction to Web Programming with Go", Units: "4"},
						{Number: "CSCI-140", Name: "Mobile Apps Using Go", Units: "4"},
					},
				},
				Spring: models.Semester{
					Term: "Spring",
					Courses: []models.Course{
						{Number: "CSCI-50", Name: "Advanced Go", Units: "5"},
						{Number: "CSCI-190", Name: "Advanced Web Programming with Go", Units: "5"},
						{Number: "CSCI-191", Name: "Advanced Mobile Apps With Go", Units: "5"},
					},
				},
			},
		}

		data = models.EducationData{
			PageData: models.PageData{
				Title:   "Education",
				Heading: "Education Information",
				Message: "This is served with Go templates.",
			},
			Years: years,
		}

		if err := templates.ExecuteTemplate(w, tmpl, data); err != nil {
			http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	case "/relax":
		tmpl = "relax.html"
		// data = models.PageData{
		// Title:   "Home",
		// Heading: "Welcome to the Home Page",
		// Message: "This is served with Go templates.",
		// }

		regions := []models.Region{
			{
				Region: "Southern",
				Hotels: []models.Hotel{
					{Name: "Hotel California", Address: "42 Sunset Boulevard", City: "Los Angeles", Zip: "95612", Region: "southern"},
					{Name: "H", Address: "4", City: "L", Zip: "95612", Region: "southern"},
				},
			},
			{
				Region: "Northern",
				Hotels: []models.Hotel{
					{Name: "Hotel North", Address: "123 Winter St", City: "San Francisco", Zip: "94101", Region: "northern"},
					{Name: "H", Address: "4", City: "L", Zip: "95612", Region: "northern"},
				},
			},
			{
				Region: "Central",
				Hotels: []models.Hotel{
					{Name: "Hotel Central", Address: "99 Main Ave", City: "Sacramento", Zip: "94203", Region: "central"},
					{Name: "H", Address: "4", City: "L", Zip: "95612", Region: "central"},
				},
			},
		}
		menu := models.Items{
			{Name: "Oatmeal", Descrip: "yum yum", Meal: "Breakfast", Price: 4.95},
			{Name: "Hamburger", Descrip: "Delicious good eating for you", Meal: "Lunch", Price: 6.95},
			{Name: "Pasta Bolognese", Descrip: "From Italy delicious eating", Meal: "Dinner", Price: 7.95},
		}

		data = models.RelaxData{
			PageData: models.PageData{
				Title:   "Relax",
				Heading: "Relax & Hotels",
				Message: "Choose your region and find hotels.",
			},
			Regions: regions,
			Menu:    menu,
		}

		if err := templates.ExecuteTemplate(w, tmpl, data); err != nil {
			http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	case "/date":
		tmpl = "data.html"
		data = models.PageData{
			Title:   "Home",
			Heading: "Welcome to the Home Page",
			Message: "This is served with Go templates.",
		}
	case "/menu":
		tmpl = "menu.html"
		m := []models.Meal{
			{
				Meal: "Breakfast",
				Item: []models.Item{
					{Name: "Oatmeal", Descrip: "yum yum", Price: 4.95},
					{Name: "Cheerios", Descrip: "American eating food traditional now", Price: 3.95},
					{Name: "Juice Orange", Descrip: "Delicious drinking in throat squeezed fresh", Price: 2.95},
				},
			},
			{
				Meal: "Lunch",
				Item: []models.Item{
					{Name: "Hamburger", Descrip: "Delicous good eating for you", Price: 6.95},
					{Name: "Cheese Melted Sandwich", Descrip: "Make cheese bread melt grease hot", Price: 3.95},
					{Name: "French Fries", Descrip: "French eat potatoe fingers", Price: 2.95},
				},
			},
			{
				Meal: "Dinner",
				Item: []models.Item{
					{Name: "Pasta Bolognese", Descrip: "From Italy delicious eating", Price: 7.95},
					{Name: "Steak", Descrip: "Dead cow grilled bloody", Price: 13.95},
					{Name: "Bistro Potatoe", Descrip: "Bistro bar wood American bacon", Price: 6.95},
				},
			},
		}
		data = m

	case "/submit":
		tmpl = "submit.html"
		if r.Method == http.MethodPost {
			r.ParseForm()
			form := make(map[string]string)
			for k, v := range r.PostForm {
				form[k] = v[0]
			}
			data = models.PageData{
				Title: "Form Submitted",
				Form:  form,
			}
		} else {
			data = models.PageData{
				Title: "Submit Form",
			}
		}

	default:
		// statusLine = "HTTP/1.1 404 Not Found\r\n"
		tmpl = "default.html"
		data = models.PageData{
			StatusLine: "HTTP/1.1 404 Not Found\r\n",
			Heading:    "Not Found",
			Message:    "The page you requested does not exist.",
		}
	}

	var buf bytes.Buffer
	// if err := templates.ExecuteTemplate(&buf, tmpl, data); err != nil {
	// fmt.Fprintf(conn, "HTTP/1.1 500 Internal Server Error\r\nContent-Type: text/plain\r\n\r\nTemplate error: %v", err)
	// return
	// }

	// fmt.Fprint(conn, statusLine)
	// fmt.Fprintf(conn, "Content-Length: %d\r\n", buf.Len())
	// fmt.Fprint(conn, "Content-Type: text/html\r\n")
	// if keepAlive {
	// fmt.Fprint(conn, "Connection: keep-alive\r\n")
	// fmt.Fprint(conn, "Keep-Alive: timeout=60\r\n")
	// } else {
	// fmt.Fprint(conn, "Connection: close\r\n")
	// }
	// fmt.Fprint(conn, "\r\n")

	// buf.WriteTo(conn)
	// }

	if err := templates.ExecuteTemplate(&buf, tmpl, data); err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Просто пишем буфер в ResponseWriter
	buf.WriteTo(w)
}
