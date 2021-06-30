package handlers

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/mgufrone/forex/internal/domains/rate"
	"strings"
	"time"
)

type mandiriWorker struct {
}

func NewMandiriWorker() IWorker {
	return &mandiriWorker{}
}

func (m *mandiriWorker) Run(ctx context.Context) ([]*rate.Rate, error) {
	uri := "https://bankmandiri.co.id/web/guest/kurs"
	ctx, cancelExec := chromedp.NewExecAllocator(context.Background(), append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36"),
		)...)
	defer cancelExec()
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()
	var html string
	err := chromedp.Run(ctx,
		chromedp.Navigate(uri),
		chromedp.WaitReady("table.table-kurs"),
		chromedp.InnerHTML("body", &html),
	)

	if err != nil {
		return nil, err
	}

	body, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		return nil, err
	}

	table := body.Find("table.table-kurs")
	rates := table.Find("tbody tr")

	if rates.Length() == 0 {
		return nil, fmt.Errorf(errElemNotFound)
	}

	entities := make([]*rate.Rate, rates.Length()*3)
	types := []string{"erate", "teller", "bank_notes"}
	idx := 0
	dateFormat := "02/01/06 - 15:04"
	dates := make([]time.Time, len(types))
	replacers := []string{"Special Rate*", "TT Counter", "Bank Notes"}

	table.Find("thead tr").Eq(0).Find("th").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}
		realIdx := i - 1
		dateString := strings.TrimSpace(strings.Replace(
			strings.Replace(selection.Text(), "WIB", "", -1),
			replacers[realIdx], "", -1),
		)
		dates[realIdx], err = time.Parse(dateFormat, strings.TrimSpace(dateString))
	})


	rates.Each(func(i int, selection *goquery.Selection) {
		code := selection.Find("td").Eq(0).Text()
		vals := selection.Find("td")
		for i2, v := range types {
			sellIdx := (i2 + 1) * 2
			buyIdx := sellIdx - 1
			buyString := strings.TrimSpace(vals.Eq(buyIdx).Text())
			sellString := strings.TrimSpace(vals.Eq(sellIdx).Text())
			buy, _ := toFloat(buyString)
			sell, _ := toFloat(sellString)
			entities[idx], _ = rate.NewRate("IDR", code, "mandiri", v, sell, buy, dates[i2])
			idx++
		}

	})

	return entities, nil
}
