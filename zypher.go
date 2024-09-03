package zypher

import (
	"errors"
	"regexp"
	"sync"
)

type Zypher struct {
	Shift       int  // number of runes/digits to shift, defaults to 3
	IterCount   int  // number of iterations to be applied to value being zyphered, defaults to 3
	Alternate   bool // if true when odd elements will reverse shift, defaults false
	IgnoreSpace bool // if true will ignore space and leave them in, defaults to false
}

func DefaultZops() *Zypher {
	return &Zypher{
		Shift:       3,
		Alternate:   false,
		IterCount:   3,
		IgnoreSpace: false,
	}
}

func NewZypher(ops ...func(*Zypher)) *Zypher {
	zy := DefaultZops()
	for _, op := range ops {
		op(zy)
	}
	return zy
}

func WithShift(shiftCount int) func(*Zypher) {
	return func(z *Zypher) {
		z.Shift = shiftCount
	}
}

func WithAlternate(wrp bool) func(*Zypher) {
	return func(z *Zypher) {
		z.Alternate = wrp
	}
}

func WithIterCount(count int) func(*Zypher) {
	return func(z *Zypher) {
		z.IterCount = count
	}
}

func WithIgnoreSpace(i bool) func(*Zypher) {
	return func(z *Zypher) {
		z.IgnoreSpace = i
	}
}

func (z Zypher) AsciZyph(arg string) (string, error) {
	isValidString := regexp.MustCompile(`^[a-zA-Z0-9 ]+$`).MatchString(arg)

	if !isValidString {
		InvalidString := errors.New("only strings containing numbers, letters, and spaces are allowd")
		return "", InvalidString
	}

	result := make([]rune, len(arg))
	var wg sync.WaitGroup

	for i := 0; i < z.IterCount; i++ {
		for indx, r := range arg {
			wg.Add(1)
			go func() {
				defer wg.Done()
				alt := false
				if z.Alternate {
					alt = indx%2 != 0
				}
				asciShift(&result[indx], r, z.Shift, alt, z.IgnoreSpace)
			}()
		}
	}
	wg.Wait()
	return string(result), nil
}

func asciShift(target *rune, r rune, shf int, alt bool, ignSpc bool) {
	switch alt {
	case true:
		if shf > 0 {
			shf = -shf
		} else {
			shf = +shf
		}
	}

	if r >= '0' && r <= '9' {
		*target = '0' + (r-'0'+rune(shf)+10)%10
	} else if r >= 'A' && r <= 'Z' {
		*target = 'A' + (r-'A'+rune(shf)+26)%26

	} else if r >= 'a' && r <= 'z' {
		*target = 'a' + (r-'a'+rune(shf)+26)%26
	} else if r == ' ' {
		if !ignSpc {
			*target = 'x'
		} else {
			*target = ' '
		}
	}
}
