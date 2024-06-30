package common

import (
	log "github.com/sirupsen/logrus"
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
