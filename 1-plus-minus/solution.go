// Package problem contains one coding problem.
package problem

const (
	target = 0

	errNotPossible = "not possible"
)

func PlusMinus(num int) string {
	if num <= 9 {
		return errNotPossible
	}

	length := 0
	for number := num; number != 0; number /= 10 {
		length++
	}

	numbers := make([]byte, length)
	for number, j := num, length-1; j >= 0; j-- {
		numbers[j] = byte(number % 10)
		number /= 10
	}

	last := length - 1

	var dfs func([]byte, int, int, byte, []byte) (string, bool)
	dfs = func(
		numbers []byte,
		index int,
		value int,
		previousOperation byte,
		buffer []byte,
	) (
		string,
		bool,
	) {
		number := int(numbers[index])

		if previousOperation == '-' {
			value -= number
		} else {
			value += number
		}

		if index == last {
			if value == target {
				return string(buffer), true
			}
		} else {
			result, ok := dfs(numbers, index+1, value, '-', append(buffer, '-'))
			if ok {
				return result, ok
			}

			return dfs(numbers, index+1, value, '+', append(buffer, '+'))
		}

		return errNotPossible, false
	}

	result, _ := dfs(numbers, 0, 0, '+', make([]byte, 0, last))
	return result
}
