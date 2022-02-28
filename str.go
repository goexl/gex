package exec

import (
	`fmt`
	`strconv`
)

func str(values ...interface{}) (to []string) {
	to = make([]string, 0, len(values))
	for _, value := range values {
		switch _value := value.(type) {
		case int8:
			to = append(to, strconv.FormatInt(int64(_value), 10))
		case *int8:
			to = append(to, strconv.FormatInt(int64(*_value), 10))
		case uint8:
			to = append(to, strconv.FormatUint(uint64(_value), 10))
		case *uint8:
			to = append(to, strconv.FormatUint(uint64(*_value), 10))
		case int:
			to = append(to, strconv.FormatInt(int64(_value), 10))
		case *int:
			to = append(to, strconv.FormatInt(int64(*_value), 10))
		case uint:
			to = append(to, strconv.FormatUint(uint64(_value), 10))
		case *uint:
			to = append(to, strconv.FormatUint(uint64(*_value), 10))
		case int32:
			to = append(to, strconv.FormatInt(int64(_value), 10))
		case *int32:
			to = append(to, strconv.FormatInt(int64(*_value), 10))
		case uint32:
			to = append(to, strconv.FormatUint(uint64(_value), 10))
		case *uint32:
			to = append(to, strconv.FormatUint(uint64(*_value), 10))
		case int64:
			to = append(to, strconv.FormatInt(_value, 10))
		case *int64:
			to = append(to, strconv.FormatInt(*_value, 10))
		case uint64:
			to = append(to, strconv.FormatUint(_value, 10))
		case *uint64:
			to = append(to, strconv.FormatUint(*_value, 10))
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
