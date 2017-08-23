# 9GAG API in Go 

[![CircleCI](https://circleci.com/gh/mlabouardy/9gag.svg?style=svg)](https://circleci.com/gh/mlabouardy/9gag) [![Go Report Card](https://goreportcard.com/badge/github.com/mlabouardy/9gag)](https://goreportcard.com/report/github.com/mlabouardy/9gag) [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

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
	for _, meme := range gag9.FindByTag("got") {
		fmt.Printf("Description: %s\nImage: %s\n", meme.Description, meme.Image)
	}
}
```

## How to use

http://www.blog.labouardy.com/create-9gag-android-application/

## Contributing

You are more than welcome to contribute to this project.  Fork and
make a Pull Request, or create an Issue if you see any problem.
