package conv

import "strconv"

func StringToInt(s string) int {
    res, _ := strconv.Atoi(s)
    return res
}

func AbsInt(x int64) int64 {

	if x < 0 {
        return -x
    }
    
    return x
}