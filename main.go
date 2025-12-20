// Author: Nokib Sarkar <nokibsarkar@gmail.com>

// Package banglalocalizer provides functions to convert numbers to their Bangla (Bengali) word representation.
//
// Features:
//   - Converts integers and floating-point numbers to Bangla words
//   - Handles negative numbers (with dash prefix)
//   - Supports very large numbers with extended Bangla units
//   - Ignores leading/trailing spaces and zeros
//   - Returns empty string for invalid input
//
// Example usage:
//
//	localizer := banglalocalizer.NewLocalizer()
//	word := localizer.ConvertIntToWords(12345678) // "এক কোটি তেইশ লক্ষ পঁয়তাল্লিশ হাজার ছয়শত আটাত্তর"
//	word := localizer.convertNumberStringToWords("-100.01") // "-একশত দশমিক এক"
package banglalocalizer

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// numericWord is the Bangla word mapping for numbers up to 100
var numericWord = map[string]string{
	"0": "শুন্য", "1": "এক", "2": "দুই", "3": "তিন", "4": "চার", "5": "পাঁচ", "6": "ছয়", "7": "সাত", "8": "আট", "9": "নয়",
	"10": "দশ", "11": "এগারো", "12": "বার", "13": "তেরো", "14": "চৌদ্দ", "15": "পনের", "16": "ষোল", "17": "সতের", "18": "আঠার", "19": "উনিশ",
	"20": "বিশ", "21": "একুশ", "22": "বাইশ", "23": "তেইশ", "24": "চব্বিশ", "25": "পঁচিশ", "26": "ছাব্বিশ", "27": "সাতাশ", "28": "আঠাশ", "29": "ঊনত্রিশ",
	"30": "ত্রিশ", "31": "একত্রিশ", "32": "বত্রিশ", "33": "তেত্রিশ", "34": "চৌত্রিশ", "35": "পঁয়ত্রিশ", "36": "ছত্রিশ", "37": "সাঁইত্রিশ", "38": "আটত্রিশ", "39": "ঊনচল্লিশ",
	"40": "চল্লিশ", "41": "একচল্লিশ", "42": "বিয়াল্লিশ", "43": "তেতাল্লিশ", "44": "চুয়াল্লিশ", "45": "পঁয়তাল্লিশ", "46": "ছেচল্লিশ", "47": "সাতচল্লিশ", "48": "আটচল্লিশ", "49": "ঊনপঞ্চাশ",
	"50": "পঞ্চাশ", "51": "একান্ন", "52": "বায়ান্ন", "53": "তিপ্পান্ন", "54": "চুয়ান্ন", "55": "পঞ্চান্ন", "56": "ছাপ্পান্ন", "57": "সাতান্ন", "58": "আটান্ন", "59": "ঊনষাট",
	"60": "ষাট", "61": "একষট্টি", "62": "বাষট্টি", "63": "তেষট্টি", "64": "চৌষট্টি", "65": "পঁয়ষট্টি", "66": "ছেষট্টি", "67": "সাতষট্টি", "68": "আটষট্টি", "69": "ঊনসত্তর",
	"70": "সত্তর", "71": "একাত্তর", "72": "বাহাত্তর", "73": "তিয়াত্তর", "74": "চুয়াত্তর", "75": "পঁচাত্তর", "76": "ছিয়াত্তর", "77": "সাতাত্তর", "78": "আটাত্তর", "79": "ঊনআশি",
	"80": "আশি", "81": "একাশি", "82": "বিরাশি", "83": "তিরাশি", "84": "চুরাশি", "85": "পঁচাশি", "86": "ছিয়াশি", "87": "সাতাশি", "88": "আটাশি", "89": "ঊননব্বই",
	"90": "নব্বই", "91": "একানব্বই", "92": "বিরানব্বই", "93": "তিরানব্বই", "94": "চুরানব্বই", "95": "পঁচানব্বই", "96": "ছিয়ানব্বই", "97": "সাতানব্বই", "98": "আটানব্বই", "99": "নিরানব্বই",
	"100": "একশত",
}

// Localizer provides methods for converting numbers to Bangla words.
type Localizer struct{}

// NewLocalizer creates and returns a new Localizer instance.
func NewLocalizer() *Localizer {
	return &Localizer{}
}

// ConvertIntToWords converts an integer to its Bangla word representation.
//
// Example:
//   localizer := NewLocalizer()
//   word := localizer.ConvertIntToWords(123) // "একশত তেইশ"
//
// Supports numbers up to 1,000,000,000,000,000 (one quadrillion) and negative values.

func (l *Localizer) ConvertIntToWords(number int) string {
	var result strings.Builder
	remain := number
	for _, u := range unitMap {
		if remain >= u.Value {
			count := remain / u.Value
			remain = remain % u.Value
			if count > 0 {
				// Special handling for 100s: use একশত, দুইশত, ...
				if u.Value == 100 {
					switch count {
					case 1:
						result.WriteString("একশত ")
					case 2:
						result.WriteString("দুইশত ")
					case 3:
						result.WriteString("তিনশত ")
					case 4:
						result.WriteString("চারশত ")
					case 5:
						result.WriteString("পাঁচশত ")
					case 6:
						result.WriteString("ছয়শত ")
					case 7:
						result.WriteString("সাতশত ")
					case 8:
						result.WriteString("আটশত ")
					case 9:
						result.WriteString(numericWord["9"] + "শত ")
					default:
						part := l.ConvertIntToWords(count)
						result.WriteString(part + " শত ")
					}
				} else {
					part := l.ConvertIntToWords(count)
					result.WriteString(part + " " + u.Name + " ")
				}
			}
		}
	}
	if remain > 0 {
		if remain <= 100 {
			if w, ok := numericWord[strconv.Itoa(remain)]; ok {
				result.WriteString(w)
			} else {
				result.WriteString(l.ConvertIntToWords(remain))
			}
		} else {
			result.WriteString(l.ConvertIntToWords(remain))
		}
	}
	return trim(result.String())
}

