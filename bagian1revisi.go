// Aplikasi Manajemen Stok Bahan Makanan 

package main

import (
	"fmt"
	"strings"
)

const NMAX int = 1024

type BahanMakanan struct {
	Nama string
	Stok int
}

type arrBahan [NMAX]BahanMakanan

var username, password string
var pilihan int
var dataBahan arrBahan
var dataCount int

func main() {
	start()
}

func start() {
	fmt.Println("  __        _______ _     ____ ___  __  __ _____   ")
	fmt.Println("  \\ \\      / / ____| |   / ___/ _ \\|  \\/  | ____| ")
	fmt.Println("   \\ \\ /\\ / /|  _| | |  | |  | | | | |\\/| |  _|    ")
	fmt.Println("    \\ V  V / | |___| |__| |__| |_| | |  | | |___     ")
	fmt.Println("     \\_/\\_/  |_____|_____\\____\\___/|_|  |_|_____|  ")
	fmt.Println()

	fmt.Println("       🌟  WELCOME TO STOK BAHAN MANAGEMENT  🌟     ")
	fmt.Println()
	fmt.Println("                     Login Account")
	fmt.Println()
	fmt.Print("                    👤 Username : ")
	fmt.Scan(&username)
	fmt.Print("                    🔑 Password : ")
	fmt.Scan(&password)

	fmt.Println("\n         =======       WELCOME ", username, "      =======")
	daftarLayanan()
}

func daftarLayanan() {
	fmt.Println("\n                     🚀 Main Menu            ")
	fmt.Println("              =============================    ")
	fmt.Println("                   1️. Lihat Stok Bahan          ")
	fmt.Println("                   2️. Tambah Bahan              ")
	fmt.Println("                   3️. Kurangi Stok                    ")
	fmt.Println("                   4️. Lihat Bahan Hampir Habis  ")
	fmt.Println("                   5️. Cari Bahan                ")
	fmt.Println("                   0️. Keluar                    ")
	fmt.Println("              =============================    ")
	fmt.Print("                   👉 Pilihan Anda : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 0:
		fmt.Println("\n      🙏 Terima kasih telah menggunakan aplikasi ini!!")
	case 1:
		lihatStok()
	case 2:
		tambahBahan()
	case 3:
		kurangiStok()
	case 4:
		lihatHampirHabis()
	case 5:
		cariBahan()
	default:
		fmt.Println("\nMasukan tidak valid, silahkan coba lagi")
		daftarLayanan()
	}
}

func lihatStok() {
	fmt.Println("\n              ===== Daftar Stok Bahan =====      ")
	for i := 0; i < dataCount; i++ {
		fmt.Printf("                   %d . %s = %d\n", i+1, dataBahan[i].Nama, dataBahan[i].Stok)
	}
	daftarLayanan()
}

func cariBahan() {
	var nama string
	fmt.Println("\n                    ==== Cari Bahan ====      ")
	fmt.Print("                      Nama Bahan : ")
	fmt.Scan(&nama)

	for i := 0; i < dataCount; i++ {
		if strings.Contains(strings.ToLower(dataBahan[i].Nama), strings.ToLower(nama)) {
			fmt.Printf("                      - %s     : %d\n", dataBahan[i].Nama, dataBahan[i].Stok)
			fmt.Println("\n               Kembali ke daftar layanan...\n")
			daftarLayanan()
			return
		}
	}

	fmt.Println("\n                      Bahan tidak ditemukan.")
	daftarLayanan()
}

func tambahBahan() {
	var nama string
	var jumlah int
	fmt.Println("\n                ===== Tambah Bahan =====      ")
	fmt.Print("                     Nama Bahan : ")
	fmt.Scan(&nama)
	fmt.Print("                     Jumlah     : ")
	fmt.Scan(&jumlah)

	for i := 0; i < dataCount; i++ {
		if strings.EqualFold(dataBahan[i].Nama, nama) {
			dataBahan[i].Stok += jumlah
			fmt.Println("                     Stok diperbarui.")
			daftarLayanan()
			return
		}
	}

	dataBahan[dataCount] = BahanMakanan{Nama: nama, Stok: jumlah}
	dataCount++
	fmt.Println()
	fmt.Println("                     Bahan ditambahkan.")
	daftarLayanan()
}

func kurangiStok() {
	var nama string
	var jumlah int
	fmt.Println("\n                ===== Kurangi Bahan =====      ")
	fmt.Print("                     Nama Bahan : ")
	fmt.Scan(&nama)
	fmt.Print("                     Jumlah     : ")
	fmt.Scan(&jumlah)

	for i := 0; i < dataCount; i++ {
		if strings.EqualFold(dataBahan[i].Nama, nama) {
			if dataBahan[i].Stok >= jumlah {
				dataBahan[i].Stok -= jumlah
				fmt.Println()
				fmt.Println("                     Stok dikurangi.")
			} else {
				fmt.Println()
				fmt.Println("                     Stok tidak cukup!")
			}
			daftarLayanan()
			return
		}
	}
	fmt.Println()
	fmt.Println("                     Bahan tidak ditemukan.")
	daftarLayanan()
}

func lihatHampirHabis() {
	fmt.Println("\n           === Bahan Hampir Habis (Stok < 3) ===   ")
	for i := 0; i < dataCount; i++ {
		if dataBahan[i].Stok < 3 {
			fmt.Printf("                      - %s      (%d)\n", dataBahan[i].Nama, dataBahan[i].Stok)
		}
	}
	daftarLayanan()
}
