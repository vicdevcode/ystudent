package v1

import (
	"errors"
	"strconv"

	"github.com/vicdevcode/ystudent/main/internal/dto"
)

func GetPage(page string, count string) (dto.Page, error) {
	var currentPage, currentCount int
	var err error
	if page == "" {
		currentPage = 1
	} else {
		currentPage, err = strconv.Atoi(page)
		if err != nil {
			return dto.Page{
				Page:  1,
				Count: 10,
			}, err
		}
		if currentPage < 1 {
			return dto.Page{
				Page:  1,
				Count: 10,
			}, errors.New("page must be greater than 0")
		}
	}
	if count == "" {
		currentCount = 10
	} else {
		currentCount, err = strconv.Atoi(count)
		if err != nil {
			return dto.Page{
				Page:  currentPage,
				Count: 10,
			}, err
		}
		if currentCount < 1 {
			return dto.Page{
				Page:  currentPage,
				Count: 10,
			}, errors.New("count must be greater than 0")
		}
	}

	return dto.Page{
		Page:  currentPage,
		Count: currentCount,
	}, nil
}
