package base62

import (
	"math"
	"strings"
)

const base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const b = 62

func ToBase62(num int) string {
	r := num % b           //il resto della divisione intera
	res := string(base[r]) // è l'indice del carattere in base62
	div := num / b         //il risultato della divisione intera
	q := int(math.Floor(float64(div)))
	for q != 0 {
		r = q % b
		temp := q / b
		q = int(math.Floor(float64(temp)))
		res = string(base[r]) + res

	}

	return string(res)
}

func ToBase10(str string) int {
	res := 0
	for _, r := range str {
		res = (b * res) + strings.Index(base, string(r))
	}
	return res
}
