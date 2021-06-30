package repository

import (
	"context"
	models2 "github.com/mgufrone/forex/delivery/grpc/rate/models"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

type dbCommandTest struct {
	suite.Suite
	db   *gorm.DB
	repo rate.ICommand
	lastEntry *rate.Rate
}

func (rp *dbCommandTest) SetupSuite() {
	rp.db, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
		QueryFields:                              true,
	})
	rp.db.AutoMigrate(&models2.Rate{})
	rp.repo = NewCommand(rp.db)
}

func (rp *dbCommandTest) TearDownSuite() {
	db, _ := rp.db.DB()
	db.Close()
}

func (rp *dbCommandTest) Test01Insert() {
	rt := rate.MustNew(
		"abc",
		"bcd",
		"somesite",
		"enote",
		0.001,
		0.002,
		time.Now(),
	)
	err := rp.repo.Create(context.Background(), rt)
	assert.Nil(rp.T(), err)
	var total int64
	rp.db.Model(&models2.Rate{}).Count(&total)
	assert.Equal(rp.T(), int64(1), total)
	rp.lastEntry = rt
}

func (rp *dbCommandTest) Test02Update() {
	rp.lastEntry.SetSymbol("cde")
	err := rp.repo.Update(context.Background(), rp.lastEntry)
	assert.Nil(rp.T(), err)
	var total int64
	rp.db.Model(&models2.Rate{}).Where("symbol = ?", "cde").Count(&total)
	assert.Equal(rp.T(), int64(1), total)
	rp.db.Model(&models2.Rate{}).Where("symbol = ?", "bcd").Count(&total)
	assert.Equal(rp.T(), int64(0), total)
	empty := rate.MustNew("base", "symb", "something", "something", 0.01, 0.02, time.Now())
	err = rp.repo.Update(context.Background(), empty)
	assert.NotNil(rp.T(), err)
}

func (rp *dbCommandTest) Test03Delete() {
	lastID := rp.lastEntry.ID()
	err := rp.repo.Delete(context.Background(), rp.lastEntry)
	assert.Nil(rp.T(), err)
	var total int64
	rp.db.Model(&models2.Rate{}).Where("id = ?", lastID).Count(&total)
	assert.Equal(rp.T(), int64(0), total)
}

func TestCommandRepository(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(dbCommandTest))
}
