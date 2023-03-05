package main

import "fmt"

// first exercise
//func main() {
//
//	for i := 0; i < 50*2; i += 2 {
//		fmt.Println(i)
//	}
//}

// second exercise
//func main() {
//	b := 0
//	for i := 0; i < 99; i += 3 {
//		fmt.Println(b)
//		b += 3
//	}
//}

//// third exercise
//func main() {
//	even := 0
//	odd := 1
//	for i := 0; i < 99; i += 2 {
//		even += 2
//		odd += 2
//		s := "ok!"
//		gg := "neok!"
//		fmt.Println(even, s)
//		fmt.Println(odd, gg)
//	}
//
//}
func main() {

	all := []int{}
	five := []int{}
	three := []int{}
	t, f, a := 0, 0, 0

	for i := 0; i < 99; i++ {
		t += 3
		f += 5
		all = append(all, i)
		three = append(three, t)
		five = append(five, f)
		return
	}
	for a, all = range all {
		switch all {
		case three:
			fmt.Printf("%v три.\n", a)
		default:
			fmt.Printf("%v ok.\n", a)

		}
	}

}
