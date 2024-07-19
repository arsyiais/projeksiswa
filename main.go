package main

import (
	"fmt"
	"siswa/view"
)

func menu() {
	for {
		fmt.Println("Menu Program")
		fmt.Println("1. Insert Siswa")
		fmt.Println("2. Update Siswa")
		fmt.Println("3. Delete Siswa")
		fmt.Println("4. Read All Siswa")
		fmt.Println("5. Search Siswa")
		fmt.Println("6. Exit")
		fmt.Println("---------------------")
		fmt.Print("masukkan menu pilihan anda: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			view.InsertSiswa()
		case 2:
			view.UpdateSiswa()
		case 3:
			view.DeleteSiswa()
		case 4:
			view.ReadAllSiswa()
		case 5:
			view.SearchSiswa()
		case 6:
			fmt.Println("exit program")
			return
		}
	}
}

func main() {
	menu()
}

// controller.InsertSiswa(1, "Habibie", 6, "Sidoajo", 75)
// controller.InsertSiswa(2, "Zanki", 1, "Surabaya", 87)
// controller.InsertSiswa(3, "Yukio", 4, "Sidoajo", 72)
// controller.InsertSiswa(4, "Yustian", 6, "Sidoajo", 90)
// controller.InsertSiswa(5, "Abi", 6, "Surabaya", 84)

// fmt.Println(controller.DeleteSiswa(2))
// siswas := controller.ReadAllSiswa()
// fmt.Println(model.SearchSiswa(2))
// siswa2 := node.Siswa{
// 	Id:     2,
// 	Nama:   "Zanki",
// 	Kelas:  2,
// 	Alamat: "Surabaya",
// 	Nilai:  94,
// }

// tes := model.UpdateSiswa(siswa2)
// fmt.Println("return update siswa :", tes)

// // model.SearchSiswa(4)
// // model.UpdateSiswa(2, "Zanki", 2, "Sidoajo", 83)

// // fmt.Println(siswas)
// // fmt.Println(siswa.Next)

// if siswas != nil {
// 	for siswas.Next != nil {
// 		fmt.Println(siswas.Next.Data)
// 		siswas = siswas.Next
// 	}
// }

// // SearchSiswa(2)
// // UpdateSiswa(5, "Abi", 7, "Sidoajo", 93)
// // DeleteSiswa(2)
// ReadAllSiswa()
