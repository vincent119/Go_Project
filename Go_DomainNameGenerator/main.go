package main

// changed to v4

import (
	"bufio"
	//"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func IsNumericOnly(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, s := range str {
		if s < '0' || s > '9' {
			return false
		}
	}

	return true
}

func randomString(l int) string {
	var pool = "abcdefghijklmnopqrstuvwxyz0123456789"
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	strFirst := (string(bytes))[0:1]
	if IsNumericOnly(strFirst) {
		return randomString(l)
	}
	return string(bytes)
}

func main() {
	rand.Seed(time.Now().UnixMilli())
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("請輸入產生domain長度: ")
	domain_int, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	domain_int1, err := strconv.Atoi(strings.Replace(domain_int, "\r\n", "", -1))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Print(domain_int1)
	fmt.Print("請輸入產生domain後綴 exp com,co : ")
	//var domain_str string
	//fmt.Scanln(&domain_str)
	reader1 := bufio.NewReader(os.Stdin)
	domain_str, err := reader1.ReadString('\r')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("請輸入產生數量:")
	reader2 := bufio.NewReader(os.Stdin)
	domain_count, err := reader2.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	domain_count1, err := strconv.Atoi(strings.Replace(domain_count, "\r\n", "", -1))
	if err != nil {
		fmt.Println(err)
	}
	//var domain_count int
	//fmt.Scanln(&domain_count)
	//domain_str1, err := strings.Replace(domain_str, "\r\n", "", -1)

	for i := 0; i < domain_count1; i++ {
		fmt.Println(randomString(domain_int1) + "." + domain_str)
	}
	fmt.Scanln()
	fmt.Println("done")
}
