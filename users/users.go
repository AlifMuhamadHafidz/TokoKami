package users

import (
	"database/sql"
	"errors"
	"log"
)

type Pegawai struct {
	ID       int
	Username string
	Password string
}

type AuthMenu struct {
	DB *sql.DB
}

func (am *AuthMenu) Login(nama string, password string) (Pegawai, error) {
	loginQry, err := am.DB.Prepare("SELECT id FROM users WHERE username = ? AND password = ?")
	if err != nil {
		log.Println("prepare insert user ", err.Error())
		return Pegawai{}, errors.New("prepare statement insert user error")
	}

	row := loginQry.QueryRow(nama, password)

	if row.Err() != nil {
		log.Println("login query ", row.Err().Error())
		return Pegawai{}, errors.New("tidak bisa login, data tidak ditemukan")
	}
	res := Pegawai{}
	err = row.Scan(&res.ID)

	if err != nil {
		log.Println("after login query ", err.Error())
		return Pegawai{}, errors.New("tidak bisa login, kesalahan setelah error")
	}

	res.Username = nama

	return res, nil
}

// Function untuk menambah pegawai baru
func (am *AuthMenu) Register(tambahPegawai Pegawai) (bool, error) {
	registerQry, err := am.DB.Prepare("INSERT INTO pegawai (username, password) values (?,?)")
	if err != nil {
		log.Println("Prepare insert pegawai ", err.Error())
		return false, errors.New("")
	}

	if am.Duplicate(tambahPegawai.Username) {
		log.Println("duplicated information")
		return false, errors.New("name sudah digunakan")
	}

	res, err := registerQry.Exec(tambahPegawai.Username, tambahPegawai.Password)
	if err != nil {
		log.Println("insert pegawai ", err.Error())
		return false, errors.New("insert pegawai error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("After insert pegawai ", err.Error())
		return false, errors.New("Error setelah insert pegawai")
	}

	if affRows <= 0 {
		log.Println("No record affected")
		return false, errors.New("No record")
	}
	return true, nil
}

func (am *AuthMenu) Duplicate(username string) bool {
	res := am.DB.QueryRow("SELECT id FROM pegawai where username = ?", username)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil {
		log.Println("Result scan error", err.Error())
		return false
	}
	return true
}
