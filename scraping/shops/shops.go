package shops

import (
	"os"
	"shops-scraping/shared"
)

type Shop struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Enabled bool   `json:"enabled"`
}

func newShop(name, code, ffVar string) Shop {
	return Shop{
		name,
		code,
		os.Getenv(ffVar) == "true",
	}
}

func GetShops() []Shop {
	return []Shop{
		newShop("H&M", shared.HM, "ENABLE_HM"),
		newShop("PULL AND BEAR", shared.PULLANDBEAR, "ENABLE_PULLNBEAR"),
		newShop("BERSHKA", shared.BERSHKA, "ENABLE_BERSHKA"),
		newShop("ZARA", shared.ZARA, "ENABLE_ZARA"),
		newShop("SHEIN", shared.SHEIN, "ENABLE_SHEIN"),
	}
}

func GetEnabledShops() []Shop {
	return shared.SlicesFilter(GetShops(), func(shop Shop) bool {
		return shop.Enabled
	})
}
