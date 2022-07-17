package main

import "fmt"

// Perm() 对 a 形成的每⼀排列调⽤ f().
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// 对索引 i 从 0 到 len(a) - 1，实现递归函数 perm().
func perm(a []rune, f func([]rune), i int) {
	// TODO

	if i == len(a) {
		f(a)
		return
	}

	for x := i; x < len(a); x++ {
		a[x], a[i] = a[i], a[x]
		perm(a, f, i+1)
		a[x], a[i] = a[i], a[x]
	}

}

func main() {
	Perm([]rune("ABC"), func(a []rune) {
		fmt.Println(string(a))
	})
}
