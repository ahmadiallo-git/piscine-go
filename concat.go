package z01

func Concat(str1 string, str2 string) string {
	faux := "Faux"
	var var1 string = str1
	var var2 string = str2
	if var1 != "" && var2 != "" {
		return (var1 + var2)
	}
	return faux
}
