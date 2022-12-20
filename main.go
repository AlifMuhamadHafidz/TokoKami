package main

import (
	"fmt"
	"toko/config"
	"toko/users"
)

func main() {

	var inputUsername string
	var inputPassword string
	var inputMenu int = 1
	var cfg = config.ReadConfig()
	var conn = config.DBConnection(*cfg)
	var authMenu = users.AuthMenu{DB: conn}

	for inputMenu != 0 {
		fmt.Println("1. Login")
		fmt.Println("0. Exit")
		fmt.Print("Masukan Pilihan : ")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			fmt.Print("Masukan Username : ")
			fmt.Scanln(&inputUsername)
			fmt.Print("Masukan Password : ")
			fmt.Scanln(&inputPassword)

			if inputUsername == "admin" && inputPassword == "admin" {
				fmt.Println("Welcome Admin")
				var isAdmin bool = true
				var menuAdmin int

				for isAdmin {
					fmt.Println("1. Register Pegawai")
					fmt.Println("9. Logout")
					fmt.Print("Masukan Pilihan : ")
					fmt.Scanln(&menuAdmin)
					switch menuAdmin {
					case 1:
						var newUser users.Pegawai
						fmt.Print("Masukkan nama : ")
						fmt.Scanln(&newUser.Username)
						fmt.Print("Masukkan password : ")
						fmt.Scanln(&newUser.Password)
						res, err := authMenu.Register(newUser)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Sukses mendaftarkan data")
						} else {
							fmt.Println("Gagal mendaftarn data")
						}
					case 9:
						isAdmin = !isAdmin
					}

				}
			} else {
				res, err := authMenu.Login(inputUsername, inputPassword)
				if err != nil {
					fmt.Println(err.Error())
				}

				if res.ID > 0 {
					isLogin := true
					inputMenuPegawai := 0
					for isLogin {
						fmt.Println("1. Tambah Barang")
						fmt.Println("9. Exit")
						fmt.Print("Masukan Pilihan : ")
						fmt.Scanln(&inputMenuPegawai)

						if inputMenuPegawai == 1 {
							fmt.Println("tambah Barang")
						} else if inputMenuPegawai == 9 {
							isLogin = false
						}
					}
				}
			}
		}
	}
}
