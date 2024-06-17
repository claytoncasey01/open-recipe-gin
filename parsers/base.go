package parsers

import "mime/multipart"

// Base generic parser interface
type Parser interface {
	Parse(file multipart.File) (any, error)
}

func ParseFile(parser Parser, file multipart.File) {
	parser.Parse(file)
}
