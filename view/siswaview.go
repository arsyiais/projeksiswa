package view

import (
	"fmt"
	"siswa/controller"
)

func InsertSiswa() {
	var id, kelas int
	var nama, alamat string
	var nilai float32
	fmt.Print("Masukkan id siswa : ")
	fmt.Scan(&id)
	fmt.Print("Masukkan nama siswa : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan kelas siswa : ")
	fmt.Scan(&kelas)
	fmt.Print("Masukkan alamat siswa : ")
	fmt.Scan(&alamat)
	fmt.Print("Masukkan nilai siswa : ")
	fmt.Scan(&nilai)
	cek := controller.InsertSiswa(id, nama, kelas, alamat, nilai)
	if cek == true {
		fmt.Println("data berhasil diinput")
	} else {
		fmt.Println("data gagal diinput")
	}
}

func UpdateSiswa() {
	var id, kelas int
	var nama, alamat string
	var nilai float32
	fmt.Print("Masukkan id siswa : ")
	fmt.Scan(&id)
	fmt.Print("Masukkan nama siswa : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan kelas siswa : ")
	fmt.Scan(&kelas)
	fmt.Print("Masukkan alamat siswa : ")
	fmt.Scan(&alamat)
	fmt.Print("Masukkan nilai siswa : ")
	fmt.Scan(&nilai)
	cek := controller.UpdateSiswa(id, nama, kelas, alamat, nilai)
	if cek == true {
		fmt.Println("Data berhasil diupdate")
	} else {
		fmt.Println("Data gagal diupdate atau sepatu tidak ditemukan")
	}
}

func ReadAllSiswa() {
	siswas := controller.ReadAllSiswa()
	if siswas != nil {
		for siswas.Next != nil {
			fmt.Println("Id Siswa : ", siswas.Next.Data.Id)
			fmt.Println("Nama Siswa : ", siswas.Next.Data.Nama)
			fmt.Println("Kelas Siswa : ", siswas.Next.Data.Kelas)
			fmt.Println("Alamat Siswa : ", siswas.Next.Data.Alamat)
			fmt.Println("Nilai Siswa : ", siswas.Next.Data.Nilai)
			fmt.Println("---------------------")
			siswas = siswas.Next
		}
	}
}

func SearchSiswa() {
	var id, kelas, nilai int
	var nama, alamat string

	fmt.Print("Masukkan id siswa yang ingin dicari: ")
	fmt.Scan(&id)
	siswa := controller.SearchSiswa(id, nama, kelas, alamat, nilai)
	if siswa != nil {
		fmt.Println("Id Siswa : ", siswa.Id)
		fmt.Println("Nama Siswa : ", siswa.Nama)
		fmt.Println("Kelas Siswa : ", siswa.Kelas)
		fmt.Println("Alamat Siswa : ", siswa.Alamat)
		fmt.Println("Nilai Siswa : ", siswa.Nilai)
	} else {
		fmt.Println("Siswa tidak ditemukan")
	}
}

func DeleteSiswa() {
	var id int
	fmt.Println("Masukkan id siswa yang ingin dihapus: ")
	fmt.Scan(&id)
	cek := controller.DeleteSiswa(id)
	if cek == true {
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Data gagal dihapus atau siswa tidak ditemukan")
	}
}
