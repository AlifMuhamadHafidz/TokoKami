package barang

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
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


// CREATE
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


//UPDATE 1
func (mb *MenuBarang) EditInfoBarang(newBarang Barang, namaBarang string) (bool, error) {
	editInfoBarangQry, err := mb.DB.Prepare("UPDATE barang SET info_barang = ? where nama_barang = ?")

	if err != nil {
		log.Println("Prepare update info barang", err.Error())
		return false, errors.New("prepare statement update info barang error")
	}

	res, err := editInfoBarangQry.Exec(newBarang.Info, namaBarang)

	if err != nil {
		log.Println("Update info Barang", err.Error())
		return false, errors.New("Update info Barang Error")
	}

	//Cek Baris yang terpengaruh query Update Info Barang
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("After Update Info Barang", err.Error())
		return false, errors.New("Error setelah Update Info Barang")
	}

	if affRows < 0 {
		log.Println("no records affected")
		return false, errors.New("no Record")
	}

	return true, nil
}

// UPDATE 2
func (mb *MenuBarang) EditStokBarang(stokBarang int, namaBarang string) (bool, error) {
	editInfoBarangQry, err := mb.DB.Prepare("UPDATE barang SET stok_barang = ? where nama_barang = ?")

	if err != nil {
		log.Println("Prepare update stok barang", err.Error())
		return false, errors.New("prepare statement update stok barang error")
	}

	res, err := editInfoBarangQry.Exec(stokBarang, namaBarang)

	if err != nil {
		log.Println("Update stok Barang", err.Error())
		return false, errors.New("Update stok Barang Error")
	}

	//Cek Baris yang terpengaruh query Update Info Barang
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("After Update stok Barang", err.Error())
		return false, errors.New("Error setelah Update stok Barang")
	}

	if affRows < 0 {
		log.Println("no records affected")
		return false, errors.New("no Record")
	}

	return true, nil
}

//READ
func (mb *MenuBarang) ListBarang() [][]string {
	rows, err := mb.DB.Query("SELECT id_barang, nama_barang, stok_barang FROM barang")
	if err != nil {
		log.Println("SELECT ERROR", err.Error())
	}
	arrBarang := []string{}
	arrBarangs := [][]string{}
	for rows.Next() {
		var id int
		var nama_barang string
		var stok int
		rows.Scan(&id, &nama_barang, &stok)
		if err != nil {
			log.Println("Error scan isi tabel pegawai", err.Error())
		}
		arrBarang = append(arrBarang, strconv.Itoa(id), nama_barang, strconv.Itoa(stok))
		arrBarangs = append(arrBarangs, arrBarang)
		arrBarang = nil
	}
	return arrBarangs
}

// DELETE
func (mb *MenuBarang) DeleteBarang(id int) (error) {
	_, err := mb.DB.Query("DELETE FROM barang where id_barang = ?", id)
	if err != nil {
		log.Println("Gagal saat menghapus", err.Error())
		return errors.New("Data gagal dihapus")
	}
	// log.Println(row)
	return nil
}
