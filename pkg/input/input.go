package input

import (
	"bytes"
	"log"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

const (
	PhoneMinLength = 10
	PhoneMaxLength = 11
)

func RemoveWithPattern(value, pattern string) string {
	tmpvalue := strings.ReplaceAll(value, " ", "")
	if len(tmpvalue) == 0 {
		return tmpvalue
	}
	return regexp.MustCompile(pattern).ReplaceAllString(tmpvalue, "")
}

func AddPhoneMask(value string) string {
	if len(value) != PhoneMinLength && len(value) != PhoneMaxLength {
		return value
	}
	var buff bytes.Buffer
	if len(value) == PhoneMinLength {
		buff.WriteString("(")
		buff.WriteString(value[0:2])
		buff.WriteString(")")
		buff.WriteString(value[2:6])
		buff.WriteString("-")
		buff.WriteString(value[6:])
	} else {
		buff.WriteString("(")
		buff.WriteString(value[0:2])
		buff.WriteString(")")
		buff.WriteString(value[2:7])
		buff.WriteString("-")
		buff.WriteString(value[7:])
	}
	return buff.String()
}

func AddCurrencyMask(valueInCents int64) string {
	if valueInCents < 0 {
		return ""
	}

	var buff bytes.Buffer
	buff.WriteString("R$ ")
	valueAsStr := strconv.FormatInt(valueInCents, 10)

	if valueInCents < 100 {
		if valueInCents > 9 && valueInCents < 100 {
			buff.WriteString("0,")
			buff.WriteString(valueAsStr)
			return buff.String()
		}
		buff.WriteString("0,0")
		buff.WriteString(valueAsStr)
		return buff.String()
	}
	buff.Reset()
	valueCents := valueAsStr[len(valueAsStr)-2:]
	buff.WriteString(valueAsStr[0 : len(valueAsStr)-2])
	buff.WriteString(".")
	buff.WriteString(valueCents)

	valueAsFloat, e := strconv.ParseFloat(buff.String(), 64)
	if e != nil {
		log.Default().Printf("Erro when parse number %v\n", e)
	}

	lang := language.MustParse("pt_BR")
	cur, _ := currency.FromTag(lang)
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(valueAsFloat, number.Scale(scale))
	p := message.NewPrinter(lang)

	return p.Sprintf("%v %v", currency.Symbol(cur), dec)
}
