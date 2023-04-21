package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
	"time"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func main() {
	r, err := whois.Whois(os.Args[1])
	if err == nil {
		p, err := whoisparser.Parse(r)
		if err == nil {
			expStr := s.Split((s.Split(p.Domain.ExpirationDate, "T")[0]), "-")
			todayStr := s.Split((s.Split((time.Now()).Format(time.RFC3339), "T")[0]), "-")

			var exp [3]int
			var today [3]int
			for i := 0; i < len(exp); i++ {
				exp[i], _ = strconv.Atoi(expStr[i])
				today[i], _ = strconv.Atoi(todayStr[i])
			}

			t1 := time.Date(exp[0], time.Month(exp[1]), exp[2], 0, 0, 0, 0, time.Local)
			t2 := time.Date(today[0], time.Month(today[1]), today[2], 0, 0, 0, 0, time.Local)
			days := t1.Sub(t2).Hours() / 24
			fmt.Print(days)
		}
	}
}
