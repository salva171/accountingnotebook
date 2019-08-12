package common


import (
	"math/rand"
	"strconv"
	"os"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


func FindElementInArray(a *[]string,element string) bool {
	for _, v := range *a {
		if v == element {
			return true
		}
	}
	return false
}

func IsNumeric(s string) (bool,float64) {
	v, err := strconv.ParseFloat(s, 64)
	return err == nil,v
 }

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}


func ExistFile(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

