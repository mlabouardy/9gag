# 9GAG API in Go


## Installing

### *go get*

    $ go get -u github.com/mlabouardy/9gag

## Example

### Getting top memes

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
### Filter by tag name

```golang
package main

import (
	"fmt"
	"github.com/mlabouardy/9gag"
)

func main() {
	gag9 := gag9.New()
	for _, meme := range gag9.FindByTag("got) {
		fmt.Printf("Description: %s\nImage: %s\n", meme.Description, meme.Image)
	}
}
```

## Contributing

You are more than welcome to contribute to this project.  Fork and
make a Pull Request, or create an Issue if you see any problem.
