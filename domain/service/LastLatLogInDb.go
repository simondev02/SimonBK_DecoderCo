package service

import (
	"SimonBK_DecoderCo/domain/models"
	"fmt"

	"gorm.io/gorm"
)

func ConsuLatLongInDb(db *gorm.DB, imei string) (float64, float64, error) {
	var record models.AvlRecord

	err := db.Where("imei = ?", imei).Order("time_stamp_event desc").First(&record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, 0, nil
		}
		return 0, 0, fmt.Errorf("error al consultar el último registro: %v", err)
	}

	return *record.Latitude, *record.Longitude, nil
}
