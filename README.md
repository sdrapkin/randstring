# randstring
## Fast cryptographic random string generators for Go. By [Stan Drapkin](https://github.com/sdrapkin/).
* Fast replacement for `crypto/rand`.`Text()`

[Go playground](https://go.dev/play/p/QocRnDsoWUM)

```go
// func Text16(length int) string
// func Text32(length int) string
// func Text32c(length int) string
// func Text64(length int) string
// func Text64URL(length int) string

package main

import (
	"fmt"

	"github.com/sdrapkin/randstring"
)

func main() {
	const (
		ROUNDS = 4
		LENGTH = 64
	)

	for range ROUNDS {
		text := randstring.Text16(LENGTH)
		fmt.Println(text)
	}
	fmt.Println("====")

	for range ROUNDS {
		text := randstring.Text32(LENGTH)
		fmt.Println(text)
	}
	fmt.Println("====")

	for range ROUNDS {
		text := randstring.Text32c(LENGTH) // Crockford Base32
		fmt.Println(text)
	}
	fmt.Println("====")

	for range ROUNDS {
		text := randstring.Text64(LENGTH)
		fmt.Println(text)
	}
	fmt.Println("====")

	for range ROUNDS {
		text := randstring.Text64URL(LENGTH)
		fmt.Println(text)
	}
	fmt.Println("====")
}
```
```
40589231573021090E5E09BB7CB11E176B50B951BF4C44E4D25A5BB577181D32
CA56C4FAA9779E92F48075E59DECC893C2780C203DAFACF3382B17C052537499
6FE30C504369A2060793874CB81A78D0D65FBB898BC914A6F9D246B8924BB582
FE2832227CFA79BF898DCB18BFF652B625207457B90206D8DC885F63290713A3
====
THCHXXSC7YWI3BKGPCDBBXRLY7M2XNDSOGH7MJFSDSHS6AGGZZNRHZDQ47LGG2FV
RVY65BE2SHZ2OXLVOX4LIJ2SBOO6736UF2X4EVBZ3G7CENQTB3VR3KK2XB54ZBW6
HMA6VR4IC46DXUOKQIYW3XW3BU3KJJFWEAJR7OZKKHPVPIXBD2O76MEMGYXO4B27
7UMH5BZXK4MJHKVWF52NPFTUITW7PDOSYCDFLL6R7EQMYAF3WVM34JEJ5D3IFQ7V
====
FDFZPQXB38SE8MAJAHTR1FMW1AWVTQAPTXSMCFSNFATB2AX9DYK3K01K3BFZYSMD
YZHQ8C7H1A286493SJZ6BSEQF457ZVAA4T6SHGSC1BK4SFA3HY9WS87ARE2B5PCP
8JYD32W354H3Y2ETTRK4SMSBQ6W2K0VSJ8WTM04GYZ4D9S46SRTNJFJ8GJNRPXG1
MZPXEQD81X0JQ03C6YVTH9S6K2ZTR7VJG2NGX3DHX2D9GSEM1RZND2BT9S442PM4
====
6cA42j+F6c1DFL21tWymhibtfI751bcPqj4BwWqj6q1tmg42WQ0/VPsNo1WPuecL
tCXNUL1Jo6CWH6QMaIvC+q0NwRUzaAZ5JM7t0W5AIm5vekp7NgnzEy3xODmqb7iP
YOx4FHQ+XQTusDSEAwVPd0Ov5btOhzJe324HDaEKBtQb1jGFznZa3gkcsLRyfk+9
XoP2b2u/x6l9lYHK9P3Yh9b+zghVg3Agvwd7uvGtI6llAngAbUhGRfY0qQwJa2LM
====
OUqGqsotiSz-Lvb6ZZK65vHiGf2FYrJ5YprzrIRmy-JOzcr-EQ6tC9FlgjE1loXw
PC1RirYAzMjb47_34tsrWoWP_OVt8_Jrm-3LWwjNLxcw9Jk2Cdq_orCHkjjygx_U
4-g_sG8ZRH8axURTxRpmcq_GMM5fHCP5iicp5ke2Wa1cnUoSWcNw9-vkVW69zsbC
irxALKRtuUVWrh4-xKEKijEHruVn_ToQTYhYolUzYIgCZC4Yk1W-j4Fr5ZyAihQG
====
```
## API

| Function                | Description |
|---|---|
| `Text16(length int)`    | Generates a random hex string of a given length (base16)               |
| `Text32(length int)`    | Generates a random base32 string of a given length (RFC 4648)          |
| `Text32c(length int)`   | Generates a random base32 Crockford string of a given length           |
| `Text64(length int)`    | Generates a random base64 string of a given length (RFC 4648)          |
| `Text64URL(length int)` | Generates a random base64 URL-safe string of a given length (RFC 4648) |

All functions panic if a negative length is provided. A length of zero returns an empty string.

## Documentation
 [![Go Reference](https://pkg.go.dev/badge/github.com/sdrapkin/randstring.svg)](https://pkg.go.dev/github.com/sdrapkin/randstring)

Full `go doc` style documentation: https://pkg.go.dev/github.com/sdrapkin/randstring

## Installation
### Using `go get`

To install the `randstring` package, run the following command:

```sh
go get -u github.com/sdrapkin/randstring
```

To use the `randstring` package in your Go project, import it as follows:

```go
import "github.com/sdrapkin/randstring"
```
