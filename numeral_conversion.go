// Author: Nokib Sarkar <nokibsarkar@gmail.com>

package banglalocalizer

import (
	"strconv"
	"strings"
)

// banglaNumerals maps English digits to Bangla numerals.
var banglaNumerals = map[string]string{
	"0": "০", "1": "১", "2": "২", "3": "৩", "4": "৪",
	"5": "৫", "6": "৬", "7": "৭", "8": "৮", "9": "৯",
}

var banglaNumeralsList = []string{"০", "১", "২", "৩", "৪", "৫", "৬", "৭", "৮", "৯"}

// ConvertIntToNumerals converts an integer to a string of Bangla numerals.
//
// Example:
//
//	localizer := NewLocalizer()
//	numerals := localizer.ConvertIntToNumerals(12345) // "১২৩৪৫"
func (l *Localizer) ConvertIntToNumerals(number int) string {
	result := strings.Builder{}

	if number < 0 {
		result.WriteString("-")
		number = -number
	}

	digits := []int{}
	if number == 0 {
		digits = append(digits, 0)
	} else {
		for number > 0 {
			digits = append(digits, number%10)
			number /= 10
		}
	}

	result.Grow(len(digits))
	for i := len(digits) - 1; i >= 0; i-- {
		result.WriteString(banglaNumeralsList[digits[i]])
	}

	return result.String()
}

// ConvertFloatToNumerals converts a float64 number to a string of Bangla numerals.
// Preserves the decimal point in the output.
//
// Example:
//
//	localizer := NewLocalizer()
//	numerals := localizer.ConvertFloatToNumerals(123.45) // "১২৩.৪৫"
func (l *Localizer) ConvertFloatToNumerals(number float64) string {
	s := strconv.FormatFloat(number, 'f', -1, 64)
	return l.convertNumberStringToNumerals(s)
}

// convertNumberStringToNumerals converts a string representing a number (integer or float)
// to its Bangla numeral representation.
//
// Example:
//
//	localizer := NewLocalizer()
//	numerals := localizer.convertNumberStringToNumerals("-123.45") // "-১২৩.৪৫"
func (l *Localizer) convertNumberStringToNumerals(numberStr string) string {
	var result strings.Builder

	for _, char := range numberStr {
		digitStr := string(char)
		if banglaChar, ok := banglaNumerals[digitStr]; ok {
			result.WriteString(banglaChar)
		} else {
			// Keep non-digit characters like '-' and '.' as is
			result.WriteString(digitStr)
		}
	}

	return result.String()
}
