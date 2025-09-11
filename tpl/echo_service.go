package tpl

func EchoServiceTemplate() []byte {
	return []byte(`/*
{{ .GetCopyright }}
*/
package service

type YourService struct {
	element1 string
	element2 int
	element3 any
}

func NewYourService(element1 string, element2 int, element3 interface{}) *YourService {
	return &YourService{
		element1,
		element2,
		element3,
	}
}

func (s *YourService) Func1() []error {
	var errs []error
	/* TODO: IMPLEMENT IT*/
	return errs
}

func (s *YourService) Func2() []error {
	var errs []error
	/* TODO: IMPLEMENT IT*/
	return errs
}

func (s *YourService) Func3() []error {
	var errs []error
	/* TODO: IMPLEMENT IT*/
	return errs
}

`)
}
