package main

import "fmt"

// error타입
// error타입은 interface로, string을 반환하는 메소드를 포함한다.
// 어떠한 struct라도 string을 반환하는 Error() 메소드만 포함하고 있으면 에러로 사용할 수 있다.
// type error interface {
// Error() string
// }

// 에러에 여러 정보를 담을 수 있다. 아래의 Error()메서드를 포함하고 있으므로, 에러로 쓰일 수 있다.
type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string { // 의문점 :reciever에 *를 달아서 주소를 받게 하면 implement를 하지 않았다는 에러 발생.
	return "암호의 길이가 짧습니다."
}

func MakePassword(password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8} // PasswordError는 Error()메서드를 가지고 있으므로,
	}
	return nil
}

func ToMinus(num int) (int, error) {
	if num < 0 {
		return num, fmt.Errorf("이미 음수입니다 : %d\n", num)
	}
	return -1 * num, nil
}

func main() {
	_, err := ToMinus(-3)
	if err != nil {
		fmt.Println("ERROR : ", err)
	}

	ret := MakePassword("wef")
	if ret != nil {
		errorInfo := ret.(PasswordError) // 인터페이스 변환. 왜 인터페이스 변환 해야하는가? 왜냐하면 MakePassword는 error타입을 반환한 것이기 때문에, 구체적으로 이것이 어떤 type인지 몰라서 Len 등의 필드를 가지는지 알 수 없다.
		fmt.Println(errorInfo.Error())
		fmt.Println("비밀번호 길이 : ", errorInfo.Len, ", 요구되는 최소 길이 : ", errorInfo.RequireLen)

	}

}
