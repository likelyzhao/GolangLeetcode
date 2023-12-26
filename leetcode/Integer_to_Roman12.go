package leecode

var thousands = [4]string{"", "M", "MM", "MMM"}
var hundreds = [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var tens = [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var digits = [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

func intToRoman(num int) string {
	//rid := num
	var res string
	//	digits := make([]int, 0)
	res = thousands[num/1000]
	num = num % 1000
	res += hundreds[num/100]
	num = num % 100
	res += tens[num/10]
	num = num % 10
	res += digits[num]

	return res
}
