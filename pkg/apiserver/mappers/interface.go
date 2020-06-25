package mappers

type Mapper interface {
	ToOutput() interface{}
	FromInput(interface{}) interface{}
}
