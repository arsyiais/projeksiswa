package controller

import (
	"siswa/model"
	"siswa/node"
)

func InsertSiswa(id int, nama string, kelas int, alamat string, nilai float32) bool {
	if id > 0 && nama != "" {
		model.InsertSiswa(id, nama, kelas, alamat, nilai)
		return true
	} else {
		return false
	}
}

func UpdateSiswa(id int, nama string, kelas int, alamat string, nilai float32) bool {
	if id > 0 && nama != "" {
		siswa := node.Siswa{
			Id:     id,
			Nama:   nama,
			Kelas:  kelas,
			Alamat: alamat,
			Nilai:  nilai,
		}
		return model.UpdateSiswa(siswa)
	}
	return false
}

func SearchSiswa(id int, nama string, kelas int, Alamat string, Nilai int) *node.Siswa {
	if id > 0 {
		addr := model.SearchSiswa(id)
		if addr != nil {
			return &addr.Data
		}
	}
	return nil
}

func DeleteSiswa(id int) bool {
	if id > 0 {
		return model.DeleteSiswa(id)
	}
	return false
}

func ReadAllSiswa() *node.TableSiswa {
	return model.ReadAllSiswa()
}
