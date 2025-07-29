// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "modernc.org/sqlite"
// )

// type Phone struct {
// 	Id      int
// 	Company string
// 	Model   string
// 	Price   int
// }

// func main() {
// 	db, err := sql.Open("sqlite", "lesson22.db")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer db.Close()
// 	rows, err := db.Query("select * from phones")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer rows.Close()
// 	phones := []Phone{}
// 	for rows.Next() {
// 		p := Phone{}
// 		err := rows.Scan(&p.Id, &p.Company, &p.Model, &p.Price)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		phones = append(phones, p)
// 	}
// 	fmt.Println(phones)
// }

// /*
// 	p := Phone{}
// 	p.Id = 1
// 	p.Company = "Apple"
// */

package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Phone struct {
	Id      int
	Company string
	Model   string
	Price   int
}

func SelectData() {
	db, err := sql.Open("sqlite", "lesson22.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from phones")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	phones := []Phone{}
	for rows.Next() {
		p := Phone{}
		err := rows.Scan(&p.Id, &p.Company, &p.Model, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		phones = append(phones, p)
	}
	fmt.Println(phones)
}
func main() {
	SelectData()
	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	_, err = db.Exec("UPDATE phones SET Price = $1 WHERE Company = $2", 2000, "Apple")
	SelectData()
}
