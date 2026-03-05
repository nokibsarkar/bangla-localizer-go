// Author: Nokib Sarkar <nokibsarkar@gmail.com>
package banglalocalizer

import (
	"testing"
)

func TestConvertNumberStringToWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Very big numbers
		// Very small numbers (fractions)
		{"0.0000001", "শুন্য দশমিক এক"},
		{"0.0000000000001", "শুন্য দশমিক এক"},
		// Negative numbers
		{"-1", "-এক"},
		{"-123", "-একশত তেইশ"},
		{"-0.5", "-শুন্য দশমিক পাঁচ"},
		{"-1000000000000000", "-এক একশত লক্ষ কোটি"},
		// Basic and boundary cases
		{"0", "শুন্য"},
		{"00", "শুন্য"},
		{"0000", "শুন্য"},
		{"1", "এক"},
		{"9", "নয়"},
		{"10", "দশ"},
		{"99", "নিরানব্বই"},
		{"100", "একশত"},
		{"101", "একশত এক"},
		{"110", "একশত দশ"},
		{"123", "একশত তেইশ"},
		{"999", "নয়শত নিরানব্বই"},
		{"1000", "এক হাজার"},
		{"1234", "এক হাজার দুইশত চৌত্রিশ"},
		{"100000", "এক লক্ষ"},
		{"1000000", "দশ লক্ষ"},
		{"10000000", "এক কোটি"},
		{"12345678", "এক কোটি তেইশ লক্ষ পঁয়তাল্লিশ হাজার ছয়শত আটাত্তর"},
		{"100000000", "দশ কোটি"},
		{"999999999", "নিরানব্বই কোটি নিরানব্বই লক্ষ নিরানব্বই হাজার নয়শত নিরানব্বই"},
		// Large numbers
		{"1000000000", "এক হাজার কোটি"},
		{"1000000001", "এক হাজার কোটি এক"},
		{"1000000010", "এক হাজার কোটি দশ"},
		{"1000000100", "এক হাজার কোটি একশত"},
		{"1000010000", "এক হাজার কোটি দশ হাজার"},
		// Negative and invalid
		{"-1", "-এক"},
		{"abc", ""},
		{"12a3", ""},
		{"", ""},
		{" ", ""},
		// Floats and decimals
		{"0.0", "শুন্য দশমিক শুন্য"},
		{"0.5", "শুন্য দশমিক পাঁচ"},
		{"1.01", "এক দশমিক এক"},
		{"1.10", "এক দশমিক এক"},
		{"1.001", "এক দশমিক এক"},
		{"123.45", "একশত তেইশ দশমিক চার পাঁচ"},
		{"100.00", "একশত দশমিক শুন্য"},
		{"100.01", "একশত দশমিক এক"},
		{"100.10", "একশত দশমিক এক"},
		{"100.001", "একশত দশমিক এক"},
		// Leading/trailing spaces and zeros
		{"  123  ", "একশত তেইশ"},
		{"000123", "একশত তেইশ"},
		{"000.000", "শুন্য দশমিক শুন্য"},
		// Edge: only fraction
		{".5", "শুন্য দশমিক পাঁচ"},
		{".01", "শুন্য দশমিক এক"},
		{".00", "শুন্য দশমিক শুন্য"},
		// Edge: very large
		{"999999999999", "নিরানব্বই শত কোটি নয় হাজার কোটি নিরানব্বই কোটি নিরানব্বই লক্ষ নিরানব্বই হাজার নয়শত নিরানব্বই"},
	}
	l := &Localizer{}
	for i, test := range tests {
		result := l.convertNumberStringToWords(test.input)
		if result != test.expected {
			t.Errorf("Test %d: convertNumberStringToWords(%q) = %q; want %q", i, test.input, result, test.expected)
		}
	}

}

