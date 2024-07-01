package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type CronExpression struct {
    Minute     string
    Hour       string
    DayOfMonth string
    Month      string
    DayOfWeek  string
    Command    string
}

type Parser interface {
    Parse(cron string) (CronExpression, error)
}

type CronParser struct{}

func (cp CronParser) Parse(cron string) (CronExpression, error) {
    fields := strings.Fields(cron)
    if len(fields) != 6 {
        return CronExpression{}, fmt.Errorf("invalid cron expression")
    }

    return CronExpression{
        Minute:     expandField(fields[0], 0, 59),
        Hour:       expandField(fields[1], 0, 23),
        DayOfMonth: expandField(fields[2], 1, 31),
        Month:      expandField(fields[3], 1, 12),
        DayOfWeek:  expandField(fields[4], 0, 6),
        Command:    fields[5],
    }, nil
}

func expandField(field string, min, max int) string {
    if field == "*" {
        return generateRange(min, max, 1)
    }

    var result []string
    parts := strings.Split(field, ",")
    for _, part := range parts {
        if strings.Contains(part, "/") {
            result = append(result, expandStepField(part, min, max)...)
        } else if strings.Contains(part, "-") {
            result = append(result, expandRangeField(part)...)
        } else {
            result = append(result, part)
        }
    }
    return strings.Join(result, " ")
}

func generateRange(min, max, step int) string {
    var result []string
    for i := min; i <= max; i += step {
        result = append(result, strconv.Itoa(i))
    }
    return strings.Join(result, " ")
}

func expandStepField(field string, min, max int) []string {
    subparts := strings.Split(field, "/")
    rangePart := subparts[0]
    step, _ := strconv.Atoi(subparts[1])
    if rangePart == "*" {
        return strings.Split(generateRange(min, max, step), " ")
    }

    return expandRangeFieldWithStep(rangePart, step)
}

func expandRangeFieldWithStep(field string, step int) []string {
    // var result []string
    var rangeMin, rangeMax int
    if strings.Contains(field, "-") {
        rangeBounds := strings.Split(field, "-")
        rangeMin, _ = strconv.Atoi(rangeBounds[0])
        rangeMax, _ = strconv.Atoi(rangeBounds[1])
    } else {
        singleValue, _ := strconv.Atoi(field)
        rangeMin = singleValue
        rangeMax = singleValue
    }
    return strings.Split(generateRange(rangeMin, rangeMax, step), " ")
}

func expandRangeField(field string) []string {
    rangeBounds := strings.Split(field, "-")
    rangeMin, _ := strconv.Atoi(rangeBounds[0])
    rangeMax, _ := strconv.Atoi(rangeBounds[1])
    return strings.Split(generateRange(rangeMin, rangeMax, 1), " ")
}