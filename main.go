package main

import(
	"net/http"
	"fmt"
	"html/template"
	"log"
	"time"
)

type refData struct{
	FirstName string
	LastName string
	ThirdName string
	Date string
	Title string
	Location string
	Publisher string
	Chapter string
	Volume string
	Page string
	Url string
}

type retVal struct{
	Apa string
}

func MakeApa(data refData) string {
	retVal := ""
	retVal = data.LastName
	if data.FirstName != "" {
		retVal+= ", "+string(data.FirstName[0])
	}
	fmt.Print(data.FirstName)
	if data.ThirdName != "" {
		retVal += ". "
		retVal+= string(data.ThirdName[0])
		retVal +=". "
	}
	if data.Date != "" {
		layout := "2006-01-02"
		t, _ := time.Parse(layout, data.Date)

		retVal += "("+t.Format("02 Jan 2006")+")"
		retVal +=". "
	}
	if data.Title != "" {
		retVal += data.Title
		retVal +=". "
	}
	if data.Location != "" {
		retVal += data.Location
		retVal +=": "
	}
	if data.Publisher != "" {
		retVal += data.Publisher
	}
	if data.Volume != "" {
		retVal +=", "
		retVal += "vol"+data.Volume
	}
	if data.Chapter != "" {
		retVal +=", "
		retVal += "ch"+data.Chapter
	}

	if data.Page != "" {
		retVal +=", "
		retVal += "page"+data.Page
	}
	if data.Url != "" {
		retVal +=". "
		retVal += "Retrieved from: "
		retVal += data.Url
	}
	return retVal
}

func getStyle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == "POST" {
			w.WriteHeader(200)
			r.ParseForm()


			var newRefData refData
			newRefData.FirstName = r.FormValue("firstname")
			newRefData.LastName = r.FormValue("lastname")
			newRefData.ThirdName = r.FormValue("thirdname")
			newRefData.Date = r.FormValue("date")
			newRefData.Location = r.FormValue("location")
			newRefData.Publisher = r.FormValue("publisher")
			newRefData.Url = r.FormValue("url")
			newRefData.Chapter = r.FormValue("chapter")
			newRefData.Volume = r.FormValue("volume")
			newRefData.Page = r.FormValue("page")
			newRefData.Title = r.FormValue("title")


			w.Header().Set("Content-Type", "text/html")
			var returnVal retVal
			returnVal.Apa = MakeApa(newRefData)
			t, _ := template.ParseFiles("index.html")
			t.Execute(w, returnVal)
		} else if r.Method == "GET" {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "text/html")
			t, _ := template.ParseFiles("index.html")
			t.Execute(w, nil)
		} else {
			w.WriteHeader(400)
			fmt.Fprintf(w,"400 ;( Bad request!")
		}
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w,"404 :( page not found")
	}
}

func main(){
	http.HandleFunc("/", getStyle)
	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		log.Fatal("Listen And Serve: ", err)
	}
}