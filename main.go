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

func bech32HrpExpand(hrp string) []int {
	var ret []int
	for _, c := range hrp {
		ret = append(ret, int(c >> 5))
	}

	ret = append(ret, 0)

	for _, c := range hrp {
		ret = append(ret, int(c & 31))
	}

	return ret
}

// Convert bytes
// 8 to 5
// [00011111 00011111]
// [00000011 00011100 00001111 00000001]


func convertBits(data []int, fromBits, toBits uint, pad bool) ([]int, error) {
	acc := 0 // bit Accumulator of shifts
	bits := uint(0) // bit position counter
	ret := []int{} // return value
	maxValue := (1 << toBits) - 1
	for idx, value := range data {
		if value < 0 || (value>>fromBits) != 0 {
			return nil, fmt.Errorf("invalid data range : data[%d]=%d (frombits=%d)", idx, value, fromBits)
		}
		acc = (acc << fromBits) | value
		bits += fromBits
		for bits >= toBits {
			bits -= toBits
			ret = append(ret, (acc>>bits)&maxValue)
		}
	}
	if pad {
		if bits > 0 {
			ret = append(ret, (acc<<(toBits-bits))&maxValue)
		}
	} else if bits >= fromBits {
		return nil, fmt.Errorf("illegal zero padding")
	} else if ((acc << (toBits - bits)) & maxValue) != 0 {
		return nil, fmt.Errorf("non-zero padding")
	}
	return ret, nil
}


func main() {
	data := []int{0xff, 0xe8, 0xa1}

	base32Data, err := convertBits(data, 8, 5, true)

	hrp := "test"

	ret, err := Encode(hrp, base32Data)

	if err!= nil {
		panic(err)
	}

	fmt.Println(ret)
}