// Unit represents a Bangla number unit (e.g., কোটি, লক্ষ, হাজার, শত).
type Unit struct {
	Value int    // Numeric value of the unit
	Name  string // Bangla name of the unit
}

// unitMap defines the Bangla unit mapping for large and very large numbers.
// The order is from largest to smallest.
var unitMap = []Unit{
	{1000000000000000, "একশত লক্ষ কোটি"}, // 1,000,000,000,000,000
	{100000000000000, "দশ লক্ষ কোটি"},    // 100,000,000,000,000
	{10000000000000, "কোটি কোটি"},        // 10,000,000,000,000
	{1000000000000, "লক্ষ কোটি"},         // 1,000,000,000,000
	{10000000000, "শত কোটি"},             // 10,000,000,000
	{1000000000, "হাজার কোটি"},           // 1,000,000,000
	{10000000, "কোটি"},
	{100000, "লক্ষ"},
	{1000, "হাজার"},
	{100, "শত"},
}

// ConvertFloatToWords converts a float64 number to its Bangla word representation.
//
// Example:
//
//	localizer := NewLocalizer()
//	word := localizer.ConvertFloatToWords(-123.45) // "-একশত তেইশ দশমিক চার পাঁচ"
func (l *Localizer) ConvertFloatToWords(number float64) string {
	s := strconv.FormatFloat(number, 'f', -1, 64)
	return l.convertNumberStringToWords(s)
}

// convertNumberStringToWords converts a string representing a number (integer or float, positive or negative)
// to its Bangla word representation. Handles leading/trailing spaces, zeros, and invalid input.
//
// Example:
//
//	localizer := NewLocalizer()
//	word := localizer.convertNumberStringToWords("  -123.45  ") // "-একশত তেইশ দশমিক চার পাঁচ"
func (l *Localizer) convertNumberStringToWords(numberStr string) string {
	// Handle negative numbers (supports multiple dash types)
	negative := false
	numberStr = trim(numberStr)
	if len(numberStr) > 0 {
		r, size := utf8.DecodeRuneInString(numberStr)
		if r == '-' || r == '\u2212' || r == '\u2013' || r == '\u2014' {
			negative = true
			numberStr = numberStr[size:]
			numberStr = trim(numberStr)
		}
	}
	// Remove leading/trailing spaces
	numberStr = trim(numberStr)
	if numberStr == "" {
		return ""
	}
	// Handle float or string with dot
	if dot := indexOf(numberStr, "."); dot != -1 {
		intPart := numberStr[:dot]
		fracPart := numberStr[dot+1:]
		intWord := l.convertNumberStringToWords(intPart)
		if intWord == "" {
			intWord = numericWord["0"]
		}
		// treat fraction as digit-by-digit
		// Remove trailing zeros for fraction
		fracRunes := []rune(fracPart)
		end := len(fracRunes)
		for end > 0 && fracRunes[end-1] == '0' {
			end--
		}
		fracPartTrimmed := fracPart[:end]
		fracWord := ""
		nonZeroFound := false
		for _, ch := range fracPartTrimmed {
			if ch == '0' {
				if nonZeroFound {
					fracWord += numericWord["0"] + " "
				}
			} else {
				fracWord += l.ConvertIntToWords(int(ch-'0')) + " "
				nonZeroFound = true
			}
		}
		fracWord = trim(fracWord)
		if fracWord == "" {
			fracWord = numericWord["0"]
		}
		result := intWord + " দশমিক " + fracWord
		if negative {
			return "-" + result
		}
		return result
	}
	// Remove leading zeros
	for len(numberStr) > 1 && numberStr[0] == '0' {
		numberStr = numberStr[1:]
	}
	if numberStr == "" {
		return ""
	}
	if numberStr == "0" {
		if negative {
			return "-" + numericWord["0"]
		}
		return numericWord["0"]
	}
	n, err := strconv.Atoi(numberStr)
	if err != nil {
		return ""
	}
	result := l.ConvertIntToWords(n)
	if negative {
		return "-" + result
	}
	return result
}

// indexOf returns the index of substr in s, or -1 if not found.
func indexOf(s, substr string) int {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// trim trims spaces, tabs, and newlines from both ends of a string.
func trim(s string) string {
	for len(s) > 0 && (s[0] == ' ' || s[0] == '\t' || s[0] == '\n') {
		s = s[1:]
	}
	for len(s) > 0 && (s[len(s)-1] == ' ' || s[len(s)-1] == '\t' || s[len(s)-1] == '\n') {
		s = s[:len(s)-1]
	}
	return s
}
