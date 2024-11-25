// Steps:

// 1. Parse Json File *
// 2. Create ds to store data *
// 3. Create html template file
// 4. http Handler for the file
// 5. Conditional Rendering

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var cyoaTemplate = template.Must(template.ParseFiles("./cyoa.html"))

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Story map[string]StoryArc

func main() {
	f, err := os.Open("./data.json")
	defer f.Close()
	if err != nil {
		fmt.Println("Err in the opening file: ", err)
		return
	}

	jsonData, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Err in the reading file: ", err)
		return
	}

	var story Story
	err = json.Unmarshal(jsonData, &story)
	if err != nil {
		fmt.Println("Error Unmarshaling JSON:", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlUnits := strings.Split(r.URL.Path, "/")
		if len(urlUnits) > 2 {
            fmt.Println("Path is not valid")
		}
        key := urlUnits[1]

        if key == "" {
            key = "intro"
        }

        data, ok:= story[key]
        if !ok {
            http.NotFound(w,r)
            return
        }

		err := cyoaTemplate.ExecuteTemplate(w, "cyoa.html", data)
		if err != nil {
			fmt.Println("Error in the rendering the template", err)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
