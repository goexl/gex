package gex

import (
	`fmt`
	`strconv`
)

func toString(values ...interface{}) (to []string) {
	to = make([]string, 0, len(values))
	for _, value := range values {
		switch value.(type) {
		case int8, *int8, uint8, *uint8, int, *int, uint, *uint, int32, *int32, uint32, *uint32, int64, *int64, uint64, *uint64:
			to = append(to, fmt.Sprintf(`%d`, value))
		case float32:
			to = append(to, strconv.FormatFloat(float64(value.(float32)), 'f', -1, 64))
		case *float32:
			to = append(to, strconv.FormatFloat(float64(*value.(*float32)), 'f', -1, 64))
		case float64:
			to = append(to, strconv.FormatFloat(value.(float64), 'f', -1, 64))
		case *float64:
			to = append(to, strconv.FormatFloat(*value.(*float64), 'f', -1, 64))
		case bool, *bool:
			to = append(to, fmt.Sprintf(`%t`, value))
		case string:
			to = append(to, value.(string))
		case *string:
			to = append(to, *value.(*string))
		case fmt.Stringer:
			to = append(to, value.(fmt.Stringer).String())
		default:
			to = append(to, fmt.Sprintf(`%v`, value))
		}
	}

	return
}
