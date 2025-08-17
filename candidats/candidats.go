package candidats

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"server-application3/models"
	"text/template"
)

func AgentHandler(res http.ResponseWriter, req *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Не удалось получить рабочую директорию:", err)
	}

	// создаём данные агента
	agent := models.DoubleZero{
		Person: models.Person{
			Name: "James Bond",
			Age:  30,
		},
		LicenseToKill: true,
	}

	// парсим шаблон
	tplPath := filepath.Join(wd, "templates", "doublezero.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Fatalln("Ошибка парсинга шаблона:", err)
	}

	// исполняем шаблон с данными
	err = tpl.Execute(res, agent)
	if err != nil {
		log.Fatalln("Ошибка выполнения шаблона:", err)
	}
}
