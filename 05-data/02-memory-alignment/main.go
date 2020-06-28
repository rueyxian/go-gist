package main

import (
	"fmt"
	"unsafe"
)

type one struct {
	val1 bool
	val2 int
	val3 []int
}

type two struct {
	val1 int16
	val2 int16
	val3 bool
}

type three struct {
	val1 bool
	val2 []int
	val3 int
}

func main() {

	s1 := one{}
	s2 := two{}
	s3 := three{}

	fmt.Println("============ one ============")
	fmt.Printf("      type \t size \t alignment \t offset \t address \n")
	fmt.Printf("%10T \t %v \t %v \t\t %v \t\t %v \n", s1, unsafe.Sizeof(s1), unsafe.Alignof(s1), unsafe.Alignof(&s1), unsafe.Pointer(&s1))
	fmt.Printf("%10T \t %v \t %v \t\t %v \t\t %v \n", s1.val1, unsafe.Sizeof(s1.val1), unsafe.Alignof(s1.val1), unsafe.Alignof(&s1.val1), unsafe.Pointer(&s1.val1))
	fmt.Printf("%10T \t %v \t %v \t\t %v \t\t %v \n", s1.val2, unsafe.Sizeof(s1.val2), unsafe.Alignof(s1.val2), unsafe.Alignof(&s1.val2), unsafe.Pointer(&s1.val2))
	fmt.Printf("%10T \t %v \t %v \t\t %v \t\t %v \n", s1.val3, unsafe.Sizeof(s1.val3), unsafe.Alignof(s1.val3), unsafe.Alignof(&s1.val3), unsafe.Pointer(&s1.val3))
	fmt.Println("")

	fmt.Println("============ two ============")
	fmt.Printf("      type \t size \t alignment \t offset \t address \n")
	fmt.Printf("%10T \t %v \t %v \t\t %v \t\t %v \n", s2, unsafe.Sizeof(s2), unsafe.Alignof(s2), unsafe.Alignof(&s2), unsafe.Pointer(&s2))
	fmt.Printf("%10T \t %v \t %v \t\t %v \t\t %v \n", s2.val1, unsafe.Sizeof(s2.val1), unsafe.Alignof(s2.val1), unsafe.Alignof(&s2.val1), unsafe.Pointer(&s2.val1))
	fmt.Printf("%10T \t %v \t %v \t\t %v \t\t %v \n", s2.val2, unsafe.Sizeof(s2.val2), unsafe.Alignof(s2.val2), unsafe.Alignof(&s2.val2), unsafe.Pointer(&s2.val2))
	fmt.Printf("%10T \t %v \t %v \t\t %v \t\t %v \n", s2.val3, unsafe.Sizeof(s2.val3), unsafe.Alignof(s2.val3), unsafe.Alignof(&s2.val3), unsafe.Pointer(&s2.val3))
	fmt.Println("")

	fmt.Println("============ two ============")
	fmt.Printf("      type \t size \t alignment \t offset \t address \n")
	fmt.Printf("%10T \t %v \t %v \t\t %v \t\t %v \n", s3, unsafe.Sizeof(s3), unsafe.Alignof(s3), unsafe.Alignof(&s3), unsafe.Pointer(&s3))
	fmt.Printf("%10T \t %v \t %v \t\t %v \t\t %v \n", s3.val1, unsafe.Sizeof(s3.val1), unsafe.Alignof(s3.val1), unsafe.Alignof(&s3.val1), unsafe.Pointer(&s3.val1))
	fmt.Printf("%10T \t %v \t %v \t\t %v \t\t %v \n", s3.val2, unsafe.Sizeof(s3.val2), unsafe.Alignof(s3.val2), unsafe.Alignof(&s3.val2), unsafe.Pointer(&s3.val2))
	fmt.Printf("%10T \t %v \t %v \t\t %v \t\t %v \n", s3.val3, unsafe.Sizeof(s3.val3), unsafe.Alignof(s3.val3), unsafe.Alignof(&s3.val3), unsafe.Pointer(&s3.val3))
	fmt.Println("")

	// fmt.Println("============ structureTwo ============")
	// fmt.Printf("%18T \t  size: %v \t alignment: %v \t alignment: %v \t address: %v \n",
	// 	s2,
	// 	unsafe.Sizeof(s2),
	// 	unsafe.Alignof(s2),
	// 	unsafe.Alignof(&s2),
	// 	unsafe.Pointer(&s2))
	// fmt.Printf("%18T \t  size: %v \t alignment: %v \t offset: %v \t address: %v \n",
	// 	s2.val1,
	// 	unsafe.Sizeof(s2.val1),
	// 	unsafe.Alignof(s2.val1),
	// 	unsafe.Offsetof(s2.val1),
	// 	unsafe.Pointer(&s2.val1))
	// fmt.Printf("%18T \t  size: %v \t alignment: %v \t offset: %v \t address: %v \n",
	// 	s2.val2,
	// 	unsafe.Sizeof(s2.val2),
	// 	unsafe.Alignof(s2.val2),
	// 	unsafe.Offsetof(s2.val2),
	// 	unsafe.Pointer(&s2.val2))
	// fmt.Printf("%18T \t  size: %v \t alignment: %v \t offset: %v \t address: %v \n",
	// 	s2.val3,
	// 	unsafe.Sizeof(s2.val3),
	// 	unsafe.Alignof(s2.val3),
	// 	unsafe.Offsetof(s2.val3),
	// 	unsafe.Pointer(&s2.val3))
	// fmt.Println("")

	// fmt.Println("============ structureThree ============")
	// fmt.Printf("%18T \t  size: %v \t alignment: %v \t alignment: %v \t address: %v \n",
	// 	s3,
	// 	unsafe.Sizeof(s3),
	// 	unsafe.Alignof(s3),
	// 	unsafe.Alignof(&s3),
	// 	unsafe.Pointer(&s3))
	// fmt.Printf("%18T \t  size: %v \t alignment: %v \t offset: %v \t address: %v \n",
	// 	s3.val1,
	// 	unsafe.Sizeof(s3.val1),
	// 	unsafe.Alignof(s3.val1),
	// 	unsafe.Offsetof(s3.val1),
	// 	unsafe.Pointer(&s3.val1))
	// fmt.Printf("%18T \t  size: %v \t alignment: %v \t offset: %v \t address: %v \n",
	// 	s3.val2,
	// 	unsafe.Sizeof(s3.val2),
	// 	unsafe.Alignof(s3.val2),
	// 	unsafe.Offsetof(s3.val2),
	// 	unsafe.Pointer(&s3.val2))
	// fmt.Printf("%18T \t  size: %v \t alignment: %v \t offset: %v \t address: %v \n",
	// 	s3.val3,
	// 	unsafe.Sizeof(s3.val3),
	// 	unsafe.Alignof(s3.val3),
	// 	unsafe.Offsetof(s3.val3),
	// 	unsafe.Pointer(&s3.val3))
	// fmt.Println("")

	// ============================================================

}