func TestConvertIntToNumerals(t *testing.T) {
	l := NewLocalizer()

	tests := []struct {
		input    int
		expected string
	}{
		{0, "০"},
		{1, "১"},
		{9, "৯"},
		{10, "১০"},
		{11, "১১"},
		{90, "৯০"},
		{909, "৯০৯"},
		{101, "১০১"},
		{1001, "১০০১"},
		{1000001, "১০০০০০১"},
		{111111111, "১১১১১১১১১"},
		{909090909, "৯০৯০৯০৯০৯"},
		{1234567890, "১২৩৪৫৬৭৮৯০"},
		{1000000000, "১০০০০০০০০০"},
		{2147483647, "২১৪৭৪৮৩৬৪৭"},
		{-1, "-১"},
		{-11, "-১১"},
		{-10, "-১০"},
		{-1010, "-১০১০"},
		{-1000001, "-১০০০০০১"},
		{-987654321, "-৯৮৭৬৫৪৩২১"},
		{-2147483648, "-২১৪৭৪৮৩৬৪৮"},
	}

	for i, test := range tests {
		result := l.ConvertIntToNumerals(test.input)
		if result != test.expected {
			t.Errorf("Test %d: ConvertIntToNumerals(%d) = %q; want %q", i, test.input, result, test.expected)
		}
	}
}

func TestConvertFloatToNumerals(t *testing.T) {
	l := NewLocalizer()

	tests := []struct {
		input    float64
		expected string
	}{
		{0.0, "০"}, // Go trims trailing .0 with FormatFloat(..., -1, ...)
		{0.1, "০.১"},
		{0.5, "০.৫"},
		{0.01, "০.০১"},
		{0.001, "০.০০১"},
		{1.5, "১.৫"},
		{2.75, "২.৭৫"},
		{12.34, "১২.৩৪"},
		{10.25, "১০.২৫"},
		{123.456, "১২৩.৪৫৬"},
		{100.00, "১০০"},
		{1.2300, "১.২৩"},
		{1000.0001, "১০০০.০০০১"},
		{0.0000001, "০.০০০০০০১"},
		{1000000.75, "১০০০০০০.৭৫"},
		{-1.5, "-১.৫"},
		{-12.34, "-১২.৩৪"},
		{-0.5, "-০.৫"},
	}

	for i, test := range tests {
		result := l.ConvertFloatToNumerals(test.input)
		if result != test.expected {
			t.Errorf("Test %d: ConvertFloatToNumerals(%v) = %q; want %q", i, test.input, result, test.expected)
		}
	}
}

func TestConvertNumberStringToNumerals(t *testing.T) {
	l := NewLocalizer()

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"0", "০"},
		{"0123456789", "০১২৩৪৫৬৭৮৯"},
		{"000001", "০০০০০১"},
		{"999999999999999", "৯৯৯৯৯৯৯৯৯৯৯৯৯৯৯"},
		{"123", "১২৩"},
		{"-123", "-১২৩"},
		{"+42", "+৪২"},
		{"++42", "++৪২"},
		{"--42", "--৪২"},
		{".75", ".৭৫"},
		{"-.75", "-.৭৫"},
		{"+.75", "+.৭৫"},
		{"000.000", "০০০.০০০"},
		{"00123.004500", "০০১২৩.০০৪৫০০"},
		{"-0123.4500", "-০১২৩.৪৫০০"},
		{" 001 ", " ০০১ "},
		{"\t123\n", "\t১২৩\n"},
		{"1,234,567.89", "১,২৩৪,৫৬৭.৮৯"},
		{"1_234_567.89", "১_২৩৪_৫৬৭.৮৯"},
		{"12a3", "১২a৩"},
		{"Room 101", "Room ১০১"},
		{"ID#2026", "ID#২০২৬"},
		{"version2.0.1", "version২.০.১"},
		{"http://v1/api/2", "http://v১/api/২"},
		{"abc", "abc"},
		{"NaN", "NaN"},
		{"+Inf", "+Inf"},
		{"-Inf", "-Inf"},
		{"০১২৩৪৫৬৭৮৯", "০১২৩৪৫৬৭৮৯"},
		{"−123", "−১২৩"}, // unicode minus (U+2212)
		{"–123", "–১২৩"}, // en dash (U+2013)
		{"—123", "—১২৩"}, // em dash (U+2014)
		{"(123)", "(১২৩)"},
		{"$123.50", "$১২৩.৫০"},
		{"12%", "১২%"},
	}

	for i, test := range tests {
		result := l.convertNumberStringToNumerals(test.input)
		if result != test.expected {
			t.Errorf("Test %d: convertNumberStringToNumerals(%q) = %q; want %q", i, test.input, result, test.expected)
		}
	}
}
