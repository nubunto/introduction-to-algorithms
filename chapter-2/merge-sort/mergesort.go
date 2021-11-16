package main

import (
	"strings"
	"fmt"
	"math"
)

func merge(A []int, p, q, r, depth int) {
	n1 := q-p+1
	n2 := r-q
	fmt.Printf("%smerge(A=%v, p=%d, q=%d, r=%d)\n", strings.Repeat("\t", depth), A, p, q, r)
	fmt.Printf("%s merge(A[p:q]=%v, n1=%d, n2=%d)\n", strings.Repeat("\t", depth), A[p:q], n1, n2)

	L := make([]int, n1+1)
	R := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		L[i] = A[(p+i)]
	}

	for j := 0; j < n2; j++ {
		R[j] = A[q+j]
	}
	fmt.Printf("%s merge(L=%v)\n", strings.Repeat("\t", depth), L)
	fmt.Printf("%s merge(R=%v)\n", strings.Repeat("\t", depth), R)

	i := 0
	j := 0
	L[n1] = math.MaxInt32
	R[n2] = math.MaxInt32

	fmt.Printf("%s merge(L=%v)\n", strings.Repeat("\t", depth), L)
	fmt.Printf("%s merge(R=%v)\n", strings.Repeat("\t", depth), R)

	fmt.Printf("%s starting K at %d\n", strings.Repeat("\t", depth), p)
	for k := p; k < r; k++ {
		if L[i] <= R[j] {
			fmt.Printf("%s  bringing L[%d]=%d to A[%d]\n", strings.Repeat("\t", depth), i, L[i], k)
			A[k] = L[i]
			i++
		} else {
			fmt.Printf("%s  bringing R[%d]=%d to A[%d]\n", strings.Repeat("\t", depth), j, R[j], k)
			A[k] = R[j]
			j++
		}
	}
}

func mergesort(A []int, p, r, depth int) {
	fmt.Printf("%smergesort(A=%v, p=%d, r=%d)\n", strings.Repeat("\t", depth), A, p, r)
	if p < r {
		q := (p+r)/2
		fmt.Printf("%sq=%d\n", strings.Repeat("\t", depth), q)
		mergesort(A, p, q, depth+1)
		mergesort(A, q+1, r, depth+1)
		merge(A, p, q, r, depth)
	}
}

func main() {
	a := []int{3, 2, 1}
	mergesort(a, 0, len(a), 0)
	fmt.Println("a>>", a)
}

