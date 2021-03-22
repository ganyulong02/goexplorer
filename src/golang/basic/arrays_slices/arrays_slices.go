package main

import (
	"fmt"
)

func main() {
	grades := [...]int{95, 97, 98}
	fmt.Printf("Grades: %v", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of Students: %v", len(students))

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{1, 1, 0}
	identityMatrix[2] = [3]int{1, 0, 1}
	fmt.Println(identityMatrix)

	// array
	a := [...]int{1, 2, 3}
	b := &a // pointer
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// slices, similar to Python list
	a_slice := []int{1, 2, 4}
	fmt.Println(a_slice)
	fmt.Printf("Length: %v\n", len(a_slice))
	fmt.Printf("Capcity: %v\n", cap(a_slice))

	b_slice := []int{1, 2, 3, 4, 5, 6}
	b_slice_all := b_slice[:]
	c := b_slice[2:4]

	fmt.Println(b_slice)
	fmt.Println(b_slice_all)
	fmt.Println(c)

	// make a slice, it is a reference
	make_a_slice := make([]int, 3, 100)
	fmt.Println(make_a_slice)
	fmt.Printf("Length: %v\n", len(make_a_slice))
	fmt.Printf("Capcity: %v\n", cap(make_a_slice))

	append_slice := []int{}
	append_slice = append(append_slice, 1, 2, 3, 4, 5)
	fmt.Println(append_slice)
	fmt.Printf("Length: %v\n", len(append_slice))
	fmt.Printf("Capcity: %v\n", cap(append_slice)) // capacity will double it previous size

	append_slice = append(append_slice, []int{6, 7, 8}...) // spead the slice and concat
	fmt.Println(append_slice)

	// stack remove from start, end , middle
	m := []int{1, 2, 3, 4, 5}
	m_leftpop := m[1:]
	m_rightpop := m[:len(m)-1]
	m_middlepop := append(m[:2], m[3:]...)
	fmt.Println((m_leftpop))
	fmt.Println((m_rightpop))
	fmt.Println((m_middlepop)) // [1 2 4 5]
}
