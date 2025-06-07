package main

import "fmt"

// type investasi struct {
//   jenis string
//   harga_saham float64
//   harga_reksa float64
//   harga_obligasi float64
// }
var firstname, lastname string
var minggu int = 1

const NMAX int = 100

type investasi struct {
	jenis    [3]string
	Saham    [NMAX]float64
	Reksa    [NMAX]float64
	Obligasi [NMAX]float64
}

var invest = investasi{
	jenis: [3]string{"Saham", "Reksa", "Obligasi"}, //0, 1, 2
}

type modaluser struct {
	Saham    [NMAX]float64
	Reksa    [NMAX]float64
	Obligasi [NMAX]float64
}

var modal modaluser

type portofoliouser struct {
	Saham    float64
	Reksa    float64
	Obligasi float64
}

var porto portofoliouser

type asetuser struct {
	Saham    float64
	Reksa    float64
	Obligasi float64
}

var aset asetuser

type iter struct {
	Saham    int
	Reksa    int
	Obligasi int
}

var iterasi iter

// for i = 0; i < iter.Saham; i++ {

// }

//fmt.Println(invest.jenis[0], invest.Saham[1]) //menampilkan saham dan harga minggu kedua

func Menu1() {
	fmt.Println("=================")
	fmt.Println(" Trading.com ")
	fmt.Println()

	// Cetak dulu prompt
	fmt.Print("First Name: ")
	fmt.Scanln(&firstname)

	fmt.Print("Last Name: ")
	fmt.Scanln(&lastname)

	fmt.Println("Hallo,", firstname, lastname, "Selamat datang di aplikasi investasi")
}

func Menu2() bool {
	var a int
	fmt.Println("====Start====")
	fmt.Println("Menu")
	fmt.Println("  1. Invest")
	fmt.Println("  2. Profil/Data")
	fmt.Println("  3. Instrumen Investasi")
	fmt.Println("  4. Exit")
	fmt.Print("Pilih menu :  ")
	fmt.Scan(&a)

	switch a {
	case 1:
		menu_invest()
	case 2:
		profil()
	case 3:
		instrumen_investasi()
	case 4:
		fmt.Println("Terimakasih telah menggunakan aplikasi ini")
		return false
	}
	return true
}

func menu_invest() {

	var opsi_invest int

	fmt.Println("================")
	fmt.Println("1. Saham")
	fmt.Println("2. Reksadana")
	fmt.Println("3. Obligasi")
	fmt.Println("4. Menu utama")
	fmt.Printf("Pilih menu : ")

	fmt.Scan(&opsi_invest)
	switch opsi_invest {

	case 1:
		fmt.Println("Investasi Saham")
		fmt.Println("Masukkan harga pasar ", iterasi.Saham)
		fmt.Print("Saham : Rp.")
		fmt.Scan(&invest.Saham[iterasi.Saham])

		invest_saham("Saham")
	case 2:
		fmt.Println("Investasi Reksadana")
		fmt.Println("Masukkan harga pasar ", iterasi.Reksa)
		fmt.Print("Reksadana : Rp.")
		fmt.Scan(&invest.Reksa[iterasi.Reksa])

		invest_saham("Reksa")
	case 3:
		fmt.Println("Investasi Obligasi")
		fmt.Println("Masukkan harga pasar ", iterasi.Obligasi)
		fmt.Print("Reksadana : Rp.")
		fmt.Scan(&invest.Obligasi[iterasi.Obligasi])
		invest_saham("Obligasi")
	case 4:
		return

	}

}

// func hitung_aset(jenis string) float64 {

//  /*
//     Menghitung total aset yang dipunyai user

//  */

//  var total float64
//  var i int
//  for i = 0; i < minggu; i++ {

//    switch jenis {

//    case "Saham":
//      total = total + porto.Saham
//    case "Reksa":
//      total = total + porto.Reksa
//    case "Obligasi":
//      total = total + porto.Obligasi
//    }
//  }

//  return total

// }

func hitung_modal(jenis string) float64 {
	var total float64
	var batas int

	switch jenis {
	case "Saham":
		batas = iterasi.Saham
		for i := 0; i <= batas; i++ {
			total += modal.Saham[i]
		}
	case "Reksa":
		batas = iterasi.Reksa
		for i := 0; i < batas; i++ {
			total += modal.Reksa[i]
		}
	case "Obligasi":
		batas = iterasi.Obligasi
		for i := 0; i < batas; i++ {
			total += modal.Obligasi[i]
		}
	}

	return total
}

