package main

// first task
//func main() {
//
//	for i := 0; i < 50*2; i += 2 {
//		fmt.Println(i)
//	}
//}

// second task
//func main() {
//	b := 0
//	for i := 0; i < 99; i += 3 {
//		fmt.Println(b)
//		b += 3
//	}
//}

// // third task
//
//	func main() {
//		even := 0
//		odd := 1
//		for i := 0; i < 99; i += 2 {
//			even += 2
//			odd += 2
//			s := "ok!"
//			gg := "neok!"
//			fmt.Println(even, s)
//			fmt.Println(odd, gg)
//		}
//
// }

//					fourth task
//func main() {
//	for i := 0; i <= 99; i++ {
//		if i%3 == 0 && i%5 == 0 {
//			fmt.Printf("%v 3 and 5 \n", i)
//		} else if i%3 == 0 {
//			fmt.Printf("%v 3 \n", i)
//		} else if i%5 == 0 {
//			fmt.Printf("%v 5 \n", i)
//		} else {
//			fmt.Printf("%v Баха крассваучик еЖЖи!\n", i)
//		}
//
//	}
//}
//func main() {
//	for i := 0; i <= 99; i++ {
//		switch {
//		case i%3 == 0 && i%5 == 0:
//			fmt.Printf("%v 3 and 5 \n", i)
//		case i%3 == 0:
//			fmt.Printf("%v 3 \n", i)
//		case i%5 == 0:
//			fmt.Printf("%v 5 \n", i)
//		default:
//			fmt.Printf("%v Баха крассаучик еЖЖи!\n", i)
//
//		}
//	}
//}

// fifth task

//func boom() error {
//	return errors.New("age entered incorrectly")
//}
//func main() {
//	err := boom()
//	var age int
//	fmt.Println("Enter your age:")
//	fmt.Scanln(&age)
//	switch {
//	case age >= 0 && age < 18:
//		fmt.Println("молодой огурчик!")
//	case age >= 0 && age < 60:
//		fmt.Println("уже не такой молодой огурчик")
//	case age >= 60:
//		fmt.Println("пажилая каструля")
//	case err != nil:
//		fmt.Println(err)
//	}
//
//}
