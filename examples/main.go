package main

import (
	"fmt"

	"github.com/uretgec/slugify"
)

func main() {
	title := slugify.Make("Hellö Wörld хелло ворлд", []string{`х`, "kh", `л`, "l"}...)
	fmt.Println(title) // "hello-world-khello-vorld"

	title = slugify.Make("影師", []string{`影`, "ying", `師`, "shi"}...)
	fmt.Println(title) // "ying-shi"

	title = slugify.Make("This & that")
	fmt.Println(title) // "this-that"

	title = slugify.Make("Çikolata Soslu Bulut Kek")
	fmt.Println(title) // "cikolata-soslu-bulut-kek"
}
