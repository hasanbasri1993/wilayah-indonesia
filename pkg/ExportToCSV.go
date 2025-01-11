package pkg

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
)

func exportTableToCSV(db *sql.DB, tableName, fileName string) error {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to query table %s: %w", tableName, err)
	}
	defer rows.Close()

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fileName, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("failed to get columns for table %s: %w", tableName, err)
	}
	if err := writer.Write(columns); err != nil {
		return fmt.Errorf("failed to write headers to CSV file %s: %w", fileName, err)
	}

	// Write rows
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	for rows.Next() {
		if err := rows.Scan(valuePtrs...); err != nil {
			return fmt.Errorf("failed to scan row for table %s: %w", tableName, err)
		}

		row := make([]string, len(columns))
		for i, val := range values {
			if val != nil {
				row[i] = fmt.Sprintf("%v", val)
			} else {
				row[i] = ""
			}
		}

		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write row to CSV file %s: %w", fileName, err)
		}
	}

	return nil
}

func ExportToCSV() {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Error closing database:", err)
		}
	}(db)

	// Export wilayah table to CSV
	fmt.Println("Exporting wilayah table to wilayah.csv...")
	if err := exportTableToCSV(db, "wilayah", "wilayah.csv"); err != nil {
		fmt.Println("Error exporting wilayah table:", err)
		return
	}
	fmt.Println("Exported wilayah table to wilayah.csv")

	// Export kodepos table to CSV
	fmt.Println("Exporting kodepos table to kodepos.csv...")
	if err := exportTableToCSV(db, "wilayah_kodepos", "wilayah_kodepos.csv"); err != nil {
		fmt.Println("Error exporting kodepos table:", err)
		return
	}
	fmt.Println("Exported kodepos table to kodepos.csv")
}
