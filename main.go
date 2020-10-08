package main

import (
	"context"
	"log"
	//"time"
	"bufio"
	"fmt"
	"github.com/chromedp/chromedp"
	"os"
	"strings"
)

func main() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	//ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	//defer cancel()

	reader := bufio.NewReader(os.Stdin)
	err := chromedp.Run(ctx, chromedp.Navigate("about:blank"))

	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Print("percolator> ")

		text, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		text = strings.Replace(text, "\n", "", -1)

		switch {
		case "help()" == text:
			fmt.Println("Welcome to the Percolator!")
		default:
			var res []byte

			err = chromedp.Run(ctx, chromedp.Evaluate(text, &res))

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(res))
		}
	}
}
