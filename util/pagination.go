package util

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// These vars have to be synced manually with kb yaml
var (
	// LimitDefault is the default amount of rows that will be retrieved
	// (if that much rows exist)
	LimitDefault = 200
	// LimitMax max amount rows that can be retrieved
	LimitMax = 400

	// OffsetDefault is the default value of the offset it it is not specified
	OffsetDefault = 0
)

// PaginationLimitParseAndValidate parses limit value and validate it against
// the constraints
func PaginationLimitParseAndValidate(limit string) (int, error) {
	if limit == "" {
		return LimitDefault, nil
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return 0, errors.Wrap(err, "unable to convert the limit value")
	}

	if limitInt < 0 || limitInt > LimitMax {
		return 0, fmt.Errorf("limit parameter with value %d does not satisfy condition: 0 < limit < %d",
			limitInt, LimitMax)
	}

	return limitInt, nil
}

// PaginationOffsetParseAndValidate parses offset value and validate it against
// the constraints
func PaginationOffsetParseAndValidate(offset string) (int, error) {
	if offset == "" {
		return OffsetDefault, nil
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return 0, errors.Wrap(err, "Problems converting the limit value")
	}

	if offsetInt < 0 {
		return 0, fmt.Errorf("limit parameter with value %d does not satisfy condition: 0 < limit ",
			offsetInt)
	}

	return offsetInt, nil
}
