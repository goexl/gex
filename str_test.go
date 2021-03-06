package gex

import (
	`reflect`
	`testing`
)

func TestStr(t *testing.T) {
	toStringTests := []struct {
		in       []interface{}
		expected []string
	}{
		{[]interface{}{0, 1, 2}, []string{`0`, `1`, `2`}},
		{[]interface{}{0.01, 1.01, 2.01}, []string{`0.01`, `1.01`, `2.01`}},
		{[]interface{}{`test`, 1, 2}, []string{`test`, `1`, `2`}},
		{[]interface{}{`test`, 1.01, 2.01}, []string{`test`, `1.01`, `2.01`}},
		{[]interface{}{`test`, 1, 2.01}, []string{`test`, `1`, `2.01`}},
		{[]interface{}{`test`, `a`, `b`}, []string{`test`, `a`, `b`}},
	}

	for _, test := range toStringTests {
		actual := str(test.in...)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("str(%v) = %v；期望：%v", test.in, actual, test.expected)
		}
	}
}
