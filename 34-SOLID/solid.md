# **객체지향 설계 5가지 원칙 SOLID**

## **1. SOLID란?**
1. 단일 책임 원칙(Single Responsibility Principle, SRP)
2. 개방-폐쇄 원칙(Open-Closed Princibple, OCP)
3. 리스코프 치환 원칙(Liskov Substitution Principle, LSP)
4. 인터페이스 분리 원칙(Interface Segregation Principle, ISP)
5. 의존 관계 역전 원칙(Dependency Inversion Principle, DIP)

이 원칙들을 지켜서 프로그램을 설계하면, 지속 가능하면서도 결합도는 낮추고 응집도는 높은 좋은 코드를 쓸 수 있고, 다른 프로젝트에서도 쓸 수 있는 경우가 많아져서 생산성이 높아진다.

## **2. 그렇다면 나쁜 설계는 무엇인가?**
1. 경직성(Rigidity) : 모듈간의 결합도가 높아서 코드를 변경하기 힘들다.
2. 취약성(Frigility) : 한 부분을 건드리면 다른 부분까지 망가지는 경우. 해당 프로그램을 수정하기 힘들게 만든다.
3. 부동성(Immobility) : 결합도가 높아서 다른 코드로 이식할 수 없는 경우.