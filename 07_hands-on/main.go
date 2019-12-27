package main

import (
	"fmt"
	"os"
	"text/template"
)

var (
	tmp *template.Template
)

func init() {
	tmp = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Restaurant struct {
	Name                     string
	Breakfast, Lunch, Dinner []Food
}

type Food struct {
	Name  string
	Price int
}

func main() {
	restaurants := []Restaurant{
		{
			Name: "Restaurant Number 1",
			Breakfast: []Food{
				{
					Name:  "Bread",
					Price: 10,
				},
				{
					Name:  "Egg",
					Price: 13,
				},
				{
					Name:  "Burger",
					Price: 14,
				},
			},
			Lunch: []Food{
				{
					Name:  "Burger",
					Price: 16,
				},
				{
					Name:  "Salad",
					Price: 17,
				},
				{
					Name:  "Chicken Nugget",
					Price: 16,
				},
			},
			Dinner: []Food{
				{
					Name:  "Steak",
					Price: 19,
				},
				{
					Name:  "Fried Fish",
					Price: 17,
				},
				{
					Name:  "Baccon",
					Price: 13,
				},
			},
		},
	}

	err := tmp.Execute(os.Stdout, restaurants)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
}
