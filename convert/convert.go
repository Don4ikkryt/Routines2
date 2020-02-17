package convert

import "strconv"

func ToUint(i interface{}) (uint, bool) {
	switch number := i.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return number.(uint), true
	case string:
		intNum, err := strconv.Atoi(number)
		if err != nil {
			return 0, false
		} else {
			return uint(intNum), true
		}

	default:
		return 0, false
	}

}
func ToInt32(i interface{}) (int32, bool) {
	switch number := i.(type) {
	case int, int8, int16, int64, float32, float64, uint, uint16, uint32, uint64:
		return number.(int32), true
	case string:
		intNum, err := strconv.Atoi(number)
		if err != nil {
			return 0, false
		} else {
			return int32(intNum), true
		}

	default:
		return 0, false
	}

}

func ToInt64(i interface{}) (int64, bool) {
	switch number := i.(type) {
	case int, int8, int16, int32, float32, float64, uint, uint16, uint32, uint64:
		return number.(int64), true
	case string:
		intNum, err := strconv.Atoi(number)
		if err != nil {
			return 0, false
		} else {
			return int64(intNum), true
		}

	default:
		return 0, false
	}

}
func ToFloat32(i interface{}) (float32, bool) {
	switch number := i.(type) {
	case int, int8, int16, int32, int64, float64, uint, uint16, uint32, uint64:
		return number.(float32), true
	case string:
		intNum, err := strconv.ParseFloat(number, 32)
		if err != nil {
			return 0, false
		} else {
			return float32(intNum), true
		}

	default:
		return 0, false
	}

}
func ToFloat64(i interface{}) (float64, bool) {
	switch number := i.(type) {
	case int, int8, int16, int32, int64, float32, uint, uint16, uint32, uint64:
		return number.(float64), true
	case string:
		intNum, err := strconv.ParseFloat(number, 32)
		if err != nil {
			return 0, false
		} else {
			return intNum, true
		}

	default:
		return 0, false
	}

}
