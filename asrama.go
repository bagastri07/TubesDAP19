package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

const nMhs = 3000
const nGdg = 5
const nKmr = 12
const nKsr = 4

//MHS Tipe data untuk mahasiswa
type MHS struct {
	Nama     string `json:"nama"`
	JKlmn    string `json:"jklmn"`
	Nim      string `json:"Nim"`
	Jurusan  string `json:"jurusan"`
	Angkatan string `json:"angkatan"`
	TAK      int    `json:"TAK"`
}

//GDG Tipe data untuk gedung
type GDG struct {
	Kode  string
	Nama  string `json:"Name"`
	Kamar [nKmr]KSR
}

type KSR struct {
	Kasur [nKsr]struct {
		Nama string
		NIM  string
	}
}
type RANK struct {
	Nama, JKlmn, Jurusan string
	TAK                  int
}
type RANKG struct {
	NamaGdg, JAsrma string
	TakTotal        float32
	Jumlah          int
}

type Mhs [nMhs]MHS
type Gdg [nGdg]GDG
type Rank [nMhs]RANK
type RankG [nGdg + nGdg]RANKG

func InputWithSpaces(data *string) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	*data = line

}

func ReadDataMHS(data *Mhs) {

	plan, _ := ioutil.ReadFile("mahasiswa.json")
	Convert := []byte(plan)
	err := json.Unmarshal(Convert, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func WriteDataMHS(data *Mhs) {
	output, err := json.MarshalIndent(data, "", "\t\t")
	if err != nil {
		fmt.Println("Gagal memperberbarui data mahasiswa:", err)
		return
	}
	err = ioutil.WriteFile("mahasiswa.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}

func ReadRankGedung(data *RankG) {

	plan, _ := ioutil.ReadFile("Peringkat Gedung.json")
	Convert := []byte(plan)
	err := json.Unmarshal(Convert, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func WriteRankGedung(data *RankG) {
	output, err := json.MarshalIndent(data, "", "\t\t")
	if err != nil {
		fmt.Println("Gagal memperberbarui data peringkat gedung:", err)
		return
	}
	err = ioutil.WriteFile("Peringkat Gedung.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}

func ReadDataGDGPutra(data *Gdg) {
	plan, _ := ioutil.ReadFile("Gedung Putra.json")
	Convert := []byte(plan)
	err := json.Unmarshal(Convert, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func WriteDataGDGPutra(data *Gdg) {
	output, err := json.MarshalIndent(data, "", "\t\t")
	if err != nil {
		fmt.Println("Gagal memperbarui data mahasiswa:", err)
		return
	}
	err = ioutil.WriteFile("Gedung Putra.json", output, 0644)
	if err != nil {
		fmt.Println("error memperbarui data file mahasiswa:", err)
		return
	}
}

func ReadDataGDGPutri(data *Gdg) {
	plan, _ := ioutil.ReadFile("Gedung Putri.json")
	Convert := []byte(plan)
	err := json.Unmarshal(Convert, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func WriteDataGDGPutri(data *Gdg) {
	output, err := json.MarshalIndent(data, "", "\t\t")
	if err != nil {
		fmt.Println("Gagal memperbarui data mahasiswa:", err)
		return
	}
	err = ioutil.WriteFile("Gedung Putri.json", output, 0644)
	if err != nil {
		fmt.Println("error memperbarui data file mahasiswa:", err)
		return
	}
}

func ReadDataRank(data *Rank) {
	plan, _ := ioutil.ReadFile("Peringkat TAK.json")
	Convert := []byte(plan)
	err := json.Unmarshal(Convert, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func WriteDataRANK(data *Rank) {
	output, err := json.MarshalIndent(data, "", "\t\t")
	if err != nil {
		fmt.Println("Gagal memperbarui data mahasiswa:", err)
		return
	}
	err = ioutil.WriteFile("Peringkat TAK.json", output, 0644)
	if err != nil {
		fmt.Println("error memperbarui data file mahasiswa:", err)
		return
	}
}

func FindEmpty(data Gdg, x, y, z *int) {
	found := false
	for !found {
		i := rand.Intn(nGdg)
		j := rand.Intn(nKmr)
		k := rand.Intn(nKsr)

		if data[i].Kamar[j].Kasur[k].NIM == "" && !found {
			*x = i
			*y = j
			*z = k
			found = true
		}
	}

}

/*func FindEmptyPutri(data GdgPutri, x, y, z *int) {
	found := false
	for i := 0; i < nGdg; i++ {
		for j := 0; j < nKmr; j++ {
			for k := 0; k < nKsr && !found; k++ {
				if data[i].Kamar[j].Kasur[k].NIM == "" {
					*x = i
					*y = j
					*z = k
					found = true
				}
			}
		}
	}
}*/

func FillNewStudent(data *Mhs) {
	for i := 0; i < nMhs; i++ {
		if data[i].Nim == "" {
			fmt.Printf("Isi data Mahasiswa baru tersebut :\n")
			fmt.Printf("\nNama Lengkap : ")
			InputWithSpaces(&data[i].Nama)
			fmt.Printf("\nCukup Ketik P atau L saja)\nJenis Kelamin : ")
			InputWithSpaces(&data[i].JKlmn)
			for data[i].JKlmn != "P" && data[i].JKlmn != "L" {
				fmt.Printf("\nCukup Ketik P atau L saja)\nJenis Kelamin : ")
				InputWithSpaces(&data[i].JKlmn)
			}
			fmt.Printf("NIM : ")
			InputWithSpaces(&data[i].Nim)
			fmt.Printf("Jurusan : ")
			InputWithSpaces(&data[i].Jurusan)
			fmt.Printf("Angkatan : ")
			InputWithSpaces(&data[i].Angkatan)
			fmt.Printf("TAK : ")
			fmt.Scanln(&data[i].TAK)
			fmt.Printf("\n\n===>Input Mahasiswa Baru Selesai<===\n")
			break
		}
	}
}

func EditDataStudent(data *Mhs, data1 Mhs) {
	var NIM string
	var Index, Input int
	fmt.Printf("\nPerubahan Data Mahasiswa\nNIM Mahasiswa Tersebut : ")
	fmt.Scanln(&NIM)
	SearchMHSfi(data1, NIM, &Index)
	if Index != -99 {
		for Input != 7 {
			fmt.Printf("\n1. Nama : %v\n2. Jenis Kelamin : %v\n3. NIM: %v\n4. Jurusan : %v\n5. Angkatan : %v\n6. TAK : %v\n7. Selesai", data[Index].Nama, data[Index].JKlmn, data[Index].Nim, data[Index].Jurusan, data[Index].Angkatan, data[Index].TAK)
			fmt.Printf("\n\nRespon : ")
			Input = Validasi(1, 7)
			switch Input {
			case 1:
				fmt.Printf("\nNama : ")
				InputWithSpaces(&data[Index].Nama)
			case 2:
				fmt.Printf("\nJenis Kelamin : ")
				InputWithSpaces(&data[Index].JKlmn)
			case 3:
				fmt.Printf("\nNIM : ")
				InputWithSpaces(&data[Index].Nim)
			case 4:
				fmt.Printf("\nJurusan : ")
				InputWithSpaces(&data[Index].Jurusan)
			case 5:
				fmt.Printf("\nAngkatan : ")
				InputWithSpaces(&data[Index].Angkatan)
			case 6:
				fmt.Printf("\nTAK : ")
				fmt.Scanln(&data[Index].TAK)
			}
		}
		fmt.Printf("\n===>Edit data mahasiswa berhasil<===")
	} else if Index == -99 {
		fmt.Printf("\nEdit data mahasiswa dibatalkan ...\n")
	}

}

func CheckIn(Mahasiswa Mhs, GPa *Gdg, GPa1 Gdg, GPI *Gdg, GPI1 Gdg) {
	var Input, Index, x, y, z int
	var NIM string

	SearchMHS(Mahasiswa, &NIM, &Index)
	if Index != -99 {
		fmt.Printf("\n\nKonfirmasi. Apakah benar ? (1. Benar 2. Tidak Benar : ")
		Input = Validasi(1, 2)
		for Input != 1 && Index != -99 {
			SearchMHS(Mahasiswa, &NIM, &Index)
			if Index != -99 {
				fmt.Printf("\n\nKonfirmasi. Apakah benar ? (1. Benar 2.Tidak Benar) : ")
				Input = Validasi(1, 2)
			}
		}
		if Index != -99 {
			isAlready := IsAlreadyFilled(Mahasiswa, GPa1, GPI1, NIM)
			if isAlready == false {
				fmt.Printf("\nKonfirmasi untuk Check in. (1. Check In 2. Batalkan) : ")
				Input = Validasi(1, 2)
				if Input == 1 {
					if Mahasiswa[Index].JKlmn == "P" {
						FindEmpty(GPI1, &x, &y, &z)
						GPI[x].Kamar[y].Kasur[z].Nama = Mahasiswa[Index].Nama
						GPI[x].Kamar[y].Kasur[z].NIM = Mahasiswa[Index].Nim
						Kamar := ConvertNoKamar(y)
						fmt.Printf("\n===>Check In berhasil<===\n\nDengan data kamar :\nNama : %v\nJurusan : %v", Mahasiswa[Index].Nama, Mahasiswa[Index].Jurusan)
						fmt.Printf("\nGedung: %v\nKamar: %v\nKasur : %v\n\n===>Check in Selesai<==\n", GPI1[x].Nama, Kamar, z+1)

					} else if Mahasiswa[Index].JKlmn == "L" {
						FindEmpty(GPa1, &x, &y, &z)
						GPa[x].Kamar[y].Kasur[z].Nama = Mahasiswa[Index].Nama
						GPa[x].Kamar[y].Kasur[z].NIM = Mahasiswa[Index].Nim
						Kamar := ConvertNoKamar(y)
						fmt.Printf("\n===>Check in berhasil<===\n\nDengan data kamar :\nNama : %v\nJurusan : %v", Mahasiswa[Index].Nama, Mahasiswa[Index].Jurusan)
						fmt.Printf("\nGedung: %v\nKamar: %v\nKasur : %v\n\n===>Check in Selesai<==\n", GPa1[x].Nama, Kamar, z+1)
					}
				} else if Input == 2 {
					fmt.Printf("\n==>Check In dibatalkan<==\n")
				}

			} else if isAlready == true {
				fmt.Printf("\nMahasiswa tersebut sudah berada dalam database gedung asrama\n")
				fmt.Printf("\n==>Check In dibatalkan<==\n")
			}
		}
	} else if Index == -99 {

	}

}

func CheckOut(Mahasiswa Mhs, GDGPutra *Gdg, GDGPutra1 Gdg, GDGPutri *Gdg, GDGPutri1 Gdg) {
	var NIM string
	var Index, Input, x, y, z int

	SearchMHS(Mahasiswa, &NIM, &Index)
	if Index != -99 {
		fmt.Printf("\n\nKonfirmasi. Apakah benar ? (1. Benar 2.Tidak Benar) : ")
		Input = Validasi(1, 2)
		for Input != 1 && Index != -99 {
			SearchMHS(Mahasiswa, &NIM, &Index)
			if Index != -99 {
				fmt.Printf("\n\nKonfirmasi. Apakah benar ? (1. Benar 2.Tidak Benar) : ")
				Input = Validasi(1, 2)
			}
		}
		if Index != -99 {
			if Mahasiswa[Index].JKlmn == "L" {
				SearchAsrama(GDGPutra1, NIM, &x, &y, &z)
				fmt.Printf("\n\nKonfirmasi Check Out. (1. Check Out, 2. Batalkan) : ")
				Input = Validasi(1, 2)
				if Input == 1 {
					GDGPutra[x].Kamar[y].Kasur[z].NIM = ""
					GDGPutra[x].Kamar[y].Kasur[z].Nama = ""
					fmt.Printf("\n ==>Check Out Berhasil<==")
				} else if Input == 2 {
					fmt.Printf("\n ==>Check Out Dibatalkan<==")
				}

			} else if Mahasiswa[Index].JKlmn == "P" {
				SearchAsrama(GDGPutri1, NIM, &x, &y, &z)
				fmt.Printf("\n\nKonfirmasi Check Out. (1. Check Out, 2. Batalkan) : ")
				Input = Validasi(1, 2)
				if Input == 1 {
					GDGPutri[x].Kamar[y].Kasur[z].NIM = ""
					GDGPutri[x].Kamar[y].Kasur[z].Nama = ""
					fmt.Printf("\n ==>Check Out Berhasil<==")
				} else if Input == 2 {
					fmt.Printf("\n ==>Check Out Dibatalkan<==")
				}
			}
		}
	} else if Index == -99 {

	}

}

func ChangeRoom(Mahasiswa Mhs, GDGPutra *Gdg, GDGPutra1 Gdg, GDGPutri *Gdg, GDGPutri1 Gdg) {
	var NIM string
	var Index, Input, x, y, z int

	SearchMHS(Mahasiswa, &NIM, &Index)
	if Index != -99 {
		fmt.Printf("\n\nKonfirmasi. Apakah benar ? (1. Benar 2.Tidak Benar) : ")
		Input = Validasi(1, 2)
		for Input != 1 && Index != -99 {
			SearchMHS(Mahasiswa, &NIM, &Index)
			if Index != -99 {
				fmt.Printf("\n\nKonfirmasi. Apakah benar ? (1. Benar 2.Tidak Benar) : ")
				Input = Validasi(1, 2)
			}
		}
		if Index != -99 {
			IsAlready := IsAlreadyFilled(Mahasiswa, GDGPutra1, GDGPutri1, NIM)
			if IsAlready == true {
				if Mahasiswa[Index].JKlmn == "L" {
					SearchAsrama(GDGPutra1, NIM, &x, &y, &z)
					fmt.Printf("\n\nKonfirmasi. (1. Change Room 2. Batalkan) : ")
					Input = Validasi(1, 2)
					if Input == 1 {
						GDGPutra[x].Kamar[y].Kasur[z].NIM = ""
						GDGPutra[x].Kamar[y].Kasur[z].Nama = ""
						FindEmpty(GDGPutra1, &x, &y, &z)
						GDGPutra[x].Kamar[y].Kasur[z].Nama = Mahasiswa[Index].Nama
						GDGPutra[x].Kamar[y].Kasur[z].NIM = Mahasiswa[Index].Nim
						Kamar := ConvertNoKamar(y)
						fmt.Printf("\n\n===>Change Room Berhasill<===\n\nDengan data kamar :\nNama : %v\nJurusan : %v", Mahasiswa[Index].Nama, Mahasiswa[Index].Jurusan)
						fmt.Printf("\nGedung: %v\nKamar: %v\nKasur : %v\n\n===>Change Room Selesai<==\n", GDGPutra1[x].Nama, Kamar, z+1)
					} else if Input == 2 {
						fmt.Printf("\n===>Change Room Dibatalkan<===")
					}
				} else if Mahasiswa[Index].JKlmn == "P" {
					SearchAsrama(GDGPutri1, NIM, &x, &y, &z)
					fmt.Printf("\n\nKonfirmasi. (1. Change Room 2. Batalkan) : ")
					Input = Validasi(1, 2)
					if Input == 1 {
						GDGPutri[x].Kamar[y].Kasur[z].NIM = ""
						GDGPutri[x].Kamar[y].Kasur[z].Nama = ""
						FindEmpty(GDGPutri1, &x, &y, &z)
						GDGPutri[x].Kamar[y].Kasur[z].Nama = Mahasiswa[Index].Nama
						GDGPutri[x].Kamar[y].Kasur[z].NIM = Mahasiswa[Index].Nim
						Kamar := ConvertNoKamar(y)
						fmt.Printf("\n\n===>Change Room Berhasill<===\n\nDengan data kamar :\nNama : %v\nJurusan : %v", Mahasiswa[Index].Nama, Mahasiswa[Index].Jurusan)
						fmt.Printf("\nGedung: %v\nKamar: %v\nKasur : %v\n\n===>Change Room Selesai<==\n", GDGPutri1[x].Nama, Kamar, z+1)
					} else if Input == 2 {
						fmt.Printf("\n===>Change Room Dibatalkan<===")
					}
				}
			} else if IsAlready == false {
				fmt.Printf("\nMahasiswa tersebut tidak terdaftar di dalam database Asrama")
			}
		}
	} else if Index == -99 {

	}

}

func IsAlreadyFilled(Mahasiswa Mhs, GDGPutra Gdg, GDGPutri Gdg, key string) bool {
	var Index int
	Already := false
	SearchMHSfi(Mahasiswa, key, &Index)
	if Mahasiswa[Index].JKlmn == "L" {
		for i := 0; i < nGdg; i++ {
			for j := 0; j < nKmr; j++ {
				for k := 0; k < nKsr; k++ {
					if GDGPutra[i].Kamar[j].Kasur[k].NIM == key {
						Already = true
						break
					}
				}
			}
		}
	} else if Mahasiswa[Index].JKlmn == "P" {
		for i := 0; i < nGdg; i++ {
			for j := 0; j < nKmr; j++ {
				for k := 0; k < nKsr; k++ {
					if GDGPutri[i].Kamar[j].Kasur[k].NIM == key {
						Already = true
						break
					}
				}
			}
		}
	}

	return Already
}

func ConvertNoKamar(z int) int {
	var kamar int
	if z >= 9 {
		kamar = z + 392
	} else if z >= 6 {
		kamar = z + 295
	} else if z >= 3 {
		kamar = z + 198
	} else {
		kamar = z + 101
	}
	return kamar
}

func Validasi(batas1 int, batas2 int) int {
	var key int

	valid := false
	for !valid {
		fmt.Scanln(&key)
		if key >= batas1 && key <= batas2 {
			valid = true
		} else if valid == false {
			fmt.Printf("Input tidak valid, silahkan diulang kembali\n Jawab (Ketik angka saja) : ")
		}
	}
	return key
}
func SearchMHS(data Mhs, nim *string, index *int) {
	var key string
	var chance int
	found := false
	for !found && chance < 4 {
		fmt.Printf("\nNIM Mahasiswa tersebut : ")
		fmt.Scanln(&key)
		for i := 0; i < nMhs && !found; i++ {
			if data[i].Nim == key {
				nama := data[i].Nama
				Jurusan := data[i].Jurusan
				Angkatan := data[i].Angkatan
				TAK := data[i].TAK
				fmt.Printf("\n==>Hasil Pencarian<===\n\nNama : %v\nJurusan : %v\nAngkatan : %v\nTAK : %v", nama, Jurusan, Angkatan, TAK)
				*index = i
				*nim = key
				found = true
			}
		}
		if found == false {
			chance = chance + 1
			if chance < 3 {
				fmt.Printf("\nMahasiswa tidak ditemukan\nUlangi Kembali. Kesempatan(%v/3)\n", chance)
			} else {
				fmt.Printf("\nMahasiswa tidak ditemukan\nUlangi Kembali. Kesempatan terakhir\n")
			}

		}
	}

	if chance > 3 {
		fmt.Printf("\nKesempatan habis. Menuju kembali ke menu ...\n")
		*index = -99
	}

}

func SearchMHSfi(data Mhs, key string, index *int) {
	var mid int
	found := false
	str := 0
	end := nMhs
	for str < end && !found {
		mid = (str + end) / 2
		if data[mid].Nim < key {
			end = mid
		} else if data[mid].Nim > key {
			str = mid + 1
		} else if data[mid].Nim == key {
			found = true
		}
	}
	if found == true {
		*index = mid
	} else if found == false {
		fmt.Printf("\nMahasiswa tidak terdaftar di database mahasiswa")
		*index = -99
	}

}

func SearchAsrama(data Gdg, key string, x, y, z *int) {
	found := false
	for i := 0; i < nGdg && !found; i++ {
		for j := 0; j < nKmr && !found; j++ {
			for k := 0; k < nKsr && !found; k++ {
				if data[i].Kamar[j].Kasur[k].NIM == key {
					*x = i
					*y = j
					*z = k
					gedung := data[i].Nama
					kamar := ConvertNoKamar(j)
					kasur := k + 1
					fmt.Printf("\n==>Data Asrama Mahasiswa Tersebut<===\n\nGedung : %v\nKamar : %v\nNo Kasur : %v", gedung, kamar, kasur)
					found = true
				}
			}
		}
	}
	if found == false {
		fmt.Printf("\nMahasiswa tersebut tidak terdaftar pada database Asrama")
	}
}

/*func SearchAsramaPutri(data GdgPutri, key string, x, y, z *int) {
	found := false
	for i := 0; i < nGdg && !found; i++ {
		for j := 0; j < nKmr && !found; j++ {
			for k := 0; k < nKsr && !found; k++ {
				if data[i].Kamar[j].Kasur[k].NIM == key {
					*x = i
					*y = j
					*z = k
					gedung := data[i].Nama
					kamar := ConvertNoKamar(j)
					kasur := k + 1
					fmt.Printf("\n==>Data Asrama Mahasiswa Tersebut<===\n\nGedung : %v\nKamar : %v\nNo Kasur : %v", gedung, kamar, kasur)
					found = true
				}
			}
		}
	}
	if found == false {
		fmt.Printf("\nMahasiswa tersebut tidak terdaftar pada database Asrama")
	}
}*/

func ShowRankMHS(data Rank) {
	fmt.Printf("==>10 Besar Mahasiswa/i Teraktif<==\n\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("Peringkat %v :\nNama :  %v\nJenis Kelamin : %v\nJurusan : %v\nPoint TAK : %v\n\n", i+1, data[i].Nama, data[i].JKlmn, data[i].Jurusan, data[i].TAK)
	}
}

func ShowRankMHSPutri(data1 Rank) {
	var j int

	fmt.Printf("==>10 Besar Mahasiswi Teraktif<==\n\n")
	for i := 0; j < 10; i++ {
		if data1[i].JKlmn == "P" {
			fmt.Printf("Peringkat %v :\nNama :  %v\nJenis Kelamin : %v\nJurusan : %v\nPoint TAK : %v\n\n", j+1, data1[i].Nama, data1[i].JKlmn, data1[i].Jurusan, data1[i].TAK)
			j++
		}
	}
}

func ShowRankMHSPutra(data2 Rank) {
	var j int

	fmt.Printf("==>10 Besar Mahasiswa Teraktif<==\n\n")
	for i := 0; j < 10; i++ {
		if data2[i].JKlmn == "L" {
			fmt.Printf("Peringkat %v :\nNama :  %v\nJenis Kelamin : %v\nJurusan : %v\nPoint TAK : %v\n\n", j+1, data2[i].Nama, data2[i].JKlmn, data2[i].Jurusan, data2[i].TAK)
			j++
		}
	}
}

func ShowRankGdg(data RankG) {
	fmt.Printf("==>Peringkat Gedung Asrama Teraktif<==\n\n")
	for i := 0; i < nGdg+nGdg; i++ {
		fmt.Printf("Peringkat %v\nNama Gedung : %v\nJenis Asrama : %v\nRata-rata TAK : %.2f\nJumlah Anggota : %v Mahasiswa\n\n", i+1, data[i].NamaGdg, data[i].JAsrma, data[i].TakTotal, data[i].Jumlah)
	}
}

func ShowRankGdgPa(data RankG) {
	var j int
	fmt.Printf("==>Peringkat Gedung Asrama Teraktif<==\n\n")
	for i := 0; i < nGdg+nGdg; i++ {
		if data[i].JAsrma == "Putra" {
			fmt.Printf("Peringkat %v\nNama Gedung : %v\nJenis Asrama : %v\nRata-rata TAK : %.2f\nJumlah Anggota : %v Mahasiswa\n\n", j+1, data[i].NamaGdg, data[i].JAsrma, data[i].TakTotal, data[i].Jumlah)
			j++
		}

	}
}

func ShowRankGdgPi(data RankG) {
	var j int
	fmt.Printf("==>Peringkat Gedung Asrama Teraktif<==\n\n")
	for i := 0; i < nGdg+nGdg; i++ {
		if data[i].JAsrma == "Putri" {
			fmt.Printf("Peringkat %v\nNama Gedung : %v\nJenis Asrama : %v\nRata-rata TAK : %.2f\nJumlah Anggota : %v Mahasiswa\n\n", j+1, data[i].NamaGdg, data[i].JAsrma, data[i].TakTotal, data[i].Jumlah)
			j++
		}

	}
}

func UpdateTAK(data *Mhs, data1 Mhs) {
	var NIM string
	var Index, Input int
	SearchMHS(data1, &NIM, &Index)
	if Index != -99 {
		fmt.Printf("\n\nKonfirmasi. Apakah benar ? (1. Benar 2.Tidak Benar) : ")
		Input = Validasi(1, 2)
		for Input != 1 && Index != -99 {
			SearchMHS(data1, &NIM, &Index)
			if Index != -99 {
				fmt.Printf("\n\nKonfirmasi. Apakah benar ? (1. Benar 2.Tidak Benar) : ")
				Input = Validasi(1, 2)
			}
		}
		if Index != -99 {
			fmt.Printf("Jumlah TAK baru : ")
			fmt.Scanln(&data[Index].TAK)
			fmt.Printf("\n\n===>Update TAK Selesai<==\n")
		}
	} else if Index == -99 {

	}

}

/*
func AutoSync(data Mhs, data2 *GdgPutri) {
	x := 0
	j := 0
	var SimpanPerempuan [170]struct {
		nama string
		nim  string
	}
	for i := 0; i < nMhs; i++ {
		if data[i].JKlmn == "P" {
			SimpanPerempuan[j].nama = data[i].Nama
			SimpanPerempuan[j].nim = data[i].Nim
			j = j + 1
		}
	}

	for i := 0; i < nGdg; i++ {
		for j := 0; j < nKmr; j++ {
			for k := 0; k < nKsr-2; k++ {
				if data2[i].Kamar[j].Kasur[k].NIM == "" {
					data2[i].Kamar[j].Kasur[k].Nama = SimpanPerempuan[x].nama
					data2[i].Kamar[j].Kasur[k].NIM = SimpanPerempuan[x].nim
					x = x + 1
				}
			}
		}
	}
}*/

func AutoSycRank(data Mhs, rank *Rank) {
	for k := 0; k < nMhs; k++ {
		rank[k].Nama = data[k].Nama
		rank[k].TAK = data[k].TAK
		rank[k].JKlmn = data[k].JKlmn
		rank[k].Jurusan = data[k].Jurusan
	}
}

func AutoSyncRankGDG(data Gdg, data1 Gdg, data2 Mhs, rank *RankG) {
	var TAKtotal, Index, nM int
	for k := 0; k < nGdg; k++ {
		nM = 0
		rank[k].NamaGdg = data[k].Nama
		rank[k].JAsrma = "Putra"
		TAKtotal = 0
		for l := 0; l < nKmr; l++ {
			for m := 0; m < nKsr; m++ {
				SearchMHSfi(data2, data[k].Kamar[l].Kasur[m].NIM, &Index)
				if data2[Index].Nim != "" {
					TAKtotal = TAKtotal + data2[Index].TAK
					nM = nM + 1
				}
			}
		}
		rank[k].TakTotal = float32(TAKtotal) / float32(nM)
		rank[k].Jumlah = nM
	}

	i := nGdg
	for k := 0; k < nGdg; k++ {
		nM = 0
		rank[i].NamaGdg = data1[k].Nama
		rank[i].JAsrma = "Putri"
		TAKtotal = 0
		for l := 0; l < nKmr; l++ {
			for m := 0; m < nKsr; m++ {
				SearchMHSfi(data2, data1[k].Kamar[l].Kasur[m].NIM, &Index)
				if data2[Index].Nim != "" {
					TAKtotal = TAKtotal + data2[Index].TAK
					nM = nM + 1
				}
			}
		}
		rank[i].TakTotal = float32(TAKtotal) / float32(nM)
		rank[i].Jumlah = nM
		i++
	}
}

func SelectSortRank(data *Rank) {
	for i := 0; i < nMhs; i++ {
		iMax := i
		for j := i + 1; j < nMhs; j++ {
			if data[iMax].TAK < data[j].TAK {
				iMax = j
			}
		}
		temp := data[iMax].TAK
		data[iMax].TAK = data[i].TAK
		data[i].TAK = temp
		temp1 := data[iMax].Nama
		data[iMax].Nama = data[i].Nama
		data[i].Nama = temp1
		temp1 = data[iMax].JKlmn
		data[iMax].JKlmn = data[i].JKlmn
		data[i].JKlmn = temp1
		temp1 = data[iMax].Jurusan
		data[iMax].Jurusan = data[i].Jurusan
		data[i].Jurusan = temp1
	}
}

func SelectSortRankG(data *RankG) {
	for i := 0; i < nGdg+nGdg; i++ {
		iMax := i
		for j := i + 1; j < nGdg+nGdg; j++ {
			if data[iMax].TakTotal < data[j].TakTotal {
				iMax = j
			}
		}
		temp := data[iMax].TakTotal
		data[iMax].TakTotal = data[i].TakTotal
		data[i].TakTotal = temp
		temp1 := data[iMax].NamaGdg
		data[iMax].NamaGdg = data[i].NamaGdg
		data[i].NamaGdg = temp1
		temp1 = data[iMax].JAsrma
		data[iMax].JAsrma = data[i].JAsrma
		data[i].JAsrma = temp1
		temp2 := data[iMax].Jumlah
		data[iMax].Jumlah = data[i].Jumlah
		data[i].Jumlah = temp2
	}
}

func InsertionSortMHS(data *Mhs) {
	i := 0
	for i < nMhs {
		t := data[i]
		j := i - 1
		for j >= 0 && data[j].Nim < t.Nim {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = t
		i = i + 1
	}
}

func SearchDataMHS(dataMahasiswa Mhs) {
	var NIM string
	var Index int

	fmt.Printf("\nNIM Mahasiswa Tersebut : ")
	fmt.Scanln(&NIM)
	SearchMHSfi(dataMahasiswa, NIM, &Index)
	fmt.Printf("\n\n===>Data Mahasiswa Tersebut<===\n\n")
	fmt.Printf("Nama : %v\nJenis Kelamin : %v\nJurusan : %v\nNIM : %v\nAngkatan : %v\nPoint TAK : %v", dataMahasiswa[Index].Nama, dataMahasiswa[Index].JKlmn, dataMahasiswa[Index].Jurusan, dataMahasiswa[Index].Nim, dataMahasiswa[Index].Angkatan, dataMahasiswa[Index].TAK)

}

func main() {
	var dataMahasiswa Mhs
	var dataGedungPa Gdg
	var dataGedungPi Gdg
	var Peringkat Rank
	var PeringkatGDG RankG
	var Input int

	/*
		//FillNewStudent(&dataMahasiswa)
		ReadDataMHS(&dataMahasiswa)
		ReadDataGDGPutra(&dataGedungPa)
		ReadDataGDGPutri(&dataGedungPi)
		//CheckIn(dataMahasiswa, &dataGedungPa, dataGedungPa, &dataGedungPi, dataGedungPi)
		//CheckOut(dataMahasiswa, &dataGedungPa, dataGedungPa, &dataGedungPi, dataGedungPi)
		//ChangeRoom(dataMahasiswa, &dataGedungPa, dataGedungPa, &dataGedungPi, dataGedungPi)
		//AutoSync(dataMahasiswa, &dataGedungPi)
		WriteDataGDGPutra(&dataGedungPa)
		WriteDataGDGPutri(&dataGedungPi)
		//FillNewStudent(&dataMahasiswa)
		//UpdateTAK(&dataMahasiswa, dataMahasiswa)
		WriteDataMHS(&dataMahasiswa)
		ReadDataRank(&Peringkat)
		AutoSycRank(dataMahasiswa, &Peringkat)
		SelectSortRank(&Peringkat)
		//ShowRankMHS(Peringkat)
		//ShowRankMHSPutra(Peringkat)
		WriteDataRANK(&Peringkat)
		ReadRankGedung(&PeringkatGDG)
		AutoSyncRankGDG(dataGedungPa, dataGedungPi, dataMahasiswa, &PeringkatGDG)
		SelectSortRankG(&PeringkatGDG)
		WriteRankGedung(&PeringkatGDG)
		//ShowRankGdg(PeringkatGDG)
		//ShowRankGdgPa(PeringkatGDG)
		ShowRankGdgPi(PeringkatGDG)
		//SearchMHS(dataMahasiswa, "1301194051")
		//fmt.Print(dataGedungPi[0].Kamar[0].Kasur[0].Nama)
		//var Input int
		//Input = Validasi(Input, 1, 2)

	*/
	ReadDataMHS(&dataMahasiswa)
	ReadDataGDGPutra(&dataGedungPa)
	ReadDataGDGPutri(&dataGedungPi)
	ReadDataRank(&Peringkat)
	ReadRankGedung(&PeringkatGDG)

	for Input != 5 {
		fmt.Printf("===>Menu Utama<===\n\n1. Info Data Asrama\n2. Info Data Mahasiswa\n3. Perubahan Data Mahasiswa\n4. Perubahan Data Asrama\n5. Keluar\n\n")
		fmt.Printf("Respon (ketik angka saja) : ")
		Input = Validasi(1, 5)

		switch Input {
		case 1:
			Input = 0
			for Input != 4 {
				fmt.Printf("\n\n===>Info Data Asrama<===\n\n1. Gedung Asrama Teraktif (Pa/Pi)\n2. Gedung Asrama Putra Teraktif\n3. Gedung Asrama Putri Teraktif\n4. Kembali\n\n")
				fmt.Printf("Respon (ketik angka saja) : ")
				Input = Validasi(1, 4)
				switch Input {
				case 1:
					ShowRankGdg(PeringkatGDG)
					time.Sleep(2 * time.Second)

				case 2:
					ShowRankGdgPa(PeringkatGDG)
					time.Sleep(2 * time.Second)
				case 3:
					ShowRankGdgPi(PeringkatGDG)
					time.Sleep(2 * time.Second)
				}
			}
		case 2:
			Input = 0
			for Input != 4 {
				fmt.Printf("\n\n===>Info Data Mahasiswa<===\n\n1. 10 Mahasiswa/i Teraktif\n2. 10 Mahasiswa Putra Teraktif\n3. 10 Mahasiswi Putri Teraktif\n4. Kembali\n\n")
				fmt.Printf("Respon (ketik angka saja) : ")
				Input = Validasi(1, 4)
				switch Input {
				case 1:
					ShowRankMHS(Peringkat)
					time.Sleep(2 * time.Second)
				case 2:
					ShowRankMHSPutra(Peringkat)
					time.Sleep(2 * time.Second)
				case 3:
					ShowRankMHSPutri(Peringkat)
					time.Sleep(2 * time.Second)
				}
			}
		case 3:
			Input = 0
			for Input != 4 {
				fmt.Printf("\n\n==>Perubahan Data Mahasiswa<===\n\n1. Update TAK Mahasiswa\n2. Form Mahasiswa Baru\n3. Edit Data Mahasiswa\n4. Kembali\n\n")
				fmt.Printf("Respon (ketik angka saja) : ")
				Input = Validasi(1, 4)
				switch Input {
				case 1:
					UpdateTAK(&dataMahasiswa, dataMahasiswa)
					InsertionSortMHS(&dataMahasiswa)
					WriteDataMHS(&dataMahasiswa)
					AutoSycRank(dataMahasiswa, &Peringkat)
					SelectSortRank(&Peringkat)
					AutoSyncRankGDG(dataGedungPa, dataGedungPi, dataMahasiswa, &PeringkatGDG)
					SelectSortRankG(&PeringkatGDG)
					time.Sleep(2 * time.Second)
				case 2:
					FillNewStudent(&dataMahasiswa)
					InsertionSortMHS(&dataMahasiswa)
					WriteDataMHS(&dataMahasiswa)
					AutoSycRank(dataMahasiswa, &Peringkat)
					SelectSortRank(&Peringkat)
					time.Sleep(2 * time.Second)
				case 3:
					EditDataStudent(&dataMahasiswa, dataMahasiswa)
					InsertionSortMHS(&dataMahasiswa)
					WriteDataMHS(&dataMahasiswa)
					AutoSycRank(dataMahasiswa, &Peringkat)
					SelectSortRank(&Peringkat)
					AutoSyncRankGDG(dataGedungPa, dataGedungPi, dataMahasiswa, &PeringkatGDG)
					SelectSortRankG(&PeringkatGDG)
					time.Sleep(2 * time.Second)
				}
			}
		case 4:
			Input = 0
			for Input != 4 {
				fmt.Printf("\n\n===>Perubahan Data Asrama<==\n\n1. Check In Asrama\n2. Check Out Asrama\n3. Change Room Kamar\n4. Kembali\n\n")
				fmt.Printf("Respon (ketik angka saja) : ")
				Input = Validasi(1, 4)

				switch Input {
				case 1:
					CheckIn(dataMahasiswa, &dataGedungPa, dataGedungPa, &dataGedungPi, dataGedungPi)
					WriteDataGDGPutra(&dataGedungPa)
					WriteDataGDGPutri(&dataGedungPi)
					AutoSyncRankGDG(dataGedungPa, dataGedungPi, dataMahasiswa, &PeringkatGDG)
					SelectSortRankG(&PeringkatGDG)
					WriteRankGedung(&PeringkatGDG)
					time.Sleep(2 * time.Second)
				case 2:
					CheckOut(dataMahasiswa, &dataGedungPa, dataGedungPa, &dataGedungPi, dataGedungPi)
					WriteDataGDGPutra(&dataGedungPa)
					WriteDataGDGPutri(&dataGedungPi)
					AutoSyncRankGDG(dataGedungPa, dataGedungPi, dataMahasiswa, &PeringkatGDG)
					SelectSortRankG(&PeringkatGDG)
					WriteRankGedung(&PeringkatGDG)
					time.Sleep(2 * time.Second)
				case 3:
					ChangeRoom(dataMahasiswa, &dataGedungPa, dataGedungPa, &dataGedungPi, dataGedungPi)
					WriteDataGDGPutra(&dataGedungPa)
					WriteDataGDGPutri(&dataGedungPi)
					AutoSyncRankGDG(dataGedungPa, dataGedungPi, dataMahasiswa, &PeringkatGDG)
					SelectSortRankG(&PeringkatGDG)
					WriteRankGedung(&PeringkatGDG)
					time.Sleep(2 * time.Second)
				}
			}

		}

	}
	InsertionSortMHS(&dataMahasiswa)
	WriteDataMHS(&dataMahasiswa)
	WriteDataGDGPutra(&dataGedungPa)
	WriteDataGDGPutri(&dataGedungPi)
	AutoSycRank(dataMahasiswa, &Peringkat)
	SelectSortRank(&Peringkat)
	WriteDataRANK(&Peringkat)
	AutoSyncRankGDG(dataGedungPa, dataGedungPi, dataMahasiswa, &PeringkatGDG)
	SelectSortRankG(&PeringkatGDG)
	WriteRankGedung(&PeringkatGDG)
}
