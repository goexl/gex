package gex

import (
	`fmt`
	`strconv`
)

func toString(values ...interface{}) (to []string) {
	to = make([]string, 0, len(values))
	for _, value := range values {
		switch _value := value.(type) {
		case int8, *int8, uint8, *uint8, int, *int, uint, *uint, int32, *int32, uint32, *uint32, int64, *int64, uint64, *uint64:
			to = append(to, fmt.Sprintf(`%d`, _value))
		case float32:
			to = append(to, strconv.FormatFloat(float64(_value), 'f', -1, 64))
		case *float32:
			to = append(to, strconv.FormatFloat(float64(*_value), 'f', -1, 64))
		case float64:
			to = append(to, strconv.FormatFloat(_value, 'f', -1, 64))
		case *float64:
			to = append(to, strconv.FormatFloat(*_value, 'f', -1, 64))
		case bool:
			to = append(to, strconv.FormatBool(_value))
		case *bool:
			to = append(to, strconv.FormatBool(*_value))
		case string:
			to = append(to, _value)
		case *string:
			to = append(to, *_value)
		case fmt.Stringer:
			to = append(to, value.(fmt.Stringer).String())
		default:
			to = append(to, fmt.Sprintf(`%v`, value))
		}
	}

	return
}
