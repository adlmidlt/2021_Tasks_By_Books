package ex10

import (
	"fmt"
	"log"
	"todo/todo/ex10/calendar"
)

func DateMain() {
	event := calendar.Event{}
	if err := event.SetTitle("Mom's birthday"); err != nil {
		log.Fatal(err)
	}

	if err := event.SetYear(2019); err != nil {
		log.Fatal(err)
	}

	if err := event.SetMonth(5); err != nil {
		log.Fatal(err)
	}

	if err := event.SetDay(22); err != nil {
		log.Fatal(err)
	}

	date := calendar.Date{}
	if err := date.SetYear(2019); err != nil {
		log.Fatal(err)
	}
	if err := date.SetMonth(10); err != nil {
		log.Fatal(err)
	}
	if err := date.SetDay(29); err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