func convert_aset(modal, harga float64) float64 {
	/*
	   Konversi aset dari modal yang diinputkan dibagi harga aset saat ini
	*/
	var aset float64

	aset = modal / harga

	return aset

}

func hitung_aset(porto, harga float64) float64 {
	/*
	   Konversi aset dari modal yang diinputkan dibagi harga aset saat ini
	*/
	var total float64

	total = porto * harga

	return total

}

func invest_saham(jenis string) {
/*
  Fungsi untuk menambah aset yang dipilih, saat ini baru Saham yang bisa dijalankan
*/


  var opsi_menu int

  fmt.Println("================")
  fmt.Println("1. Tambah investasi")
  fmt.Println("2. Ubah investasi")
  fmt.Println("3. Hapus investasi")
  fmt.Println("4. Kembali")
  fmt.Println("5. Menu utama")
  fmt.Printf("Pilih Menu : ")

  fmt.Scan(&opsi_menu)
  switch opsi_menu {
    
    case 1:
        if jenis == "Saham" {
          tambah_invest()          
        } else if jenis == "Reksa" {
          tambah_invest_reksa()
        } else if jenis == "Obligasi" {
          tambah_invest_obligasi()     
        }
    case 2:
        if jenis == "Saham" {
          ubahInvest("Saham")       
        } else if jenis == "Reksa" {
          ubahInvest("Reksa")
        } else if jenis == "Obligasi" {
          ubahInvest("Obligasi")       
        }
    case 3:
        if jenis == "Saham" {
          hapus_invest()       
        } else if jenis == "Reksa" {
          hapus_invest_reksa()
        } else if jenis == "Obligasi" {
          hapus_invest_obligasi()       
        }


    case 4:
      menu_invest()
    case 5:
      return
  }

}

func ubahInvest(jenisDari string) {
    var toJenis string
    var dana float64

    // 1. Tentukan konteks jenis asal
    var hargaDari float64
    var portoDari *float64
    var asetDari *float64
    var modalDari *[NMAX]float64
    var idxDari int

    switch jenisDari {
    case "Saham":
        idxDari     = iterasi.Saham
        hargaDari   = invest.Saham[idxDari]
        portoDari   = &porto.Saham
        asetDari    = &aset.Saham
        modalDari   = &modal.Saham
    case "Reksa":
        idxDari     = iterasi.Reksa
        hargaDari   = invest.Reksa[idxDari]
        portoDari   = &porto.Reksa
        asetDari    = &aset.Reksa
        modalDari   = &modal.Reksa
    case "Obligasi":
        idxDari     = iterasi.Obligasi
        hargaDari   = invest.Obligasi[idxDari]
        portoDari   = &porto.Obligasi
        asetDari    = &aset.Obligasi
        modalDari   = &modal.Obligasi
    default:
        fmt.Println("Jenis asal tidak dikenal.")
        return
    }

    // 2. Input target jenis dan jumlah dana
    fmt.Print("Pindahkan ke (Saham/Reksa/Obligasi): ")
    fmt.Scan(&toJenis)
    fmt.Print("Masukkan jumlah dana yang ingin dipindah: Rp.")
    fmt.Scan(&dana)

    // 3. Validasi harga tujuan, paksa input jika masih nol
    var hargaKe float64
    switch toJenis {
    case "Saham":
        hargaKe = invest.Saham[iterasi.Saham]
    case "Reksa":
        hargaKe = invest.Reksa[iterasi.Reksa]
    case "Obligasi":
        hargaKe = invest.Obligasi[iterasi.Obligasi]
    default:
        fmt.Println("Jenis tujuan tidak dikenal.")
        return
    }

    if hargaKe <= 0 {
        fmt.Printf("Harga %s minggu ini belum diinput. Masukkan harga baru: Rp.", toJenis)
        fmt.Scan(&hargaKe)
        // Simpan ke array invest sesuai jenis
        switch toJenis {
        case "Saham":
            invest.Saham[iterasi.Saham] = hargaKe
        case "Reksa":
            invest.Reksa[iterasi.Reksa] = hargaKe
        case "Obligasi":
            invest.Obligasi[iterasi.Obligasi] = hargaKe
        }
    }

    // 4. Hitung lembar yang dijual dan validasi kepemilikan
    lembarDijual := dana / hargaDari
    if *portoDari < lembarDijual {
        fmt.Println("Tidak cukup lembar untuk dijual. Batal.")
        return
    }

    // 5. Kurangi data jenis asal
    *portoDari -= lembarDijual
    *asetDari  -= lembarDijual * hargaDari
    modalDari[idxDari] -= dana

    // 6. Siapkan konteks jenis tujuan (pointer ke porto, aset, modal, idxKe)
    var portoKe *float64
    var asetKe  *float64
    var modalKe *[NMAX]float64
    var idxKe   int

    switch toJenis {
    case "Saham":
        idxKe   = iterasi.Saham
        portoKe = &porto.Saham
        asetKe  = &aset.Saham
        modalKe = &modal.Saham
    case "Reksa":
        idxKe   = iterasi.Reksa
        portoKe = &porto.Reksa
        asetKe  = &aset.Reksa
        modalKe = &modal.Reksa
    case "Obligasi":
        idxKe   = iterasi.Obligasi
        portoKe = &porto.Obligasi
        asetKe  = &aset.Obligasi
        modalKe = &modal.Obligasi
    }

    // 7. Konversi dana ke lembar baru dan update
    lembarBaru := dana / hargaKe
    *portoKe += lembarBaru
    *asetKe  += lembarBaru * hargaKe
    modalKe[idxKe] += dana

    // 8. Tampilkan ringkasan hasil
    fmt.Printf("Berhasil memindahkan Rp%.0f dari %s ke %s.\n", dana, jenisDari, toJenis)
    fmt.Printf("- Lembar %s dijual: %.2f → sisa: %.2f\n", jenisDari, lembarDijual, *portoDari)
    fmt.Printf("- Lembar %s baru : %.2f → total: %.2f\n", toJenis, lembarBaru, *portoKe)
}

