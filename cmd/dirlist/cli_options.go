package main

import (
	"fmt"
	"strconv"

	"github.com/ayushg3112/dirlist"
	"github.com/ayushg3112/dirlist/sort"
)

type cliOptions struct {
	rootDir   string
	sortOrder string
	sortField string
	port      int
	help      bool
	usages    string
}

func (c *cliOptions) validate() []string {
	valErrors := make([]string, 0)

	if c.rootDir == "" {
		valErrors = append(valErrors, "--root-dir is required")
	}

	if c.port == 0 {
		valErrors = append(valErrors, "--port cannot be 0")
	}

	_, err := sort.ToField(c.sortField)
	if err != nil {
		valErrors = append(valErrors, fmt.Sprintf("--sort-field error: %s", err.Error()))
	}

	_, err = sort.ToOrder(c.sortOrder)
	if err != nil {
		valErrors = append(valErrors, fmt.Sprintf("--sort-order error: %s", err.Error()))
	}

	return valErrors
}

func (c *cliOptions) toProcessingOptions() dirlist.ProcessingOptions {
	sortingField, err := sort.ToField(c.sortField)

	if err != nil {
		panic(err)
	}

	sortingOrder, err := sort.ToOrder(c.sortOrder)

	if err != nil {
		panic(err)
	}

	return dirlist.ProcessingOptions{
		RootDir:   c.rootDir,
		SortField: sortingField,
		SortOrder: sortingOrder,
		HTTPPort:  strconv.Itoa(c.port),
	}
}
