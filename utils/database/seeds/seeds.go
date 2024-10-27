package seeds

import (
	"gorm.io/gorm"
	"starter-template/utils/database/seed"
)

func All() []seed.Seed {
	var seeds []seed.Seed = []seed.Seed{
		{
			Name: "Create Admin",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "admin", "admin@gmail.com", "0123456789")
			},
		},
	}
	return seeds
}
