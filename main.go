package main

import (
	"fmt"
)

// Program ini berisi beberapa utilitas dan mini-game sederhana

// KOLEKSI CLASS DAN FUNCTION YANG DIBUAT OLEH AZIZ HEHANUSSA :)
// Mengubah integer menjadi string (pembungkus fmt.Sprint untuk konsistensi)
func intToString(val int) string {
	return fmt.Sprint(val)
}

// Mengubah float32 menjadi string
func float32ToString(val float32) string {
	return fmt.Sprint(val)
}

// Mengubah float64 menjadi string
func float64ToString(val float64) string {
	return fmt.Sprint(val)
}

// Mengubah boolean menjadi string "true"/"false"
func boolToString(val bool) string {
	if val {
		return "true"
	} else {
		return "false"
	}
}

// Membersihkan terminal dengan escape sequence (tidak guaranteed cross-platform)
func clearTerminal() {
	// Sesuatu-sesuatu yang ditemukan di Internet dan komunitas GoLang
	fmt.Print("\033[H\033[2J")
}

// Membulatkan float64 ke bawah (floor) dan menangani nilai negatif secara benar
// Contoh: mathfloor(1.9) -> 1, mathfloor(-1.1) -> -2
func mathfloor(x float64) int {
	var intPart = int(x)
	if x < 0 && float64(intPart) != x {
		return int(intPart - 1)
	}
	return int(intPart)
}

// Membatasi nilai minimum: jika val < min maka kembalikan min
func mathboundmin(val float64, min float64) float64 {
	if val < min {
		return min
	}
	return val
}

// Membatasi nilai maksimum: jika val > max maka kembalikan max
func mathboundmax(val float64, max float64) float64 {
	if val > max {
		return max
	}
	return val
}

// Membatasi nilai pada rentang [min, max]
func mathbound(val float64, min float64, max float64) float64 {
	return mathboundmax(mathboundmin(val, min), max)
}

// Memutar (wrap) nilai integer ke rentang [min, max] (modular wrap)
func mathwrap(val int, min int, max int) int {
	var temprange int = max - min + 1
	var tempval = val
	if val < min {
		// Menambahkan beberapa kali rentang supaya berada dalam range
		tempval = val + (temprange * mathfloor((float64(min)-float64(val))/float64(temprange)+1))
	}
	return (min + (tempval-min)%temprange)
}

// Seed untuk PRNG sederhana
var randomseed uint64 = 1234567890123456789

// Fungsi randomInt menghasilkan bilangan bulat pseudo-random dalam rentang [min,max].
// Implementasi menggunakan xorshift64* (fast PRNG) tanpa dependensi eksternal.
// Pastikan min <= max, fungsi akan menukar jika perlu.
func randomInt(min int, max int) int {
	// pastikan rentang valid
	if min > max {
		min, max = max, min
	}
	// xorshift64* (operasi bit untuk menghasilkan pseudo-random)
	randomseed ^= randomseed >> 12
	randomseed ^= randomseed << 25
	randomseed ^= randomseed >> 27
	val := randomseed * 2685821657736338717

	span := uint64(max - min + 1)
	if span == 0 {
		// rentang nol (min==max) -> kembalikan min
		return min
	}
	return min + int(val%span)
}

// Variabel global untuk interaksi dan statistik pemain
var curInput int = -1
var curTitle string
var playerScores, playerWins, playerGame int

// Membaca input integer dari pengguna ke curInput.
// Juga mereset randomseed menggunakan randomInt supaya seed sedikit berubah tiap input.
func plsInput() {
	randomseed = uint64(randomInt(123456, 999999999))
	fmt.Print(">> ")
	fmt.Scan(&curInput)
}

// Menangani input tak valid: menampilkan pesan dan mengatur curInput ke -1
func invalidInputYikes() {
	fmt.Println("INPUT INVALID atau INPUT TIDAK ADA PILIHAN.")
	curInput = -1
}

