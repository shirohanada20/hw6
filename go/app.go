package app

import (
	"fmt"
	"net/http"
)



func init() {
	http.HandleFunc("/", handlePata)
}

func handlePata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	first_word := r.FormValue("first_word")
	second_word := r.FormValue("second_word")
	result := ""


	fmt.Fprintf(w, `<!DOCTYPE html>
		<html>
		<head>
		<title>パタトクカシー</title>
		</head>
		<body>`)

	first_rune := []rune(first_word)
	second_rune := []rune(second_word)
	first_len := len(first_rune)
	second_len := len(second_rune)
	count := 0

	for i:= 0; i < first_len; i++ {
		count = i
		result = result + string(first_rune[i]) + string(second_rune[i])
		if i == second_len - 1 {
			break
		}
	}

	if count < second_len - 1 {
		for i := count + 1; i < second_len; i++ {
			result += string(second_rune[i])
		}
	}

	if count < first_len -1 {
		for i := count + 1; i < first_len; i++ {
			result += string(first_rune[i])
		}
	}

	//fmt.Fprintf(w, "%d\n", count)

	fmt.Fprintf(w, "%s\n", result)

	fmt.Fprintf(w, `<form>
		<input type = text name = first_word><br>
		<input type = text name = second_word><br>
		<input type = submit>
		</form>
		</body>
		</html>`)
}
