package myModule

import "fmt"

var age int = 18
var isStudent bool = true

func CheckUserAge() {
	age := 18

	if age >= 18 {
		fmt.Println("Anda dewasa.")
	} else {
		fmt.Println("Anda masih anak-anak.")
	}
}

// CheckStudentAge
func CheckStudentAge() {

	if age >= 18 && isStudent {
		fmt.Println("Anda adalah mahasiswa dewasa.")
	} else {
		fmt.Println("Anda bukan mahasiswa dewasa.")
	}
}

// CheckWeekend
func CheckWeekend() {
	isWeekend := true
	isHoliday := false

	if isWeekend || isHoliday {
		fmt.Println("Hari ini adalah hari libur atau akhir pekan.")
	} else {
		fmt.Println("Hari ini adalah hari kerja biasa.")
	}
}
