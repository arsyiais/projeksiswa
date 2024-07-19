package node

type Siswa struct {
	Id     int
	Nama   string
	Kelas  int
	Alamat string
	Nilai  float32
}

type TableSiswa struct {
	Data Siswa
	Next *TableSiswa
}
