package pkg

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	wilayah = "https://raw.githubusercontent.com/cahyadsn/wilayah/refs/heads/master/db/wilayah.sql"
	kodepos = "https://raw.githubusercontent.com/cahyadsn/wilayah_kodepos/refs/heads/main/db/wilayah_kodepos.sql"
)

func downloadFile(url string, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file from %s: %w", url, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)

	file, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", dest, err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", dest, err)
	}

	return nil
}

func executeSQLFile(db *sql.DB, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file %s: %w", filePath, err)
	}

	_, err = db.Exec(string(data))
	if err != nil {
		return fmt.Errorf("failed to execute SQL file %s: %w", filePath, err)
	}

	return nil
}

func DownloadSql() {
	// Temporary files to save SQL scripts
	wilayahFile := "wilayah.sql"
	kodeposFile := "kodepos.sql"

	// Download wilayah.sql
	fmt.Println("Downloading wilayah.sql...")
	if err := downloadFile(wilayah, wilayahFile); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Downloaded wilayah.sql")

	// Download kodepos.sql
	fmt.Println("Downloading kodepos.sql...")
	if err := downloadFile(kodepos, kodeposFile); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Downloaded kodepos.sql")

	// Create SQLite database
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println("Error creating database:", err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Error closing database:", err)
		}
	}(db)

	// Execute wilayah.sql
	fmt.Println("Executing wilayah.sql...")
	if err := executeSQLFile(db, wilayahFile); err != nil {
		fmt.Println("Error executing wilayah.sql:", err)
		return
	}
	fmt.Println("Executed wilayah.sql")

	// Execute kodepos.sql
	fmt.Println("Executing kodepos.sql...")
	if err := executeSQLFile(db, kodeposFile); err != nil {
		fmt.Println("Error executing kodepos.sql:", err)
		return
	}
	fmt.Println("Executed kodepos.sql")

	fmt.Println("Database setup completed.")

}
