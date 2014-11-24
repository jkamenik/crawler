package main

type HttpError struct {
	original string
}

func (self HttpError) Error() string {
	return self.original
}
