package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"encoding/hex"
	"encoding/base64"
)

func main() {
	LinebyLineScan()
}

func htob(s string) string{
	b, _ := hex.DecodeString(s)
	return base64.RawStdEncoding.EncodeToString(b)
}

func LinebyLineScan() {
	var (
		file *os.File
		err error
	)
	file = os.Stdin //NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	if len(os.Args) > 1 {
		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
//	fmt.Println(os.Args, len(os.Args))
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := strings.Split(scanner.Text(), ",")
		if len(x[0]) == 64{
			fmt.Println("SET", htob(x[1]), htob(x[0]))
			continue
		}
		fmt.Println("SET", htob(x[1]), x[0])
	}
	if err := scanner.Err(); err != nil {
//		log.Fatal(err)
	}
}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//
//	splitter := func(s, sep string) (a string, b string) {
//		x := strings.Split(s, sep)
//		return x[0], x[1]
//	}
//
//	for scanner.Scan() {
//		line := scanner.Text()
//		if strings.Contains(line, ",") {
//			seed, hash := splitter(scanner.Text(), ",")
//			if len(seed) > 1 && len(seed) <= 64 && len(hash) == 64 {
//				fmt.Println("SET", hash, seed)
//			}
//		}
//	}
//}
