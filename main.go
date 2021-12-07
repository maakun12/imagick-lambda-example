package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/gographics/imagick.v2/imagick"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	log.Println("start...")
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	log.Println(mw)
	// ...etc

	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
