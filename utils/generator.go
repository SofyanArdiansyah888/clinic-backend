package utils

import (
	"backend/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func GenerateID(db *gorm.DB, model string, update bool) string {
	now := time.Now()
	year, month := now.Format("06"), now.Format("01")
	yearMonth := year + month

	var seq models.MonthlySequence
	err := db.Where("model = ? AND year_month = ?", model, yearMonth).First(&seq).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		seq = models.MonthlySequence{
			Model:     model,
			YearMonth: yearMonth,
			Counter:   1,
		}
		if update {
			if err := db.Create(&seq).Error; err != nil {
				return ""
			}
		}
	} else if err != nil {
		return ""
	} else {
		// Always increment counter
		seq.Counter++
		if update {
			if err := db.Model(&seq).Update("counter", seq.Counter).Error; err != nil {
				return ""
			}
		}
	}

	id := fmt.Sprintf("%s-%s%s%05d", model, year, month, seq.Counter)
	return id
}
