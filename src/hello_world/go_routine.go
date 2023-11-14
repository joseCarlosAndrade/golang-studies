package main

import (
	"fmt"
	"time"
	"io"
	"strings"
)

func main() {
	// alo("main routine")
	go alo("go routine")
	alo("main")


	// messing with channels ///////////////////
	s := []int{1, 2, 3 ,4 ,5}

	c := make(chan int)
	go sum(s[0:1], c) // 17
	go sum(s[1:2], c) // -5
	go sum(s[2:3], c) // 12
	go sum(s[3:4], c) // 12
	go sum(s[4:5], c) // 12
	x, y, z, j, k := <-c, <-c, <-c, <-c, <-c // receive from c

	fmt.Println(x, y, z, j, k)
	//////////////////////////////////




	// readers ////////////////////////////
	r := strings.NewReader("reading from this string")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("number of bytes read %d. b=  %q\n", n, b)
	}

	a := 'A' +1
	fmt.Printf("%c\n", a)
} // main

// example of io.Read implementation
type MyReader struct{}

func (m MyReader) Read(b []byte) (int, error) {
	
	for i :=0 ;i < len(b) ; i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func alo(s string) {
	for i:=0 ; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}

}

/////////////// for messing with channels
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	m := time.Duration(sum*100)
	time.Sleep(m*time.Millisecond)
	c <- sum // send sum to c
}

func wait(s []int, c chan int) {
	fmt.Println("waiting..")	
	time.Sleep(400*time.Millisecond)
	
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

///////////////////////////////////////