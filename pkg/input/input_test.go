package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveWithPattern(t *testing.T) {
	strWithoutSpace := RemoveWithPattern("9 87458-854214 741", "[^0-9]+")
	assert.Equal(t, "987458854214741", strWithoutSpace)
}

func TestAddPhoneMask(t *testing.T) {
	phoneWithTenDigitsMask := AddPhoneMask("2125987741")
	assert.Equal(t, "(21)2598-7741", phoneWithTenDigitsMask)

	phoneWithElevenDigitsMask := AddPhoneMask("11997454215")
	assert.Equal(t, "(11)99745-4215", phoneWithElevenDigitsMask)

	phoneWithNineDigitsMask2 := AddPhoneMask("119974542")
	assert.Equal(t, "119974542", phoneWithNineDigitsMask2)
}

func TestAddCurrencyMask(t *testing.T) {
	res := AddCurrencyMask(int64(-10))
	assert.Equal(t, "", res)

	res = AddCurrencyMask(int64(-1000))
	assert.Equal(t, "", res)

	res = AddCurrencyMask(int64(0))
	assert.Equal(t, "R$ 0,00", res)

	res = AddCurrencyMask(int64(1))
	assert.Equal(t, "R$ 0,01", res)

	res = AddCurrencyMask(int64(123))
	assert.Equal(t, "R$ 1,23", res)

	res = AddCurrencyMask(int64(5110))
	assert.Equal(t, "R$ 51,10", res)

	res = AddCurrencyMask(int64(99996))
	assert.Equal(t, "R$ 999,96", res)

	res = AddCurrencyMask(int64(199996))
	assert.Equal(t, "R$ 1.999,96", res)

	res = AddCurrencyMask(int64(9199996))
	assert.Equal(t, "R$ 91.999,96", res)

	res = AddCurrencyMask(int64(29199996))
	assert.Equal(t, "R$ 291.999,96", res)

	res = AddCurrencyMask(int64(829199996))
	assert.Equal(t, "R$ 8.291.999,96", res)

	res = AddCurrencyMask(int64(7829199996))
	assert.Equal(t, "R$ 78.291.999,96", res)

	res = AddCurrencyMask(int64(17829199996))
	assert.Equal(t, "R$ 178.291.999,96", res)

	res = AddCurrencyMask(int64(917829199996))
	assert.Equal(t, "R$ 9.178.291.999,96", res)
}
