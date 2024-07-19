package model

import (
	"siswa/database"
	"siswa/node"
)

func InsertSiswa(id int, nama string, kelas int, alamat string, nilai float32) {
	newDataSiswa := node.TableSiswa{
		Data: node.Siswa{id, nama, kelas, alamat, nilai},
		Next: nil,
	}
	var tempLL *node.TableSiswa
	tempLL = &database.DBsiswa
	if database.DBsiswa.Next == nil {
		database.DBsiswa.Next = &newDataSiswa
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newDataSiswa
	}
}
func ReadAllSiswa() *node.TableSiswa {
	var tempLL *node.TableSiswa
	tempLL = &database.DBsiswa
	if database.DBsiswa.Next == nil {
		return nil
	} else {
		return tempLL

	}
}

func SearchSiswa(id int) *node.TableSiswa {
	var tempLL *node.TableSiswa
	tempLL = database.DBsiswa.Next
	cek := false
	if database.DBsiswa.Next == nil {
		return nil
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				return tempLL
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			return nil
		}
	}
	return nil
}

func UpdateSiswa(siswa node.Siswa) bool {
	addr := SearchSiswa(siswa.Id)
	if addr == nil {
		return false
	} else {
		addr.Data.Nama = siswa.Nama
		addr.Data.Kelas = siswa.Kelas
		addr.Data.Alamat = siswa.Alamat
		addr.Data.Nilai = siswa.Nilai
		return true
	}
}

func DeleteSiswa(id int) bool {
	var tempLL *node.TableSiswa
	tempLL = &database.DBsiswa
	if tempLL.Next == nil {
		return false
	} else {
		for tempLL.Next != nil {
			if tempLL.Next.Data.Id == id {
				tempLL.Next = tempLL.Next.Next
				return true
			}
			tempLL = tempLL.Next
		}
		return false
	}

}
