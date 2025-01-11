package pkg

import "os"

func RemoveTemp() {
	var files = []string{"data.db", "wilayah.sql", "kodepos.sql", "wilayah.csv", "wilayah_kodepos.csv"}
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			panic(err)
		}
	}
}
