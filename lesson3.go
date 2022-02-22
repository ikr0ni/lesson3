package lesson3

import (
	"fmt"
	"github.com/pkg/errors"
	_ "github.com/valyala/fasthttp"
	"log"
)

func Main() {
	fmt.Println("Hello World")
	log.Println("Hello World")
	errors.New("Hello World")
}
