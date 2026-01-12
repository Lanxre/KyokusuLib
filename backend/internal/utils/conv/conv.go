package conv

import "strconv"

func StringToInt(s string) int {
    res, _ := strconv.Atoi(s)
    return res
}