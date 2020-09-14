# Exercise 4.11

## Description

Build a tool that creates an offline index for xkcd comics that can be searched for specific words.


## Usage

`go build -o xkcdgo` // compile

### Create offline index

`./xkcdgo populate` // will write json formatted index to file xkcdgo in the current directory

### Search for a word

* Create the offline index with `./xkcdgo populate`
* Run `./xkcdgo search WORD`








