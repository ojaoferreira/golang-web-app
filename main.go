package main

import (
	"database/sql"
	"net/http"
	"os"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

// Post entity
type Post struct {
	ID    int
	Title string
	Body  string
}

func main() {
	appPort := os.Getenv("APP_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbCharSet := os.Getenv("DB_CHARSET")

	if dbHost == "" || dbPort == "" || dbName == "" || dbUser == "" || dbPass == "" {
		panic("Ops! Algo deu errado. Favor configurar as variáveis de ambientes.")
	}

	if dbCharSet == "" {
		dbCharSet = "utf8"
	}
	var stringConexao = dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharSet

	var db, err = sql.Open("mysql", stringConexao)
	if err != nil {
		panic("Ops! Não foi possível abrir a conexão com o banco de dados.")
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("templates/js"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("templates/fonts"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/index.html"))
		t.ExecuteTemplate(w, "index.html", nil)
	})

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		if r.Method != "POST" {
			http.Redirect(w, r, "/", 301)
		}

		post := Post{Title: strings.Join(r.Form["title"], ""), Body: strings.Join(r.Form["body"], "")}

		sqlPost, err := db.Prepare("insert into posts (title, body) values (?, ?)")
		if err != nil {
			panic(err)
		}

		_, err = sqlPost.Exec(post.Title, post.Body)
		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/", 301)
	})

	if appPort == "" {
		appPort = "8080"
	}

	http.ListenAndServe(":"+appPort, nil)

}
