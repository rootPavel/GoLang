// database/sql - пакет для работы с базами, но необходим еще driver для работы с конкретной бд

package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Phone struct {
	Id     int
	Vendor string
	Model  string
	Price  int
}

func main() {
	db, err := sql.Open("sqlite", "Lesson21.db")
	if err != nil {
		log.Println("Cannot open db!")
		return
	}
	defer db.Close()

	// rows, err := db.Query("SELECT * FROM Phones")
	// rows, err := db.Query("SELECT * FROM Phones WHERE Vendor = 'Xiaomi'")
	// rows, err := db.Query(fmt.Sprintf("SELECT * FROM Phones WHERE Vendor = '%s'", "Xiaomi"))
	rows, err := db.Query("SELECT * FROM Phones WHERE Vendor = ?", "Xiaomi")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	Phones := []Phone{}

	// SELECT
	for rows.Next() {
		p := Phone{}
		err := rows.Scan(&p.Id, &p.Vendor, &p.Model, &p.Price)
		if err != nil {
			log.Println(err)
			continue
		}
		Phones = append(Phones, p)
	}
	// for _, v := range Phones {
	// 	fmt.Println(v)
	// }

	// END OF SELECT

	//------------------------

	// // INSERT
	// result, err := db.Exec(`INSERT INTO Phones (Vendor, Model, Price) VALUES ("Samsung", "Galaxy", 1500)`)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Println(result.LastInsertId())
	// log.Println(result.RowsAffected())
	// // END INSERT

	//-------------------------

	// UPDATE
	result, err := db.Exec(`UPDATE Phones SET Price = 4500 WHERE Model = "Galaxy"`)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())
	// END UPDATE

	//-------------------------

	// UPDATE
	result, err = db.Exec(`DELETE FROM Phones WHERE id = 6`)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(result.RowsAffected())
	// END UPDATE

}
