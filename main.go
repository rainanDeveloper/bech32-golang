package main

import (
	"bytes"
	"fmt"
)

var charset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

func Encode(hrp string, data []int) (string, error) {

	var ret bytes.Buffer

	ret.WriteString(hrp)
	ret.WriteString("1")

	for idx, p := range data {
		if p < 0 || p >= len(charset) {
			return "", fmt.Errorf("invalid data: data[%d]=%d", idx, p)
		}

		ret.WriteByte(charset[p])
	}

	return ret.String(), nil
}

func convertHexToData(str string) []int {
	var ret []int
	for _, c := range str {
		ret = append(ret, int(c >> 5))
	}

	ret = append(ret, 0)

	for _, c := range str {
		ret = append(ret, int(c & 31))
	}

	return ret
}

func main() {
	name := "f1f1f1f1"

	data := convertHexToData(name)

	encoded, err := Encode("data", data)

	if err!= nil {
		panic(err)
	}

	fmt.Println(encoded)
}