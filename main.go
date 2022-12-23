package main

import (
	"bufio"
	f "fmt"
	"os"
	"strconv"
	"toko/barang"
	"toko/config"
	"toko/nota"
	"toko/pelanggan"
	"toko/transaksi"
	"toko/users"
)

// test

func main() {

	var inputUsername string
	var inputPassword string
	var inputMenu int = 1
	var cfg = config.ReadConfig()
	var conn = config.DBConnection(*cfg)
	var authMenu = users.AuthMenu{DB: conn}
	var iniBarang = barang.MenuBarang{DB: conn}
	var iniPelanggan = pelanggan.AuthPelanggan{DB: conn}
	var iniTransaksi = transaksi.TransMenu{DB: conn}
	var iniNota = nota.NotaMenu{DB: conn}

	// membuat scan kalimat
	scanner := bufio.NewScanner(os.Stdin)

	for inputMenu != 0 {
		f.Println("\n======= TokoKami =======")
		f.Print(">> Halaman Utama\n\n")
		f.Println("1. Login")
		f.Println("0. Exit")
		f.Print("Masukan Pilihan: ")
		f.Scanln(&inputMenu)
		if inputMenu == 1 {
			f.Println("\n======= TokoKami =======")
			f.Print(">> Halaman Login\n\n")
			f.Print("Masukan Username : ")
			f.Scanln(&inputUsername)
			f.Print("Masukan Password : ")
			f.Scanln(&inputPassword)

			if inputUsername == "admin" && inputPassword == "admin" {
				f.Println("Welcome Admin")
				var isAdmin bool = true
				var menuAdmin int

				for isAdmin {
					f.Println("\n======= TokoKami for Admin =======")
					f.Print(">> Halaman Menu Admin\n\n")
					f.Println("1. Pegawai") //sudah
					f.Println("2. Barang")  //sudah
					f.Println("3. Pelanggan")
					f.Println("4. Show Transaksi")
					f.Println("5. Hapus Transaksi")
					f.Println("0. Logout") //sudah
					f.Print("Masukan Pilihan : ")
					f.Scanln(&menuAdmin)

					switch menuAdmin {
					// START MENUADMIN ==> PEGAWAI
					case 1:
						menuAdminPegawai := 1
						for menuAdminPegawai != 0 {
							f.Println("\n======= TokoKami for Admin =======")
							f.Print(">> Halaman Admin Menu Pegawai\n\n")
							f.Println("1. Register Pegawai") //sudah
							f.Println("2. List Pegawai")     //sudah
							// f.Println("3. Update Pegawai")
							f.Println("3. Delete Pegawai") //sudah
							f.Println("0. <<= Back")       //sudah
							f.Print("Masukan Pilihan: ")
							f.Scanln(&menuAdminPegawai)

							// variabel menampung hasil select data pegawai
							listPegawai := authMenu.ListPegawai()

							switch menuAdminPegawai {
							// Create Pegawai
							case 1:
								var newUser users.Pegawai
								f.Println("\n======= TokoKami for Admin =======")
								f.Print(">> Halaman Register Pegawai\n\n")
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

							// Read Pegawai
							case 2:
								f.Println("\n======= TokoKami for Admin =======")
								f.Print(">> Halaman Daftar Pegawai\n\n")
								f.Println("No | Nama")
								f.Println("_____________")
								// listPegawai := authMenu.ListPegawai()
								for i := 0; i < len(listPegawai); i++ {
									f.Print(" ", i+1, " | ")
									for j := 1; j < len(listPegawai[i]); j++ {
										f.Print(listPegawai[i][j])
									}
									f.Println()
								}
								f.Println("_____________")

							// Delete Pegawai
							case 3:
								var inputNomor int
								f.Println("\n======= TokoKami for Admin =======")
								f.Print(">> Halaman Hapus Pegawai\n\n")
								f.Print("Masukkan nomor pegawai yang akan dihapus: ")
								f.Scanln(&inputNomor)
								// listPegawai := authMenu.ListPegawai()
								idDelete, _ := strconv.Atoi(listPegawai[inputNomor-1][0])
								err := authMenu.DeletePegawai(idDelete)
								if err != nil {
									f.Println("Data gagal dihapus", err.Error())
								} else {
									f.Println("Data berhasil dihapus")
								}
							}
						}
					// END MENUADMIN ==> PEGAWAI

					// START MENUADMIN ==> BARANG
					case 2:
						menuAdminBarang := 1
						for menuAdminBarang != 0 {
							f.Println("\n======= TokoKami for Admin =======")
							f.Print(">> Halaman Admin Menu Barang\n\n")
							// f.Println("1. Create Barang")
							f.Println("1. Show List Barang") //sudah
							// f.Println("3. Update Barang")
							f.Println("2. Delete Barang") //sudah
							f.Println("0. <<= Back")      //sudah
							f.Print("Masukan Pilihan: ")
							f.Scanln(&menuAdminBarang)

							// variabel untuk menyimpan hasil select daftar barang
							listBarang := iniBarang.ListBarang()

							switch menuAdminBarang {

							// Read Barang
							case 1:
								f.Println("\n======= TokoKami for Admin =======")
								f.Print(">> Halaman Daftar Barang\n\n")
								f.Println("No |   Nama  | Stok")
								f.Println("______________________")
								// listBarang := iniBarang.ListBarang()
								for i := 0; i < len(listBarang); i++ {
									f.Print(" ", i+1, " | ")
									for j := 1; j < len(listBarang[i]); j++ {
										f.Print(listBarang[i][j], " | ")
									}
									f.Println()
								}
								f.Println("______________________")

							// Delete Barang
							case 2:
								var inputNomor int
								f.Println("\n======= TokoKami for Admin =======")
								f.Print(">> Halaman Hapus Barang\n\n")
								f.Print("Masukkan nomor barang yang akan dihapus: ")
								f.Scanln(&inputNomor)
								// listBarang := iniBarang.ListBarang()
								idDelete, _ := strconv.Atoi(listBarang[inputNomor-1][0])
								err := iniBarang.DeleteBarang(idDelete)
								if err != nil {
									f.Println("Barang gagal dihapus", err.Error())
								} else {
									f.Println("Barang berhasil dihapus")
								}
							}
						}

						// END MENUADMIN ==> BARANG

					// START MENUADMIN ==> PELANGGAN

					// 22.20
					case 3:
						menuAdminPelanggan := 1
						for menuAdminPelanggan != 0 {
							f.Println("\n======= TokoKami for Admin =======")
							f.Print(">> Halaman Admin Menu Pelanggan\n\n")
							// f.Println("1. Create Pelanggan")
							f.Println("1. List Pelanggan") //sudah
							// f.Println("3. Update Pelanggan")
							f.Println("2. Delete Pelanggan") //sudah
							f.Println("0. <<= Back")         //sudah
							f.Print("Masukan Pilihan: ")
							f.Scanln(&menuAdminPelanggan)

							// variabel untuk menyimpan hasil select daftar pelanggan
							listPelanggan := iniPelanggan.ListPelanggan()

							switch menuAdminPelanggan {

							// Read Pelanggan
							case 1:
								f.Println("\n======= TokoKami for Admin =======")
								f.Print(">> Halaman Daftar Pelanggan\n\n")
								f.Println("No | Nama")
								f.Println("______________________")
								for i := 0; i < len(listPelanggan); i++ {
									f.Print(" ", i+1, " | ")
									for j := 1; j < len(listPelanggan[i]); j++ {
										f.Print(listPelanggan[i][j])
									}
									f.Println()
								}
								f.Println("______________________")

							// Delete Pelanggan
							case 2:
								var inputNomor int
								f.Println("\n======= TokoKami for Admin =======")
								f.Print(">> Halaman Hapus Pelanggan\n\n")
								f.Print("Masukkan nomor pelanggan yang akan dihapus: ")
								f.Scanln(&inputNomor)
								idDelete, _ := strconv.Atoi(listPelanggan[inputNomor-1][0])
								err := iniPelanggan.DeletePelanggan(idDelete)
								if err != nil {
									f.Println("Gagal menghapus pelanggan", err.Error())
								} else {
									f.Println("Pelanggan berhasil dihapus")
								}
							}
						}
						// END MENUADMIN ==> PELANGGAN

						// START MENUADMIN ==> TRANSAKSI
					case 4:
						f.Println("\n======= TokoKami for Admin =======")
						f.Print(">> Halaman Admin Show Transaksi\n\n")
						var showTransaksi transaksi.Transaksi

						res, err := iniTransaksi.ShowTransaksi(showTransaksi)

						if err != nil {
							f.Println(err.Error())
						}

						if res {
							f.Println("Sukses Show Transaksi")
						} else {
							f.Println("Gagal Show Transaksi")
						}

					case 5:
						f.Println("\n======= TokoKami for Admin =======")
						f.Print(">> Halaman Admin Hapus Transaksi\n\n")
						var deleteTransaksi transaksi.Transaksi
						f.Print("Masukan id Transaksi yang ingin dihapus: ")
						f.Scanln(&deleteTransaksi.ID)

						res, err := iniTransaksi.DeleteTransaksi(deleteTransaksi)

						if err != nil {
							f.Println(err.Error())
						}

						if res {
							f.Println("Sukses menghapus Transaksi")
						} else {
							f.Println("Gagal menghapus Transaksi")
						}

					// START MENUADMIN ==> NOTA

					// END MENUADMIN ==> NOTA

					// MENUADMIN LOGOUT
					case 0:
						isAdmin = !isAdmin
					}
				}
				// END MENUADMIN

				// START MENUPEGAWAI
			} else {
				res, err := authMenu.Login(inputUsername, inputPassword)
				if err != nil {
					f.Println(err.Error())
				}

				if res.ID > 0 {
					isLogin := true
					inputMenuPegawai := 0
					for isLogin {
						f.Println("\n======= TokoKami for Pegawai =======")
						f.Print(">> Halaman Menu Pegawai\n\n")
						f.Println("1. Tambah Barang")
						f.Println("2. Edit Info Barang")
						f.Println("3. Edit Stok Barang")
						f.Println("4. Tambah Transaksi")
						f.Println("5. Tambah Pelanggan")
						f.Println("6. Show List Barang")
						f.Println("9. Exit")
						f.Print("Masukan Pilihan : ")
						f.Scanln(&inputMenuPegawai)

						if inputMenuPegawai == 1 {
							f.Println("\n======= TokoKami for Pegawai =======")
							f.Print(">> Halaman Tambah Barang\n\n")
							insertBarang := barang.Barang{}
							f.Print("Nama Barang: ")
							scanner.Scan()
							insertBarang.Nama = scanner.Text()
							f.Print("Info Barang: ")
							scanner.Scan()
							insertBarang.Info = scanner.Text()
							f.Print("Jumlah barang: ")
							f.Scanln(&insertBarang.Stok)
							insertBarang.ID_Pegawai = res.ID

							_, err := iniBarang.TambahBarang(insertBarang)
							if err != nil {
								f.Println("Barang Gagal Ditambahkan", err.Error())
							}
							f.Println("Barang Berhasil Ditambahkan")

						} else if inputMenuPegawai == 2 {
							updateBarang := barang.Barang{}
							f.Println("\n======= TokoKami for Pegawai =======")
							f.Print(">> Halaman Update Info Barang\n\n")
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
							f.Println("\n======= TokoKami for Pegawai =======")
							f.Print(">> Halaman Update Stok Barang\n\n")
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

						} else if inputMenuPegawai == 4 {

							newTransaksi := transaksi.Transaksi{}
							f.Println("\n======= TokoKami =======")
							f.Print(">> Halaman Tambah Transaksi\n\n")
							f.Print("Masukan id pelanggan: ")
							f.Scanln(&newTransaksi.ID_Pelanggan)
							newTransaksi.ID_Pegawai = res.ID

							_, err := iniTransaksi.AddTransaksi(newTransaksi)

							if err != nil {
								f.Println("Gagal Menambahkan Transaksi", err.Error())
							} else {
								f.Println("Sukses Menambahkan Transaki")
							}

							inputNota := 1
							for inputNota != 0 {
								f.Println("\n======= TokoKami =======")
								f.Println(">> Halaman Daftar Transaksi")
								f.Print("____________________________\n\n")
								listTranNota := iniNota.ListTranNota()
								for i := 0; i < 1; i++ {
									f.Println("No Trans : ", listTranNota[0])
									f.Println("Tanggal  : ", listTranNota[1])
									f.Println("Pegawai  : ", listTranNota[2])
									f.Println("Pembeli  : ", listTranNota[3])
								}
								f.Println("____________________________")
								id, _ := strconv.Atoi(listTranNota[0])
								listNota := iniNota.ListNota(id)
								f.Println("No  |  Barang  |  Jumlah")
								f.Println("-------------------")
								for i := 0; i < len(listNota); i++ {
									f.Print(i+1, " | ")
									for j := 0; j < len(listNota[0]); j++ {
										f.Print(listNota[i][j], " | ")
									}
									f.Println()
								}
								f.Println("____________________________")

								insertNota := nota.Nota{}
								f.Println("======================")
								f.Print("Masukan id barang: ")
								f.Scanln(&insertNota.ID_Barang)
								f.Print("Masukan jumlah: ")
								f.Scanln(&insertNota.Jumlah)
								insertNota.ID_Trans, _ = strconv.Atoi(listTranNota[0])

								//
								stokBarang := iniNota.SelectStok(insertNota.ID_Barang)
								stok, _ := strconv.Atoi(stokBarang[1])

								if stok-insertNota.Jumlah < 0 {
									f.Println("Maaf stok hanya tersisa: ", stok)
									continue
								}
								_, err := iniNota.AddNota(insertNota)
								sisaStok := stok - insertNota.Jumlah
								isStokUpdated, err := iniBarang.EditStokBarang(sisaStok, stokBarang[0])
								if err != nil {
									f.Println(err.Error())
								}
								if isStokUpdated {
								}
								//

								if err != nil {
									f.Println("Nota Gagal Ditambahkan", err.Error())
								}
								f.Println("Nota Berhasil Ditambahkan")

								f.Println("\n======================")
								f.Println("1. Belanja Lagi")
								f.Println("0. Checkout & Cetak Nota")
								f.Print("Tentukan pilihanmu: ")
								f.Scanln(&inputNota)

								// Selesai transaksi dan cetak nota
								if inputNota == 0 {
									f.Println("\n======= TokoKami =======")
									f.Println(">> Halaman Nota Oke")
									f.Print("____________________________\n\n")
									listTransNota := iniNota.ListTranNota()
									for i := 0; i < 1; i++ {
										f.Println("No Trans : ", listTransNota[0])
										f.Println("Tanggal  : ", listTransNota[1])
										f.Println("Pegawai  : ", listTransNota[2])
										f.Println("Pembeli  : ", listTransNota[3])
									}
									f.Println("____________________________")
									id, _ := strconv.Atoi(listTransNota[0])
									listNotaOke := iniNota.ListNota(id)
									f.Println("No  |  Barang  |  Jumlah")
									f.Println("-------------------")
									for i := 0; i < len(listNotaOke); i++ {
										f.Print(i+1, " | ")
										for j := 0; j < len(listNotaOke[0]); j++ {
											f.Print(listNotaOke[i][j], " | ")
										}
										f.Println()
									}
									f.Println("____________________________")
								}
							}

						} else if inputMenuPegawai == 5 {
							f.Println("\n======= TokoKami for Pegawai =======")
							f.Print(">> Halaman Register Pelanggan\n\n")
							insertPelanggan := pelanggan.Pelanggan{}
							f.Print("Nama Pelanggan: ")
							scanner.Scan()
							insertPelanggan.Nama = scanner.Text()
							insertPelanggan.ID_Pegawai = res.ID

							_, err := iniPelanggan.TambahPelanggan(insertPelanggan)
							if err != nil {
								f.Println("Pelanggan Gagal Ditambahkan", err.Error())
							} else {
								f.Println("Pelanggan Berhasil Ditambahkan")
							}

						} else if inputMenuPegawai == 6 {
							listBarang := iniBarang.ListBarang()
							f.Print(">> Halaman Daftar Barang\n\n")
							f.Println("No |   Nama  | Stok")
							f.Println("______________________")
							// listBarang := iniBarang.ListBarang()
							for i := 0; i < len(listBarang); i++ {
								f.Print(" ", i+1, " | ")
								for j := 1; j < len(listBarang[i]); j++ {
									f.Print(listBarang[i][j], " | ")
								}
								f.Println()
							}
							f.Println("______________________")

						} else if inputMenuPegawai == 9 {
							isLogin = false
						}

					}
				}
			}
		}
	}
}
