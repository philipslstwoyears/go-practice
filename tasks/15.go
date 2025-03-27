package main

var justString string // не стоит создавать глобальную переменную

func createHugeString(length int) string {
	sl := make([]byte, length)
	for i := range sl {
		sl[i] = 'a'
	}
	return string(sl)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100])) // не ссылаюсь на часть глобальной переменной, а создаю сразу новую строку чтобы не было утечки данных и чтобы освободить память
}

func main() {
	someFunc()
}
