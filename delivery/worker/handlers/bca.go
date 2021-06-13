package handlers

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mgufrone/forex/internal/domains/rate"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	errElemNotFound = "element not found"
)

type bcaWorker struct {
	client *http.Client
}

func toFloat(val string) (v float64, err error) {
	val = strings.Replace(val, ".", "", -1)
	val = strings.Replace(val, ",", ".", -1)
	return strconv.ParseFloat(val, 64)
}
func monthReplacer(val string) string {
	dict := []string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"July",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	for i, v := range dict {
		m := time.Date(2006, time.Month(i + 1), 1, 0, 0, 0, 0, time.UTC)
		val = strings.Replace(val, v, m.Month().String(), -1)
	}
	return val
}
func NewBcaWorker() IWorker {
	client := http.DefaultClient
	return &bcaWorker{ client}
}


func (b *bcaWorker) Run(ctx context.Context) ([]*rate.Rate, error) {
	uri := "https://www.bca.co.id/kurs"
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)

	if err != nil {
		return nil, err
	}

	res, err := b.client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := goquery.NewDocumentFromReader(res.Body)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	rates := body.Find("table.m-table-kurs tbody tr")

	if rates.Length() == 0 {
		return nil, fmt.Errorf(errElemNotFound)
	}

	entities := make([]*rate.Rate, rates.Length() * 3)
	types := []string{"erate", "teller", "bank_notes"}
	idx := 0
	updatedAtString := strings.TrimSpace(body.Find(".o-kurs-refresh-description span.refresh-date").Text())
	updatedAt, err := time.Parse("02 January 2006 15.04", monthReplacer(updatedAtString))

	if err != nil {
		return nil, err
	}

	rates.Each(func(i int, selection *goquery.Selection) {
		code, ok := selection.Attr("code")
		if !ok {
			return
		}
		vals := selection.Find("td:not([class])")
		for i2, v := range types {
			sellIdx := (i2 * 2) + 1
			buyIdx := sellIdx - 1
			buyString := strings.TrimSpace(vals.Eq(buyIdx).Text())
			sellString := strings.TrimSpace(vals.Eq(sellIdx).Text())
			buy, _ := toFloat(buyString)
			sell, _ := toFloat(sellString)
			entities[idx] = rate.NewRate("IDR", code, "bca", v, sell, buy, updatedAt)
			idx ++
		}

	})

	return entities, nil
}

