package main

import ("fmt")

func tarzan(n int) {
	for i := 0; i < n; i ++ {
		fmt.Printf("타잔이 %d원짜리 팬티를 입고, %d원짜리 칼을 차고 노래를 한다.\n", i*10, (i+1)*10)
	}
}

func main() {
	tarzan(3)
}