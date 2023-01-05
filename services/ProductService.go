package services

import (
	"fmt"
	"parte2/models"
	"regexp"
	"time"
)


func ValidateCodeValue(products []models.Product, codeValue string) (isPresent bool) {
	for _, product := range products {
		if(product.CodeValue == codeValue) {
			isPresent = true
			break
		}
	}
	return
}

func ValidateDate(products []models.Product, date string) bool {

	expresion := "[0-3][0-9]/[0-1][0-9]/[0-2][0-9][0-9][0-9]"
	match, _ := regexp.MatchString(expresion, date)
	
	if !match {
		return true
	} else {
		_, err := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", date[6:10], date[3:5], date[0:2]))

		fmt.Println(fmt.Sprintf("%s-%s-%s", date[6:10], date[3:5], date[0:2]))
		
		if err != nil {
			return true
		}
	
	}
	return false
	
}