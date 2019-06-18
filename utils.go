package cobranca

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PadType string

const (
	StrPadLeft  PadType = "str_pad_left"
	StrPadRight PadType = "str_pad_right"
)

var (
	transliterations = map[string]*regexp.Regexp{
		"A":  regexp.MustCompile(`À|Á|Â|Ã|Ä|à|á|â|ã|ä`),
		"AA": regexp.MustCompile(`Å|å`),
		"AE": regexp.MustCompile(`Æ|æ`),
		"C":  regexp.MustCompile(`Ç|ç`),
		"E":  regexp.MustCompile(`È|É|Ê|Ë|è|é|ê|ë`),
		"D":  regexp.MustCompile(`Ð|ð`),
		"I":  regexp.MustCompile(`Ì|Í|Î|Ï|ì|í|î|ï`),
		"L":  regexp.MustCompile(`Ł|ł`),
		"N":  regexp.MustCompile(`Ñ|ñ|ń`),
		"O":  regexp.MustCompile(`Ò|Ó|Ô|Õ|Ö|ò|ó|ô|õ|ö|ō`),
		"OE": regexp.MustCompile(`Œ|Ø|œ|ø`),
		"Th": regexp.MustCompile(`Þ`),
		"U":  regexp.MustCompile(`Ù|Ú|Û|Ü|ù|ú|û|ü|ũ|ū|ŭ|ů|ű|ų`),
		"Y":  regexp.MustCompile(`Ý|ý|ÿ`),
		"S":  regexp.MustCompile(`ś`),
		"SS": regexp.MustCompile(`ß`),
		"Z":  regexp.MustCompile(`ż`),
		"TH": regexp.MustCompile(`þ`),
	}
)

func Brancos(s string, l int) string {
	s = Sanitize(s)
	if len(s) > l {
		return s[0:l]
	}
	return StrPad(s, l, " ", StrPadRight)
}

func BrancosLeft(s string, l int) string {
	s = Sanitize(s)
	if len(s) > l {
		return s[len(s)-l:]
	}
	return StrPad(s, l, " ", StrPadLeft)
}

func OnlyNumbers(s string) string {
	reg, _ := regexp.Compile("[^0-9]+")
	return reg.ReplaceAllString(s, "")
}

// Sanitize replaces non-ASCII characters with an ASCII approximation, or if none exists, to “?”.
func Sanitize(word string) string {
	for repl, regex := range transliterations {
		word = regex.ReplaceAllString(word, repl)
	}

	var safe string
	for _, r := range word {
		if isAscii(r) {
			safe += string(r)
		} else {
			safe += "?"
		}
	}
	return strings.ToUpper(safe)
}

func isAscii(s rune) bool {
	return int(s) >= 32 && int(s) <= 126
}

func SemMascara(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, "/", "", -1)
	return s
}

func StrPad(s string, padLen int, padStr string, padType PadType) string {
	LenS := len(s)
	if LenS == padLen {
		return s
	}

	if LenS > padLen {
		if padType == StrPadLeft {
			return s[LenS-padLen:]
		} else {
			return s[0:padLen]
		}
	}

	c := (padLen - len(s)) / len(padStr)
	r := strings.Repeat(padStr, c)
	if padType == StrPadLeft {
		return r + s
	} else {
		return s + r
	}
}

func Zeros(s string, l int) string {
	if len(s) > l {
		return s[len(s)-l:]
	}
	f := "%0" + strconv.Itoa(l) + "s"
	return fmt.Sprintf(f, s)
}
