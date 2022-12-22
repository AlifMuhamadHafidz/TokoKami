package transaksi

import (
	"database/sql"
	"errors"
	"fmt"
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
	addTransQry, err := tm.DB.Prepare("INSERT INTO transaksi(id_pegawai, id_pelanggan) values (?,?);")
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

func (tm *TransMenu) DeleteTransaksi(deleteTransaksi Transaksi) (bool, error) {

	delTransQry, err := tm.DB.Prepare("DELETE FROM transaksi WHERE id_transaksi=?")
	if err != nil {
		log.Println("prepare delete transaksi ", err.Error())
		return false, errors.New("prepare statement delete transaksi error")
	}

	res, err := delTransQry.Exec(deleteTransaksi.ID)
	if err != nil {
		log.Println("delete transaksi ", err.Error())
		return false, errors.New("delete transaksi error")
	}
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after Delete transaksi ", err.Error())
		return false, errors.New("error after delete transaksi")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (tm *TransMenu) ShowTransaksi(newTransaksi Transaksi) (bool, error) {

	res, err := tm.DB.Query("SELECT * FROM transaksi")
	if err != nil {
		log.Println("select transaksi", err.Error())
		return false, errors.New("select transaksi error")
	}

	for res.Next() {
		var showTransaksi Transaksi

		err := res.Scan(&showTransaksi.ID, &showTransaksi.ID_Pegawai, &showTransaksi.ID_Pelanggan, &showTransaksi.ID_Barang)

		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%v\n", showTransaksi)
	}

	return true, nil
}
