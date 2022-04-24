package thirdapp

import (
	"2miner-monitoring/config"
	"2miner-monitoring/data"
	"2miner-monitoring/es"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func DualType(e *colly.HTMLElement) []byte {
	var dual = regexp.MustCompile(`([a-zA-Z]{0,}?\+[a-zA-Z]{0,})([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,}\/s)([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,}\/s)([0-9]{0,})w([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,})([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,}\/w)([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,}\/w)\$([0-9]{0,}?\.[0-9]{0,})\$([0-9]{0,}?\.[0-9]{0,}) ([0-9]{0,}) ([a-zA-Z]{0,})([0-9]{0,}) ([a-zA-Z]{0,})`)
	if dual.MatchString(e.Text) {
		DualCrawled := data.Card{}
		coin := dual.ReplaceAllString(e.Text, `$1`)
		coinsName := strings.Split(coin, "+")
		DualCrawled.CoinPrimary = coinsName[0]
		DualCrawled.CoinAlt = coinsName[1]
		DualCrawled.CoinName = coin
		DualCrawled.Conso, _ = strconv.Atoi(dual.ReplaceAllString(e.Text, `$6`))
		DualCrawled.Income, _ = strconv.ParseFloat(dual.ReplaceAllString(e.Text, `$13`), 64)
		DualCrawled.HashrateAlt, _ = strconv.ParseFloat(dual.ReplaceAllString(e.Text, `$4`), 64)
		DualCrawled.HashratePrimary, _ = strconv.ParseFloat(dual.ReplaceAllString(e.Text, `$2`), 64)
		DualCrawled.UnitPrimary = dual.ReplaceAllString(e.Text, `$3`)
		DualCrawled.UnitAlt = dual.ReplaceAllString(e.Text, `$5`)
		DualCrawled.Card = e.Request.URL.Path[1:]
		DualCrawled.Timestamp = time.Now().Format(time.RFC3339)
		JsonDualCrawled, _ := json.Marshal(DualCrawled)
		return JsonDualCrawled
	}
	return nil
}

func Singletype(e *colly.HTMLElement) []byte {
	var single = regexp.MustCompile(`([a-zA-Z]{0,})([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,}\/s)([0-9]{0,})w([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,})([0-9]{0,}?\.[0-9]{0,}) ([a-zA-Z]{0,}\/w)\$([0-9]{0,}?\.[0-9]{0,})\$([0-9]{0,}?\.[0-9]{0,}) ([0-9]{0,}) ([a-zA-Z]{0,})([0-9]{0,}) ([a-zA-Z]{0,})`)
	if single.MatchString(e.Text) {
		SingleCrawled := data.Card{}
		coin := single.ReplaceAllString(e.Text, `$1`)
		SingleCrawled.CoinPrimary = coin
		SingleCrawled.CoinName = coin
		SingleCrawled.HashratePrimary, _ = strconv.ParseFloat(single.ReplaceAllString(e.Text, `$2`), 64)
		SingleCrawled.UnitPrimary = single.ReplaceAllString(e.Text, `$3`)
		SingleCrawled.Conso, _ = strconv.Atoi(single.ReplaceAllString(e.Text, `$4`))
		SingleCrawled.Income, _ = strconv.ParseFloat(single.ReplaceAllString(e.Text, `$9`), 64)
		SingleCrawled.Card = e.Request.URL.Path[1:]
		SingleCrawled.Timestamp = time.Now().Format(time.RFC3339)
		JsonSingleCrawled, _ := json.Marshal(SingleCrawled)
		return JsonSingleCrawled
	}
	return nil
}

func DispatchType(e *colly.HTMLElement) []byte {
	res := DualType(e)
	if res == nil {
		res = Singletype(e)
	}
	return res
}

func RunCrawler() (int, string) {
	var CardsInfo [][]byte
	c := colly.NewCollector(
		colly.AllowedDomains("www.hashrate.no", "hashrate.no"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Houston nous avons un problème : ", err)
	})
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		cardInfo := DispatchType(e)
		if cardInfo != nil {
			CardsInfo = append(CardsInfo, cardInfo)
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Printf("harvesting : %s", r.URL.String())
	})
	c.OnScraped(func(s *colly.Response) {
		log.Printf("harvested : %s", s.Request.URL.String())
	})
	for _, elem := range config.Cards.CardsList {
		url := fmt.Sprintf("https://www.hashrate.no/%s", elem)
		c.Visit(url)
	}
	es.BulkData("2miners-hashrate_no", CardsInfo)
	return 200, "Card Stats harvested"
}
