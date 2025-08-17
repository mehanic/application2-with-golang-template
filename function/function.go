package function

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"server-application3/models"
	"strconv"
	"strings"
	"time"
)

func FirstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func MonthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func Double(x int) int {
	return x + x
}

func Square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func SqRoot(x float64) float64 {
	return math.Sqrt(x)
}

func Prs(filePath string) []models.Record {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]models.Record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue // пропускаем заголовок
		}
		date, err := time.Parse("2006-01-02", row[0])
		if err != nil {
			log.Printf("ошибка парсинга даты: %v", err)
			continue
		}
		open, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			log.Printf("ошибка парсинга числа: %v", err)
			continue
		}

		records = append(records, models.Record{
			Date: date,
			Open: open,
		})
	}

	return records
}
