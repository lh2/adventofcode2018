#!/bin/sh

if [ "x$1" = "x" ] || [ "x$2" = "x" ]; then
	echo 1>&2 USAGE: aoc.sh DAY TASK
	exit 1
fi
if [ ! -d "d$1" ]; then
	echo 1>&2 Sadly, there is currently no solution for day $1
	exit 1
fi
if [ $2 -lt 1 ] || [ $2 -gt 2 ]; then
	echo 1>&2 TASK has to be 1 or 2
	exit 1
fi
if [ ! -e input/$1.txt ]; then
	./getinput.sh $1
fi

cat <<EOF > d$1/gen.go
package main // import "entf.net/adventofcode2018/d$1"

import (
	"bufio"
	"fmt"
	"os"
)

func inAsSlice(in chan string) []string {
	list := make([]string, 0)
	for line := range in {
		list = append(list, line)
	}
	return list
}

func main() {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	fmt.Println(task$2(ch))	
}
EOF

go build -o bin ./d$1 && ./bin < input/$1.txt
rm bin 2> /dev/null
rm d$1/gen.go 2>/dev/null
