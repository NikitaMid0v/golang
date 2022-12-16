package utils

import (
	"errors"
	"strconv"
	"strings"
)

func ParseIdFromUrl(Path string) (int, error) {
	id := strings.TrimPrefix(Path, "/promotions/")

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("id must be integer")
	}

	return convertedId, nil
}
