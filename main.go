package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Lütfen bir URL girin.")
		return
	}
	siteAdresi := os.Args[1]

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var htmlIcerik string
	var ekranGoruntusu []byte
	var linkler []string

	err := chromedp.Run(ctx,
		chromedp.Navigate(siteAdresi),
		chromedp.OuterHTML("html", &htmlIcerik),
		chromedp.FullScreenshot(&ekranGoruntusu, 90),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('a')).map(a => a.href)`, &linkler),
	)

	if err != nil {
		log.Fatal("Hata oluştu:", err)
	}

	ioutil.WriteFile("output.txt", []byte(htmlIcerik), 0644)

	ioutil.WriteFile("screenshot.png", ekranGoruntusu, 0644)

	f, _ := os.Create("urls.txt")
	for _, link := range linkler {
		f.WriteString(link + "\n")
	}
	f.Close()

	fmt.Println("İşlem tamam. Dosyalar kaydedildi.")
}
