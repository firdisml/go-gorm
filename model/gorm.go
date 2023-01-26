package model

func GetAllGorms() ([]Gorm, error) {
	var gorms []Gorm

	tx := db.Find(&gorms)

	if tx.Error != nil {
		return []Gorm{}, tx.Error
	}

	return gorms, nil
}

func GetGorm(id uint64) (Gorm, error) {
	var gorm Gorm

	tx := db.Where("id = ?", id).First(&gorm)

	if tx.Error != nil {
		return Gorm{}, tx.Error
	}

	return gorm, nil
}

func CreateGorm(gorm Gorm) error {
	tx := db.Create(&gorm)
	return tx.Error
}

func UpdateGorm(gorm Gorm) error {

	tx := db.Save(&gorm)
	return tx.Error
}

func DeleteGorm(id uint64) error {

	tx := db.Unscoped().Delete(&Gorm{}, id)
	return tx.Error
}

func FindByGormUrl(url string) (Gorm, error) {
	var gorm Gorm
	tx := db.Where("gorm = ?", url).First(&gorm)
	return gorm, tx.Error
}
