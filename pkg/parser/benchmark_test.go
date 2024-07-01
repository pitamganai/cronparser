package parser

import (
	"testing"
)

func BenchmarkParse(b *testing.B) {
    cronExpr := "*/15 0 1,15 * 1-5 /usr/bin/find"

    for i := 0; i < b.N; i++ {
        _, err := CronParser{}.Parse(cronExpr)
        if err != nil {
            b.Fatalf("unexpected error: %v", err)
        }
    }
}