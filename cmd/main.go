package main

import (
	"github.com/divilla/events.git"
	"log"
)

type (
	One struct {
		Id   int
		Name string
	}
)

func main() {
	e := events.NewEventsManager(log.Default())

	// prints:
	// 1 One
	// aaa
	e.Subscribe("one-create", func(target interface{}, data events.Map) error {
		t := target.(*One)
		println(t.Id, t.Name)
		println(data["some-data"].(string))
		data["other-data"] = "bbb"

		return nil
	})

	// prints:
	// 1 One
	// bbb
	e.Subscribe("one-create", func(target interface{}, data events.Map) error {
		t := target.(*One)
		println(t.Id, t.Name)
		println(data["other-data"].(string))

		return nil
	})

	one := &One{
		Id:   1,
		Name: "One",
	}

	if err := e.Dispatch("one-create", one, events.Map{"some-data": "aaa"}); err != nil {
		log.Fatal(err)
	}
}
