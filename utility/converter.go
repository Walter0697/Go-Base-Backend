package utility

import (
	"strconv"
)

func ConvertStringToUint(input string) (uint, error) {
	output, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(output), nil
}

func IsAuthorized(userrole uint, authorized_role uint) bool {
	return userrole | authorized_role == authorized_role
}