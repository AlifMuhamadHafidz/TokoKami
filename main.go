package main

import (
	"bufio"
	f "fmt"
	"os"
	"strconv"
	"toko/barang"
	"toko/config"
	"toko/pelanggan"
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
	var iniPelanggan = pelanggan.AuthPelanggan{DB: conn}

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
					f.Println("====================")
					f.Println("1. Register Pegawai")
					f.Println("2. List Pegawai")
					f.Println("3. Hapus Pegawai")
					f.Println("4. Hapus Barang")
					f.Println("5. Hapus Pelanggan")
					f.Println("6. Hapus Transaksi")
					f.Println("7. Hapus Nota")
					f.Println("9. Logout")
					f.Print("Masukan Pilihan : ")
					f.Scanln(&menuAdmin)

					switch menuAdmin {

					case 1:
						var newUser users.Pegawai
						f.Println("Halaman Tambah Pegawai")
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

					// List pegawai
					case 2:
						f.Println("<<-- List Pegawai -->>")
						f.Println("No. Nama")
						listPegawai := authMenu.ListPegawai()
						for i := 0; i < len(listPegawai); i++ {
							f.Print(" ", i+1, ". ")
							for j := 1; j < len(listPegawai[i]); j++ {
								f.Print(listPegawai[i][j])
							}
							f.Println()
						}

					// Hapus pegawai
					case 3:
						var inputNomor int
						f.Println("========================")
						f.Println("~~ Halaman Hapus Pegawai ~~")
						f.Print("Masukkan nomor pegawai yang akan dihapus: ")
						f.Scanln(&inputNomor)
						listPegawai := authMenu.ListPegawai()
						idDelete, _ := strconv.Atoi(listPegawai[inputNomor-1][0])
						err := authMenu.DeletePegawai(idDelete)
						if err != nil {
							f.Println("Data gagal dihapus", err.Error())
						} else {
							f.Println("Data berhasil dihapus")
						}

					case 4:
						f.Println("Hapus Pelanggan")

					case 5:
						f.Println("Hapus Transaksi")

					case 6:
						f.Println("Hapus Nota")

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
						f.Println("3. Edit Stok Barang")
						f.Println("8. Tambah Pelanggan")
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
							f.Print("Masukan Nama Barang : ")
							scanner.Scan()
							updateBarang.Nama = scanner.Text()
							f.Print("Masukan info Barang Terbaru : ")
							scanner.Scan()
							updateBarang.Info = scanner.Text()

							isUpdated, err := iniBarang.EditInfoBarang(updateBarang, updateBarang.Nama)
							if err != nil {
								f.Println(err.Error())
							}

							if isUpdated {
								f.Println("Berhasil Update Info Barang")
							}

						} else if inputMenuPegawai == 3 {
							var updateStokBarang int
							updateBarang := barang.Barang{}
							f.Print("Masukan Stok barang terbaru : ")
							f.Scanln(&updateStokBarang)
							f.Print("Masukan Nama Barang : ")
							scanner.Scan()
							updateBarang.Nama = scanner.Text()

							isStokUpdated, err := iniBarang.EditStokBarang(updateStokBarang, updateBarang.Nama)

							if err != nil {
								f.Println(err.Error())
							}

							if isStokUpdated {
								f.Println("Berhasil Update Stok Barang")
							}

						} else if inputMenuPegawai == 8 {
							insertPelanggan := pelanggan.Pelanggan{}
							f.Print("Nama Pelanggan: ")
							scanner.Scan()
							insertPelanggan.Nama = scanner.Text()

							_, err := iniPelanggan.TambahPelanggan(insertPelanggan)
							if err != nil {
								f.Println("Pelanggan Gagal Ditambahkan", err.Error())
							} else {
								f.Println("Pelanggan Berhasil Ditambahkan")
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
