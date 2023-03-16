package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Member struct {
	Name string
}

type Groupe struct {
	ID           int
	Name         string
	Image        string
	Description  string
	CreationDate string
	FirstAlbum   string
}

type Concert struct {
	ID       int
	Artist   string
	Date     string
	Country  string
	Location string
}

type Result struct {
	Membres  []Member
	Groupes  []Groupe
	Concerts []Concert
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/band/{name}", bandHandler)
	r.HandleFunc("/", homeHandler)

	r.HandleFunc("/concerts", concertHandler)
	r.HandleFunc("/results", resultsHandler)

	fs := http.FileServer(http.Dir("./static"))
	staticDir := "/static/"
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, fs))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Server running ; head to http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Ouverture de la DB
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/groupes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Création de la requête et envoi à la DB
	query := "SELECT g.id, g.name, g.image, g.creationDate, g.firstAlbum FROM groupes g"

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Création de la liste de data récupérées depuis la DB
	var groupes []Groupe
	for rows.Next() {
		var groupe Groupe
		err := rows.Scan(&groupe.ID, &groupe.Name, &groupe.Image, &groupe.CreationDate, &groupe.FirstAlbum)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		groupes = append(groupes, groupe)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parsing de la page
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	// Exécution du template avec les retrieved data de la DB
	err = tmpl.Execute(w, groupes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func concertHandler(w http.ResponseWriter, r *http.Request) {
	// Ouverture de la DB
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/groupes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Création de la requête et envoi à la DB
	query := "SELECT l.name, p.name, d.date, g.name FROM `concerts` c JOIN `lieux` l ON c.id_lieu = l.id JOIN `pays` p ON l.id_pays = p.id JOIN `dates` d ON c.id_date = d.id JOIN `groupes_concerts` gc ON c.id_concert = gc.id_concerts JOIN `groupes` g ON gc.id_groupes = g.id WHERE 1 = 1 ORDER BY d.date DESC"

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Création de la liste de data récupérées depuis la DB
	var concerts []Concert
	for rows.Next() {
		var concert Concert
		var dateStr string
		err := rows.Scan(&concert.Location, &concert.Country, &dateStr, &concert.Artist)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		concert.Date = parsedDate.Format("2 Jan 2006")

		concerts = append(concerts, concert)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	// Parsing de la page
	tmpl, err := template.ParseFiles("static/concerts.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.Header().Add("Content-Type", "text/css")

	// Exécution du template avec les retrieved data de la DB
	err = tmpl.Execute(w, concerts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func bandHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	// Ouverture de la DB
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/groupes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// MEMBRES
	// Création de la requête et envoi à la DB
	query := "SELECT m.name FROM membres m INNER JOIN groupes_membres gm ON gm.id_membres = m.id INNER JOIN groupes g ON g.id = gm.id_groupes WHERE g.name LIKE '%" + name + "%'"

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Création de la liste de data récupérées depuis la DB
	var membres []Member
	for rows.Next() {
		var membre Member
		err := rows.Scan(&membre.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		membres = append(membres, membre)
	}

	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// GROUPES
	// Création de la requête et envoi à la DB
	query_groupes := "SELECT g.name, g.image, g.description, g.firstAlbum, g.creationDate FROM groupes g WHERE g.name LIKE '%" + name + "%'"

	rows3, err := db.Query(query_groupes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows3.Close()

	// Création de la liste de data récupérées depuis la DB
	var groupes []Groupe
	for rows3.Next() {
		var groupe Groupe
		var dateStr string
		err := rows3.Scan(&groupe.Name, &groupe.Image, &groupe.Description, &dateStr, &groupe.CreationDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		groupe.FirstAlbum = parsedDate.Format("2 Jan 2006")

		groupes = append(groupes, groupe)
	}

	err = rows3.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// CONCERTS

	query2 := "SELECT l.name, p.name, d.date FROM concerts c JOIN lieux l ON c.id_lieu = l.id JOIN pays p ON l.id_pays = p.id JOIN dates d ON c.id_date = d.id JOIN groupes_concerts gc ON c.id_concert = gc.id_concerts JOIN groupes g ON gc.id_groupes = g.id WHERE g.name LIKE '%" + name + "%'"
	rows2, err := db.Query(query2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows2.Close()

	// Création de la liste de data récupérées depuis la DB
	var concerts []Concert
	for rows2.Next() {
		var concert Concert
		var dateStr string
		err := rows2.Scan(&concert.Location, &concert.Country, &dateStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		concert.Date = parsedDate.Format("2 Jan 2006")

		concerts = append(concerts, concert)
	}

	err = rows2.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	results := Result{
		Membres:  membres,
		Groupes:  groupes,
		Concerts: concerts,
	}

	// Parsing de la page
	tmpl, err := template.ParseFiles("static/band.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	tmpl.Execute(w, results)
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/groupes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	name := r.FormValue("name")
	log.Println(r.Form)

	query := "SELECT name, image, creationDate, firstAlbum FROM groupes WHERE name LIKE '%" + name + "%'"

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var groupes []Groupe
	for rows.Next() {
		var groupe Groupe
		err := rows.Scan(&groupe.Name, &groupe.Image, &groupe.CreationDate, &groupe.FirstAlbum)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		groupes = append(groupes, groupe)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/results.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	err = tmpl.Execute(w, groupes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
