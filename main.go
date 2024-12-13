package main

import (
	"fmt"
	"github.com/theheftyfine/adventofcode2024/day1"
	"github.com/theheftyfine/adventofcode2024/day10"
	"github.com/theheftyfine/adventofcode2024/day11"
	"github.com/theheftyfine/adventofcode2024/day12"
	"github.com/theheftyfine/adventofcode2024/day2"
	"github.com/theheftyfine/adventofcode2024/day3"
	"github.com/theheftyfine/adventofcode2024/day4"
	"github.com/theheftyfine/adventofcode2024/day5"
	"github.com/theheftyfine/adventofcode2024/day6"
	"github.com/theheftyfine/adventofcode2024/day7"
	"github.com/theheftyfine/adventofcode2024/day8"
	"github.com/theheftyfine/adventofcode2024/day9"
	"github.com/theheftyfine/adventofcode2024/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var days = []model.DayRunner{
	day1.NewDay("input/input1.txt"),
	day2.NewDay("input/input2.txt"),
	day3.NewDay("input/input3.txt"),
	day4.NewDay("input/input4.txt"),
	day5.NewDay("input/input5.txt"),
	day6.NewDay("input/input6.txt"),
	day7.NewDay("input/input7.txt"),
	day8.NewDay("input/input8.txt"),
	day9.NewDay("input/input9.txt"),
	day10.NewDay("input/input10.txt"),
	day11.NewDay("input/input11.txt"),
	day12.NewDay("input/input12.txt"),
}

const port = "8080"

func main() {
	//err := clipboard.Init()
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//a := app.New()
	//w := a.NewWindow("Advent of Code 2024")
	//w.Resize(fyne.NewSize(800, 600))
	//content := container.NewAppTabs()
	//for i, s := range screens {
	//	content.Append(container.NewTabItem("Day "+strconv.Itoa(i+1), s))
	//}
	//rows := container.NewVBox(content)
	//w.SetContent(rows)
	//
	//w.ShowAndRun()
	fs := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/day/", dayHandler)
	http.HandleFunc("/days", daysHandler)
	http.HandleFunc("/sidebar", sidebarHandler)
	http.Handle("/", fs)
	fmt.Println("Starting server at port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func sidebarHandler(w http.ResponseWriter, r *http.Request) {
	ids := []int{}
	for i := range days {
		ids = append(ids, i+1)
	}
	t, err := template.ParseFiles("./templates/sidebar.html")
	if err == nil {
		err = t.Execute(w, ids)
	} else {
		fmt.Println(err)
	}
}

func daysHandler(w http.ResponseWriter, r *http.Request) {
	ids := []int{}
	for i := range days {
		ids = append(ids, i+1)
	}
	t, err := template.ParseFiles("./templates/days.html")
	if err == nil {
		err = t.Execute(w, ids)
	} else {
		fmt.Println(err)
	}
}

func dayHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path[5:], "/")
	if len(parts) == 2 {
		day, err := strconv.Atoi(parts[0])
		part, err := strconv.Atoi(parts[1])
		if err == nil && day > 0 && day <= len(days) && part > 0 && part <= len(days[day-1].Parts()) {
			_, err = fmt.Fprintf(w,
				"<span id=\"result-%d\">Result: %d</span>",
				day, days[day-1].Parts()[part-1]())
			if err != nil {
				log.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else if len(parts) == 1 {
		i, err := strconv.Atoi(parts[0])
		fmt.Println(err)
		if err == nil {
			t, err := template.ParseFiles("./templates/day.html")
			fmt.Println(err)
			if err == nil {
				t.Execute(w, i)
			}
		}
	}
}
