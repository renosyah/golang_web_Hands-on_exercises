package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

var (
	tmp *template.Template
)

type DataCSV struct {
	Date                                     string
	Open, High, Low, Close, Volume, AdjClose float64
}

func init() {
	tmp = template.Must(template.ParseFiles("tpl.gohtml"))
}

func showData(res http.ResponseWriter, req *http.Request) {

	var data []DataCSV

	file, err := os.Open("table.csv")
	if err != nil {
		fmt.Println("Error open csv", err)
		os.Exit(1)
	}

	defer file.Close()

	csvr := csv.NewReader(file)
	if _, err := csvr.Read(); err != nil { //read header
		fmt.Println("Error read csv at header", err)
		os.Exit(1)
	}

	for {
		row, err := csvr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error read csv at some line", err)
			os.Exit(1)
		}

		one := DataCSV{}
		one.Date = row[0]
		one.Open, err = strconv.ParseFloat(row[1], 64)
		if err != nil {
			fmt.Println("Error parse float open", err)
		}
		one.High, err = strconv.ParseFloat(row[2], 64)
		if err != nil {
			fmt.Println("Error parse float High", err)
		}
		one.Low, err = strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Println("Error parse floatLow", err)
		}
		one.Close, err = strconv.ParseFloat(row[4], 64)
		if err != nil {
			fmt.Println("Error parse float Close", err)
		}
		one.Volume, err = strconv.ParseFloat(row[5], 64)
		if err != nil {
			fmt.Println("Error parse float Volume", err)
		}
		one.AdjClose, err = strconv.ParseFloat(row[6], 64)
		if err != nil {
			fmt.Println("Error parse float AdjClose", err)
		}

		data = append(data, one)

	}

	err = tmp.Execute(res, data)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", showData)
	http.ListenAndServe(":8000", nil)
}
