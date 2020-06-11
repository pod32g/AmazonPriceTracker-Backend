package scrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

func ExtractPrice(URL string) (float32, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return 0.0, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36 Edg/83.0.478.45")

	resp, err := client.Do(req)

	if err != nil {
		return 0.0, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return 0.0, err
	}

	text := doc.Find("span#priceblock_ourprice").First().Text()

	fmt.Println(text)

	text = strings.Replace(text, "$", "", 1)
	text = strings.Replace(text, ",", "", 1)

	price, err := strconv.ParseFloat(text, 32)

	if err != nil {
		return 0.0, err
	}

	return float32(price), nil
}

func ExtractTitle(URL string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36 Edg/83.0.478.45")

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return "", err
	}

	title := doc.Find("span#productTitle").First().Text()

	title = strings.ReplaceAll(title, "\n", "")

	fmt.Println(title)

	return title, nil
}
