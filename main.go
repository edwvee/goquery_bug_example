package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	f, err := os.Open("openwrt-ar71xx-mikrotik-rb-nor-flash-16M-initramfs-kernel.bin")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		panic(err)
	}

	doc.Find("h3").Each(func(i int, sel *goquery.Selection) {
		fmt.Printf("found h3 tag with length in %d bytes\n", len(sel.Text()))
	})
}
