package patient

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Patient struct {
	Name    string
	Surname string
	Age     int
	Gender  string
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/globomantics")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, surname, age, gender FROM patients WHERE surname LIKE '%" + query + "%'")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []Patient
	for rows.Next() {
		var p Patient
		err := rows.Scan(&p.Name, &p.Surname, &p.Age, &p.Gender)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		results = append(results, p)
	}

	tmpl := template.Must(template.ParseFiles("templates/results.html"))
	err = tmpl.Execute(w, results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
