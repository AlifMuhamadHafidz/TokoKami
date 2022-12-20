package barang

import (
	"database/sql"
	"errors"
	"log"
)

type Barang struct {
	ID   int
	Nama string
	Info string
	Stok string
}

type MenuBarang struct {
	DB *sql.DB
}

func (mb *MenuBarang) TambahBarang(newBarang Barang) (int, error) {
	// menyiapakn query untuk insert
	addBarangQry, err := mb.DB.Prepare("INSERT INTO barang (nama_barang, info_barang, stok_barang) values (?,?,?)")
	if err != nil {
		log.Println("Prepare Insert Barang", err.Error())
		return 0, errors.New("Prepare statement insert barang error")
	}

	// menjalankan query dengan parameter tertentu
	res, err := addBarangQry.Exec(newBarang.Nama, newBarang.Info, newBarang.Stok)
	if err != nil {
		log.Println("Insert barang", err.Error())
		return 0, errors.New("Insert barang error")
	}

	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()
	if err != nil {
		log.Println("After insert barang ", err.Error())
		return 0, errors.New("error setelah insert")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return 0, errors.New("no record")
	}

	id, _ := res.LastInsertId()

	return int(id), nil
}
