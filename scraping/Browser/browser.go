package Browser

import (
	"github.com/go-rod/rod"
	"log"
)

var browser *rod.Browser

func GetInstance() *rod.Browser {
	if browser == nil {
		CreateInstance()
	}
	return browser
}

func CreateInstance() {
	browser = rod.New().MustConnect()
	log.Println("browser instance created")
}
