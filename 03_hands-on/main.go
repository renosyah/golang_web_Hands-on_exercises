package main

import (
	"fmt"
	"os"
	"text/template"
)

var (
	tmp *template.Template
)

type Hotel struct {
	Name, Address, City, Region string
	Zip                         int
}

func init() {
	tmp = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []Hotel{
		{
			Name:    "Hotel 1",
			Address: "jl street 1 no 104",
			City:    "los angles",
			Region:  "Southern",
			Zip:     12345,
		},
		{
			Name:    "Hotel 2",
			Address: "jl street 3 no 104",
			City:    "los angles",
			Region:  "Southern",
			Zip:     12346,
		},
		{
			Name:    "Hotel 3",
			Address: "jl street 3 no 104",
			City:    "los angles",
			Region:  "Southern",
			Zip:     12347,
		},
	}

	err := tmp.Execute(os.Stdout, hotels)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
}
