package helper

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//`gorm:"foreignKey:ShipmentID"`
type Shipment struct {
	gorm.Model
	Packages []Package
	Data     string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Package struct {
	gorm.Model
	Data       string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
	ShipmentID uint
}

func (Shipment) TableName() string {
	return "Shipment"
}

func (Package) TableName() string {
	return "Package"
}

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=git_user password=P4p3r1n0 dbname=mydb port=5432 sslmode=disable TimeZone=Europe/Rome"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// errAutoMigrate := db.AutoMigrate(&Shipment{}, &Package{})

	errMigrate0 := db.Migrator().CreateTable(&Package{})

	if errMigrate0 != nil {
		return nil, errMigrate0
	}
	errMigrate1 := db.Migrator().CreateTable(&Shipment{})
	if errMigrate1 != nil {
		return nil, errMigrate1
	}
	return db, nil

}
