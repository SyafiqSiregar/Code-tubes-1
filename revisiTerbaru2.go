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
	fmt.Println("  __        _______ _     ____ ___  __  __ _____   ")
	fmt.Println("  \\ \\      / / ____| |   / ___/ _ \\|  \\/  | ____| ")
	fmt.Println("   \\ \\ /\\ / /|  _| | |  | |  | | | | |\\/| |  _|    ")
	fmt.Println("    \\ V  V / | |___| |__| |__| |_| | |  | | |___     ")
	fmt.Println("     \\_/\\_/  |_____|_____\\____\\___/|_|  |_|_____|  ")
	fmt.Println()

	fmt.Println("       🌟  WELCOME TO STOK BAHAN MANAGEMENT  🌟     ")
	fmt.Println()
	fmt.Println("                    1️. Register")
	fmt.Println("                    2️. Login")
	fmt.Println("                    0️. Keluar")
	fmt.Print("\n                 👉 Pilihan Anda : ")
	fmt.Scan(&pilihan)
	
	
	switch pilihan {
    case 1:
        register()
    case 2:
        login()
    case 0:
        fmt.Println("\n      🙏 Terima kasih telah menggunakan aplikasi ini!!")
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
	fmt.Println("\n                     🚀 Main Menu            ")
	fmt.Println("              =============================    ")
	fmt.Println("                   1️. Lihat Stok Bahan          ")
	fmt.Println("                   2️. Tambah Bahan              ")
	fmt.Println("                   3️. Kurangi Stok                    ")
	fmt.Println("                   4️. Lihat Bahan Hampir Habis  ")
	fmt.Println("                   5️. Cari Bahan                ")
	fmt.Println("                   6️. Lihat total Pengeluaran   ")
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
    var username, password string
    fmt.Println("\n               ===== Register Akun =====")
	fmt.Println()
    fmt.Print("                  👤 Username : ")
    fmt.Scan(&username)
    fmt.Print("                  🔑 Password : ")
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
    var inputUsername, inputPassword string
    fmt.Println("\n               ======= Login Akun =======")
	fmt.Println()
    fmt.Print("                    👤 Username : ")
    fmt.Scan(&inputUsername)
    fmt.Print("                    🔑 Password : ")
    fmt.Scan(&inputPassword)

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
	fmt.Println("\n              ===== Daftar Stok Bahan =====      ")
	for i = 0; i < dataCount; i++ {
		fmt.Printf("                   %d . %s = %d\n", i+1, dataBahan[i].Nama, dataBahan[i].Stok)
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
	fmt.Println("\n                    ==== Cari Bahan ====      ")
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

	fmt.Println("\n                      Bahan tidak ditemukan.")
	daftarLayanan()
}

func tambahBahan() {
	/*
		IS: totalPengeluaran bertipe float64 dan telah terdefinisi sebagai total pengeluaran yang tercatat.
		FS: Menampilkan total pengeluaran bulanan ke layar.
*/
	var nama string
	var jumlah,i int
	var harga float64
	
	fmt.Println("\n                ===== Tambah Bahan =====      ")
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
	fmt.Println()
	fmt.Println("                     Bahan ditambahkan.")
	daftarLayanan()
}

func kurangiStok() {
	/*
	IS: dataBahan terdefinisi dengan dataCount sebagai jumlahnya.
	FS: Mengurangi stok bahan sesuai input jika tersedia.
*/
	var nama string
	var jumlah,i int
	fmt.Println("\n                ===== Kurangi Bahan =====      ")
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
	fmt.Println("\n           === Bahan Hampir Habis (Stok < 3) ===   ")
	for i = 0; i < dataCount; i++ {
		if dataBahan[i].Stok < 3 {
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
	fmt.Println("\n          ===== Total Pengeluaran Bulanan =====")
	fmt.Println()
	fmt.Printf("                     Total: Rp %.2f\n", totalPengeluaran)
	daftarLayanan()
}
