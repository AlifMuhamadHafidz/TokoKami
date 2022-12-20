package main

import "fmt"

func main() {

	var inputUsername string
	var inputPassword string
	var inputMenu int = 1

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
						fmt.Println("regis Karyawan")
					case 9:
						isAdmin = !isAdmin
					}
				}

			} else {
				fmt.Println("Bukan Admin")
				break
			}
		}
	}

}
