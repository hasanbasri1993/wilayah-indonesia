package model

type Provinsi struct {
	KodeProvinsi string
	Nama         string
}

type Kabupaten struct {
	KodeProvinsi  string
	Nama          string
	KodeKabupaten string
}

type Kecamatan struct {
	KodeKecamatan string
	Nama          string
	KodeKabupaten string
}

type Kelurahan struct {
	KodeKelurahan string
	Nama          string
	KodeKecamatan string
	KodePos       string
}
