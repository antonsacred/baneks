# Go baneks

Random joke of category B from the site [baneks.ru](https://baneks.ru/)

Случайный анекдот категории Б с сайта [baneks.ru](https://baneks.ru/)

## Usage

```go
package main

import (
	"fmt"
	"github.com/antonsacred/baneks"
)

func main() {
	anek, err := baneks.RandomBAnek()
	if err != nil {
		// hande error
	}
	fmt.Println(anek)
}

```
