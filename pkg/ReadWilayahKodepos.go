package pkg

import (
	"encoding/csv"
	"os"
	"strings"
)

func ReadWilayahKodepos(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	kodeposMap := make(map[string]string)
	for _, record := range records {
		kode := strings.ReplaceAll(record[0], ".", "")
		kodepos := record[1]
		kodeposMap[kode] = kodepos
	}
	return kodeposMap, nil
}
