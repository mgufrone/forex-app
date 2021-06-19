package repository

import (
	"context"
	"github.com/mgufrone/forex/delivery/grpc/models"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

type dbRepositoryTest struct {
	suite.Suite
	db   *gorm.DB
	repo rate.IQuery
}

func (rp *dbRepositoryTest) SetupSuite() {
	rp.db, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
		QueryFields:                              true,
	})
	rp.db.AutoMigrate(&models.Rate{})
	rp.repo = NewQuery(rp.db)
}

func (rp *dbRepositoryTest) TearDownSuite() {
	db, _ := rp.db.DB()
	db.Close()
}

func (rp *dbRepositoryTest) Test01GetAllEmpty() {
	res, err := rp.repo.GetAll(context.Background())
	assert.Len(rp.T(), res, 0)
	assert.Nil(rp.T(), err)
}
func (rp *dbRepositoryTest) Test02GetAllSome() {
	rp.db.Create(&models.Rate{
		ID:         "random",
		Base:       "abc",
		Symbol:     "cde",
		Source:     "somesite",
		SourceType: "sometime",
		Sell:       0.003,
		Buy:        0.004,
		UpdatedAt:  time.Now(),
	})
	res, err := rp.repo.GetAll(context.Background())
	assert.Len(rp.T(), res, 1)
	assert.Nil(rp.T(), err)
	assert.Equal(rp.T(), "random", res[0].ID())
	assert.Equal(rp.T(), "abc", res[0].Base())
	assert.Equal(rp.T(), "cde", res[0].Symbol())
	assert.Equal(rp.T(), 0.004, res[0].Buy())
	assert.Equal(rp.T(), 0.003, res[0].Sell())
	cb := rp.repo.CriteriaBuilder().Where(rate.WhereSymbol("USD"))
	res, err = rp.repo.Apply(cb).GetAll(context.Background())
	assert.Len(rp.T(), res, 0)
	assert.Nil(rp.T(), err)
}

func (rp *dbRepositoryTest) Test03Count() {
	cb := rp.repo.CriteriaBuilder().Where(rate.WhereSymbol("BTC"))
	total, err := rp.repo.Apply(cb).Count(context.Background())
	assert.Equal(rp.T(), int64(0), total)
	assert.Nil(rp.T(), err)

	cb = rp.repo.CriteriaBuilder().Where(rate.WhereSource("somesite"))
	total, err = rp.repo.Apply(cb).Count(context.Background())
	assert.Equal(rp.T(), int64(1), total)
	assert.Nil(rp.T(), err)
}
func (rp *dbRepositoryTest) Test04FindByID() {
	res, err := rp.repo.FindByID(context.Background(), "something")
	assert.Nil(rp.T(), res)
	assert.NotNil(rp.T(), err)

	res, err = rp.repo.FindByID(context.Background(), "random")
	assert.NotNil(rp.T(), res)
	assert.Nil(rp.T(), err)
}

func TestQuery(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(dbRepositoryTest))
}
