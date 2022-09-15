# **객체지향 설계 5가지 원칙 SOLID**

## **1. SOLID란?**
1. 단일 책임 원칙(Single Responsibility Principle, SRP)
2. 개방-폐쇄 원칙(Open-Closed Princibple, OCP)
3. 리스코프 치환 원칙(Liskov Substitution Principle, LSP)
4. 인터페이스 분리 원칙(Interface Segregation Principle, ISP)
5. 의존 관계 역전 원칙(Dependency Inversion Principle, DIP)

이 원칙들을 지켜서 프로그램을 설계하면, 지속 가능하면서도 결합도는 낮추고 응집도는 높은 좋은 코드를 쓸 수 있고, 다른 프로젝트에서도 쓸 수 있는 경우가 많아져서 생산성이 높아진다.

### **1. Single Responsibility Principle, SRP**
모든 객체는 하나의 책임을 져야만 한다. 보고서는 보고만 하면 되지, 보고서를 보내는 기능을 가질 필요는 없다. 예시 코드를 읽어보면서 Strategy pattern 처럼 구성 위주의 코드를 짜는게 이 원칙을 지키는데 유리하다는 인상을 받았다.

### **2. Open-Closed Principle**
확장에는 열려있고 변경에는 닫혀있어야 한다. 즉, 새로운 기능을 확장할 때 기존의 코드를 변경하는 것을 지양하도록 설계해야 한다. 마찬가지로 전략 패턴과 같이 상속 위주의 패턴보다 구성 위주의 코드를 짜는게 이 원칙을 지키는데 유리하다는 생각이 들었다. 또 1. Single responsibility principle과 밀접한 관계가 있다는 생각이 들었다.

### **3. Liscov Substitution Principle**
q(x)를 타입 T의 객체 x에 대해 증명할 수 있는 속성이라고 할 때, 타입 T의 하위 타입 S에 대하여 타입 S 객체 y도 q(y)가 성립하여야 한다. 즉, 어떠한 특정 의도대로 만들어진 함수(일종의 함수 구현자와 호출자의 합의된 계약 성립)가, T의 하위 객체에 대하여도 제대로 이행되어야 한다는 뜻.

```Go
type T interface {
    String() string
}

type S type {
    firstField int
}

type SS type {
    secondField int
}

func (rec *S) String() string { // 인터페이스 T를 구현
    return rec.firstField
}

func (rec *SS) String() string { // 인터페이스 T를 구현
    return rec.secondField
}

// q의 구현은, 이 함수의 호출자와 구현자의 합의된 의도에 의해 일종의 "계약"이 발생한다고 볼 수 있다.
func q(t T) {
    if _, ok := t.(*SS); ok { // 그런데 T의 하위 객체  t(SS)가 의도된 대로 Println을 이행하지 않고 panic해버린다면 리스코프 치환 법칙을 위반하는 코드이다.
        panic("Can't be SS")
    }

    fmt.Println(t.String()) // T를 구현한 모든 타입에 대해 
}

x := &S{"firstField"}
y := &SS{"secondField"}

// 모두 T를 구현하였으므로 q(x), q(y)가 잘 동작해야 한다.
q(x)
q(y)

```



## **2. 그렇다면 나쁜 설계는 무엇인가?**
1. 경직성(Rigidity) : 모듈간의 결합도가 높아서 코드를 변경하기 힘들다.
2. 취약성(Frigility) : 한 부분을 건드리면 다른 부분까지 망가지는 경우. 해당 프로그램을 수정하기 힘들게 만든다.
3. 부동성(Immobility) : 결합도가 높아서 다른 코드로 이식할 수 없는 경우.

