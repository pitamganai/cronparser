package main

import (
	"fmt"
	"os"

	"github.com/pitamganai/cronparser/pkg/parser"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: cron-parser <cron-expression>")
        return
    }

    cron := os.Args[1]
    var p parser.Parser = parser.CronParser{}
    parsed, err := p.Parse(cron)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("minute        %s\n", parsed.Minute)
    fmt.Printf("hour          %s\n", parsed.Hour)
    fmt.Printf("day of month  %s\n", parsed.DayOfMonth)
    fmt.Printf("month         %s\n", parsed.Month)
    fmt.Printf("day of week   %s\n", parsed.DayOfWeek)
    fmt.Printf("command       %s\n", parsed.Command)
}
