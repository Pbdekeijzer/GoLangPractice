package main

import (
	"fmt"
	"math"
)

var count int

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	root *Node
}

func (l *LinkedList) AddToList(value int) {
	l.root = &Node{value, l.root}
}

func (l *LinkedList) RemoveFromList(value int) {
	l.root = l.root._removeFromNodesRecursive(value)
}

func (n *Node) _removeFromNodesRecursive(value int) *Node {
	if n == nil {
		return n
	}
	if value == n.value {
		return n.next._removeFromNodesRecursive(value)
	} else {
		return &Node{n.value, n.next._removeFromNodesRecursive(value)}
	}
}

func (l *LinkedList) Map(f func(int) int) {
	root := l.root
	if root != nil {
		root.value = f(root.value)
	}
	for root.next != nil {
		root.next.value = f(root.next.value)
		root = root.next
	}
}

func main() {
	linkedlist := LinkedList{&Node{15, nil}}
	linkedlist.AddToList(5)
	linkedlist.AddToList(10)
	linkedlist.AddToList(29)

	linkedlist.RemoveFromList(10)
	fmt.Println(linkedlist.root.next.value)

	x := func(x int) int {
		return x * 10
	}
	linkedlist.Map(x)
	fmt.Println(linkedlist.root.value)

	testMatrix := [][]int{
		[]int{-12, -6, 3, 5, 9},
		[]int{-5, -2, 9, 12, 15},
		[]int{-3, 5, 14, 16, 18},
		[]int{0, 6, 17, 22, 23},
		[]int{5, 7, 19, 24, 28},
	}

	var value int
	value = getNegativeValuesInMatrix(testMatrix)
	fmt.Println(value)

	fmt.Println(testMatrix)

	//Thom 1
	// z := false
	// testArray := []int{5, 2, 8, 7}
	// z = findNumber(9, testArray)
	// fmt.Println(z)

	//Barld 1
	// var testArray []int
	// testArray = findPrimeBelowNumber(100)
	// for _, tester := range testArray {
	// 	fmt.Println(tester)
}

func getNegativeValuesInMatrix(matrix [][]int) int {
	count := 0
	x := len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < x; j++ {
			fmt.Println(j)
			if matrix[i][j] < 0 {
				count++
			} else {
				x = j
				break
			}
		}
	}
	return count
}

func findSumValuesInArray(sum int, intArray []int) bool {
	x := false
	for i := 0; i < len(intArray); i++ {
		for j := 0; j < len(intArray); j++ {
			if i != j {
				if intArray[i]+intArray[j] == sum {
					x = true
				}
			}
		}
	}
	return x
}

func findPrimeBelowNumber(n int) []int {
	var primeArray []int
	primeArray = append(primeArray, 2)

	for i := 3; i < n; i++ {
		isPrimeNumber := true
		sqrtI := int(math.Sqrt(float64(i)))
		for _, prime := range primeArray {
			if prime > sqrtI {
				break
			}
			if i%prime == 0 {
				isPrimeNumber = false
				break
			}
		}

		if isPrimeNumber {
			primeArray = append(primeArray, i)
		}
	}

	return primeArray
	// for i := 2; i < n; i++ {
	// 	ispriem := true
	// 	for j := 3; j < i; j++ {
	// 		if i%j == 0 {
	// 			ispriem = false
	// 			break
	// 		}
	// 	}
	// 	if ispriem {
	// 		priemArray = append(priemArray, i)
	// 	}
	// }
}
