package main

import (
	"os"
	"fmt"
	"gosqlite.googlecode.com/hg/sqlite"
)

func main() {
	fmt.Println("********* Reading from a SQLite3 Database **********")
	db := "test.db"

	conn, err := sqlite.Open(db)
	if err != nil {
		fmt.Println("Unable to open the database: %s", err)
		os.Exit(1)
	}

	defer conn.Close()

	conn.Exec("CREATE TABLE articles(id INTEGER PRIMARY KEY AUTOINCREMENT, title VARCHAR(200), body TEXT, date TEXT);")

	insertSql := `INSERT INTO articles(title, body, date) VALUES("This is a Test Article Title.",
        "MO, dates are a pain.  I spent considerable time trying to decide how best to
        store dates in my app(s), and eventually chose to use Unix times (integers).
        It seemed an easy choice as I program in Perl and JavaScript.

        Lately, I've begun to regret the choice I made.  Every ad-hoc query I need to
        do (select * from mytable...) becomes an exercise in using SQLite date
        functions.  If I had it to do over, I would probably store my datetimes as
        YYYY-MM-DD HH:MM:SS strings.",
        "12/05/2010");`

	err = conn.Exec(insertSql)
	if err != nil {
		fmt.Println("Error while Inserting: %s", err)
	}

	selectStmt, err := conn.Prepare("SELECT id, title, body, date FROM articles;")
	err = selectStmt.Exec()
	if err != nil {
		fmt.Println("Error while Selecting: %s", err)
	}

	if selectStmt.Next() {
		var id int 
		var title string 
		var body string
		var date string
		
		err = selectStmt.Scan(&id, &title, &body, &date)
		if err != nil {
            fmt.Printf("Error while getting row data: %s\n", err)
            os.Exit(1)
		}
		fmt.Printf("Id => %s\n", id)
		fmt.Printf("Title => %s\n", title)
		fmt.Printf("Body => %s\n", body)
		fmt.Printf("Date => %s\n", date)
	}
}
