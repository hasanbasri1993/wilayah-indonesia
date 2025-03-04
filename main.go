package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
	"wilayah/model"
	"wilayah/pkg"
)

func main() {

	pkg.DownloadSql()
	pkg.ExportToCSV()
	file, err := os.Open("wilayah.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	var provinsiData []model.Provinsi
	var kabupatenData []model.Kabupaten
	var kecamatanData []model.Kecamatan
	var kelurahanData []model.Kelurahan

	for _, record := range records[1:] {
		kode := record[0]
		nama := record[1]

		switch len(kode) {
		case 2:
			provinsiData = append(provinsiData, model.Provinsi{KodeProvinsi: kode, Nama: nama})
		case 5:
			kodeProvinsi := kode[:2]
			kabupatenData = append(kabupatenData, model.Kabupaten{KodeKabupaten: kode, Nama: nama, KodeProvinsi: kodeProvinsi})
		case 8:
			kodeKabupaten := kode[:5]
			kecamatanData = append(kecamatanData, model.Kecamatan{KodeKecamatan: kode, Nama: nama, KodeKabupaten: kodeKabupaten})
		case 13:
			kodeKecamatan := kode[:8]
			kelurahanData = append(kelurahanData, model.Kelurahan{KodeKelurahan: kode, Nama: nama, KodeKecamatan: kodeKecamatan})
		}
	}

	fmt.Println("Provinsi. Writing to CSV...")
	err = pkg.WriteCSV("provinsi.csv", provinsiData)
	if err != nil {
		return
	}
	fmt.Println("Done")
	time.Sleep(1 * time.Second)

	fmt.Println("Kabupaten. Writing to CSV...")
	err = pkg.WriteCSV("kabupaten.csv", kabupatenData)
	if err != nil {
		return
	}
	fmt.Println("Done")
	time.Sleep(1 * time.Second)

	fmt.Println("Kecamatan. Writing to CSV...")
	err = pkg.WriteCSV("kecamatan.csv", kecamatanData)
	if err != nil {
		return
	}
	fmt.Println("Done")
	time.Sleep(1 * time.Second)

	fmt.Println("Kelurahan. Writing to CSV...")
	err = pkg.WriteCSV("kelurahan.csv", kelurahanData)
	if err != nil {
		return
	}

	// Read wilayah_kodepos.csv
	kodeposMap, err := pkg.ReadWilayahKodepos("wilayah_kodepos.csv")
	if err != nil {
		fmt.Println("Error reading wilayah_kodepos.csv:", err)
		return
	}

	// Merge kelurahanData with kodeposMap
	for i := range kelurahanData {
		if kodepos, found := kodeposMap[strings.ReplaceAll(kelurahanData[i].KodeKelurahan, ".", "")]; found {
			kelurahanData[i].KodePos = kodepos
		}
	}

	fmt.Println("Update kodepos. Writing to CSV...")
	err = pkg.WriteCSV("kelurahan.csv", kelurahanData)
	if err != nil {
		return
	}
	fmt.Println("Update kodepos. Done.")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
	pkg.RemoveTemp()
}
