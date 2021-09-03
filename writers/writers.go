package writers

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/Ex0dIa-dev/ssh-honeypot-go/writers/colors"
)

const Welcome = "Hello Hacker! You're a new bee in my honeypot! Bye :D"

func Write(w io.Writer, str string) {

	chars := strings.Split(str, "")

	for _, ch := range chars {
		fmt.Fprint(w, ch)
		time.Sleep(30 * time.Millisecond)
	}

}

func ColorWrite(w io.Writer, str, color string) {
	Write(w, color+str+colors.Reset)
}

func PrintEnd(w io.Writer, ends int) {

	for i := 0; i < ends; i++ {
		fmt.Fprint(w, "\n")
	}

}
