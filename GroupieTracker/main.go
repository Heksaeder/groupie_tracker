package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Groupe struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	CreationDate int    `json:"creationDate"`
	FirstAlbum   string `json:"firstAlbum"`
}

type Membre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Concert struct {
	ID     int
	IDLieu int
	IDDate int
	Lieu   string
	Date   string
}

type Pays struct {
	ID   int
	Name string
}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/groupes")
	if err != nil {
		panic(err.Error())
	}

	// Defer the close till after the main function has finished
	defer db.Close()

	// query := "SELECT g.name, g.image, g.creationDate, g.firstAlbum, m.name FROM membres m, groupes g WHERE m.id IN (SELECT id_membres FROM groupes_membres WHERE id_groupes =" + strconv.Itoa(id) + ") AND g.id IN (SELECT id_groupes FROM groupes_membres WHERE id_groupes =" + strconv.Itoa(id) + ")"

	// Récupère la liste des membres par groupe
	// query_three := "SELECT m.name FROM membres m WHERE m.id IN (SELECT id_membres FROM groupes_membres gm WHERE gm.id_groupes = 1)"

	// Récupère date et lieu d'un concert pour un id donné
	// query_four := "SELECT l.name, d.date FROM lieux l, dates d WHERE l.id IN (SELECT id_lieu FROM concerts c WHERE id_concert = 17) AND d.id IN (SELECT id_date FROM concerts c WHERE id_concert = 17)"

	// Récupère lieu, pays, date, groupe selon id concert
	query_five := "SELECT l.name, p.name, d.date, g.name FROM lieux l, pays p, dates d, groupes g WHERE l.id IN (SELECT id_lieu FROM concerts c WHERE id_concert = 17) AND p.id IN (SELECT id_pays FROM lieux l WHERE l.id IN (SELECT id_lieu FROM concerts WHERE id_concert = 17)) AND d.id IN (SELECT id_date FROM concerts WHERE id_concert = 17) AND g.id IN (SELECT id_groupes FROM groupes_concerts WHERE id_concerts = 17)"

	readTable, err := db.Query(query_five)
	if err != nil {
		panic(err.Error())
	}

	for readTable.Next() {
		var groupe Groupe
		// var membre Membre
		var concert Concert
		var pays Pays

		// for each row, scan the result into our groupe composite object
		err = readTable.Scan(&concert.Lieu, &pays.Name, &concert.Date, &groupe.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Print out the band's name, the band's creation date, the band's first album and the member's names
		fmt.Println(groupe.Name+" - "+concert.Lieu, concert.Date, pays.Name)
	}

	// be careful deferring Queries if you are using transactions
	defer readTable.Close()
}
