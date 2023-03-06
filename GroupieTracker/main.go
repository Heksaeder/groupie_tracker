package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Groupe struct {
	ID           int
	Name         string
	Image        string
	CreationDate string
	FirstAlbum   string
	Member       string
}

type Concert struct {
	ID       int
	Artist   string
	Date     string
	Country  string
	Location string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/band", bandHandler)
	http.HandleFunc("/concerts", concertHandler)
	http.HandleFunc("/results", resultsHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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
	query := "SELECT l.name, p.name, d.date, g.name FROM `concerts` c JOIN `lieux` l ON c.id_lieu = l.id JOIN `pays` p ON l.id_pays = p.id JOIN `dates` d ON c.id_date = d.id JOIN `groupes_concerts` gc ON c.id_concert = gc.id_concerts JOIN `groupes` g ON gc.id_groupes = g.id WHERE 1 = 1"

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

	// Parsing de la page
	tmpl, err := template.ParseFiles("static/concerts.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	// Exécution du template avec les retrieved data de la DB
	err = tmpl.Execute(w, concerts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func bandHandler(w http.ResponseWriter, r *http.Request) {
	// Ouverture de la DB
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/groupes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	name := r.FormValue("name")
	log.Println(r.Form)

	// Création de la requête et envoi à la DB
	query := "SELECT g.name, m.name FROM membres m INNER JOIN groupes_membres gm ON gm.id_membres = m.id INNER JOIN groupes g ON g.id = gm.id_groupes WHERE g.name LIKE '%" + name + "%'"
	// ADD A POST METHOD TO GET THE BAND NAME FROM LINK

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
		err := rows.Scan(&groupe.Name, &groupe.Member)
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
	tmpl, err := template.ParseFiles("static/band.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécution du template avec les retrieved data de la DB
	err = tmpl.Execute(w, groupes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
