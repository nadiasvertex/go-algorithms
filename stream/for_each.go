package stream

import "github.com/nadiasvertex/algorithms/common"

// ForEach processes each item in the stream. It will not stop until the input
// stream terminates.
func ForEach[T any](s Stream[T], op common.UnaryOp[T]) {
	for {
		if v, atEnd := s.Next(); atEnd {
			return
		} else {
			op(v)
		}
	}
}
