// Issues prints a table of GitHub issues matching the search terms and reports
// the results in age categories of this month, this year, and older than a year.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"andr.io/ch4/ex4_10/github"
)

func monthsSince(createdAtTime time.Time) int {
	now := time.Now()
	months := 0
	month := createdAtTime.Month()
	for createdAtTime.Before(now) {
		createdAtTime = createdAtTime.Add(time.Hour * 24)
		nextMonth := createdAtTime.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	return months
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// Uses space to store the strings in categories. Simpler than sorting the
	// results.
	var month []string
	var year []string
	var older []string

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		timeSince := monthsSince(item.CreatedAt)
		s := fmt.Sprintf(
			"#%-5d %9.9s %d %.55s",
			item.Number,
			item.User.Login,
			timeSince,
			item.Title)
		if timeSince <= 1 {
			month = append(month, s)
		} else if timeSince > 1 && timeSince <= 12 {
			year = append(year, s)
		} else if timeSince > 12 {
			older = append(older, s)
		}
	}

	fmt.Println("This month:")
	for _, i := range month {
		fmt.Println(i)
	}

	fmt.Println("This year:")
	for _, i := range year {
		fmt.Println(i)
	}

	fmt.Println("Older:")
	for _, i := range older {
		fmt.Println(i)
	}
}

// Search for 'bbc':

// 15621 issues:
// This month:
// #114     uckthis 1 Is BBC Working?
// #8       thlllll 1 United Kingdom - BBC 6Music
// #594      nektro 1 www.bbc.com
// #22948 NellieThe 1 [💥] http://bbc.co.uk
// #23695 NellieThe 1 [💥] https://www.bbc.co.uk
// #97    radiocont 0 Repo & homepage URLs for psammead-headings are incorrec
// #1609    guevara 1 Twilight of the Taj - BBC News
// #31      pietrop 1 BBC Kaldi Converter [to export content back into Kaldi
// #38        sareh 1 Add InlineLink Component to Psammead
// #1     michael-s 1 This is an extraction from github.com/bbc/coals
// #54    ChrisBAsh 1 Put Copyright component into Psammead
// #57    ChrisBAsh 1 Put Caption component into Psammead
// #214    jefro108 0 BBC Planet Earth programmes only downloading in SD
// #841   tomwinnin 1 [PROGRAMMES-6728] streamable versions 500 error
// #281   sbeaumont 1 <abbr> tag uses "title" and currently fails the title a
// #128   liamwalsh 1  added hotstepper & BBC Civilisations to apps
// #11      jaustin 1 Feature request: Nordic UART/BBC micro:bit support
// #851   tomwinnin 0 [PROGRAMMES-6752] fixed dates for podcasts
// This year:
// #7     wpoernomo 2 test BBC
// #1     bobsterne 2 BBC player not working.
// #63    bobsterne 2 BBC i-player unusable
// #209     MGoeppl 2 BBC iPlayer Radio moved to BBC Sounds, unable to downlo
// #10          dr3 3 Add support for BBC Grandstand
// #1     shani-ace 2 added BBC sports
// #68    WPetterss 2 Add support for BBC iPlayer
// #275    mikeposh 2 Error parsing `.some_class_name:disabled::before`
// #52    jasoncart 2 BBC scraper no longer works
// #643   jwillough 2 category_urls() gets ".com" and ".co.uk", as well as "h
// #426      b16r05 3 [Package Request] BBC iplayer
// Older:
// #8     marieALap 16 bbc wildlife
