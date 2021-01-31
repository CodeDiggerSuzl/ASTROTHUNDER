package packaging

import (
	"bytes"
	"fmt"
)

func main() {
	bytebuffer := bytes.NewBuffer([]byte{})
	if err := Encode(bytebuffer, "hello world 0!!!"); err != nil {
		panic(err)
	}
	if err := Encode(bytebuffer, "hello world 1!!!"); err != nil {
		panic(err)
	}
	for {
		if bt, err := Decode(bytebuffer); err == nil {
			fmt.Println(string(bt))
			continue
		}
		break
	}
}
