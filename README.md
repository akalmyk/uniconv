# uniconv
uniconv helps to convert Unicode characters in strings or files and get a regular text instead of those ugly \uXXXXs.

Usage with files:

```
package main

import (
	"log"

	"github.com/akalmyk/uniconv"
)

func main() {
	inputFile := "path/to/input/file"
	outputFile := "path/to/output/file"

	err := uniconv.ConvFileToFile(inputFile, outputFile)
	if err != nil {
		log.Fatal(err)
	}
}
```

Usage File to string:

```
package main

import (
	"log"

	"github.com/akalmyk/uniconv"
)

func main() {
	inputFile := "path/to/input/file"

	convText, err := uniconv.ConvFileToString(inputFile)
	if err != nil {
		log.Fatal(err)
	}

  log.Println(convText)
}
```

Usage with strings:

```
package main

import (
	"log"

	"github.com/akalmyk/uniconv"
)

func main() {
	text := "Monse\u00c3\u00b1or \u00c3\u0093scar Arnulfo Romero International Airport"

	convText := uniconv.ConvText(text)
	log.Println(convText)
}
```

Output:
```
Monseñor Óscar Arnulfo Romero International Airport
```
