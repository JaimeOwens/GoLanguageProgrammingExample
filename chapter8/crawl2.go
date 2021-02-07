package chapter8

import (
	"fmt"
	"gopl/chapter5"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := chapter5.Extract(url)
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return list
}

func Crawl2() {
	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl2(link)
				}(link)
			}
		}
	}
}
