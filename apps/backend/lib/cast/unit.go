package cast

import "strconv"

func Unit(value string) (uint, error) {
	if value == "" {
		return 0, nil
	}
	unit, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return uint(unit), nil
}