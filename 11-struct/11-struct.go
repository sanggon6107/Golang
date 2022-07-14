package main

import "fmt"

func main() {
	// 변수의 배치 : memory alignment와 memory padding에 조심하면 메모리 효율성을 증가시킬 수 있다.
	type Student struct { // 타입의 이름이 대문자이면 패키지 외부로 공개되는 타입.
		Name  string // 필드명과 필드 타입
		Class int
		No    int
		Score float64
	}

	type VipStudent struct {
		Student    Student
		ParentName string
		Score      float64
	}

	var a Student // 초기값을 생략하면 모든 필드가 기본값으로 초기화 된다.
	a.Name = "Kim"
	a.Class = 1
	a.No = 1
	a.Score = 80.0

	// 모든 값 초기화
	var b Student = Student{"Lee", 1, 2, 90.0}
	fmt.Println(b.Score)

	// 일부 값 초기화
	var c Student = Student{Name: "Choi", Score: 85.1}
	fmt.Println(c.Name)

	d := c // 변수명 대입으로 복사할 수 있다.
	fmt.Println(d.Name)

	v := VipStudent{Student{"Choi", 1, 3, 100.0}, "Kim", 90.0}
	// fmt.Println(v.No)
	fmt.Println(v.Score) // 포함된 필드의 이름과 바깥 구조체의 필드 이름이 겹칠 때는 호출한 변수의 타입이 가진 쪽이 선택된다.
	// 만약 Student의 Score을 호출하고 싶다면 v.Student.Score

}