func tambah_invest() {

	var jumlah_lembar float64
	// var jumlah_aset float64
	var jenis string
	var opsi_menu int
	var modal_user float64

	jenis = invest.jenis[0]
	fmt.Println("================")
	fmt.Println("Modal Anda", hitung_modal(jenis))
	fmt.Println("Aset Anda", porto.Saham, "Lembar")
	fmt.Println("================")

	fmt.Println("Harga per lembar : Rp.", invest.Saham[iterasi.Saham])
	fmt.Printf("Dana Investasi : Rp.")
	fmt.Scan(&modal_user)

	jumlah_lembar = convert_aset(modal_user, invest.Saham[iterasi.Saham])

	porto.Saham = porto.Saham + jumlah_lembar
	modal.Saham[iterasi.Saham] = modal.Saham[iterasi.Saham] + modal_user
	fmt.Println("Lembar saham yang didapatkan :", jumlah_lembar)

	fmt.Println("Total lembar", porto.Saham)

	// jumlah_aset = hitung_aset(jumlah_lembar, invest.Saham[minggu-1])

	// aset.Saham = aset.Saham + jumlah_aset

	aset.Saham = hitung_aset(porto.Saham, invest.Saham[iterasi.Saham])

	fmt.Println("Total Nilai Aset", aset.Saham)
	// fmt.Println("Modal yang telah dikeluarkan : ", hitung_modal(modal_total))

	fmt.Println("================")
	fmt.Println("1. Kembali")
	fmt.Println("2. Menu Utama")
	fmt.Printf("Pilih Menu : ")

	fmt.Scan(&opsi_menu)
	switch opsi_menu {

	// minggu++

	case 1:
		invest_saham("Saham")
	case 2:
		iterasi.Saham = iterasi.Saham + 1
		return

	}

}

