package db

// This file is responsible for providing the database functions
// The REST controllers handling methods will call these

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sevren/test/models"
	log "github.com/sirupsen/logrus"
)

type Dao struct {
	DB *gorm.DB
}

//StoreUsedLicenses - Stores each license recieved from the rabbitMQ exchange
func (d *Dao) StoreUsedLicenses(license string) error {

	l := models.Used_licenses{License: license}
	var res *gorm.DB
	db := d.DB.Table("used_licenses")

	existing_item := models.Used_licenses{}
	// Check if the database already contains that license
	if db.Where("license = ?", license).Take(&existing_item).RecordNotFound() {
		res = db.Create(l)

		if res.Error != nil {
			log.Error(res.Error)
			return res.Error
		} else {
			log.Infof("Inserted %s in the used_licenses table\n", license)
		}
	}
	return nil
}

//GetLicenses - Retrieves the Users licenses from the database
func (d *Dao) GetLicenses(user string) ([]string, error) {
	u := models.User_licenses{}
	if err := d.DB.Where(&models.User_licenses{Username: user}).First(&u).Error; err != nil {
		return nil, errors.New("User: " + user + "has no licenses")
	}
	return u.Lics, nil
}

//GetLicenses - Retrieves the Users licenses from the database
func (d *Dao) GetUsedLicenses() ([]string, error) {
	l := []models.Used_licenses{}
	db := d.DB.Table("used_licenses")
	if err := db.Find(&l).Error; err != nil {
		return nil, errors.New("No licenses used")
	}
	var usedLicenses = []string{}
	for _, license := range l {
		usedLicenses = append(usedLicenses, license.License)
	}

	return usedLicenses, nil
}
