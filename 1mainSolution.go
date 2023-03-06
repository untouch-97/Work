package main

// Во-первых, каждую задачку можно вынести в отдельную папку, где создать файл с package main и func main(),
// и тогда не надо будет комментить код, и можно будет описание к задаче в этом же файле/папке поместить.

// first task - тут ОК
//func main() {
//
//	for i := 0; i < 50*2; i += 2 {
//		fmt.Println(i)
//	}
//}

// second task - тут ОК, но можно без переменной b было обойтись
//func main() {
//	b := 0
//	for i := 0; i < 99; i += 3 {
//		fmt.Println(b)
//		b += 3
//	}
//}

// // third task - НЕ ОК вообще: читай задание
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

//					fourth task - через switch case норм, но сам итератор не надо было выводить. читай задание
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

// fifth task - неправильно, читай задание. логика должна быть в функции (не в main), а в main ты вызываешь эту функцию
// скан stdin не обязательно было, но норм
// почитай про default в switch операторе

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