// Menampilkan judul halaman dengan membersihkan terminal.
// Jika txt kosong, gunakan curTitle sebelumnya.
func dashTitleHehe(txt string) {
	clearTerminal()
	if txt == "" {
		fmt.Println("-----", curTitle, "-----")
	} else {
		curTitle = txt
		fmt.Println("-----", txt, "-----")
	}
}

// Titik masuk program
func main() {
	mainMenu()
}

// Menu utama sederhana
func mainMenu() {
	curInput = -1
	dashTitleHehe("HALAMAN UTAMA")
	fmt.Println("1. Main Game")
	fmt.Println("2. Keluar")
	for curInput < 0 {
		plsInput()
		if curInput == 1 {
			menuGame()
			return
		} else if curInput == 2 {
			break
		} else {
			invalidInputYikes()
		}
	}
}

// Menu untuk memilih mini-game
func menuGame() {
	var simpleGameChoose = func(name string, playedCount int) string {
		if playedCount > 0 {
			return name + "     (" + fmt.Sprint(playedCount) + "x Dimainkan)"
		} else {
			return name
		}
	}
	curInput = -1
	dashTitleHehe("MENU GAME")
	fmt.Println("Total poin:", playerScores)
	fmt.Println("Total game dimainin:", playerGame)
	fmt.Println("Total menang:", playerWins)
	fmt.Println(simpleGameChoose("\n1. [+POIN] \"Tebak angka dalam Matematika\"", playedMathQuiz))
	fmt.Println(simpleGameChoose("2. \"Simulasi Kalkulator\"", playedCalculator))
	fmt.Println(simpleGameChoose("3. [+POIN] \"Gunting Batu Kertas\"", playedRPS))
	fmt.Println("4. Balik ke menu awal")
	fmt.Println("\nTolong beri input apa yang anda ingin pilih.")
	for curInput < 0 {
		plsInput()
		if curInput == 1 {
			gameMathQuiz()
			return
		} else if curInput == 2 {
			gameCalculator()
			return
		} else if curInput == 3 {
			gameRPS()
			return
		} else if curInput == 4 {
			mainMenu()
			return
		} else {
			invalidInputYikes()
		}
	}
}

// Berapa kali game ini dimainkan
var playedMathQuiz int

// Game: Tebak angka matematika (penjumlahan, pengurangan, perkalian)
func gameMathQuiz() {
	var questioncount int = 0
	var questioncorrect int = 0
	curInput = -1
	dashTitleHehe("MATH QUIZ")
	fmt.Println("Apakah anda yakin ingin memainkan game ini?\n1. Ya\n2. Tidak")
	for curInput < 0 {
		plsInput()
		if curInput == 1 {
			break
		} else if curInput == 2 {
			menuGame()
			return
		} else {
			invalidInputYikes()
		}
	}
	dashTitleHehe("")
	curInput = -1
	fmt.Println("Mau berapa banyak quiz nya?")
	plsInput()
	// Pastikan minimal 1 pertanyaan
	questioncount = int(mathboundmin(float64(curInput), 1))
	dashTitleHehe("")
	curInput = -1
	fmt.Println("Apakah anda udah siap?\n1. Ya\n2. Tidak")
	for curInput < 0 {
		plsInput()
		if curInput == 1 {
			break
		} else if curInput == 2 {
			menuGame()
			return
		} else {
			invalidInputYikes()
		}
	}
	var i int
	curInput = -1
	for i = 1; i <= questioncount; i++ {
		var x, y, z int
		var result, answer float64
		var zs string
		// Pilih angka acak dan operasi
		x = randomInt(-50, 50)
		y = randomInt(-50, 50)
		z = randomInt(1, 3)
		if z == 1 {
			result = float64(x) + float64(y)
			zs = "+"
		} else if z == 2 {
			result = float64(x) - float64(y)
			zs = "-"
		} else if z == 3 {
			result = float64(x) * float64(y)
			zs = "x"
		}
		dashTitleHehe("")
		fmt.Println("Pertanyaan", i, "\nBerapa hasil dari", x, zs, y, "= ?")
		fmt.Print(">> ")
		fmt.Scan(&answer)
		if answer == result {
			questioncorrect++
		}
	}
	dashTitleHehe("")
	fmt.Println("Hasil nya sudah dihitung, dan nilai anda adalah...")
	// Menghitung persentase dengan dua desimal (cara pembulatan sederhana)
	var grade = float64(int((float64(questioncorrect)/float64(questioncount))*10000) / 100)
	fmt.Println("Nilai:", grade, "%")
	fmt.Println("Soal Benar:", questioncorrect, "/", questioncount)
	fmt.Println("\nApakah anda terima dengan hasil ini?\n1. Ya")
	curInput = -1
	for curInput < 0 {
		plsInput()
		if curInput == 1 {
			playerScores += int(grade)
			playerGame++
			playedMathQuiz++
			if grade >= 50 {
				playerWins++
			}
			menuGame()
			return
		} else {
			invalidInputYikes()
		}
	}
}

