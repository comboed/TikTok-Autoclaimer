package main

import (
	"encoding/hex"
	"crypto/md5"
	"math/rand"
	"math/big"
	"strconv"
	"strings"
	"time"
	"fmt"
)

func convertHex(str string) *big.Int {
	var hexData = big.NewInt(0)
	for _, strip := range strings.ToUpper(str) {
		if strip <= 57 {
			hexData.Lsh(hexData, 4)
			hexData.Add(hexData, big.NewInt(int64(strip - 48)))
		} else if 65 <= strip && strip <= 70 {
			hexData.Lsh(hexData, 4)
			hexData.Add(hexData, big.NewInt(int64(strip - 65 + 10)))
		}
	}
	return hexData
}

func convertBytes(str string) []int64 {
	var slice []int64
	for i := 0; i < len(str); i += 2 {
		var a string = string(str[i])
		var b string = string(str[1 + i])
		var c *big.Int = big.NewInt(0)
		c.Add(c.Lsh(convertHex(a), 4), convertHex(b))
		slice = append(slice, c.Int64())
	}
	return slice
}

func generateMD5Hash(str string) string {
	var hash [16]byte = md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func generateSigHash(params string) string {
	var MD5 string = "00000000000000000000000000000000"
	var hash string = generateMD5Hash(params)
	hash += strings.Repeat(MD5, 3)
	return hash
}

func initiateXGorgon(time int64, bytes []int64) []string {
	var xGorgon []string
	for i := 0; i < 4; i++ {
		var temp string = strconv.FormatInt(bytes[i], 16)
		xGorgon = append(xGorgon, temp)
	}
	for i := 0; i < 4; i++ {
		xGorgon = append(xGorgon, "0")
	}
	for i := 0; i < 4; i++ {
		xGorgon = append(xGorgon, strconv.FormatInt(bytes[i + 32], 16))
	}
	for i := 0; i < 4; i++ {
		xGorgon = append(xGorgon, "0")
	}
	for i := 0; i < 4; i++ {
		xGorgon = append(xGorgon, strconv.FormatInt(time, 16)[i * 2:2 * i + 2])
	}
	return xGorgon
}

func convertXGorgon(xGorgon []string) []string {
	var byteTable = strings.Split("D6 28 3B 71 70 76 BE 1B A4 FE 19 57 5E 6C BC 21 B2 14 37 7D 8C A2 FA 67 55 6A 95 E3 FA 67 78 ED 8E 55 33 89 A8 CE 36 B3 5C D6 B2 6F 96 C4 34 B9 6A EC 34 95 C4 FA 72 FF B8 42 8D FB EC 70 F0 85 46 D8 B2 A1 E0 CE AE 4B 7D AE A4 87 CE E3 AC 51 55 C4 36 AD FC C4 EA 97 70 6A 85 37 6A C8 68 FA FE B0 33 B9 67 7E CE E3 CC 86 D6 9F 76 74 89 E9 DA 9C 78 C5 95 AA B0 34 B3 F2 7D B2 A2 ED E0 B5 B6 88 95 D1 51 D6 9E 7D D1 C8 F9 B7 70 CC 9C B6 92 C5 FA DD 9F 28 DA C7 E0 CA 95 B2 DA 34 97 CE 74 FA 37 E9 7D C4 A2 37 FB FA F1 CF AA 89 7D 55 AE 87 BC F5 E9 6A C4 68 C7 FA 76 85 14 D0 D0 E5 CE FF 19 D6 E5 D6 CC F1 F4 6C E9 E7 89 B2 B7 AE 28 89 BE 5E DC 87 6C F7 51 F2 67 78 AE B3 4B A2 B3 21 3B 55 F8 B3 76 B2 CF B3 B3 FF B3 5E 71 7D FA FC FF A8 7D FE D8 9C 1B C4 6A F9 88 B5 E5", " ")
	var xGorgonBase int64 = 254
	for i := 0; i < len(xGorgon); i++ {
		var hexData int64
		if i == 0 {
			var tempData, _ = strconv.ParseInt(byteTable[0], 16, 64)
			hexData, _ = strconv.ParseInt(byteTable[tempData-1], 16, 64)
			byteTable[i] = strconv.FormatInt(hexData, 16)
		} else if i == 1 {
			hexData, _ = strconv.ParseInt(byteTable[253], 16, 64)
			byteTable[i] = strconv.FormatInt(hexData, 16)
		} else {
			var hexByteTMP, _ = strconv.ParseInt(byteTable[i], 16, 64)
			xGorgonBase += hexByteTMP
			if xGorgonBase > 256 {
				xGorgonBase -= 256
			}
			hexData, _ = strconv.ParseInt(byteTable[xGorgonBase-1], 16, 64)
			byteTable[i] = strconv.FormatInt(hexData, 16)
		}
		if hexData*2 > 256 {
			hexData = hexData*2 - 256
		} else {
			hexData = hexData * 2
		}
		var data1, _ = strconv.ParseInt(byteTable[hexData-1], 16, 64)
		var data2, _ = strconv.ParseInt(xGorgon[i], 16, 64)
		xGorgon[i] = strconv.FormatInt(data1^data2, 16)
	}
	return xGorgon
}

func cleanXGorgon(xGorgon []string) []string {
	for i := 0; i < len(xGorgon); i++ {
		var baseByte string = xGorgon[i]
		if len(baseByte) < 2 {
			baseByte += strings.Repeat("0", 1)
		} else {
			baseByte = string(xGorgon[i][1]) + string(xGorgon[i][0])
		}
		if i < len(xGorgon) - 1 {
			var hexTMP, _ = strconv.ParseInt(xGorgon[i+1], 16, 64)
			var hexTMP2, _ = strconv.ParseInt(baseByte, 16, 64)
			baseByte = strconv.FormatInt(hexTMP^hexTMP2, 16)
		} else {
			var hexTemp, _ = strconv.ParseInt(baseByte, 16, 64)
			var hexTemp2, _ = strconv.ParseInt(xGorgon[0], 16, 64)
			baseByte = strconv.FormatInt(hexTemp^hexTemp2, 16)
		}
		var absString, _ = strconv.ParseInt(baseByte, 16, 64)
		absString = absString & 170 / 2

		var byteTMP, _ = strconv.ParseInt(baseByte, 16, 64)
		var byteHandle = 85&byteTMP*2 | absString

		byteHandle = byteHandle&51*4 | byteHandle&204/4
		var reverseByte string = strconv.FormatInt(byteHandle, 16)
		if len(reverseByte) > 1 {
			reverseByte = reverseString(reverseByte)
		} else {
			reverseByte += strings.Repeat("0", 1)
		}
		var byteFinal, _ = strconv.ParseInt(reverseByte, 16, 64)
		byteFinal = byteFinal ^ 255 ^ 20

		xGorgon[i] = strconv.FormatInt(byteFinal, 16)
	}
	return xGorgon
}

func generateXGorgon(xKhronos int64, bytes []int64) string {
	var xGorgonSlice []string = []string{"3", "61", "41", "10", "80", "0"}
	var xGorgonString string

	var data []string = initiateXGorgon(xKhronos, bytes)
	data = convertXGorgon(data)
	data = cleanXGorgon(data)

	for i := 0; i < len(data); i++ {
		xGorgonSlice = append(xGorgonSlice, data[i])
	}
	for i := 0; i < len(xGorgonSlice); i++ {
		var char string = xGorgonSlice[i] + ""
		if len(char) > 1 {
			xGorgonString += char
		} else {
			xGorgonString += "0"
			xGorgonString += char
		}
	}
	return xGorgonString
}

func generateSession() string {
    rand.Seed(time.Now().UnixNano())
    var bytes []byte = make([]byte, 32)
    rand.Read(bytes)
    return fmt.Sprintf("sessionid=%x", bytes)[:32]
}