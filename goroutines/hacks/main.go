package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Min returns the min
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max returns the max
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// BigInt is the generic BigInt struct
type BigInt struct {
	// the array data
	d []uint
	// the base
	b uint
}

// Multiply multiplies to BigInts
func (x BigInt) Multiply(y BigInt) (BigInt, error) {
	if x.b != y.b {
		return BigInt{}, fmt.Errorf("bases of x and y are not the same: %d, %d", x.b, y.b)
	}

	total := uint(0)

	p, q := len(x.d), len(y.d)
	base := x.b

	prod := make([]uint, p+q)

	for i := 0; i < p+q-1; i++ {
		jMin := Max(0, i-p+1)
		jMax := Min(i, q-1)

		for j := jMin; j <= jMax; j++ {
			total += (x.d[i-j] * y.d[j])
		}

		prod[i] = total % base
		total /= base
	}

	prod[p+q-1] = total % base

	return BigInt{prod, base}, nil
}

// String returns the string version of the number
func (x BigInt) String() string {
	s := ""

	for i := len(x.d) - 1; i >= 0; i-- {
		s += strconv.Itoa(int(x.d[i]))
	}

	return strings.TrimLeftFunc(s, func(r rune) bool {
		if r == '0' {
			return true
		}

		return false
	})
}

// InitBigInt inits a BigInt from a string
func InitBigInt(s string, b uint) (BigInt, error) {
	if b <= 1 {
		return BigInt{}, fmt.Errorf("base is not positive: %v", b)
	}

	a := make([]uint, len(s))

	for i := 0; i < len(a); i++ {
		x, err := strconv.Atoi(fmt.Sprintf("%c", s[i]))

		if err != nil {
			return BigInt{}, err
		}

		a[len(a)-i-1] = uint(x)
	}

	return BigInt{a, uint(b)}, nil
}

func factorial(n int) (BigInt, error) {
	base := uint(10)
	x := BigInt{[]uint{1}, base}

	for i := 2; i <= n; i++ {
		y, err := InitBigInt(strconv.Itoa(i), base)
		if err != nil {
			return BigInt{}, err
		}

		z, err := x.Multiply(y)
		if err != nil {
			return BigInt{}, err
		}

		x = z
	}

	return x, nil
}

func digits(n int) int64 {
	if n < 0 {
		return int64(0)
	}

	if n <= 1 {
		return int64(1)
	}

	var digits int64
	digits = 0

	for i := 2; i <= n; i++ {
		digits += int64(math.Log10(float64(i)))
	}

	return int64(math.Floor(float64(digits))) + 1
}

func main() {
	file, err := os.Open("in0")
	check(err)

	r := bufio.NewReader(file)

	l, _, err := r.ReadLine()
	check(err)

	n, err := strconv.Atoi(string(l))
	check(err)

	var sum int64
	sum = 0

	for i := 0; i < n; i++ {
		t, _, err := r.ReadLine()
		check(err)

		d, err := strconv.Atoi(string(t))
		check(err)

		f, err := factorial(d)
		check(err)

		// fmt.Println(d, int64(len(f.String())))
		sum += int64(len(f.String()))
	}

	fmt.Println("Phase 0:", sum)

	file.Close()

	file, err = os.Open("in1")
	check(err)

	r = bufio.NewReader(file)

	l, _, err = r.ReadLine()
	check(err)

	n, err = strconv.Atoi(string(l))
	check(err)

	str := ""

	for i := 0; i < n; i++ {
		t, _, err := r.ReadLine()
		check(err)

		tt := strings.Split(string(t), " ")

		d := make([]int, len(tt))

		for j := range tt {
			tmp, err := strconv.Atoi(tt[j])
			check(err)

			d[j] = tmp
		}

		f, err := factorial(d[0])
		check(err)

		fs := f.String()

		start := len(fs) - d[1] + 1
		end := len(fs) - d[2]

		// fmt.Println(i, d, len(fs), start-end, fs, fs[end:start])
		str += fs[end:start]
	}

	file.Close()

	fmt.Println("Phase 1:", str)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}