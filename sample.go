package main

import "fmt"

// ----------------------------------
// STRUKTUR DATA (STRUCT)
// ----------------------------------

type Lapangan struct {
	ID int
	Nama string
	HargaSewa float64
}

type Penyewa struct {
	ID int
	Nama string
	NoTelepon string
}

type JadwalSewa struct {
	ID int
	IDLapangan int
	IDPenyewa int
	JamMulai int
	JamSelesai int
	Bulan string
	Status string
	HargaTotal float64
}

var listLapangan []Lapangan
var listPenyewa []Penyewa
var listJadwal []JadwalSewa

var nextIDLapangan = 1
var nextIDPenyewa = 1
var nextIDJadwal = 1

// ----------------------------------
// MENU UTAMA
// ----------------------------------

func main() {
	for {
		fmt.Println("\n========================================")
		fmt.Println("             FUTSAL-BOOK               ")
		fmt.Println("     ---BOOKING LAPANGAN FUTSAL---      ")
		fmt.Println("========================================")
		fmt.Println("1. Kelola Data Lapangan & Penyewa")
		fmt.Println("2. Transaksi / Cek Ketersediaan Jam")
		fmt.Println("3. Pencarian")
		fmt.Println("4. Sorting")
		fmt.Println("5. Statistik Pendapatan & Jam Terpopuler")
		fmt.Println("6. Keluar")
		fmt.Println("========================================")
		fmt.Print("Pilih menu (1-6): ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			kelolaLapangan()
		case 2:
			menuTransaksi()
		case 3:
			menuPencarian()
		case 4:
			menuPengurutan()
		case 5:
			tampilkanStatistik()
		case 6:
			fmt.Println("Terima kasih telah menggunakan Futsal-Book!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// ----------------------------------
// 1. LAPANGAN & PENYEWA
// ----------------------------------

func kelolaLapangan() {
	fmt.Println("\n--- KELOLA DATA ---")
	fmt.Println("1. Tambah Lapangan")
	fmt.Println("2. Lihat/Hapus/Ubah Lapangan")
	fmt.Println("3. Tambah Penyewa")
	fmt.Println("4. Lihat/Hapus/Ubah Penyewa")
	fmt.Print("Pilih opsi: ")

	var opsi int
	fmt.Scanln(&opsi)

	switch opsi {
		
	case 1:
		var lapangan Lapangan
		lapangan.ID = nextIDLapangan
		fmt.Print("Nama Lapangan (Menggunakan Huruf Kecil Semua): ")
		fmt.Scanln(&lapangan.Nama)
		fmt.Print("Harga Sewa per Jam: ")
		fmt.Scanln(&lapangan.HargaSewa)
		listLapangan = append(listLapangan, lapangan)
		nextIDLapangan++
		fmt.Println("Lapangan berhasil ditambahkan!")

	case 2:
		fmt.Println("\nData Lapangan Saat Ini:")
		if len(listLapangan) == 0 {
			fmt.Println("   (Belum ada data lapangan)")
			return
		}
		for i := 0; i < len(listLapangan); i++ {
			fmt.Printf("   [%d] ID: %d | Nama: %s | Harga/Jam: %.2f\n", i, listLapangan[i].ID, listLapangan[i].Nama, listLapangan[i].HargaSewa)
		}

		fmt.Print("\nPilih Aksi (1: Ubah, 2: Hapus, 3: Kembali): ")
		var aksi int
		fmt.Scanln(&aksi)

		if aksi == 1 {
			fmt.Print("Masukkan Nomor Indeks Lapangan yang ingin diubah: ")
			var indeks int
			fmt.Scanln(&indeks)

			if indeks >= 0 && indeks < len(listLapangan) {
				fmt.Print("Nama Lapangan Baru: ")
				fmt.Scanln(&listLapangan[indeks].Nama)
				fmt.Print("Harga Sewa Baru per Jam: ")
				fmt.Scanln(&listLapangan[indeks].HargaSewa)
				fmt.Println("-> Data lapangan berhasil diperbarui!")
			} else {
				fmt.Println("-> Indeks tidak valid!")
			}

		} else if aksi == 2 {
			fmt.Print("Masukkan Nomor Indeks Lapangan yang ingin dihapus: ")
			var indeks int
			fmt.Scanln(&indeks)

			if indeks >= 0 && indeks < len(listLapangan) {
				listLapangan = append(listLapangan[:indeks], listLapangan[indeks+1:]...)
				fmt.Println("-> Lapangan berhasil dihapus!")
			} else {
				fmt.Println("-> Indeks tidak valid!")
			}
		}
	
	case 3:
		var penyewa Penyewa
		penyewa.ID = nextIDPenyewa
		fmt.Print("Nama Penyewa: ")
		fmt.Scanln(&penyewa.Nama)
		fmt.Print("Nomor Telepon: ")
		fmt.Scanln(&penyewa.NoTelepon)
		listPenyewa = append(listPenyewa, penyewa)
		nextIDPenyewa++
		fmt.Println("Penyewa berhasil ditambahkan!")
	
	case 4:
		fmt.Println("\nData Penyewa Saat Ini:")

		if len(listPenyewa) == 0 {
			fmt.Println("   (Belum ada data penyewa)")
			return
		}

		for i := 0; i < len(listPenyewa); i++ {
			fmt.Printf("   [%d] ID: %d | Nama: %s | No.Telp: %s\n", i, listPenyewa[i].ID, listPenyewa[i].Nama, listPenyewa[i].NoTelepon)
		}

		fmt.Print("\nPilih Aksi (1: Ubah, 2: Hapus, 3: Kembali): ")
		var aksi int
		fmt.Scanln(&aksi)

		if aksi == 1 {
			fmt.Print("Masukkan Nomor Indeks Penyewa yang ingin diubah: ")
			var indeks int
			fmt.Scanln(&indeks)

			if indeks >= 0 && indeks < len(listPenyewa) {
				fmt.Print("Nama Penyewa Baru (Huruf kecil): ")
				fmt.Scanln(&listPenyewa[indeks].Nama)
				fmt.Print("Nomor Telepon Baru: ")
				fmt.Scanln(&listPenyewa[indeks].NoTelepon)
				fmt.Println("-> Data penyewa berhasil diperbarui!")
			} else {
				fmt.Println("-> Indeks tidak valid!")
			}

		} else if aksi == 2 {
			fmt.Print("Masukkan Nomor Indeks Penyewa yang ingin dihapus: ")
			var indeks int
			fmt.Scanln(&indeks)

			if indeks >= 0 && indeks < len(listPenyewa) {
				// Mekanisme Hapus Data di Slice
				listPenyewa = append(listPenyewa[:indeks], listPenyewa[indeks+1:]...)
				fmt.Println("-> Penyewa berhasil dihapus!")
			} else {
				fmt.Println("-> Indeks tidak valid!")
			}
		}
	}
}

// ----------------------------------
// 2. TRANSAKSI
// ----------------------------------

func menuTransaksi() {
	var idLap, idUsr, jam, durasi int
	var bulan string

	fmt.Println("\n--- CATAT TRANSAKSI / BOOKING ---")
	fmt.Print("Masukkan ID Lapangan: ")
	fmt.Scanln(&idLap)
	fmt.Print("Masukkan ID Penyewa: ")
	fmt.Scanln(&idUsr)
	fmt.Print("Jam Mulai (Contoh: 15 untuk jam 3 sore): ")
	fmt.Scanln(&jam)
	fmt.Print("Durasi (Jam): ")
	fmt.Scanln(&durasi)
	fmt.Print("Bulan Transaksi (Contoh: Januari): ")
	fmt.Scanln(&bulan)

	tersedia := true

	for i := 0; i < len(listJadwal); i++ {
		jadwal := listJadwal[i]

		if jadwal.IDLapangan == idLap && jadwal.Bulan == bulan && jadwal.Status == "Dipesan" {

			if jam < jadwal.JamSelesai && (jam+durasi) > jadwal.JamMulai {
				tersedia = false
				break
			}
		}
	}

	if !tersedia {
		fmt.Println("Maaf, lapangan pada jam tersebut sudah dipesan!")
		return
	}

	var hargaPerJam float64

	for i := 0; i < len(listLapangan); i++ {

		l := listLapangan[i]

		if l.ID == idLap {
			hargaPerJam = l.HargaSewa
		}
	}

	// Simpan Transaksi Booking
	total := hargaPerJam * float64(durasi)
	baru := JadwalSewa{
		ID:         nextIDJadwal,
		IDLapangan: idLap,
		IDPenyewa:  idUsr,
		JamMulai:   jam,
		JamSelesai: jam + durasi,
		Bulan:      bulan,
		Status:     "Dipesan",
		HargaTotal: total,
	}

	listJadwal = append(listJadwal, baru)
	nextIDJadwal++
	fmt.Printf("Booking Berhasil! Total Biaya: IDR %.2f\n", total)

}

// ----------------------------------
// 3. PENCARIAN
// ----------------------------------

func menuPencarian() {
	fmt.Println("\n--- CARI DATA PENYEWA ---")
	fmt.Println("1. Cari Berdasarkan Nama Tepat (Sequential Search)")
	fmt.Println("2. Cari Berdasarkan No Telepon (Binary Search)")
	fmt.Print("Pilih metode: ")
	var opsi int
	fmt.Scanln(&opsi)

	if opsi == 1 {
		fmt.Print("Masukkan Nama Penyewa (Gunakan huruf kecil semua): ")
		var cariNama string
		fmt.Scanln(&cariNama)

		found := false

		for i := 0; i < len(listPenyewa); i++ {

			if listPenyewa[i].Nama == cariNama {
				fmt.Printf("[Ketemu] ID: %d | Nama: %s | Telp: %s\n", listPenyewa[i].ID, listPenyewa[i].Nama, listPenyewa[i].NoTelepon)
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Data nama penyewa tidak ditemukan.")
		}

	} else if opsi == 2 {
		fmt.Print("Masukkan No Telepon yang dicari: ")
		var cariTelp string
		fmt.Scanln(&cariTelp)

		n := len(listPenyewa)
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if listPenyewa[j].NoTelepon > listPenyewa[j+1].NoTelepon {

					listPenyewa[j], listPenyewa[j+1] = listPenyewa[j+1], listPenyewa[j]
				}
			}
		}

		low := 0
		high := len(listPenyewa) - 1
		found := false

		for low <= high {
			mid := (low + high) / 2
			if listPenyewa[mid].NoTelepon == cariTelp {
				fmt.Printf("[Ketemu via Binary] ID: %d | Nama: %s | Telp: %s\n", listPenyewa[mid].ID, listPenyewa[mid].Nama, listPenyewa[mid].NoTelepon)
				found = true
				break
			} else if listPenyewa[mid].NoTelepon < cariTelp {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		if !found {
			fmt.Println("Nomor telepon tidak ditemukan.")
		}
	}
}

// ----------------------------------
// 4. SORTING
// ----------------------------------

func menuPengurutan() {
    jadwalKosong := []JadwalSewa{
        {IDLapangan: 1, JamMulai: 13, HargaTotal: 150000, Status: "Tersedia"},
        {IDLapangan: 1, JamMulai: 7,  HargaTotal: 100000, Status: "Tersedia"},
        {IDLapangan: 2, JamMulai: 19, HargaTotal: 200000, Status: "Tersedia"},
        {IDLapangan: 2, JamMulai: 10, HargaTotal: 120000, Status: "Tersedia"},
    }

    fmt.Println("\n--- URUTKAN JADWAL KOSONG ---")
    fmt.Println("1. Urutkan berdasarkan Jam Mulai (Selection Sort)")
    fmt.Println("2. Urutkan berdasarkan Harga Sewa (Insertion Sort)")
    fmt.Print("Pilih opsi: ")
    var opsi int
    fmt.Scanln(&opsi)

    if opsi == 1 {
        var t JadwalSewa
        var i, j, idx_min int

        n := len(jadwalKosong)
        i = 1

        for i <= n-1 {
            idx_min = i - 1
            j = i

            for j < n {
                if jadwalKosong[idx_min].JamMulai > jadwalKosong[j].JamMulai {
                    idx_min = j
                }
                j++
            }

            t = jadwalKosong[idx_min]
            jadwalKosong[idx_min] = jadwalKosong[i-1]
            jadwalKosong[i-1] = t

            i++
        }
        fmt.Println("\nHasil Urut Jam Mulai (Selection Sort):")

    } else if opsi == 2 {
        var i, j int
        var temp JadwalSewa

        n := len(jadwalKosong)
        i = 1

        for i <= n-1 {
            j = i
            temp = jadwalKosong[j]

            for j > 0 && temp.HargaTotal < jadwalKosong[j-1].HargaTotal {
                jadwalKosong[j] = jadwalKosong[j-1]
                j--
            }

            jadwalKosong[j] = temp
            i++
        }
        fmt.Println("\nHasil Urut Harga Sewa (Insertion Sort):")
    }

    for i := 0; i < len(jadwalKosong); i++ {
        jk := jadwalKosong[i]
        fmt.Printf("Lapangan ID: %d | Jam Mulai: %02d:00 | Estimasi Harga: IDR %.2f\n", jk.IDLapangan, jk.JamMulai, jk.HargaTotal)
    }
}

// ----------------------------------
// 5. STATISTIK
// ----------------------------------

type RekapPendapatan struct {
	Bulan string
	Total float64
}

type RekapJam struct {
	Jam       int
	Frekuensi int
}

func tampilkanStatistik() {
	fmt.Println("\n--- STATISTIK LAPANGAN ---")

	var listPendapatan []RekapPendapatan
	var listHitungJam []RekapJam

	for i := 0; i < len(listJadwal); i++ {
		j := listJadwal[i]

		if j.Status == "Dipesan" {
			

			ketemuBulan := false

			for k := 0; k < len(listPendapatan); k++ {
				if listPendapatan[k].Bulan == j.Bulan {

					listPendapatan[k].Total += j.HargaTotal
					ketemuBulan = true
					break
				}
			}

			if !ketemuBulan {
				baru := RekapPendapatan{Bulan: j.Bulan, Total: j.HargaTotal}
				listPendapatan = append(listPendapatan, baru)
			}

			
			for jamAktif := j.JamMulai; jamAktif < j.JamSelesai; jamAktif++ {
				ketemuJam := false
				
				for k := 0; k < len(listHitungJam); k++ {
					if listHitungJam[k].Jam == jamAktif {
				
						listHitungJam[k].Frekuensi++
						ketemuJam = true
						break
					}
				}
				
				if !ketemuJam {
					baru := RekapJam{Jam: jamAktif, Frekuensi: 1}
					listHitungJam = append(listHitungJam, baru)
				}
			}
		}
	}

	fmt.Println("Total Pendapatan Bulanan:")
	if len(listPendapatan) == 0 {
		fmt.Println("  Belum ada transaksi.")
	}
	for i := 0; i < len(listPendapatan); i++ {
		fmt.Printf("  - %s: IDR %.2f\n", listPendapatan[i].Bulan, listPendapatan[i].Total)
	}
	jamTerpopuler := -1
	maxFrekuensi := 0

	for i := 0; i < len(listHitungJam); i++ {
		if listHitungJam[i].Frekuensi > maxFrekuensi {
			maxFrekuensi = listHitungJam[i].Frekuensi
			jamTerpopuler = listHitungJam[i].Jam
		}
	}

	fmt.Println("\nJam Operasional Paling Sering Dipesan:")
	if jamTerpopuler != -1 {
		fmt.Printf("  - Jam %02d:00 s/d %02d:00 (Dipesan sebanyak %d kali)\n", jamTerpopuler, jamTerpopuler+1, maxFrekuensi)
	} else {
		fmt.Println("  Belum ada data pemesanan jam.")
	}
}
