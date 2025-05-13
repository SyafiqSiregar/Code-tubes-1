// Aplikasi Manajemen Stok Bahan Makanan

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const NMAX int = 1024

type BahanMakanan struct {
	Nama  string
	Stok  int
	Harga float64
}
type User struct {
	Username string
	Password string
}

type arrBahan [NMAX]BahanMakanan

var users []User
var username, password string
var pilihan int
var dataBahan arrBahan
var dataCount int
var totalPengeluaran float64

func main() {
	start()
}

func start() {
	clearScreen()
	fmt.Println("  ╔══════════════════════════════════════════════════════╗")
	fmt.Println("  ║   __        _______ _     ____ ___  __  __ _____     ║")
	fmt.Println("  ║   \\ \\      / / ____| |   / ___/ _ \\|  \\/  | ____|    ║")
	fmt.Println("  ║    \\ \\ /\\ / /|  _| | |  | |  | | | | |\\/| |  _|      ║")
	fmt.Println("  ║     \\ V  V / | |___| |__| |__| |_| | |  | | |___     ║")
	fmt.Println("  ║      \\_/\\_/  |_____|_____\\____\\___/|_|  |_|_____|    ║")
	fmt.Println("  ╚══════════════════════════════════════════════════════╝")

	fmt.Println("      ╔═══════════════════════════════════════════╗")
	fmt.Println("      ║          🚀 Aplikasi Manajemen            ║")
	fmt.Println("      ║    Created by Syafiq Yusuf Ikhsan &       ║")
	fmt.Println("      ║             Dafa Izul Haq                 ║")
	fmt.Println("      ║      Algoritma Pemrograman 2025           ║")
	fmt.Println("      ╠═══════════════════════════════════════════╣")
	fmt.Println("      ║   🌟 APLIKASI MANAGEMENT STOK BAHAN 🌟    ║")
	fmt.Println("      ╠═══════════════════════════════════════════╣")
	fmt.Println("      ║                1. Register                ║")
	fmt.Println("      ║                2. Login                   ║")
	fmt.Println("      ║                0. Keluar                  ║")
	fmt.Println("      ╚═══════════════════════════════════════════╝")
	fmt.Print("                   👉 Pilihan Anda : ")

	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		register()
	case 2:
		login()
	case 0:
		fmt.Println("   ╔═══════════════════════════════════════════════════╗")
		fmt.Println("   ║ 🙏 Terima kasih telah menggunakan aplikasi ini!!  ║")
		fmt.Println("   ╚═══════════════════════════════════════════════════╝")
	default:
		fmt.Println("\n      ❌ Masukan tidak valid, silahkan coba lagi.")
		start()
	}
}

