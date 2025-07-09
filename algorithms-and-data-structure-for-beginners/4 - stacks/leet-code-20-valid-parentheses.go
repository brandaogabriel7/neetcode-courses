package stacks

func isValid(s string) bool {
	parentheses_map := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	st := make([]rune, 0)
	for _, ch := range s {
		if _, exists := parentheses_map[ch]; exists {
			st = append(st, ch)
		} else {
			if len(st) == 0 {
				return false
			}

			last_ch := st[len(st)-1]
			if parentheses_map[last_ch] == ch {
				st = st[:len(st)-1]
			} else {
				return false
			}
		}
	}

	return len(st) == 0
}

type stack []rune

func (s stack) Push(v rune) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, rune) {
	l := len(s)

	if l == 0 {
		return nil, 0
	}

	return s[:l-1], s[l-1]
}

func (s stack) Peek() rune {
	l := len(s)
	if l == 0 {
		return 0
	}

	return s[l-1]
}

func isValid2(s string) bool {
	// map each open parentheses to its corresponding closing parentheses
	parentheses_map := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	st := make(stack, 0)
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			st = st.Push(ch)
		} else {
			last_ch := st.Peek()
			if parentheses_map[last_ch] == ch {
				st, _ = st.Pop()
			} else {
				return false
			}
		}
	}
	return len(st) == 0
}
