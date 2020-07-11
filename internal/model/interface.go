package model

type Mapper interface {
	ToMap() map[string]interface{}
}
