package days

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func D04() {
	const input = "iwrupvqb"
	result1 := 0
	result2 := 0

	for i := 1; result1 == 0 || result2 == 0; i++ {
		secret := input + strconv.Itoa(i)
		hash := md5Converter(secret)
		if strings.HasPrefix(hash, "00000") && result1 == 0 {
			result1 = i
		}
		if strings.HasPrefix(hash, "000000") && result2 == 0 {
			result2 = i
		}
	}
	fmt.Println(result1, result2)
}

func md5Converter(input string) string {
	data := []byte(input)
	return fmt.Sprintf("%x", md5.Sum(data))
}
