package dirlist

import "github.com/ayushg3112/dirlist/sort"

type ProcessingOptions struct {
	RootDir   string
	HTTPPort  string
	SortOrder sort.Order
	SortField sort.Field
}
