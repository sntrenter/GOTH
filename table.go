package main

import (
	"database/sql"
	"html/template"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func addRow(url string) string {
	//TODO: add logging
	log.Println("adding row, url: ", url)
	//add row to sqlite db filename is "url.db" table name is "url" the first column is an integer (automatic) and the second column is a string
	db, err := sql.Open("sqlite3", "url.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var existingURL string
	err = db.QueryRow("SELECT url FROM url WHERE url = ?", url).Scan(&existingURL)
	if err == nil {
		log.Println("Row already exists")
		return "Row already exists"
	}
	if err != sql.ErrNoRows {
		log.Fatal(err)
	}
	result, err := db.Exec("INSERT INTO url (url) VALUES (?)", url)
	if err != nil {
		log.Fatal(err)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	// Optionally, you can fetch and return the newly created row if needed
	var newRowURL string
	err = db.QueryRow("SELECT url FROM url WHERE rowid = ?", lastInsertID).Scan(&newRowURL)
	if err != nil {
		log.Fatal(err)
	}

	return "Row added: " + newRowURL
}

func getRowsAsHTML(num int) string {
	// Initialize an empty HTML string
	html := ""
	var rows *sql.Rows
	var err error
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "url.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query the database for rows

	if num == 0 {
		rows, err = db.Query("SELECT rowid, url FROM url")
	} else {
		rows, err = db.Query("SELECT rowid, url FROM url LIMIT ?", num)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate over the rows and build the HTML table rows
	for rows.Next() {
		var id int64
		var url string
		if err := rows.Scan(&id, &url); err != nil {
			log.Fatal(err)
		}
		html += "<tr><td>" + strconv.FormatInt(id, 10) + "</td><td>" + template.HTMLEscapeString(url) + "</td></tr>"
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Return the generated HTML
	return html
}
