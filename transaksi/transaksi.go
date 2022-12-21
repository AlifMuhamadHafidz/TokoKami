package transaksi

import (
	"database/sql"
	"errors"
	"log"
)

type Transaksi struct {
	ID           int
	ID_Pegawai   int
	ID_Pelanggan int
	ID_Barang    int
}

type TransMenu struct {
	DB *sql.DB
}

func (tm *TransMenu) AddTransaksi(newTransaksi Transaksi) (bool, error) {
	addTransQry, err := tm.DB.Prepare("INSERT into transaksi(id_pegawai, id_pelanggan) values (?, ?)")
	if err != nil {
		log.Println("insert transaksi prepare", err.Error())
		return false, errors.New("prepare insert transaksi error")
	}

	res, err := addTransQry.Exec(newTransaksi.ID_Pegawai, newTransaksi.ID_Pelanggan)
	if err != nil {
		log.Println("insert transaksi", err.Error())
		return false, errors.New("insert transaksi error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert transaksi", err.Error())
		return false, errors.New("after insert transaksi error")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return true, errors.New("no record")
	}

	return true, nil
}
