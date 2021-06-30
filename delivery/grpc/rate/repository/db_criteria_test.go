package repository

import (
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/mgufrone/forex/internal/shared/criteria"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"testing"
	"time"
)

func TestSimpleCriteria01(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
		QueryFields:                              true,
		DryRun:                                   true,
	})
	cb := dbCriteriaBuilder{}
	db = cb.
		Where(criteria.NewCondition("something", criteria.Eq, "something")).
		Paginate(1, 10).
		Order("something", "desc").(dbCriteriaBuilder).apply(db)
	var res map[string]interface{}
	stmt := db.Table("something").Find(&res).Statement.SQL
	assert.Equal(t,
		"select * from `something` where something = ? order by `something` desc limit 10",
		strings.ToLower(stmt.String()))
}
func TestComplexCriteria01(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
		QueryFields:                              true,
		DryRun:                                   true,
	})
	cb := dbCriteriaBuilder{}
	db = cb.Or(
		cb.Where(criteria.NewCondition("something", criteria.Eq, "something")),
		cb.And(
			cb.Where(
				criteria.NewCondition("some1", criteria.Between, []time.Time{time.Now(), time.Now().Add(time.Hour)}),
				criteria.NewCondition("some2", criteria.Like, time.Now().Format(rate.DateLayout)),
			),
		)).
		Paginate(10, 10).
		Order("something", "desc").(dbCriteriaBuilder).apply(db)
	var res map[string]interface{}
	stmt := db.Table("something").Find(&res).Statement.SQL
	assert.Equal(t,
		"select * from `something` where (something = ? or ((some1 between ? and ?) and some2 like ?)) order by `something` desc limit 10 offset 90",
		strings.ToLower(stmt.String()))
}
