package days

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func D04() {
	const input = "iwrupvqb"
	result := 0

	for i := 1; result == 0; i++ {
		secret := input + strconv.Itoa(i)
		hash := md5Converter(secret)
		if strings.HasPrefix(hash, "00000") {
			result = i
		}
	}
	fmt.Println(result)
}

func md5Converter(input string) string {
	data := []byte(input)
	return fmt.Sprintf("%x", md5.Sum(data))
}
