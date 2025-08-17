package parsing

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"server-application3/models"
	"strconv"
	"text/template"
	"time"
)

func ParsingHandler(res http.ResponseWriter, req *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Не удалось получить рабочую директорию:", err)
	}

	// читаем CSV
	records := Prs(filepath.Join(wd, "templates", "table.csv"))
	// **выводим в лог для проверки**
	//	log.Printf("Records: %+v", records)
	tpl, err := template.ParseFiles(filepath.Join(wd, "templates", "hw.gohtml"))
	if err != nil {
		log.Fatalln("Ошибка парсинга шаблона:", err)
	}

	// исполняем шаблон с данными
	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}

// функция парсинга CSV

func Prs(filePath string) []models.Record {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Ошибка открытия CSV:", err)
	}
	defer src.Close()

	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("Ошибка чтения CSV:", err)
	}

	records := make([]models.Record, 0, len(rows))

	for i, row := range rows {
		if i == 0 { // пропускаем заголовок
			continue
		}

		date, err := time.Parse("2006-01-02", row[0])
		if err != nil {
			log.Printf("строка %d: ошибка парсинга даты '%s': %v", i, row[0], err)
			continue
		}

		open, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			log.Printf("строка %d: ошибка парсинга Open '%s': %v", i, row[1], err)
			continue
		}

		high, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			log.Printf("строка %d: ошибка парсинга High '%s': %v", i, row[2], err)
			continue
		}

		low, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			log.Printf("строка %d: ошибка парсинга Low '%s': %v", i, row[3], err)
			continue
		}

		closePrice, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			log.Printf("строка %d: ошибка парсинга Close '%s': %v", i, row[4], err)
			continue
		}

		volume, err := strconv.ParseInt(row[5], 10, 64)
		if err != nil {
			log.Printf("строка %d: ошибка парсинга Volume '%s': %v", i, row[5], err)
			continue
		}

		adjClose, err := strconv.ParseFloat(row[6], 64)
		if err != nil {
			log.Printf("строка %d: ошибка парсинга Adj Close '%s': %v", i, row[6], err)
			continue
		}

		records = append(records, models.Record{
			Date:     date,
			Open:     open,
			High:     high,
			Low:      low,
			Close:    closePrice,
			Volume:   volume,
			AdjClose: adjClose,
		})
	}

	return records
}
