package main

import "fmt"

var firstname, lastname string
var minggu int = 1

const NMAX int = 100

type investasi struct {
	// HARGA INVESTASI PASAR
	jenis    [3]string
	Saham    [NMAX]float64
	Reksa    [NMAX]float64
	Obligasi [NMAX]float64
}

var invest = investasi{
	// JENIS ASET
	jenis: [3]string{"Saham", "Reksa", "Obligasi"}, //0, 1, 2
}

type modaluser struct {
	//  MODAL YANG DIPUNYA USER
	Saham    [NMAX]float64
	Reksa    [NMAX]float64
	Obligasi [NMAX]float64
}

var modal modaluser

type portofoliouser struct {
	// ASET USER DALAM BENTUR LEMBAR/UNIT	
	Saham    float64
	Reksa    float64
	Obligasi float64
}

var porto portofoliouser

type asetuser struct {
	// PORTOFOLIO USER (UNTUNG RUGI)
	Saham    float64
	Reksa    float64
	Obligasi float64
}

var aset asetuser

type iter struct {
	// ITERASI INPUTAN HARGA PASAR KE X
	Saham    int
	Reksa    int
	Obligasi int
}

var iterasi iter

type summary struct {
    Jenis      string
    Nilai      float64  // total nilai aset
    Persentase float64  // profit (%) 
}

func Menu1() {
	//MENU AWAL
	fmt.Println("=================")
	fmt.Println(" marinvest.id ")
	fmt.Println()

	// Masukan nama user
	fmt.Print("First Name: ")
	fmt.Scanln(&firstname)

	fmt.Print("Last Name: ")
	fmt.Scanln(&lastname)

	fmt.Println("Hallo,", firstname, lastname, "Selamat datang di aplikasi investasi")
}

func Menu2() bool {
	//MENU UTAMA
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
	// MENU PILIHAN INVEST
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
		fmt.Println("\nInvestasi Saham")
		fmt.Print("Masukkan harga pasar Saham : Rp.")
		fmt.Scan(&invest.Saham[iterasi.Saham])

		opsi_investasi("Saham")
	case 2:
		fmt.Println("\nInvestasi Reksadana")
		fmt.Print("Masukkan harga pasar Reksadana : Rp.")
		fmt.Scan(&invest.Reksa[iterasi.Reksa])

		opsi_investasi("Reksa")
	case 3:
		fmt.Println("\nInvestasi Obligasi")
		fmt.Print("Masukkan harga pasar Obligasi : Rp.")
		fmt.Scan(&invest.Obligasi[iterasi.Obligasi])
		opsi_investasi("Obligasi")
	case 4:
		return

	}

}

func hitung_modal(jenis string) float64 {
	// FUNCTION HITUNG MODAL
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
	   hitung aset dari porto yang dipunya dikali harga aset saat ini
	*/
	var total float64

	total = porto * harga

	return total

}

