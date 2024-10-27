package seed

import "gorm.io/gorm"

type Seed struct {
	Name string
	Run  func(db *gorm.DB) error
}
