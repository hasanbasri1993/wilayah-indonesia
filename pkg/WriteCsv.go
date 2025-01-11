package pkg

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"wilayah/model"
)

func WriteCSV(filename string, data interface{}) error {
	var records [][]string

	switch v := data.(type) {
	case []model.Provinsi:
		records = append(records, []string{"KodeProvinsi", "Nama"})
		for _, p := range v {
			records = append(records, []string{p.KodeProvinsi, p.Nama})
		}
	case []model.Kabupaten:
		records = append(records, []string{"KodeKabupaten", "Nama", "KodeProvinsi"})
		for _, k := range v {
			records = append(records, []string{strings.ReplaceAll(k.KodeKabupaten, ".", ""), k.Nama, k.KodeProvinsi})
		}
	case []model.Kecamatan:
		records = append(records, []string{"KodeKecamatan", "Nama", "KodeKabupaten"})
		for _, kec := range v {
			records = append(records, []string{strings.ReplaceAll(kec.KodeKecamatan, ".", ""), kec.Nama, strings.ReplaceAll(kec.KodeKabupaten, ".", "")})
		}
	case []model.Kelurahan:
		records = append(records, []string{"KodeKelurahan", "Nama", "KodeKecamatan", "KodePos"})
		for _, kel := range v {
			records = append(records, []string{strings.ReplaceAll(kel.KodeKelurahan, ".", ""), kel.Nama, strings.ReplaceAll(kel.KodeKecamatan, ".", ""), kel.KodePos})
		}
	default:
		return fmt.Errorf("unsupported data type")
	}

	file, err := os.Create("output/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}
