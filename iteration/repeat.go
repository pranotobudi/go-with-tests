package iteration

func Repeat(s string) string {
	res := ""
	for i := 0; i < 5; i++ {
		res = res + s
	}
	return res
}
