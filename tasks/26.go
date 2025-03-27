// package main
//
// import (
//
//	"fmt"
//	"log"
//	"strings"
//
// )
//
//	func main() {
//		var a, b string
//		fmt.Println("Введите строку в которой вы хотите проверить содержание символов из строки строки номер 2")
//		fmt.Scan(&a)
//		fmt.Println("Введите строку номер 2")
//		fmt.Scan(&b)
//		result := make([]string, 0)
//		str1 := strings.ToLower(a)
//		str2 := strings.ToLower(b)
//		for _, i2 := range str2 {
//			for _, i3 := range str1 {
//				if strings.Contains(string(i2), string(i3)) {
//					result = append(result, string(i3))
//				}
//			}
//		}
//		if len(result) == len(str1) {
//			log.Println("true")
//		} else {
//			log.Println("false")
//		}
//	}

//Разработать программу, которая проверяет, что все символы в строке уникальные
//(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isUnique("abCdefabCdefAafabCdefAafAaf"))
}

func isUnique(str string) bool {
	alf := []rune(strings.ToLower(str))
	m := map[rune]struct{}{}
	for _, v := range alf {
		_, ok := m[v]
		if ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}
