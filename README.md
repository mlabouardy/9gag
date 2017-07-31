# 9GAG Golang API

```golang
package main

import (
	"fmt"
	"github.com/mlabouardy/9gag"
)

func main() {
	gag9 := gag9.New()
	for _, meme := range gag9.Find() {
		fmt.Printf("Description: %s\nImage: %s\n", meme.Description, meme.Image)
	}
}
```
