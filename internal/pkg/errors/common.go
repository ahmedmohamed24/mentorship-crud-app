package errors

import "errors"

var ErrInvalidDocumentID = errors.New("invalid document ID")
var ErrDocumentNotFound = errors.New("document not found")
var ErrInternal = errors.New("internal server error")
var ErrInvalidPage = errors.New("invalid page number")
