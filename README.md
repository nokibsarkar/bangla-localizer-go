# Bangla Localizer Go

Bangla Localizer Go is a lightweight Go module for converting numbers to:

- **Bangla words** (e.g., `123` → `একশত তেইশ`)
- **Bangla numerals** (e.g., `123` → `১২৩`)

It supports integers, floats, negatives, and large values.

## Features

- Convert `int` and `float64` to **Bangla words**
- Convert `int` and `float64` to **Bangla numerals** (`০১২৩৪৫৬৭৮৯`)
- Handles negative values
- Handles very large numbers with Bangla units (`লক্ষ`, `কোটি`, etc.)
- String-based internal parsing handles spaces/zeros for word conversion

## Installation

```bash
go get github.com/nokibsarkar/bangla-localizer-go
```

## Quick Start

```go
package main

import (
	"fmt"

	banglalocalizer "github.com/nokibsarkar/bangla-localizer-go"
)

func main() {
	l := banglalocalizer.NewLocalizer()

	// Bangla words
	fmt.Println(l.ConvertIntToWords(123))         // একশত তেইশ
	fmt.Println(l.ConvertFloatToWords(-123.45))   // -একশত তেইশ দশমিক চার পাঁচ

	// Bangla numerals
	fmt.Println(l.ConvertIntToNumerals(1234567890)) // ১২৩৪৫৬৭৮৯০
	fmt.Println(l.ConvertFloatToNumerals(10.25))    // ১০.২৫
}
```

## API

- `NewLocalizer() *Localizer`
- `ConvertIntToWords(number int) string`
- `ConvertFloatToWords(number float64) string`
- `ConvertIntToNumerals(number int) string`
- `ConvertFloatToNumerals(number float64) string`

## Examples

### Words

| Input | Output |
|---|---|
| `123` | `একশত তেইশ` |
| `-123` | `-একশত তেইশ` |
| `0.5` | `শুন্য দশমিক পাঁচ` |
| `1000000000` | `এক হাজার কোটি` |

### Numerals

| Input | Output |
|---|---|
| `0` | `০` |
| `1234567890` | `১২৩৪৫৬৭৮৯০` |
| `-987654321` | `-৯৮৭৬৫৪৩২১` |
| `10.25` | `১০.২৫` |
| `0.0000001` | `০.০০০০০০১` |

## Behavior Notes

- `ConvertFloatToWords` and `ConvertFloatToNumerals` use Go float formatting (`strconv.FormatFloat(..., 'f', -1, 64)`).
  - Example: `100.00` becomes `100` before conversion.
- Word conversion trims input and applies number parsing rules internally.
- Numeral conversion maps only ASCII digits `0-9` to Bangla digits `০-৯`, while preserving other characters.

## Testing

```bash
go test ./...
```

## License

MIT License. See [LICENSE](LICENSE).