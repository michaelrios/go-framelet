package models

type Viewable interface {
	Bytes() ([]byte, error)
}