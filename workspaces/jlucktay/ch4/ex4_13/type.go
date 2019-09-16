package main

type SearchResult struct {
	Movies       []Movie `json:"Search"`
	TotalResults int     `json:"totalResults,string"`
	Response     string  `json:"Response"`
}

type Movie struct {
	Title  string `json:"Title"`
	Year   int    `json:"Year,string"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
