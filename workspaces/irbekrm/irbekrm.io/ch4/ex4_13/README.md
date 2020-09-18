# Exercise 4.11

## Description

Build a tool that retrieves the poster for a selected movie from Open Movie Database

## Usage

* `go build -o omdbgo` // compile

* Retrieve **Patreon** API key from [omdapi](http://www.omdbapi.com/apikey.aspx) and export it as `OMD_API_KEY`

* Run `./omdbgo MOVIE` // will create `poster.png` file in the current directory

#### Example
```
go build -o omdbgo
EXPORT OMD_API_KEY=[YOUR_KEY]
./omdbgo "Ali: Fear eats the soul"
```





