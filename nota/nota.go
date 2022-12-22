package nota

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	// "strconv"
)

type Nota struct {
	ID_Trans 	int
	ID_Barang	int
	Jumlah		int
}

type NotaMenu struct {
	DB *sql.DB
}

func (nm *NotaMenu) AddNota(newNota Nota) (bool, error) {
	addNotaQry, err := nm.DB.Prepare("INSERT INTO nota(id_transaksi, id_barang, qty) values (?,?,?);")
	if err != nil {
		log.Println("insert nota prepare", err.Error())
		return false, errors.New("prepare insert transaksi error")
	}

	res, err := addNotaQry.Exec(newNota.ID_Trans, newNota.ID_Barang, newNota.Jumlah)
	if err != nil {
		log.Println("insert nota", err.Error())
		return false, errors.New("insert nota error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert nota", err.Error())
		return false, errors.New("after insert nota error")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return true, errors.New("no record")
	}

	return true, nil
}



func (nm *NotaMenu) ListNota(id int) [][]string {
	rows, err := nm.DB.Query("SELECT b.nama_barang, n.qty FROM barang b, nota n WHERE n.id_barang = b.id_barang AND n.id_transaksi = ?", id)
	if err != nil {
		log.Println("SELECT ERROR", err.Error())
	}
	arrNota := []string{}
	arrNotas := [][]string{}
	for rows.Next() {
		var nama string
		var jumlah int
		rows.Scan(&nama, &jumlah)
		if err != nil {
			log.Println("Error scan isi tabel nota", err.Error())
		}
		arrNota = append(arrNota, nama, strconv.Itoa(jumlah))
		arrNotas = append(arrNotas, arrNota)
		arrNota = nil
	}
	return arrNotas
}

func (nm *NotaMenu) ListTranNota() []string {
	rows, err := nm.DB.Query("SELECT t.id_transaksi, t.tgl_transaksi, p.nama , p2.username FROM transaksi t, pelanggan p, pegawai p2 WHERE p2.id = t.id_pegawai AND p.id_pelanggan = t.id_pelanggan ORDER BY 1 DESC LIMIT 1;")
	if err != nil {
		log.Println("SELECT ERROR", err.Error())
	}
	arrTran := []string{}
	// arrTrans := [][]string{}
	for rows.Next() {
		var id_tran int
		var tgl_trans string
		var nama_cus string
		var nama_peg string
		rows.Scan(&id_tran, &tgl_trans, &nama_cus, &nama_peg)
		if err != nil {
			log.Println("Error scan isi tabel transaksi", err.Error())
		}
		arrTran = append(arrTran, strconv.Itoa(id_tran), tgl_trans, nama_peg, nama_cus)
		// arrTrans = append(arrTrans, arrTran)
		// arrTran = nil
	}
	// log.Println(arrTran)
	return arrTran
}