package common

import (
	"github.com/go-rod/rod"
	log "github.com/sirupsen/logrus"
	"math"
	"regexp"
	"strconv"
)

func GetPrice(txt string) float32 {
	re := regexp.MustCompile("\\d+\\.\\d+")
	results := re.FindStringSubmatch(txt)

	if len(results) == 0 {
		return 0
	}
	price, err := strconv.ParseFloat(results[0], 32)

	if err != nil {
		log.Errorf("Error during price extract: %v", err)
		return 0
	}

	return float32(price)
}

func CloseCookieDialog(page *rod.Page) {
	cookiesContainer := "#onetrust-consent-sdk"
	page.MustElement(cookiesContainer).MustRemove()
}

func WaitForLoad(page *rod.Page, resultsSelector, emptySelector string) (hasResults bool) {
	var ch = make(chan bool, 1)

	go func() {
		err := page.WaitElementsMoreThan(resultsSelector, 0)
		if err == nil {
			ch <- true
		}
	}()
	go func() {
		err := page.WaitElementsMoreThan(emptySelector, 0)
		if err == nil {
			ch <- false
		}
	}()

	hasResults = <-ch

	return
}

func Group[T comparable](collection []T, grpSize int) (grp [][]T) {
	turns := math.Ceil(float64(float32(len(collection)) / float32(grpSize)))
	grpNumber, _ := strconv.Atoi(strconv.FormatFloat(turns, 'f', 0, 32))

	for i := 0; i < grpNumber; i++ {
		if grpSize*(i+1) > cap(collection) {
			grp = append(grp, collection[i*grpSize:])
			continue
		}
		grp = append(grp, collection[i*grpSize:grpSize*(i+1)])
	}
	return
}
