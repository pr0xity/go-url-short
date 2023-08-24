package model

func GetAllShrts() ([]Shrt, error) {
	var shrts []Shrt
	tx := db.Find(&shrts)
	if tx.Error != nil {
		return []Shrt{}, tx.Error
	}

	return shrts, nil

}

func GetShrtById(id int) (Shrt, error) {
	var shrt Shrt
	tx := db.First(&shrt, id)

	if tx.Error != nil {
		return Shrt{}, tx.Error
	}

	return shrt, nil
}

func CreateShrt(shrt Shrt) error {
	tx := db.Create(&shrt)
	return tx.Error
}

func UpdateShrt(shrt Shrt) error {
	tx := db.Save(&shrt)
	return tx.Error
}

func DeleteShrt(id int) error {
	tx := db.Unscoped().Delete(&Shrt{}, id)
	return tx.Error
}

func FindByShrtUrl(url string) (Shrt, error) {
	var shrt Shrt
	tx := db.Where("shrt = ?", url).First(&shrt)
	return shrt, tx.Error
}
