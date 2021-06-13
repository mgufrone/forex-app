package handlers

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mgufrone/forex/internal/domains/rate"
	"net/http"
	"strings"
	"time"
)

type bniWorker struct {
	client *http.Client
}

func NewBniWorker() IWorker {
	client := http.DefaultClient
	return &bniWorker{client: client}
}

func (b *bniWorker) Run(ctx context.Context) ([]*rate.Rate, error) {
	uri := "https://www.bni.co.id/id-id/beranda/informasivalas"
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)

	if err != nil {
		return nil, fmt.Errorf("error %w", err)
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

	rates := body.Find(`div[id*="upValasInfo"] > div[id*="ValasInfo"]`)

	if rates.Length() == 0 {
		return nil, fmt.Errorf(errElemNotFound)
	}

	var entities []*rate.Rate

	types := []string{"erate", "teller", "bank_notes"}

	rates.Each(func(i int, selection *goquery.Selection) {
		dateString := selection.Find(".date-infoView span").Text()
		dateString = strings.Replace(dateString, "Pembaharuan Terakhir:", "", -1)
		dateString = strings.Replace(dateString, " WIB (GMT+07:00).", "", -1)
		updatedAtString := strings.TrimSpace(dateString)
		updatedAt, _ := time.Parse("02/01/2006 15:04:05", updatedAtString)

		selection.Find(".content-infoView table.table tbody tr").Each(func(i2 int, selection *goquery.Selection) {
			tds := selection.Find("td")
			code := tds.Eq(0).Text()
			buyString := strings.TrimSpace(tds.Eq(1).Text())
			sellString := strings.TrimSpace(tds.Eq(2).Text())
			buy, _ := toFloat(buyString)
			sell, _ := toFloat(sellString)
			curRate := rate.NewRate("IDR", code, "bca", types[i], sell, buy, updatedAt)
			entities = append(entities, curRate)
		})
	})

	return entities, nil
}
