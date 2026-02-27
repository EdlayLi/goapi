package stat

import (
	"apigo/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	Database *db.Db
}

func NewStatRepository(database *db.Db) *StatRepository {
	return &StatRepository{
		Database: database,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	carrentDate := datatypes.Date(time.Now())

	repo.Database.Find(&stat, "link_id = ? and date = ?", linkId, carrentDate)
	if stat.ID == 0 {
		repo.Database.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   carrentDate,
		})
	} else {
		stat.Clicks += 1
		repo.Database.Save(&stat)
	}
}
