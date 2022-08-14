# Slugify
URL-friendly slugify with your languages support.

## Install

```
go get -u github.com/uretgec/slugify
```

## Examples
Simple One

```go
package main

import (
	"fmt"
	"github.com/uretgec/slugify"
)

func main() {
	title := slugify.Make("Hellö Wörld хелло ворлд", []string{`х`, "kh", `л`, "l"}...) // not good example but works
	fmt.Println(title) // "hello-world-khello-vorld"

	title = slugify.Make("影師", []string{`影`, "ying", `師`, "shi"}...) // not good example but works
	fmt.Println(title) // "ying-shi"

	title = slugify.Make("This & that")
	fmt.Println(title) // "this-that"

	title = slugify.Make("Çikolata Soslu Bulut Kek")
	fmt.Println(title) // "cikolata-soslu-bulut-kek"

	title = slugify.Make("Çikolata'nın Sosunda /Kendi ?Halinde -- Ne Üdüğü Belirsiz == -- - \" Bulut Kek")
	fmt.Println(title) // "cikolatanin-sosunda-kendi-halinde-ne-udugu-belirsiz-bulut-kek"
}
```

## Tests
```
go test -timeout 30s github.com/uretgec/slugify
```