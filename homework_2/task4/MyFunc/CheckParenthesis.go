package myfunc

import "strconv"

func CheckParenthesis(formula string) string {
	var openParenthesis int
	var closeParenthesis int
	var positionParenthesis int
	var message string

	for _, v := range formula {
		if v == '(' {
			openParenthesis++
			positionParenthesis++
		} else if v == ')' {
			closeParenthesis++
			positionParenthesis--

			if positionParenthesis < 0 {
				message = "Скобки неверно установлены! "
			}
		}
	}
	if openParenthesis == 0 && closeParenthesis == 0 {
		message = "В вашей формуле нет скобок!"
	} else if openParenthesis == closeParenthesis && positionParenthesis == 0 {
		message += "Количество верно, " + strconv.Itoa(openParenthesis) + " открывающие, " + strconv.Itoa(closeParenthesis) + " закрывающие"
	} else {
		message += "Количество неверно, " + strconv.Itoa(openParenthesis) + " открывающие, " + strconv.Itoa(closeParenthesis) + " закрывающие"
	}

	return message
}
