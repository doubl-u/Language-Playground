package main

import (
	//	"encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type data_struct struct {
	Username string  `json: "username"`
	Password string  `json: "password"`
	Balance  int     `json: "balance"`
	History  [10]int `json: "history"`
}

func makeflag() func(string) {
	counter := 0
	return func(s string) {
		counter += 1
		fmt.Printf("<flag.%s=%d>\n", s, counter)
	}
}

func getJSON() (*data_struct, *os.File) {
	jfile_path := "jfile.json"

	jFile, err := os.OpenFile(jfile_path, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("failed to open jsonfile:", err)
	}
	jRaw, err := io.ReadAll(jFile)
	if err != nil {
		log.Fatal("failed to read jsonfile:", err)
	}
	jData := data_struct{}
	json.Unmarshal(jRaw, &jData)
	return &jData, jFile
}

func writeJSON(jData *data_struct, jFile *os.File) {
	jRaw, err := json.Marshal(jData)
	if err != nil {
		log.Fatal("failed to Marshal:", err)
	}
	jFile.Seek(0, 0)
	jFile.Truncate(0)
	jFile.Write(jRaw)

}

func closeJSON(jFile *os.File) {
	jFile.Close()
}

func showBalance(p1 *data_struct) {
	fmt.Printf("balance:[%d]\n", p1.Balance)
}

func deposit(p1 *data_struct) {
	fmt.Println("how much monies")
	var amount int
	_, err := fmt.Scan(&amount)
	if err != nil {
		log.Fatal("invalid input type", err)
	}
	p1.Balance += amount
	updateHistory(p1, amount)

}

func withdraw(p1 *data_struct) {
	fmt.Println("how much monies")
	var amount int
	_, err := fmt.Scan(&amount)
	if err != nil {
		log.Fatal("invalid input type", err)
	}
	p1.Balance -= amount
	updateHistory(p1, -1*amount)
}

func updateHistory(p1 *data_struct, amount int) {
	for i, _ := range p1.History {
		if i > 0 {
			p1.History[i-1] = p1.History[i]
		}
	}
	p1.History[len(p1.History)-1] = amount

}

func showHistory(p1 *data_struct) {
	d1 := p1.History
	total := p1.Balance
	for i := len(p1.History) - 1; i >= 0; i-- {
		d1[i] = total
		total -= p1.History[i]
	}
	draw(&d1)
}

func draw(data *[10]int) {
	var points [100]int
	shape := [4]rune{'.', 'o', '^', '■'}
	var max int
	for _, v := range data {
		if max < v {
			max = v
		}
	}

	var last_value int
	for i, v := range data {
		temp_point := int(math.Ceil(10.0 * float64(v) / float64(max)))

		var temp_val int
		if temp_point == 0 && last_value == 0 {
			temp_val = 0
		} else if temp_point != 0 && last_value == 0 {
			temp_val = 3
		} else if temp_point > last_value {
			temp_val = 2
		} else {
			temp_val = 1
		}
		points[temp_point+10*i-1] = temp_val

		if v != 0 {
			last_value = temp_point
		}
	}

	for row := range 10 {
		for col := range 10 {
			fmt.Printf(" %c ", shape[points[((col+1)*10-row)-1]])
		}
		fmt.Printf("\n")
	}
}

func main() {
	//f1 := makeflag()
	p1, p2 := getJSON()
	defer closeJSON(p2)

	//temp_arr := [10]int{10, 56, 48, 4, 99, 78, 40, 65, 44, 66}
	//draw(&temp_arr)
	//log.Fatal()

	fmt.Println("what you want?")
	for exitFlag := false; exitFlag == false; {
		var option string
		fmt.Printf("\n1.show balance\n2.deposit\n3.withdraw\n4.show history\n0.exit\n")
		fmt.Scan(&option)
		switch option {
		case "0":
			fmt.Println("Exiting program.")
			exitFlag = true
		case "1":
			showBalance(p1)
		case "2":
			deposit(p1)
			writeJSON(p1, p2)
		case "3":
			withdraw(p1)
			writeJSON(p1, p2)
		case "4":
			showHistory(p1)
		default:
			fmt.Println("invalid input.")
		}
	}
}
