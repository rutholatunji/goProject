package main

import "fmt"

func comicBook() {
	var publisher, writer, artist, title string
	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."

	var year, pageNumber int
	year = 1997
	pageNumber = 14

	var grade float32
	grade = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "Year:", year, "Page numbers:", pageNumber, "Grade:", grade)

	title = "Epic Vol.1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, ",Year:", year, ",Page numbers:", pageNumber, ",Grade:", grade)

}
