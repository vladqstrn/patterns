package main

type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee SpecificRequest"
}

type Target interface {
	Request() string
}

type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{
		adaptee: adaptee,
	}
}

//Шаблон адаптера — это шаблон проектирования, который позволяет объектам с несовместимыми интерфейсами работать вместе.
//Шаблон включает в себя создание нового класса, который обертывает существующий объект и сопоставляет его методы с другим интерфейсом.
//Это позволяет объектам, которые ранее были несовместимы, беспрепятственно работать вместе.

//Adaptee — это существующий класс с определенным интерфейсом (SpecificRequest). Класс адаптера реализует целевой интерфейс, который требуется клиентскому коду.
//Адаптер сопоставляет метод Request с методом SpecificRequest Adaptee, позволяя клиентскому коду использовать Adaptee через адаптер, как если бы он был целью.
//Это позволяет клиентскому коду использовать Adaptee, даже если он изначально не был разработан для удовлетворения требований клиента.
