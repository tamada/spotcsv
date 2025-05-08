package spotcsv

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Indexes struct {
	Indexces []uint
}

func NewIndexes(indexces []uint) *Indexes {
	return &Indexes{
		Indexces: indexces,
	}
}

func NewIndexesFromString(str string) (*Indexes, error) {
	errs := []error{}
	result := []uint{}
	for _, s := range strings.Split(str, ",") {
		if strings.Contains(s, ":") {
			r, err := newRangeFromString(s)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			result = append(result, r...)
		} else {
			index, err := strconv.Atoi(s)
			if err != nil || index < 0 {
				errs = append(errs, fmt.Errorf("invalid range: %s (each index should be positive number)", str))
			}
			result = append(result, uint(index))
		}
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return NewIndexes(result), nil
}

func newRangeFromString(str string) ([]uint, error) {
	list := strings.Split(str, ":")
	if len(list) != 2 {
		return nil, fmt.Errorf("invalid range: %s (range must be <from:to>)", str)
	}
	from, err1 := strconv.Atoi(list[0])
	to, err2 := strconv.Atoi(list[1])
	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("invalid range: %s (both <from> and <to> should be positive number <from:to>)", str)
	}
	if from >= to {
		return nil, fmt.Errorf("invalid range: %s (<from> should be less than <to>)", str)
	}
	r := []uint{}
	for i := from; i <= to; i++ {
		r = append(r, uint(i))
	}
	return r, nil
}
