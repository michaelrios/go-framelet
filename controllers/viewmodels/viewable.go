package viewmodels

type Viewable interface {
	Bytes() ([]byte, error)
}