func hapus_invest() {

	var pencairan_aset float64
	var dana_cair float64
	var status bool
	var konfirmasi, setelah int

	aset.Saham = hitung_aset(porto.Saham, invest.Saham[iterasi.Saham])

	status = false

	fmt.Println("Harga Saham : ", invest.Saham[iterasi.Saham])
	fmt.Println("Asset Anda Saat ini : ")
	fmt.Println(porto.Saham, "Lembar")
	fmt.Println("Nilai Aset", aset.Saham)

	// fmt.Scan(&dana_cair)

	for status == false {

		fmt.Println("2. Kembali")
		fmt.Println("3. Menu Utama")
		fmt.Printf("Masukan dana yang akan dicairkan : Rp.")
		fmt.Scan(&dana_cair)

		if dana_cair == 2 {

			invest_saham("Saham")

		} else if dana_cair == 3 {

			return

		} else if dana_cair < 0 {

			fmt.Println("Dana tidak bisa negatif !")

		} else {

			if dana_cair > aset.Saham {

				fmt.Println("Dana kurang !")

			} else {

				pencairan_aset = convert_aset(dana_cair, invest.Saham[iterasi.Saham])

				fmt.Println("Anda akan mencairkan sebanyak Rp.", dana_cair)
				fmt.Println(pencairan_aset, " Lembar saham akan dikurangi dari rekening Anda")
				fmt.Println("Anda yakin ? (ketik 1)")
				fmt.Println("1. Ya")
				fmt.Println("2. Kembali")
				fmt.Println("3. Menu Utama")

				fmt.Scan(&konfirmasi)

				switch konfirmasi {

				case 1:
					fmt.Println("Selamat dana sebesar Rp.", dana_cair, " akan masuk ke rekening Anda")
					aset.Saham = aset.Saham - dana_cair
					porto.Saham = porto.Saham - pencairan_aset
					status = true

					fmt.Println("1. Kembali")
					fmt.Println("2. Menu Utama")
					fmt.Scan(&setelah)

					if setelah == 1 {
						invest_saham("Saham")
					} else if setelah == 2 {
						return
					}

				case 2:
					invest_saham("Saham")
				case 3:
					menu_invest()
				}

			}

		}
	}

}

func tambah_invest_reksa() {

	fmt.Println("Berhasil masuk")

	var jumlah_unit float64
	var jenis string
	var opsi_menu int
	var modal_user float64

	jenis = invest.jenis[1]
	fmt.Println("================")
	fmt.Println("Modal Anda", hitung_modal(jenis))
	fmt.Println("Aset Anda", porto.Reksa, "Lembar")
	fmt.Println("================")

	fmt.Println("Harga per Unit : Rp.", invest.Reksa[iterasi.Reksa])
	fmt.Printf("Dana Investasi : Rp.")
	fmt.Scan(&modal_user)

	jumlah_unit = convert_aset(modal_user, invest.Reksa[iterasi.Reksa])

	porto.Reksa = porto.Reksa + jumlah_unit
	modal.Reksa[iterasi.Reksa] = modal.Reksa[iterasi.Reksa] + modal_user
	fmt.Println("Unit Obligasi yang didapatkan :", jumlah_unit)

	fmt.Println("Total Unit", porto.Reksa)
	aset.Reksa = hitung_aset(porto.Reksa, invest.Reksa[iterasi.Reksa])

	fmt.Println("Total Nilai Aset", aset.Reksa)

	fmt.Println("================")
	fmt.Println("1. Kembali")
	fmt.Println("2. Menu Utama")
	fmt.Printf("Pilih Menu : ")

	fmt.Scan(&opsi_menu)
	switch opsi_menu {

	// minggu++

	case 1:
		invest_saham("Reksa")
	case 2:
		iterasi.Reksa = iterasi.Reksa + 1
		return

	}

}

