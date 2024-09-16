package piscine

func ConcatParams(args []string) string {
	var res_args string
	for i := range args {
		res_args += args[i]
		if i != len(args)-1 {
			res_args += "\n"
		}
	}
	return res_args
}
