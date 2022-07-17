package main

import "fmt"

func main() {
	// string은 필드가 2개(메모리를 가리키는 포인터와 길이)인 구조체이다.
	// 즉, string은 불변의 리터럴을 가리키는 포인터, 그리고 길이로 이루어져 있다.
	// 따라서 string을 복사 생성하면 string 내용물 자체가 복사되는 것이 아니라,
	// string 객체의 필드(가리키는 주소, 길이)가 복사된다.

	// 백쿼트 키는 특수문자가 동작하지 않게 한다.
	str1 := "Tab\tworks!"
	str2 := `Tab\tdoesn't\twork..`
	fmt.Println(str1)
	fmt.Println(str2)

	// 문자열간 연산도 가능
	str1Clone := str1 + str2
	fmt.Println(str1Clone)

	if str1Clone == str1 {
		fmt.Println("str1Clone == str1")
	}

	// 백쿼트키는 여러줄 적는데 \n을 쓰지 않아도 된다.
	str3 := "Sometimes, \nit is the very people \nno one imagines anything of,"
	str4 := `who do the things that no one can imagine.`
	fmt.Println(str3)
	fmt.Println(str4)

	// Go는 UTF-8를 표준 문자코드로 쓴다. 따라서, 영어와 한국어의 바이트 단위가 다른데,
	// 아래처럼 len()함수를 string 그대로 쓰면 바이트의 길이를 출력한다.
	// 또, 문자별 바이트가 다르기 때문에 문자열의 각 문자에 접근하기도 쉽지 않다.
	str5 := "스트링string"
	fmt.Println(len(str5))

	// 그렇다면 어떻게 문자의 길이를 계산하고, 각 문자에 접근할 수 있을까?
	// -> rune 슬라이스를 쓸 수 있다. rune은 int32의 별칭. 즉, 문자열을 int32의 슬라이스(동적 배열)
	// 로 나타내면, 그들의 요소 하나하나가 utf-8의 코드가 될 것이므로, rune 슬라이스의 길이는
	// 곧 문자열의 길이가 된다.
	// Go는 강타입 언어이지만, 편의를 위해 string과 []rune 타입의 변환을 지원한다.
	fmt.Println(len([]rune(str5)))

	str6 := []rune("string")
	for i := 0; i < len(str6); i++ { // 또는 for _, v := range str6
		fmt.Print(str6[i])
	}

	fmt.Println("-----------")
	str7 := "str7"
	str8 := str7
	str7 = "changed"
	fmt.Println("str7 : ", str7, "str8 : ", str8) // 복사 생성을 할 때 주소와 길이만 복사된다고 해도, str7만 변경된다면
	// 의도대로 str7만 변경한다. 이는 문자열이 불변 값을 가지기 때문에 가능하다. 즉, 지금까지는 사용된 적 없는 새로운 문자열을 가리키게 되면
	// 새로운 곳에 메모리를 할당한 다음 str7가 그 문자를 가리키게 한다.
	// 따라서, string타입의 합 연산을 너무 자주 사용하면 메모리 상 낭비가 발생하게 된다.
	// 이럴 때는 strings.Builder 를 사용할 수 있다.
}