func hapus_invest_reksa() {

	var pencairan_aset float64
	var dana_cair float64
	var status bool
	var konfirmasi, setelah int

	aset.Reksa = hitung_aset(porto.Reksa, invest.Reksa[iterasi.Reksa])

	status = false

	fmt.Println("Harga Reksa : ", invest.Reksa[iterasi.Reksa])
	fmt.Println("Asset Anda Saat ini : ")
	fmt.Println(porto.Reksa, "Unit")
	fmt.Println("Nilai Aset", aset.Reksa)

	// fmt.Scan(&dana_cair)

	for status == false {

		fmt.Println("2. Kembali")
		fmt.Println("3. Menu Utama")
		fmt.Printf("Masukan dana yang akan dicairkan : Rp.")
		fmt.Scan(&dana_cair)

		if dana_cair == 2 {

			invest_saham("Reksa")

		} else if dana_cair == 3 {

			return

		} else if dana_cair < 0 {

			fmt.Println("Dana tidak bisa negatif !")

		} else {

			if dana_cair > aset.Reksa {

				fmt.Println("Dana kurang !")

			} else {

				pencairan_aset = convert_aset(dana_cair, invest.Reksa[iterasi.Reksa])

				fmt.Println("Anda akan mencairkan sebanyak Rp.", dana_cair)
				fmt.Println(pencairan_aset, " Unit Reksadana akan dikurangi dari rekening Anda")
				fmt.Println("Anda yakin ? (ketik 1)")
				fmt.Println("1. Ya")
				fmt.Println("2. Kembali")
				fmt.Println("3. Menu Utama")

				fmt.Scan(&konfirmasi)

				switch konfirmasi {

				case 1:
					fmt.Println("Selamat dana sebesar Rp.", dana_cair, " akan masuk ke rekening Anda")
					aset.Reksa = aset.Reksa - dana_cair
					porto.Reksa = porto.Reksa - pencairan_aset
					status = true

					fmt.Println("1. Kembali")
					fmt.Println("2. Menu Utama")
					fmt.Scan(&setelah)

					if setelah == 1 {
						invest_saham("Reksa")
					} else if setelah == 2 {
						return
					}

				case 2:
					invest_saham("Reksa")
				case 3:
					menu_invest()
				}

			}

		}
	}

}

func tambah_invest_obligasi() {

	fmt.Println("Berhasil masuk Obligasi")

	var jumlah_unit float64
	var jenis string
	var opsi_menu int
	var modal_user float64

	jenis = invest.jenis[2]
	fmt.Println("================")
	fmt.Println("Modal Anda", hitung_modal(jenis))
	fmt.Println("Aset Anda", porto.Obligasi, "Unit")
	fmt.Println("================")

	fmt.Println("Harga per Unit : Rp.", invest.Obligasi[iterasi.Obligasi])
	fmt.Printf("Dana Investasi : Rp.")
	fmt.Scan(&modal_user)

	jumlah_unit = convert_aset(modal_user, invest.Obligasi[iterasi.Obligasi])

	porto.Obligasi = porto.Obligasi + jumlah_unit
	modal.Obligasi[iterasi.Obligasi] = modal.Obligasi[iterasi.Obligasi] + modal_user
	fmt.Println("Unit Obligasi yang didapatkan :", jumlah_unit)

	fmt.Println("Total Unit", porto.Obligasi)
	aset.Obligasi = hitung_aset(porto.Obligasi, invest.Obligasi[iterasi.Obligasi])

	fmt.Println("Total Nilai Aset", aset.Obligasi)

	fmt.Println("================")
	fmt.Println("1. Kembali")
	fmt.Println("2. Menu Utama")
	fmt.Printf("Pilih Menu : ")

	fmt.Scan(&opsi_menu)
	switch opsi_menu {

	// minggu++

	case 1:
		invest_saham("Obligasi")
	case 2:
		iterasi.Obligasi = iterasi.Obligasi + 1
		return

	}

}

