package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func sumDecipher(code, key []byte) int {
	res := 0
	for i := range code {
		res += int(code[i] ^ key[i%3])
	}
	return res
}

func getCode(bytes []byte) []byte {
	strings := strings.Split(string(bytes), ",")
	var ret []byte
	for _, str := range strings {
		num, _ := strconv.Atoi(str)
		ret = append(ret, byte(num))
	}
	return ret
}

func decipher(code, key []byte) string {
	res := ""
	for i := range code {
		res += string(code[i] ^ key[i%3])
	}
	return res
}

func main() {
	f, _ := os.Open("p059_cipher.txt")
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	bytes = bytes[0 : len(bytes)-1] //newline
	code := getCode(bytes)
	for a := 97; a <= 122; a++ {
		for b := 97; b <= 122; b++ {
			for c := 97; c <= 122; c++ {
				key := []byte{byte(a), byte(b), byte(c)}
				decoded := decipher(code, key)
				if strings.Contains(decoded, "the") && strings.Contains(decoded, "in") && strings.Contains(decoded, "and") {
					fmt.Println(decoded)
					fmt.Println(sumDecipher(code, key))
				}
			}
		}
	}

}