func daftarLayanan() {
	/*
		IS: Pengguna diberikan pilihan menu layanan yang tersedia.
		FS: Menjalankan fungsi sesuai pilihan pengguna (lihat stok, tambah bahan, kurangi stok, dll.).
	*/
	fmt.Println("        ╔════════════════════════════════════════╗")
	fmt.Println("        ║             🚀 Main Menu               ║")
	fmt.Println("        ╠════════════════════════════════════════╣")
	fmt.Println("        ║  1. Lihat Stok Bahan                   ║")
	fmt.Println("        ║  2. Tambah Bahan                       ║")
	fmt.Println("        ║  3. Kurangi Stok                       ║")
	fmt.Println("        ║  4. Lihat Bahan Hampir Habis           ║")
	fmt.Println("        ║  5. Cari Bahan                         ║")
	fmt.Println("        ║  6. Lihat Total Pengeluaran            ║")
	fmt.Println("        ║  0. Keluar                             ║")
	fmt.Println("        ╚════════════════════════════════════════╝")
	fmt.Print("                   👉 Pilihan Anda : ")
	fmt.Scan(&pilihan)
	clearScreen()

	switch pilihan {
	case 0:
		fmt.Println("  ╔═══════════════════════════════════════════════════╗")
		fmt.Println("  ║ 🙏 Terima kasih telah menggunakan aplikasi ini!!  ║")
		fmt.Println("  ╚═══════════════════════════════════════════════════╝")
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
	case 6:
		lihatPengeluaranBulanan()
	default:
		fmt.Println("\nMasukan tidak valid, silahkan coba lagi")
		daftarLayanan()
	}
}
func register() {
	/*
		IS: Daftar pengguna (users) terdefinisi.
		FS: Menambahkan akun baru ke daftar pengguna jika username belum digunakan.
	*/
	clearScreen()
	var username, password string
	fmt.Println("              ╔═════════════════════════════╗")
	fmt.Println("              ║      🚀 Register Akun       ║")
	fmt.Println("              ╚═════════════════════════════╝")
	fmt.Print("                   👤 Username : ")
	fmt.Scan(&username)
	fmt.Print("                   🔑 Password : ")
	fmt.Scan(&password)
	for _, user := range users {
		if user.Username == username {
			fmt.Println("\n                 Username sudah digunakan!")
			start()
			return
		}
	}
	users = append(users, User{Username: username, Password: password})
	fmt.Println("\n              Akun berhasil didaftarkan!")
	start()
}
func login() {
	/*
		IS: Daftar pengguna (users) terdefinisi.
		FS: Mengakses menu layanan jika username dan password sesuai, atau menampilkan pesan kesalahan.
	*/
	clearScreen()
	var inputUsername, inputPassword string
	fmt.Println("              ╔═════════════════════════════╗")
	fmt.Println("              ║        🚀 Login Akun        ║")
	fmt.Println("              ╚═════════════════════════════╝")
	fmt.Print("                   👤 Username : ")
	fmt.Scan(&inputUsername)
	fmt.Print("                   🔑 Password : ")

	fmt.Scan(&inputPassword)
	clearScreen()

	for _, user := range users {
		if user.Username == inputUsername && user.Password == inputPassword {
			fmt.Println("\n              ======= WELCOME ", user.Username, " =======")
			daftarLayanan()
			return
		}
	}

	fmt.Println("\n           ❌ Username atau password salah!")
	start()
}

func lihatStok() {
	/*
		IS: Data bahan makanan tersedia.
		FS: Daftar bahan makanan ditampilkan dalam urutan stok terendah ke tertinggi(selection sort).
	*/
	var i, j, minIdx int
	var tempBahan BahanMakanan

	// Selection Sort untuk mengurutkan stok bahan
	for i = 0; i < dataCount-1; i++ {
		minIdx = i
		for j = i + 1; j < dataCount; j++ {
			if dataBahan[j].Stok < dataBahan[minIdx].Stok {
				minIdx = j
			}
		}
		// Swap dataBahan[i] dengan dataBahan[minIdx]
		tempBahan = dataBahan[i]
		dataBahan[i] = dataBahan[minIdx]
		dataBahan[minIdx] = tempBahan
	}
	// Menampilkan daftar stok bahan yang sudah diurutkan
	fmt.Println("              ╔═════════════════════════════╗")
	fmt.Println("              ║      Daftar Stok Bahan :    ║")
	fmt.Println("              ╚═════════════════════════════╝")
	for i = 0; i < dataCount; i++ {
		fmt.Printf("                        %d . %s = %d\n", i+1, dataBahan[i].Nama, dataBahan[i].Stok)
	}

	daftarLayanan()
}
func cariBahan() {
	/*
		IS: dataBahan terdefinisi dengan dataCount sebagai jumlahnya.
		FS: Menampilkan bahan yang sesuai dengan nama yang dicari (jika ditemukan).
	*/
	var nama string
	var i int
	fmt.Println("              ╔═════════════════════════════╗")
	fmt.Println("              ║         Cari Bahan :        ║")
	fmt.Println("              ╚═════════════════════════════╝")
	fmt.Print("                      Nama Bahan : ")
	fmt.Scan(&nama)

	for i = 0; i < dataCount; i++ {
		if strings.Contains(strings.ToLower(dataBahan[i].Nama), strings.ToLower(nama)) {
			fmt.Printf("                      - %s     : %d\n", dataBahan[i].Nama, dataBahan[i].Stok)
			fmt.Println("\n               Kembali ke daftar layanan...\n")
			daftarLayanan()
			return
		}
	}

	fmt.Println("\n                   Bahan tidak ditemukan.")
	daftarLayanan()
}

