package internal

import "fmt"

type ParseError struct {
	want tokenId
	got  token
}

func (e ParseError) Error() string {
	return fmt.Sprintf("%d:%d - expected %v, found %v", e.got.row, e.got.col, e.want, e.got.id)
}