func opsi_investasi(jenis string) {
/*
  Fungsi untuk mengubah aset yang dipilih
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
          tambah_invest("Saham")          
        } else if jenis == "Reksa" {
          tambah_invest("Reksa")
        } else if jenis == "Obligasi" {
          tambah_invest("Obligasi")
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
          hapus_invest("Saham")

        } else if jenis == "Reksa" {
          hapus_invest("Reksa")

        } else if jenis == "Obligasi" {
          hapus_invest("Obligasi")
        }


    case 4:
      menu_invest()
    case 5:
      return
  }

}

func ubahInvest(jenisDari string) {
	// MENU UBAH JENIS INVEST (KONVERSI)
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


    fmt.Print("Pindahkan ke (Saham/Reksa/Obligasi): ")
    fmt.Scan(&toJenis)
    fmt.Print("Masukkan jumlah dana yang ingin dipindah: Rp.")
    fmt.Scan(&dana)


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

        switch toJenis {
        case "Saham":
            invest.Saham[iterasi.Saham] = hargaKe
        case "Reksa":
            invest.Reksa[iterasi.Reksa] = hargaKe
        case "Obligasi":
            invest.Obligasi[iterasi.Obligasi] = hargaKe
        }
    }


    lembarDijual := dana / hargaDari
    if *portoDari < lembarDijual {
        fmt.Println("Tidak cukup lembar untuk dijual. Batal.")
        return
    }


    *portoDari -= lembarDijual
    *asetDari  -= lembarDijual * hargaDari
    modalDari[idxDari] -= dana


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


    lembarBaru := dana / hargaKe
    *portoKe += lembarBaru
    *asetKe  += lembarBaru * hargaKe
    modalKe[idxKe] += dana


    fmt.Printf("Berhasil memindahkan Rp%.0f dari %s ke %s.\n", dana, jenisDari, toJenis)
    fmt.Printf("- Lembar %s dijual: %.2f → sisa: %.2f\n", jenisDari, lembarDijual, *portoDari)
    fmt.Printf("- Lembar %s baru : %.2f → total: %.2f\n", toJenis, lembarBaru, *portoKe)
}


func tambah_invest(jenisDari string) {
	// TAMBAH ASSET INVESTASI
	var jumlah_aset_dibeli float64
    var opsi_menu int
    var modal_user float64

    switch jenisDari {
    case invest.jenis[0]:
    	fmt.Println("================")
    	fmt.Println("Modal Anda 	 : Rp.", hitung_modal(invest.jenis[0]))
    	fmt.Println("Aset Anda        : ", porto.Saham, "Lembar")
    	fmt.Println("Harga per lembar : Rp.", invest.Saham[iterasi.Saham])    	
    	fmt.Println("================")

    	fmt.Printf("Dana Investasi               : Rp.")
    	fmt.Scan(&modal_user)

    	jumlah_aset_dibeli = convert_aset(modal_user, invest.Saham[iterasi.Saham])
    	porto.Saham = porto.Saham + jumlah_aset_dibeli
    	modal.Saham[iterasi.Saham] = modal.Saham[iterasi.Saham] + modal_user

    	fmt.Printf("Lembar saham yang didapatkan : %.4f\n", jumlah_aset_dibeli)
    	fmt.Printf("Total lembar                 : %.4f\n", porto.Saham)

    	aset.Saham = hitung_aset(porto.Saham, invest.Saham[iterasi.Saham])
     	fmt.Printf("Total Nilai Aset             : Rp. %.2f\n", aset.Saham)    	

	    fmt.Println("================")
	    fmt.Println("1. Kembali")
	    fmt.Println("2. Menu Utama")
	    fmt.Printf("Pilih Menu : ")

	    fmt.Scan(&opsi_menu)
	    switch opsi_menu {

	    case 1:
	        opsi_investasi("Saham")
	    case 2:
	        iterasi.Saham = iterasi.Saham + 1
	        return    	
	    }

    case invest.jenis[1]:
	    fmt.Println("================")
	    fmt.Println("Modal Anda 	 : Rp.", hitung_modal(invest.jenis[1]))
	    fmt.Println("Aset Anda        : ", porto.Reksa, "Unit")
	    fmt.Println("Harga per Unit   : Rp.", invest.Reksa[iterasi.Reksa])
	    fmt.Println("================")


	    fmt.Printf("Dana Investasi                 : Rp.")
	    fmt.Scan(&modal_user)	    

    	jumlah_aset_dibeli = convert_aset(modal_user, invest.Reksa[iterasi.Reksa])
	    porto.Reksa = porto.Reksa + jumlah_aset_dibeli
	    modal.Reksa[iterasi.Reksa] = modal.Reksa[iterasi.Reksa] + modal_user

	    fmt.Printf("Unit Reksadana yang didapatkan : %.4f\n", jumlah_aset_dibeli)
    	fmt.Printf("Total Unit                     : %.4f\n", porto.Reksa)

     	aset.Reksa = hitung_aset(porto.Reksa, invest.Reksa[iterasi.Reksa])
     	fmt.Printf("Total Nilai Aset               : Rp. %.2f\n", aset.Reksa)

	    fmt.Println("================")
	    fmt.Println("1. Kembali")
	    fmt.Println("2. Menu Utama")
	    fmt.Printf("Pilih Menu : ")

	    fmt.Scan(&opsi_menu)
	    switch opsi_menu {

	    case 1:
	        opsi_investasi("Reksa")
	    case 2:
	        iterasi.Reksa = iterasi.Reksa + 1
	        return

	    }   	

    case invest.jenis[2]:
	    fmt.Println("================")
	    fmt.Println("Modal Anda     : Rp.", hitung_modal(invest.jenis[2]))
	    fmt.Println("Aset Anda      : ", porto.Obligasi, "Unit")
	    fmt.Println("Harga per Unit : Rp.", invest.Obligasi[iterasi.Obligasi])	    
	    fmt.Println("================")

	    fmt.Printf("Dana Investasi                : Rp.")
	    fmt.Scan(&modal_user)

	    jumlah_aset_dibeli = convert_aset(modal_user, invest.Obligasi[iterasi.Obligasi])
	    porto.Obligasi = porto.Obligasi + jumlah_aset_dibeli
	    modal.Obligasi[iterasi.Obligasi] = modal.Obligasi[iterasi.Obligasi] + modal_user

    	fmt.Printf("Unit Obligasi yang didapatkan : %.4f\n", jumlah_aset_dibeli)
    	fmt.Printf("Total Unit                    : %.4f\n", porto.Obligasi)

     	aset.Obligasi = hitung_aset(porto.Obligasi, invest.Obligasi[iterasi.Obligasi])   	
    	fmt.Printf("Total Nilai Aset              : Rp. %.2f\n", aset.Obligasi)

	    fmt.Println("================")
	    fmt.Println("1. Kembali")
	    fmt.Println("2. Menu Utama")
	    fmt.Printf("Pilih Menu : ")

	    fmt.Scan(&opsi_menu)
	    switch opsi_menu {

	    case 1:
	        opsi_investasi("Obligasi")
	    case 2:
	        iterasi.Obligasi = iterasi.Obligasi + 1
	        return    	


    default:
    	fmt.Println("Jenis aset tidak ditemukan.")
    	return
    }	


}


}


func hapus_invest(jenisDari string) {
	// HAPUS ASET INVESTASI
    var danacair float64
    var pilih int
    var asetcair float64

    switch jenisDari {
    case invest.jenis[0]:
	    aset.Saham = hitung_aset(porto.Saham, invest.Saham[iterasi.Saham])
	    fmt.Println("================")
	    fmt.Printf("Harga Saham       : Rp%.0f\n", invest.Saham[iterasi.Saham])
	    fmt.Printf("Aset Anda Saat Ini: %.4f Lembar\n", porto.Saham)
	    fmt.Printf("Nilai Aset        : Rp%.0f\n\n", aset.Saham)

	    for {

	        fmt.Println("Masukkan dana yang akan dicairkan:")
	        fmt.Println("  (2. Kembali, 3. Menu Utama)")
	        fmt.Print("Rp. ")

	        fmt.Scan(&danacair)

	        if danacair == 2 {
	            opsi_investasi("Saham")
	            return
	        } else if danacair == 3 {
	            return
	        } else if danacair < 0 {
	            fmt.Println("Dana tidak bisa negatif!")
	        } else if danacair > aset.Saham {
	            fmt.Println("Dana melebihi nilai aset!")
	        } else {
	            // 3. Konfirmasi pencairan
	            asetcair = convert_aset(danacair, invest.Saham[iterasi.Saham])
	            fmt.Printf("\nAnda akan mencairkan Rp%.0f → %.4f lembar\n", danacair, asetcair)
	            fmt.Println("Anda yakin?")
	            fmt.Println("1. Ya")
	            fmt.Println("2. Kembali")
	            fmt.Println("3. Menu Utama")
	            fmt.Print("Pilihan: ")

	            fmt.Scan(&pilih)

	            if pilih == 1 {

	                porto.Saham -= asetcair
	                aset.Saham  -= danacair
	                fmt.Printf("\nDana Rp.%.0f berhasil dicairkan. Sisa lembar: %.4f\n\n", danacair, porto.Saham)

	                // 5. Setelah pencairan
	                fmt.Println("1. Kembali ke Saham")
	                fmt.Println("2. Menu Utama")
	                fmt.Print("Pilihan: ")
	                fmt.Scan(&pilih)
	                if pilih == 1 {
	                    opsi_investasi("Saham")
	                }
	                return

	            } else if pilih == 2 {
	                // Kembali ke pencarian

	            } else if pilih == 3 {
	                return

	            } else {
	                fmt.Println("Pilihan tidak valid.")
	            }
	        }

	    }



    case invest.jenis[1]:
	    aset.Reksa = hitung_aset(porto.Reksa, invest.Reksa[iterasi.Reksa])
	    fmt.Println("================")
	    fmt.Printf("Harga Reksa       : Rp%.0f\n", invest.Reksa[iterasi.Reksa])
	    fmt.Printf("Aset Anda Saat Ini: %.4f Unit\n", porto.Reksa)
	    fmt.Printf("Nilai Aset        : Rp%.0f\n\n", aset.Reksa)

	    for {

	        fmt.Println("Masukkan dana yang akan dicairkan:")
	        fmt.Println("  (2. Kembali, 3. Menu Utama)")
	        fmt.Print("Rp. ")

	        fmt.Scan(&danacair)

	        if danacair == 2 {
	            opsi_investasi("Reksa")
	            return
	        } else if danacair == 3 {
	            return
	        } else if danacair < 0 {
	            fmt.Println("Dana tidak bisa negatif!")

	        } else if danacair > aset.Reksa {
	            fmt.Println("Dana melebihi nilai aset!")

	        } else {

	            asetcair = convert_aset(danacair, invest.Reksa[iterasi.Reksa])
	            fmt.Printf("\nAnda akan mencairkan Rp%.0f → %.4f unit\n", danacair, asetcair)
	            fmt.Println("Anda yakin?")
	            fmt.Println("1. Ya")
	            fmt.Println("2. Kembali")
	            fmt.Println("3. Menu Utama")
	            fmt.Print("Pilihan: ")

	            fmt.Scan(&pilih)

	            if pilih == 1 {

	                porto.Reksa -= asetcair
	                aset.Reksa  -= danacair
	                fmt.Printf("\nDana Rp.%.0f berhasil dicairkan. Sisa Unit: %.4f\n\n", danacair, porto.Reksa)

	                fmt.Println("1. Kembali ke Reksa")
	                fmt.Println("2. Menu Utama")
	                fmt.Print("Pilihan: ")
	                fmt.Scan(&pilih)
	                if pilih == 1 {
	                    opsi_investasi("Reksa")
	                }
	                return

	            } else if pilih == 2 {
	                // Kembali ke pencarian

	            } else if pilih == 3 {
	                return

	            } else {
	                fmt.Println("Pilihan tidak valid.")
	            }
	        }

	    }


    case invest.jenis[2]:
	    aset.Obligasi = hitung_aset(porto.Obligasi, invest.Obligasi[iterasi.Obligasi])
	    fmt.Println("================")
	    fmt.Printf("Harga Obligasi       : Rp%.0f\n", invest.Obligasi[iterasi.Obligasi])
	    fmt.Printf("Aset Anda Saat Ini: %.4f Lembar\n", porto.Obligasi)
	    fmt.Printf("Nilai Aset        : Rp%.0f\n\n", aset.Obligasi)

	    for {

	        fmt.Println("Masukkan dana yang akan dicairkan:")
	        fmt.Println("  (2. Kembali, 3. Menu Utama)")
	        fmt.Print("Rp. ")

	        fmt.Scan(&danacair)

	        if danacair == 2 {
	            opsi_investasi("Obligasi")
	            return
	        } else if danacair == 3 {
	            return
	        } else if danacair < 0 {
	            fmt.Println("Dana tidak bisa negatif!")
	        } else if danacair > aset.Obligasi {
	            fmt.Println("Dana melebihi nilai aset!")
	        } else {
	            // 3. Konfirmasi pencairan
	            asetcair = convert_aset(danacair, invest.Obligasi[iterasi.Obligasi])
	            fmt.Printf("\nAnda akan mencairkan Rp%.0f → %.4f lembar\n", danacair, asetcair)
	            fmt.Println("Anda yakin?")
	            fmt.Println("1. Ya")
	            fmt.Println("2. Kembali")
	            fmt.Println("3. Menu Utama")
	            fmt.Print("Pilihan: ")

	            fmt.Scan(&pilih)

	            if pilih == 1 {

	                porto.Obligasi -= asetcair
	                aset.Obligasi  -= danacair
	                fmt.Printf("\nDana Rp.%.0f berhasil dicairkan. Sisa unit: %.4f\n\n", danacair, porto.Obligasi)

	                // 5. Setelah pencairan
	                fmt.Println("1. Kembali ke Obligasi")
	                fmt.Println("2. Menu Utama")
	                fmt.Print("Pilihan: ")
	                fmt.Scan(&pilih)
	                if pilih == 1 {
	                    opsi_investasi("Obligasi")
	                }
	                return

	            } else if pilih == 2 {
	                // Kembali ke pencarian

	            } else if pilih == 3 {
	                return

	            } else {
	                fmt.Println("Pilihan tidak valid.")
	            }
	        }

	    }


    default:
    	fmt.Println("Jenis aset tidak ditemukan.")
    	return
    }	


}

func profil() {
	// MENU PROFIL USER
	var jenis string
	var modaltotal, pct float64
	var i, maxIdx, j int

    fmt.Println("===================")
    fmt.Println("==== Profil Diri ====")
    fmt.Printf("Nama       : %s %s\n\n", firstname, lastname)
    fmt.Println("=== Ringkasan Aset ===")

    for i = 0; i < 3; i++ {
        jenis = invest.jenis[i]
        modaltotal = float64(hitung_modal(jenis))
        fmt.Printf("\n--- %s ---\n", jenis)
        fmt.Printf("Modal       : Rp%.0f\n", modaltotal)

        switch jenis {
        case "Saham":
            if iterasi.Saham > 0 {
                fmt.Printf("Aset        : %.4f lembar\n", porto.Saham)
                fmt.Printf("Nilai Aset  : Rp%.2f\n", aset.Saham)
                fmt.Printf("Profit      : Rp%.2f\n", aset.Saham - modaltotal)
            } else {
                fmt.Println("Belum memiliki aset saham")
            }
        case "Reksa":
            if iterasi.Reksa > 0 {
                fmt.Printf("Aset        : %.4f unit\n", porto.Reksa)
                fmt.Printf("Nilai Aset  : Rp%.2f\n", aset.Reksa)
                fmt.Printf("Profit      : Rp%.2f\n", aset.Reksa - modaltotal)
            } else {
                fmt.Println("Belum memiliki aset reksa dana")
            }
        case "Obligasi":
            if iterasi.Obligasi > 0 {
                fmt.Printf("Aset        : %.4f unit\n", porto.Obligasi)
                fmt.Printf("Nilai Aset  : Rp%.2f\n", aset.Obligasi)
                fmt.Printf("Profit      : Rp%.2f\n", aset.Obligasi - modaltotal)
            } else {
                fmt.Println("Belum memiliki aset obligasi")
            }
        }
    }

 
    type summary struct {
    	// RINGKASAN ASET USER
        Jenis      string
        Nilai      float64
        Persentase float64
    }
    var list [3]summary
    for i = 0; i < 3; i++ {
        jenis := invest.jenis[i]
        var nilai, modaltotal float64
        switch jenis {
        case "Saham":
            nilai = aset.Saham
            modaltotal = float64(hitung_modal("Saham"))
        case "Reksa":
            nilai = aset.Reksa
            modaltotal = float64(hitung_modal("Reksa"))
        case "Obligasi":
            nilai = aset.Obligasi
            modaltotal = float64(hitung_modal("Obligasi"))
        }
        pct = 0.0
        if modaltotal > 0 {
            pct = (nilai - modaltotal) / modaltotal * 100
        }
        list[i] = summary{Jenis: jenis, Nilai: nilai, Persentase: pct}
    }

    fmt.Println("\n=== Submenu Urutkan Aset ===")
    fmt.Println("1. Urutkan berdasarkan Nilai Aset")
    fmt.Println("2. Urutkan berdasarkan Persentase Profit")
    fmt.Println("3. Kembali")
    fmt.Print("Pilih: ")
    var pilih int
    fmt.Scan(&pilih)

    if pilih == 1 {
        // Selection sort descending by Nilai
        for i = 0; i < 2; i++ {
            maxIdx = i
            for j = i + 1; j < 3; j++ {
                if list[j].Nilai > list[maxIdx].Nilai {
                    maxIdx = j
                }
            }
            list[i], list[maxIdx] = list[maxIdx], list[i]
        }
        fmt.Println("\n-- Hasil Urutkan dari Nilai Aset --")
    } else if pilih == 2 {
        // Insertion sort descending by Persentase
        for i = 1; i < 3; i++ {
            key := list[i]
            j = i - 1
            for j >= 0 && list[j].Persentase < key.Persentase {
                list[j+1] = list[j]
                j--
            }
            list[j+1] = key
        }
        fmt.Println("\n-- Hasil Urutkan dari Persentase Profit --")
    } else {
        return
    }

    for i = 0; i < 3; i++ {
        fmt.Printf("%-8s | Nilai: Rp%.2f | Profit: %.2f%%\n",
            list[i].Jenis, list[i].Nilai, list[i].Persentase)
    }

    fmt.Println("\nTekan Enter untuk kembali...")
    fmt.Scanln()
    fmt.Scanln()
}



func riwayat_investasi(jenis string) {
	// MENU HITUNG RIWAYAT INVESTASI
	var status int
	fmt.Println("==================")
	fmt.Println("Riwayat", jenis)

	switch jenis {
	case "Saham":
		if iterasi.Saham == 0 {
			fmt.Println("Tidak ada riwayat investasi", jenis)
		} else {
			fmt.Printf("%-10s | %-25s | %-25s\n", "No", "Modal", "Harga")
			for i := 0; i <= iterasi.Saham; i++ {
				fmt.Printf("%-10d | Rp.%-22.2f | Rp.%-22.2f\n", i+1, modal.Saham[i], invest.Saham[i])
			}
		}
	case "Reksa":
		if iterasi.Reksa == 0 {
			fmt.Println("Tidak ada riwayat investasi", jenis)
		} else {
			fmt.Printf("%-10s | %-25s | %-25s\n", "No", "Modal", "Harga")
			for i := 0; i <= iterasi.Reksa; i++ {
				fmt.Printf("%-10d | Rp.%-22.2f | Rp.%-22.2f\n", i+1, modal.Reksa[i], invest.Reksa[i])
			}
		}
	case "Obligasi":
		if iterasi.Obligasi == 0 {
			fmt.Println("Tidak ada riwayat investasi", jenis)
		} else {
			fmt.Printf("%-10s | %-25s | %-25s\n", "No", "Modal", "Harga")
			for i := 0; i <= iterasi.Obligasi; i++ {
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
	// MENU RIWAYAT
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
	// MENU MENCARI ASET INVESTASI
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
	// MENU ASET INVESTASI
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
	// MAIN
	Menu1()
	runing := true //untuk menghentikan program
	for runing {
		runing = Menu2()
	}
}