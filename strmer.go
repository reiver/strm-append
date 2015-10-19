package verboten


import (
	"github.com/reiver/go-strm/driver"
)


const (
	// APPEND is a (string) constant that this Beginner driver
	// is registered under.
	APPEND = "APPEND"
)


func init() {
	strmDriver := newStrmer()

	strmdriver.RegisterStrmer(APPEND, strmDriver)
}


type internalStrmer struct{}


func newStrmer() strmdriver.Strmer {
	strmDriver := internalStrmer{

	}

	return &strmDriver
}



func (strmDriver *internalStrmer) Strm(src <-chan []interface{}, dst chan<- []interface{}, args ...interface{}) {

	// Execute.
	for datum := range src {

		dst <- datum

	}
	dst <- args
	close(dst)
}
