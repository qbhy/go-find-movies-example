package main

import (
	"bufio"
	"os"
	"io"
	"strings"
	"encoding/json"
	"fmt"
	"github.com/96qbhy/go-find-movies"
)

type Params struct {
	Limit   int
	Keyword string
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	var param Params
	for {
		s, err := inputReader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		s = strings.TrimSpace(s)
		if s != "" {
			json.Unmarshal([]byte(s), &param)
			s = findMovies.Find(param.Keyword, param.Limit)
			fmt.Println(s)
		} else {
			fmt.Println("get empty \n")
		}
	}
}
