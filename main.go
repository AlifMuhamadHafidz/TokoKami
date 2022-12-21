package main

import (
	"bufio"
	"fmt"
	f "fmt"
	"os"
	"toko/barang"
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
	var iniBarang = barang.MenuBarang{DB: conn}

	// membuat scan kalimat
	scanner := bufio.NewScanner(os.Stdin)

	for inputMenu != 0 {
		f.Println("1. Login")
		f.Println("0. Exit")
		f.Print("Masukan Pilihan : ")
		f.Scanln(&inputMenu)
		if inputMenu == 1 {
			f.Print("Masukan Username : ")
			f.Scanln(&inputUsername)
			f.Print("Masukan Password : ")
			f.Scanln(&inputPassword)

			if inputUsername == "admin" && inputPassword == "admin" {
				f.Println("Welcome Admin")
				var isAdmin bool = true
				var menuAdmin int

				for isAdmin {
					f.Println("1. Register Pegawai")
					f.Println("9. Logout")
					f.Print("Masukan Pilihan : ")
					f.Scanln(&menuAdmin)
					switch menuAdmin {
					case 1:
						var newUser users.Pegawai
						f.Print("Masukkan nama : ")
						f.Scanln(&newUser.Username)
						f.Print("Masukkan password : ")
						f.Scanln(&newUser.Password)
						res, err := authMenu.Register(newUser)
						if err != nil {
							f.Println(err.Error())
						}
						if res {
							f.Println("Sukses mendaftarkan data")
						} else {
							f.Println("Gagal mendaftarn data")
						}
					case 9:
						isAdmin = !isAdmin
					}

				}
			} else {
				res, err := authMenu.Login(inputUsername, inputPassword)
				if err != nil {
					f.Println(err.Error())
				}

				if res.ID > 0 {
					isLogin := true
					inputMenuPegawai := 0
					for isLogin {
						f.Println("1. Tambah Barang")
						f.Println("2. Edit Info Barang")
						f.Println("9. Exit")
						f.Print("Masukan Pilihan : ")
						f.Scanln(&inputMenuPegawai)

						if inputMenuPegawai == 1 {
							insertBarang := barang.Barang{}
							f.Print("Nama Barang: ")
							scanner.Scan()
							insertBarang.Nama = scanner.Text()
							f.Print("Deskripsi barang: ")
							scanner.Scan()
							insertBarang.Info = scanner.Text()
							f.Print("Jumlah barang: ")
							f.Scanln(&insertBarang.Stok)

							_, err := iniBarang.TambahBarang(insertBarang)
							if err != nil {
								f.Println("Barang Gagal Ditambahkan", err.Error())
							}
							f.Println("Barang Berhasil Ditambahkan")

						} else if inputMenuPegawai == 2 {
							updateBarang := barang.Barang{}
							fmt.Print("Masukan Nama Barang : ")
							scanner.Scan()
							updateBarang.Nama = scanner.Text()
							fmt.Print("Masukan info Barang Terbaru : ")
							scanner.Scan()
							updateBarang.Info = scanner.Text()

							isUpdated, err := iniBarang.EditInfoBarang(updateBarang, updateBarang.Nama)
							if err != nil {
								fmt.Println(err.Error())
							}

							if isUpdated {
								fmt.Println("Berhasil Update Info Barang")
							}

						} else if inputMenuPegawai == 9 {
							isLogin = false
						}
					}
				}
			}
		}
	}
}