func hapus_invest_obligasi() {

	var pencairan_aset float64
	var dana_cair float64
	var status bool
	var konfirmasi, setelah int

	aset.Obligasi = hitung_aset(porto.Obligasi, invest.Obligasi[iterasi.Obligasi])

	status = false

	fmt.Println("Harga Obligasi : ", invest.Obligasi[iterasi.Obligasi])
	fmt.Println("Asset Anda Saat ini : ")
	fmt.Println(porto.Obligasi, "Unit")
	fmt.Println("Nilai Aset", aset.Obligasi)

	// fmt.Scan(&dana_cair)

	for status == false {

		fmt.Println("2. Kembali")
		fmt.Println("3. Menu Utama")
		fmt.Printf("Masukan dana yang akan dicairkan : Rp.")
		fmt.Scan(&dana_cair)

		if dana_cair == 2 {

			invest_saham("Obligasi")

		} else if dana_cair == 3 {

			return

		} else if dana_cair < 0 {

			fmt.Println("Dana tidak bisa negatif !")

		} else {

			if dana_cair > aset.Obligasi {

				fmt.Println("Dana kurang !")

			} else {

				pencairan_aset = convert_aset(dana_cair, invest.Obligasi[iterasi.Obligasi])

				fmt.Println("Anda akan mencairkan sebanyak Rp.", dana_cair)
				fmt.Println(pencairan_aset, " Unit Obligasi akan dikurangi dari rekening Anda")
				fmt.Println("Anda yakin ? (ketik 1)")
				fmt.Println("1. Ya")
				fmt.Println("2. Kembali")
				fmt.Println("3. Menu Utama")

				fmt.Scan(&konfirmasi)

				switch konfirmasi {

				case 1:
					fmt.Println("Selamat dana sebesar Rp.", dana_cair, " akan masuk ke rekening Anda")
					aset.Obligasi = aset.Obligasi - dana_cair
					porto.Obligasi = porto.Obligasi - pencairan_aset
					status = true

					fmt.Println("1. Kembali")
					fmt.Println("2. Menu Utama")
					fmt.Scan(&setelah)

					if setelah == 1 {
						invest_saham("Obligasi")
					} else if setelah == 2 {
						return
					}

				case 2:
					invest_saham("Obligasi")
				case 3:
					menu_invest()
				}

			}

		}
	}

}

func profil() {
	var status int
	var jenis string
	fmt.Println("===================")
	fmt.Println("====Profil Diri====")
	fmt.Println(" ")
	fmt.Print("Nama: ")
	fmt.Println(firstname, lastname)
	fmt.Println("Ringkasan Aset: ")
	fmt.Println("")
	for i := 0; i < 3; i++ {
		jenis = invest.jenis[i]
		fmt.Println("===", jenis, "===")
		fmt.Println("Modal", jenis, "Anda Rp. ", float64(hitung_modal(jenis)))
		switch jenis {
		case "Saham":
			if iterasi.Saham > 0 {
				fmt.Println("Aset: ", porto.Saham, "lembar")
				fmt.Println("Nilai Aset: ", float64(porto.Saham*invest.Saham[iterasi.Saham-1]))
				fmt.Println("Profit: ", float64(porto.Saham*invest.Saham[iterasi.Saham-1])-float64(hitung_modal(jenis)))
			} else {
				fmt.Println("Belum memiliki aset saham")
			}
		case "Reksa":
			if iterasi.Reksa > 0 {
				fmt.Println("Aset: ", porto.Reksa, "lembar")
				fmt.Println("Nilai Aset: ", float64(porto.Reksa*invest.Reksa[iterasi.Reksa-1]))
				fmt.Println("Profit: ", float64(porto.Reksa*invest.Reksa[iterasi.Reksa-1])-float64(hitung_modal(jenis)))
			} else {
				fmt.Println("Belum memiliki aset reksadana")
			}
		case "Obligasi":
			if iterasi.Obligasi > 0 {
				fmt.Println("Aset: ", porto.Obligasi, "lembar")
				fmt.Println("Nilai Aset: ", float64(porto.Obligasi*invest.Obligasi[iterasi.Obligasi-1]))
				fmt.Println("Profit: ", float64(porto.Obligasi*invest.Obligasi[iterasi.Obligasi-1])-float64(hitung_modal(jenis)))
			} else {
				fmt.Println("Belum memiliki aset obligasi")
			}

		}
		fmt.Println("")
	}
	fmt.Print("Kembali ke homepage? (ketik 1) ")
	fmt.Scan(&status)
	if status == 1 {
		return
	} else {
		fmt.Println("Inputan Tidak Valid, Silahkan ulangi lagi")
		fmt.Scan(&status)
	}
}
func riwayat_investasi(jenis string) {
	var status int
	fmt.Println("==================")
	fmt.Println("Riwayat", jenis)

	switch jenis {
	case "Saham":
		if iterasi.Saham == 0 {
			fmt.Println("Tidak ada riwayat investasi", jenis)
		} else {
			fmt.Printf("%-10s | %-25s | %-25s\n", "No", "Modal", "Harga")
			for i := 0; i < iterasi.Saham; i++ {
				fmt.Printf("%-10d | Rp.%-22.2f | Rp.%-22.2f\n", i+1, modal.Saham[i], invest.Saham[i])
			}
		}
	case "Reksa":
		if iterasi.Reksa == 0 {
			fmt.Println("Tidak ada riwayat investasi", jenis)
		} else {
			fmt.Printf("%-10s | %-25s | %-25s\n", "No", "Modal", "Harga")
			for i := 0; i < iterasi.Reksa; i++ {
				fmt.Printf("%-10d | Rp.%-22.2f | Rp.%-22.2f\n", i+1, modal.Reksa[i], invest.Reksa[i])
			}
		}
	case "Obligasi":
		if iterasi.Obligasi == 0 {
			fmt.Println("Tidak ada riwayat investasi", jenis)
		} else {
			fmt.Printf("%-10s | %-25s | %-25s\n", "No", "Modal", "Harga")
			for i := 0; i < iterasi.Obligasi; i++ {
				fmt.Printf("%-10d | Rp.%-22.2f | Rp.%-22.2f\n", i+1, modal.Obligasi[i], invest.Obligasi[i])
			}
		}
	}

	fmt.Println("Kembali? (ketik 1)")
	fmt.Scan(&status)
	if status == 1 {
		riwayat()
	} else {
		fmt.Println("Input tidak valid. Silakan ulangi.")
		riwayat_investasi(jenis)
	}
}

