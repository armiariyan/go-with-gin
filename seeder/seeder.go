package seeder

import (
	"gitlab.com/armiariyan/intern_golang/entity"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func Opt_house(db *gorm.DB) []entity.Opt_house {
	return []entity.Opt_house{
		{
			ID: 1,
			Label: "belum punya",
		},
		{
			ID: 2,
			Label: "punya",
		},
		{
			ID: 3,
			Label: "tidak ada data",
		},
	}
}

func Opt_payment_frequency(db *gorm.DB) []entity.Opt_payment_frequency {
	return []entity.Opt_payment_frequency{
		{
			ID: 1,
			Label: "harian",
		},
		{
			ID: 2,
			Label: "mingguan",
		},
		{
			ID: 3,
			Label: "bulanan",
		},
		{
			ID: 4,
			Label: "triwulan",
		},
		{
			ID: 5,
			Label: "semester",
		},
		{
			ID: 6,
			Label: "tahunan",
		},
	}
}

func Opt_payment_type(db *gorm.DB) []entity.Opt_payment_type {
	return []entity.Opt_payment_type{
		{
			ID: 1,
			Label: "Installment",
		},
		{
			ID: 2,
			Label: "Bullet Payment",
		},
		{
			ID: 3,
			Label: "discount",
		},
		{
			ID: 4,
			Label: "grace period",
		},
		{
			ID: 5,
			Label: "lainnya",
		},
	}
}

func Opt_status(db *gorm.DB) []entity.Opt_status {
	return []entity.Opt_status{
		{
			ID: 1,
			Label: "lunas",
		},
		{
			ID: 2,
			Label: "belum lunas",
		},
	}
}

func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: Opt_house(db)},
		{Seeder: Opt_payment_frequency(db)},
		{Seeder: Opt_payment_type(db)},
		{Seeder: Opt_status(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}

	return nil
}