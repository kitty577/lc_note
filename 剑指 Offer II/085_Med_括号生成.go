package main

func generateParenthesis(n int) []string {
	res := []string{}
	var find func(restLeft, restRight int, tmp []byte)
	find = func(restLeft, restRight int, tmp []byte) {
		if restLeft < 0 || restRight < 0 || restLeft > restRight {
			return
		}

		if restLeft == 0 && restRight == 0 {
			res = append(res, string(tmp))
			return
		}

		find(restLeft-1, restRight, append(tmp, '('))
		find(restLeft, restRight-1, append(tmp, ')'))
	}

	find(n, n, []byte{})

	return res
}
