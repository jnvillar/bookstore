package scrapper

import (
	"testing"
)

func TestScrapAll(t *testing.T) {
	scrapBooks()
}

func TestScrap(t *testing.T) {
	scrapBotanica(nil, "distribuidoralabotica.json")
}

func TestMeli(t *testing.T) {
	scrapMeli(nil, "meli.json")
}
