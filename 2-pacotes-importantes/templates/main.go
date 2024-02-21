package main

import (
	"html/template"
	"net/http"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int64
}

type Cursos []Curso

func main() {
	curso := Curso{"TI", 40}
	temp := template.New("Curso Template")
	temp, _ = temp.Parse("Nome: {{.Nome}} - Carga Horaria: {{.CargaHoraria}}")
	err := temp.Execute(os.Stdout,curso)
	if err!=nil{
		panic(err)
	}

	// Template.Must une a criacao e o parse do template
	t := template.Must(template.New("Curso Template").Parse("Nome: {{.Nome}} - Carga Horaria: {{.CargaHoraria}}"))
	err = t.Execute(os.Stdout,curso)
	if err!=nil{
		panic(err)
	}

	tHtml := template.Must(template.New("template.html").ParseFiles("template.html"))
	data := Cursos{
		{"go", 40},
		{"java", 20},
		{"node", 10},
	}
	err = tHtml.Execute(os.Stdout,data)
	if err!=nil{
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		tHtml := template.Must(template.New("template.html").ParseFiles("template.html"))
		data := Cursos{
			{"go", 40},
			{"java", 20},
			{"node", 10},
		}
		err = tHtml.Execute(w,data)
	})
	http.ListenAndServe(":8080",nil)

}