// Berapa kali game ini dimainkan
var playedCalculator int

// Game: Kalkulator formula sederhana
// Catatan penting:
// - form menyimpan urutan tokens (angka/operator) sebagai []interface{}
// - numbers menyimpan angka-angka (float64) sebagai []interface{}
// - operators menyimpan operator sebagai []interface{}
// Saat menghitung akhir, program melakukan type assertion numbers[i].(float64)
// Pastikan nilai yang dimasukkan memang float64 (fmt.Scan dengan var float64).
func gameCalculator() {
	curInput = -1
	dashTitleHehe("CALCULATOR")
	fmt.Println("Apakah anda yakin ingin memainkan game ini?\n1. Ya\n2. Tidak")
	for curInput < 0 {
		plsInput()
		if curInput == 1 {
			break
		} else if curInput == 2 {
			menuGame()
			return
		} else {
			invalidInputYikes()
		}
	}
	curInput = -1
	var form = []interface{}{}
	var numbers = []interface{}{}
	var operators = []interface{}{}
	var issymbol bool = false   // apakah input selanjutnya harus operator?
	var isfinished bool = false // true jika user menekan '='
	for !isfinished {
		dashTitleHehe("")
		fmt.Print("Formula: ")
		var i int
		for i = 0; i < len(form); i++ {
			fmt.Print(fmt.Sprint(form[i]) + " ")
		}
		if issymbol {
			// Minta operator (+ - x :)
			var inputt string
			fmt.Println("\nMasukkan operator. (Inputnya + / - / x / :)\n(Jika sudah selesai bisa input (=) supaya hasil muncul.)")
			for inputt == "" {
				fmt.Print(">> ")
				fmt.Scan(&inputt)
				if inputt != "+" && inputt != "-" && inputt != "x" && inputt != ":" && inputt != "=" {
					invalidInputYikes()
					inputt = ""
				} else if inputt == "=" {
					// Selesai memasukkan formula, lanjut hitung
					isfinished = true
					break
				} else {
					form = append(form, inputt)
					operators = append(operators, inputt)
					issymbol = false
				}
			}
		} else {
			// Minta angka (float64)
			var inputt float64
			fmt.Println("\nMasukkan angka.")
			fmt.Print(">> ")
			fmt.Scan(&inputt)
			form = append(form, inputt)
			numbers = append(numbers, inputt)
			issymbol = true
		}
	}
	if isfinished {
		dashTitleHehe("")
		var formula string
		var result float64
		formula = "Formula:"
		var i int
		for i = 0; i < len(form); i++ {
			formula += " " + fmt.Sprint(form[i])
		}
		fmt.Print("Formulanya: ")
		fmt.Println(formula)
		fmt.Println("Maka hasilnya... ")
		// Ambil angka pertama (type assertion dari interface{} ke float64)
		result = numbers[0].(float64)
		if len(operators) > 0 {
			var i int
			for i = 0; i < len(operators); i++ {
				// Operator disimpan sebagai string, sehingga perbandingan langsung valid
				if operators[i] == "+" {
					result += numbers[i+1].(float64)
				} else if operators[i] == "-" {
					result -= numbers[i+1].(float64)
				} else if operators[i] == "x" {
					result *= numbers[i+1].(float64)
				} else if operators[i] == ":" {
					result /= numbers[i+1].(float64)
				}
			}
		}
		fmt.Print(result)
		fmt.Println("\nApakah sudah puas dengan hasilnya?\n1. Ya")
		curInput = -1
		for curInput < 0 {
			plsInput()
			if curInput == 1 {
				playerGame++
				playedCalculator++
				menuGame()
				return
			} else {
				invalidInputYikes()
			}
		}
	}
}

