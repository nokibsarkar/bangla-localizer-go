# Author: Nokib Sarkar <nokibsarkar@gmail.com>

# Bangla Localizer Go

Bangla Localizer Go is a Golang module to convert numbers (integer, float, negative, very large) into Bangla (Bengali) words.

## Features

- Convert integers and floating-point numbers to Bangla words
- Handles negative numbers (with dash prefix)
- Supports very large numbers with extended Bangla units
- Ignores leading/trailing spaces and zeros
- Returns empty string for invalid input

## Installation

```sh
go get github.com/nokibsarkar/bangla-localizer-go
```


## Usage

```go
import "github.com/nokibsarkar/bangla-localizer-go"

func main() {
	localizer := banglalocalizer.NewLocalizer()
	fmt.Println(localizer.ConvertIntToWords(123))         // একশত তেইশ
	fmt.Println(localizer.ConvertFloatToWords(-123.45))   // -একশত তেইশ দশমিক চার পাঁচ
	fmt.Println(localizer.ConvertIntToWords(1000000000))  // এক হাজার কোটি
	fmt.Println(localizer.ConvertIntToWords(-1))          // -এক
	fmt.Println(localizer.ConvertFloatToWords(0.01))      // শুন্য দশমিক এক
}
```

### API

- `NewLocalizer() *Localizer` — Create a new localizer instance
- `ConvertIntToWords(number int) string` — Convert integer to Bangla words
- `ConvertFloatToWords(number float64) string` — Convert float to Bangla words


### Examples

```go
localizer := banglalocalizer.NewLocalizer()
fmt.Println(localizer.ConvertIntToWords(123))           // একশত তেইশ
fmt.Println(localizer.ConvertIntToWords(-123))          // -একশত তেইশ
fmt.Println(localizer.ConvertIntToWords(1000000000))    // এক হাজার কোটি
fmt.Println(localizer.ConvertFloatToWords(0.5))         // শুন্য দশমিক পাঁচ
fmt.Println(localizer.ConvertFloatToWords(-0.5))        // -শুন্য দশমিক পাঁচ
fmt.Println(localizer.ConvertIntToWords(1000000000000000)) // এক একশত লক্ষ কোটি
fmt.Println(localizer.convertNumberStringToWords("  123.45  ")) // একশত তেইশ দশমিক চার পাঁচ
fmt.Println(localizer.convertNumberStringToWords("abc")) // (empty string)
```

| Input                | Output                                      |
|----------------------|---------------------------------------------|
| 123                  | একশত তেইশ                                  |
| -123                 | -একশত তেইশ                                 |
| 1000000000           | এক হাজার কোটি                               |
| 0.5                  | শুন্য দশমিক পাঁচ                             |
| -0.5                 | -শুন্য দশমিক পাঁচ                           |
| 1000000000000000     | এক একশত লক্ষ কোটি                           |
| "  123.45  "         | একশত তেইশ দশমিক চার পাঁচ                   |
| "abc"                | (empty string)                              |

## Limits

- **Integer range:** Supports up to 1,000,000,000,000,000 (one quadrillion) and negative values of similar magnitude.
- **Float range:** Fractional part is read digit-by-digit in Bangla, trailing zeros are ignored.
- **Input type:** Accepts int, float64, or string (via convertNumberStringToWords).
- **Edge cases:**
	- Leading/trailing spaces and zeros are ignored.
	- Negative numbers return dash-prefixed Bangla words.
	- Invalid or non-numeric input returns an empty string.
	- For floats, both sides of the decimal are converted as integers.

## Testing

Run all tests:

```sh
go test ./...
```


## Documentation

Full GoDoc documentation is available in the source and on [pkg.go.dev](https://pkg.go.dev/github.com/nokibsarkar/bangla-localizer-go).


## License

MIT License. See [LICENSE](LICENSE).