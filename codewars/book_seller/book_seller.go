package book_seller

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(books []string, categories []string) string {
	result := make(map[string]int) 
	for _, book := range books {
		splited := strings.Split(book, " ")
		firstChar := string(splited[0][0])
		amount, err := strconv.Atoi(splited[1])
		if err != nil {
			fmt.Println(err)
			return ""
		}
		result[firstChar] += amount
	}

	if len(result) != 0 {
		resultList := []string {}
		for _, cate := range categories {
			resultList = append(resultList, fmt.Sprintf("(%v : %v)", cate, result[cate]))
		}
		return strings.Join(resultList, " - ")
	}

	
	return ""
}