// Berapa kali game ini dimainkan
var playedRPS int

// Game: Gunting-Batu-Kertas (best of X)
// Menggunakan randomInt untuk keputusan AI
func gameRPS() {
	curInput = -1
	dashTitleHehe("ROCK PAPER SCISSOR")
	fmt.Println("Apakah anda yakin ingin memainkan game ini?\n1. Ya\n2. Tidak")
	for curInput < 0 {
		plsInput()
		if curInput == 1 {
			break
		} else if curInput == 2 {
			menuGame()
			return
		} else {
			invalidInputYikes()
		}
	}
	curInput = -1
	dashTitleHehe("")
	var turns, wins, ties, loses int
	fmt.Println("Mau bermain gunting batu kertas sampai berapa kali?\n1. 1 kali\n2. 3 kali\n3. 5 kali")
	for curInput < 0 {
		plsInput()
		if curInput == 1 {
			turns = 1
			break
		} else if curInput == 2 {
			turns = 3
			break
		} else if curInput == 3 {
			turns = 5
			break
		} else {
			invalidInputYikes()
		}
	}
	var rpss = []string{"Gunting", "Batu", "Kertas"}
	var i int
	for i = 1; i <= turns; i++ {
		var aiPick int = randomInt(1, 3)
		dashTitleHehe("")
		fmt.Println("Ronde", i, "\n1. Gunting\n2. Batu\n3. Kertas\n\nPilih salah satu.")
		curInput = -1
		for curInput < 0 {
			plsInput()
			if curInput >= 1 && curInput <= 3 {
				fmt.Println("Anda pilih:", rpss[curInput-1], "\nMusuh pilih:", rpss[aiPick-1])
				// Menentukan pemenang ronde berdasarkan aturan standar
				if curInput == aiPick {
					ties++
					fmt.Println("= Seri!")
					break
				} else if (curInput == 1 && aiPick == 3) || (curInput == 2 && aiPick == 1) || (curInput == 3 && aiPick == 2) {
					wins++
					fmt.Println("= Anda menang!")
					break
				} else {
					loses++
					fmt.Println("= Anda kalah!")
					break
				}
			} else {
				invalidInputYikes()
			}
		}
		curInput = -1
		if i == turns {
			fmt.Println("\nInput 1 untuk lanjut ke hasilnya.")
			for curInput < 0 {
				plsInput()
				if curInput == 1 {
					break
				} else {
					invalidInputYikes()
				}
			}
		} else {
			fmt.Println("\nInput 1 untuk lanjut ke ronde berikutnya.")
			for curInput < 0 {
				plsInput()
				if curInput == 1 {
					break
				} else {
					invalidInputYikes()
				}
			}
		}
	}
	// Menghitung skor akhir sebagai persentase (menang=1, seri=0.5)
	var result float64 = 0
	result += 1 * float64(wins)
	result += 0.5 * float64(ties)
	result /= float64(turns)
	var grade float64 = float64(int(result*10000) / 100)
	dashTitleHehe("")
	curInput = -1
	fmt.Println("Nilai anda adalah", grade, "%")
	fmt.Println("Menang:", wins, "\nSeri:", ties, "\nKalah:", loses)
	fmt.Println("\nApakah anda terima dengan hasil ini?\n1. Ya")
	for curInput < 0 {
		plsInput()
		if curInput == 1 {
			playerScores += int(grade)
			playerGame++
			playedRPS++
			if grade >= 50 {
				playerWins++
			}
			menuGame()
			return
		} else {
			invalidInputYikes()
		}
	}
}
