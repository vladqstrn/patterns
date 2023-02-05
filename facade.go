package main

type SubSystemA struct{}

func (s *SubSystemA) MethodA() string {
	return "SubSystemA MethodA"
}

type SubSystemB struct{}

func (s *SubSystemB) MethodB() string {
	return "SubSystemB MethodB"
}

type Facade struct {
	subA *SubSystemA
	subB *SubSystemB
}

func NewFacade() *Facade {
	return &Facade{
		subA: &SubSystemA{},
		subB: &SubSystemB{},
	}
}

func (f *Facade) MethodFacade() string {
	return f.subA.MethodA() + " " + f.subB.MethodB()
}

//Шаблон Facade — это шаблон проектирования, используемый в разработке программного обеспечения,
//который обеспечивает упрощенный интерфейс для сложной системы классов, библиотек или сервисов.
//Идея шаблона состоит в том, чтобы скрыть детали реализации подсистемы за единым унифицированным интерфейсом.
//Это упрощает использование подсистемы и помогает уменьшить связь между подсистемой и клиентским кодом, который ее использует.

//SubSystemA и SubSystemB — это две отдельные подсистемы с разными методами.
//Структура Facade действует как фасад, предоставляя упрощенный интерфейс для этих подсистем с помощью метода MethodFacade.
//Это упрощает использование подсистем, поскольку клиентскому коду нужно взаимодействовать только с Фасадом, и ему не нужно знать об отдельных подсистемах.
