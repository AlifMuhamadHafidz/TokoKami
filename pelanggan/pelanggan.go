package pelanggan

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type Pelanggan struct {
	ID		int
	Nama	string
}

type AuthPelanggan struct {
	DB *sql.DB
}

// Function untuk menambah pelanggan baru oleh pegawai
func (ape *AuthPelanggan) TambahPelanggan(newPelanggan Pelanggan) (bool, error) {
	addPelangganQry, err := ape.DB.Prepare("INSERT INTO pelanggan (nama) values (?)")
	if err != nil {
		log.Println("Prepare insert pelanggan ", err.Error())
		return false, errors.New("insert pelanggan error")
	}

	if ape.DuplicatePelanggan(newPelanggan.Nama) {
		log.Println("duplicated information")
		return false, errors.New("name sudah digunakan")
	}

	res, err := addPelangganQry.Exec(newPelanggan.Nama)
	if err != nil {
		log.Println("insert pelanggan ", err.Error())
		return false, errors.New("insert pelanggan error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("After insert pelanggan ", err.Error())
		return false, errors.New("error setelah insert pelanggan")
	}

	if affRows <= 0 {
		log.Println("No record affected")
		return false, errors.New("no record")
	}
	return true, nil
}

func (ape *AuthPelanggan) DuplicatePelanggan(nama string) bool {
	res := ape.DB.QueryRow("SELECT id_pelanggan FROM pelanggan where nama = ?", nama)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			log.Println("Result scan error", err.Error())
		}
		return false
	}
	return true
}

// 22.20
// function list daftar pelanggan
func (ape *AuthPelanggan) ListPelanggan() [][]string {
	rows, err := ape.DB.Query("SELECT id_pelanggan, nama FROM pelanggan")
	if err != nil {
		log.Println("Error select pelanggan", err.Error())
	}

	arrPelanggan := []string{}
	arrPelanggans := [][]string{}
	for rows.Next() {
		var id int
		var nama string
		rows.Scan(&id, &nama)
		if err != nil {
			log.Println("Error scan isi tabel pelanggan", err.Error())
		}
		arrPelanggan = append(arrPelanggan, strconv.Itoa(id), nama)
		arrPelanggans = append(arrPelanggans, arrPelanggan)
		arrPelanggan = nil
	}
	return arrPelanggans
}


// function hapus pelanggan
func (ape *AuthPelanggan) DeletePelanggan(id int) (error) {
	_, err := ape.DB.Query("DELETE FROM pelanggan where id_pelanggan = ?", id)
	if err != nil {
		log.Println("Gagal saat menghapus", err.Error())
		return errors.New("data gagal dihapus")
	}
	// log.Println(row)
	return nil
}
// 22.20