func tambahBahan() {
	/*
		IS: totalPengeluaran bertipe float64 dan telah terdefinisi sebagai total pengeluaran yang tercatat.
		FS: Menampilkan total pengeluaran bulanan ke layar.
	*/
	var nama string
	var jumlah, i int
	var harga float64

	fmt.Println("              ╔═════════════════════════════╗")
	fmt.Println("              ║        Tambah Bahan :       ║")
	fmt.Println("              ╚═════════════════════════════╝")
	fmt.Print("                     Nama Bahan : ")
	fmt.Scan(&nama)
	fmt.Print("                     Jumlah     : ")
	fmt.Scan(&jumlah)
	fmt.Print("                     Harga per Satuan: Rp ")
	fmt.Scan(&harga)

	for i = 0; i < dataCount; i++ {
		if strings.EqualFold(dataBahan[i].Nama, nama) {
			dataBahan[i].Stok += jumlah
			dataBahan[i].Harga = harga
			totalPengeluaran += float64(jumlah) * harga
			fmt.Println("                     Stok diperbarui.")
			daftarLayanan()
			return
		}
	}

	dataBahan[dataCount] = BahanMakanan{Nama: nama, Stok: jumlah, Harga: harga}
	dataCount++
	totalPengeluaran += float64(jumlah) * harga
	clearScreen()
	fmt.Println("                     Bahan ditambahkan.")
	daftarLayanan()

}

func kurangiStok() {
	/*
		IS: dataBahan terdefinisi dengan dataCount sebagai jumlahnya.
		FS: Mengurangi stok bahan sesuai input jika tersedia.
	*/
	var nama string
	var jumlah, i int
	fmt.Println("              ╔═════════════════════════════╗")
	fmt.Println("              ║        Kurangi Bahan :      ║")
	fmt.Println("              ╚═════════════════════════════╝")
	fmt.Print("                     Nama Bahan : ")
	fmt.Scan(&nama)
	fmt.Print("                     Jumlah     : ")
	fmt.Scan(&jumlah)

	for i = 0; i < dataCount; i++ {
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
	/*
		IS: dataBahan terdefinisi dengan dataCount sebagai jumlahnya.
		FS: Menampilkan bahan dengan stok < 3.
	*/
	var i int
	fmt.Println("        ╔════════════════════════════════════════╗")
	fmt.Println("        ║      Bahan Hampir Habis (Stok <= 5)    ║")
	fmt.Println("        ╚════════════════════════════════════════╝")

	for i = 0; i < dataCount; i++ {
		if dataBahan[i].Stok <= 5 {
			fmt.Printf("                      - %s      (%d)\n", dataBahan[i].Nama, dataBahan[i].Stok)
		}
	}
	daftarLayanan()
}

func lihatPengeluaranBulanan() {
	/*
		IS: totalPengeluaran terdefinisi.
		FS: menampilkan Total pengeluaran bulanan.
	*/
	fmt.Println("        ╔════════════════════════════════════════╗")
	fmt.Println("        ║        Total Pengeluaran Bulanan       ║")
	fmt.Println("        ╚════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("                     Total: Rp %.2f\n", totalPengeluaran)
	daftarLayanan()
}
func clearScreen() {
	/* IS: -
	   FS: Mengosongkan layar.
	*/
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd = exec.Command("clear")
	} else {
		fmt.Println("Platform tidak didukung.")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