func riwayat() { 
	var opsi_menu int
	fmt.Println("=================")
	fmt.Println("Riwayat Investasi")
	fmt.Println("1. Saham")
	fmt.Println("2. Reksadana")
	fmt.Println("3. Obligasi")
	fmt.Println("4. Kembali")
	fmt.Println("5. Menu utama")
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&opsi_menu)

	switch opsi_menu {
	case 1:
		riwayat_investasi("Saham")
	case 2:
		riwayat_investasi("Reksa")
	case 3:
		riwayat_investasi("Obligasi")
	case 4:
		instrumen_investasi()
	case 5:
		return
	}
}
func pencarian() {
	var opsi_menu int
	var minimum_saham, minimum_reksa, minimum_obligasi float64
	var maksimum_saham, maksimum_reksa, maksimum_obligasi float64
	var modal_saham_minimum, modal_reksa_minimum, modal_obligasi_minimum float64
	var modal_saham_maksimum, modal_reksa_maksimum, modal_obligasi_maksimum float64
	fmt.Println("==================")
	fmt.Println("=====Pencarian====")
	fmt.Println("1. Harga Investasi Terendah")
	fmt.Println("2. Harga Investasi Tertinggi")
	fmt.Println("3. Modal Terendah")
	fmt.Println("4. Modal Tertinggi")
	fmt.Println("5. Kembali")
	fmt.Println("6. Menu utama")
	fmt.Print("Pilih Menu: ")
	fmt.Scan(&opsi_menu)
	fmt.Println("")

	switch opsi_menu {
	case 1:
		minimum_saham = invest.Saham[0]
		for i := 0; i < iterasi.Saham; i++ {
			if invest.Saham[i] < minimum_saham {
				minimum_saham = invest.Saham[i]
			}
		}
		fmt.Println("Nilai Minimum Saham: ", minimum_saham)

		minimum_reksa = invest.Reksa[0]
		for j := 0; j < iterasi.Reksa; j++ {
			if invest.Reksa[j] < minimum_reksa {
				minimum_reksa = invest.Reksa[j]
			}
		}
		fmt.Println("Nilai Minimum Reksadana: ", minimum_reksa)

		minimum_obligasi = invest.Obligasi[0]
		for k := 0; k < iterasi.Obligasi; k++ {
			if invest.Obligasi[k] < minimum_obligasi {
				minimum_obligasi = invest.Obligasi[k]
			}
		}
		fmt.Println("Nilai Minimum Obligasi: ", minimum_obligasi)
		fmt.Println("")
		pencarian()
	case 2:
		maksimum_saham = invest.Saham[0]
		for i := 0; i < iterasi.Saham; i++ {
			if invest.Saham[i] > maksimum_saham {
				maksimum_saham = invest.Saham[i]
			}
		}
		fmt.Println("Nilai Maksimum Saham: ", maksimum_saham)

		maksimum_reksa = invest.Reksa[0]
		for j := 0; j < iterasi.Reksa; j++ {
			if invest.Reksa[j] > maksimum_reksa {
				maksimum_reksa = invest.Reksa[j]
			}
		}
		fmt.Println("Nilai Maksimum Reksadana: ", maksimum_reksa)

		maksimum_obligasi = invest.Obligasi[0]
		for k := 0; k < iterasi.Obligasi; k++ {
			if invest.Obligasi[k] > maksimum_obligasi {
				maksimum_obligasi = invest.Obligasi[k]
			}
		}
		fmt.Println("Nilai Maksimum Obligasi: ", maksimum_obligasi)
		fmt.Println("")
		pencarian()
	case 3:
		modal_saham_minimum = modal.Saham[0]
		for i := 0; i < iterasi.Saham; i++ {
			if modal.Saham[i] < modal_saham_minimum {
				modal_saham_minimum = modal.Saham[i]
			}
		}
		fmt.Println("Modal Minimum Saham: ", modal_saham_minimum)

		modal_reksa_minimum = modal.Reksa[0]
		for i := 0; i < iterasi.Reksa; i++ {
			if modal.Reksa[i] < modal_reksa_minimum {
				modal_reksa_minimum = modal.Reksa[i]
			}
		}
		fmt.Println("Modal Minimum Reksadana: ", modal_reksa_minimum)

		modal_obligasi_minimum = modal.Obligasi[0]
		for i := 0; i < iterasi.Obligasi; i++ {
			if modal.Obligasi[i] < modal_obligasi_minimum {
				modal_obligasi_minimum = modal.Obligasi[i]
			}
		}
		fmt.Println("Modal Minimum Obligasi: ", modal_obligasi_minimum)
		fmt.Println("")
		pencarian()
	case 4:
		modal_saham_maksimum = modal.Saham[0]
		for i := 0; i < iterasi.Saham; i++ {
			if modal.Saham[i] > modal_saham_maksimum {
				modal_saham_maksimum = modal.Saham[i]
			}
		}
		fmt.Println("Modal Maksimum Saham: ", modal_saham_maksimum)

		modal_reksa_maksimum = modal.Reksa[0]
		for i := 0; i < iterasi.Reksa; i++ {
			if modal.Reksa[i] > modal_reksa_maksimum {
				modal_reksa_maksimum = modal.Reksa[i]
			}
		}
		fmt.Println("Modal Maksimum Reksadana: ", modal_reksa_maksimum)

		modal_obligasi_maksimum = modal.Obligasi[0]
		for i := 0; i < iterasi.Obligasi; i++ {
			if modal.Obligasi[i] > modal_obligasi_maksimum {
				modal_obligasi_maksimum = modal.Obligasi[i]
			}
		}
		fmt.Println("Modal Maksimum Obligasi: ", modal_obligasi_maksimum)
		fmt.Println("")
		pencarian()
	case 5:
		instrumen_investasi()
	case 6:
		return
	}

}
func instrumen_investasi() {
	var opsi_menu int
	fmt.Println("=========================")
	fmt.Println("===Instrumen Investasi===")
	fmt.Println("1. Riwayat")
	fmt.Println("2. Pencarian")
	fmt.Println("3. Menu Utama")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&opsi_menu)

	switch opsi_menu {
	case 1:
		riwayat()
	case 2:
		pencarian()
	case 3:
		return
	}
}
func main() {
	Menu1()
	runing := true //untuk menghentikan program
	for runing {
		runing = Menu2()
	}
}

/*Tampilan :

1
Masukkan nama :

2
Hallo,... Selamat datang di aplikasi investasi
=start=
Minggu 1

input harga saham, reksa, obligasi

Input harga saham :
Input harga reksadana :
Input harga obligasi :

Menu
  1. Invest
  2. Profil / data
  3. Instrumen investasi
  4. Akhiri minggu

(fitur utama)
if pilih Menu 1.
  1. Tambah (pilih : reksa, saham, Obligasi) saham beres
  2. ubah (pilih : reksa, saham, Obligasi)
  3. hapus (pilih : reksa, saham, Obligasi)


else if menu 2.
  1. Cetak laporan mingguan
  2. Cetak portofolio

else if menu 3.
  1. Harga saham waktu ke waktu
  2. Harga reksa waktu ke waktu
  3. Harga obligasi waktu ke waktu

else if menu 4.
  Lanjut ke minggu berikutnya, loop ke bagian start*/
