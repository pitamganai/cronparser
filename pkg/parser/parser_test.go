package parser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
    cases := []struct {
        input    string
        expected CronExpression
        wantErr  bool
    }{
        {"*/15 0 1,15 * 1-5 /usr/bin/find", CronExpression{"0 15 30 45", "0", "1 15", "1 2 3 4 5 6 7 8 9 10 11 12", "1 2 3 4 5", "/usr/bin/find"}, false},
        {"0 12 * * 0 /usr/bin/backup", CronExpression{"0", "12", "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31", "1 2 3 4 5 6 7 8 9 10 11 12", "0", "/usr/bin/backup"}, false},
        {"*/10 0-5 1,15 1,6 1-5 /usr/bin/find", CronExpression{"0 10 20 30 40 50", "0 1 2 3 4 5", "1 15", "1 6", "1 2 3 4 5", "/usr/bin/find"}, false},
        {"* * * * * /bin/echo", CronExpression{"0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59", "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23", "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31", "1 2 3 4 5 6 7 8 9 10 11 12", "0 1 2 3 4 5 6", "/bin/echo"}, false},
        {"*/15 0-23 1-31 1-12 0-6 /usr/bin/test", CronExpression{"0 15 30 45", "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23", "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31", "1 2 3 4 5 6 7 8 9 10 11 12", "0 1 2 3 4 5 6", "/usr/bin/test"}, false},
    }

    for _, c := range cases {
        result, err := CronParser{}.Parse(c.input)
        if (err != nil) != c.wantErr {
            t.Errorf("Parse() error = %v, wantErr %v", err, c.wantErr)
            continue
        }
        if !reflect.DeepEqual(result, c.expected) {
            t.Errorf("Parse() = %v, want %v", result, c.expected)
        }
    }
}
