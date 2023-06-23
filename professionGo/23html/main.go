package main

import (
	"html/template"
	"os"
)

// Поиск определенного шаблона
func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, &Kayak)
}
func main() {
	allTemplates, err := template.ParseGlob("templates/*.html")
	if err == nil {
		selectedTemplated := allTemplates.Lookup("template.html")
		err = Exec(selectedTemplated)
	}
	if err != nil {
		Printfln("Error: %v %v", err.Error())
	}
}

/*
//Перечисление загруженных шаблонов
func main() {
	allTemplates, err := template.ParseGlob("templates/*.html")
	if err == nil {
		for _, t := range allTemplates.Templates() {
			Printfln("Template name: %v", t.Name())
		}
	} else {
		Printfln("Error: %v %v", err.Error())
	}
}

/*
//Загрузка нескольких шаблонов
func main() {
	allTemplates, err1 := template.ParseFiles("templates/template.html",
		"templates/extras.html")
	if err1 == nil {
		allTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
		os.Stdout.WriteString("\n")
		allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
	} else {
		Printfln("Error: %v %v", err1.Error())
	}
}

/*
	t1, err1 := template.ParseFiles("templates/template.html")
	t2, err2 := template.ParseFiles("templates/extras.html")
	if err1 == nil && err2 == nil {
		t1.Execute(os.Stdout, &Kayak)
		os.Stdout.WriteString("\n")
		t2.Execute(os.Stdout, &Kayak)
	} else {
		Printfln("Error: %v %v", err1.Error(), err2.Error())
	}
}*/
