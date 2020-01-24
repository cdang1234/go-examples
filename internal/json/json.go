package json

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Craft string `json:"craft"`
	Name  string `json:"name"`
}

type people struct {
	Number int `json:"number"`
	People []person
}

func DecodeIntoJSON() {
	text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`
	textBytes := []byte(text)

	people1 := people{}
	err := json.Unmarshal(textBytes, &people1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people1)
}

// type movie struct {
// 	Title string
// }

// type Movies struct {
// 	Data        []movie
// 	Total_pages int `json:"total_pages"`
// }

// func DecodeIntoJSONMovies() {
// 	text := `{"page":1,"per_page":10,"total":15,"total_pages":2,"data":[{"Title":"Italian Spiderman","Year":2007,"imdbID":"tt2705436"},{"Title":"Superman, Spiderman or Batman","Year":2011,"imdbID":"tt2084949"},{"Title":"Spiderman","Year":1990,"imdbID":"tt0100669"},{"Title":"Spiderman","Year":2010,"imdbID":"tt1785572"},{"Title":"Fighting, Flying and Driving: The Stunts of Spiderman 3","Year":2007,"imdbID":"tt1132238"},{"Title":"Spiderman and Grandma","Year":2009,"imdbID":"tt1433184"},{"Title":"The Amazing Spiderman T4 Premiere Special","Year":2012,"imdbID":"tt2233044"},{"Title":"Amazing Spiderman Syndrome","Year":2012,"imdbID":"tt2586634"},{"Title":"Hollywood's Master Storytellers: Spiderman Live","Year":2006,"imdbID":"tt2158533"},{"Title":"Spiderman 5","Year":2008,"imdbID":"tt3696826"}]}`
// 	textBytes := []byte(text)

// 	movies := Movies{}
// 	err := json.Unmarshal(textBytes, &movies)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(movies)
// }
