package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/camelcase"
	"github.com/gocolly/colly"
)

type Member struct {
	name      string
	birthday  time.Time
	bloodType string
	horoscope string
	height    int
	nickname  string
}

func main() {
	fmt.Println(getMemberName())
}

func getMemberName() []string {
	c := colly.NewCollector()

	memberNames := make([]string, 0)
	c.OnHTML(".entry-member__name > a", func(h *colly.HTMLElement) {
		memberNames = append(memberNames, strings.Join(camelcase.Split(h.Text), " "))
	})

	c.Visit("https://jkt48.com/member/list?lang=id")
	return memberNames
}
