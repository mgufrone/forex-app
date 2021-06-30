package repository

import (
	"context"
	"fmt"
	models2 "github.com/mgufrone/forex/delivery/grpc/rate/models"
	"github.com/mgufrone/forex/internal/domains/rate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
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
	rp.db.AutoMigrate(&models2.Rate{})
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
	rp.db.Create(&models2.Rate{
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
func (rp *dbRepositoryTest) TestComplexQueryLatestRates() {
	rp.db.Delete(&models2.Rate{}, "id != 0")
	datasets := []*models2.Rate{
		{"1", "abc", "cde", "bank1", "enote", 0.02, 0.03, time.Now().Add(-time.Hour)},
		{"2", "abc", "cde", "bank1", "enote", 0.01, 0.02, time.Now().Add(time.Hour)},
		{"3", "abc", "usd", "bank1", "banknote", 0.03, 0.01, time.Now().Add(-time.Minute)},
	}
	rp.db.Create(datasets)
	dateLayout := "2006-01-02"
	date, _ := time.Parse(dateLayout, time.Now().Format(dateLayout))
	res, err := rp.repo.Latest(context.Background(), date)
	assert.Nil(rp.T(), err)
	assert.Len(rp.T(), res, 2)
	assert.Equal(rp.T(), datasets[1].ID, res[0].ID())
	cb := rp.repo.CriteriaBuilder().Where(rate.WhereSourceType("enote"))
	res, err = rp.repo.Apply(cb).Latest(context.Background(), date)
	assert.NotNil(rp.T(), res)
	assert.Nil(rp.T(), err)
	assert.Len(rp.T(), res, 1)
	assert.Equal(rp.T(), datasets[1].ID, res[0].ID())
}

func (rp *dbRepositoryTest) TestHistory01Empty() {
	rp.db.Delete(&models2.Rate{}, "id != 0")
	span := rate.FiveMinute
	dateFormat := "2006-01-02 03:04"
	now := time.Now()
	res, err := rp.repo.History(context.Background(), span, time.Now(), time.Now().Add(-time.Hour))
	assert.Nil(rp.T(), err)
	assert.Len(rp.T(), res, 0)
	for idx, r := range res {
		assert.Equal(rp.T(), now.Add(-(time.Minute * time.Duration(5 * idx))).Format(dateFormat), r.UpdatedAt().Format(dateFormat))
	}
}
func (rp *dbRepositoryTest) TestHistory00Simple() {
	rp.db.Delete(&models2.Rate{}, "id != 0")
	span := rate.FiveMinute
	dateFormat := "2006-01-02 03:04"
	totalDatasets := 60
	datasets := make([]*models2.Rate, totalDatasets)
	now := time.Now()
	for i := 0; i < totalDatasets; i++ {
		datasets[i] = &models2.Rate{
			ID:         strconv.Itoa(i + 1),
			Base:       "idr",
			Symbol:     "usd",
			Source:     "bank1",
			SourceType: "enote",
			Sell:       0.1 * float64(i + 1),
			Buy:        0.2 * float64(i + 1),
			UpdatedAt:  now.Add(-(time.Hour - (time.Minute * time.Duration(i)))),
		}
	}
	rp.db.Create(datasets)
	res, err := rp.repo.History(context.Background(), span, time.Now(), time.Now().Add(-time.Hour))
	assert.Nil(rp.T(), err)
	assert.Len(rp.T(), res, 12)
	for idx, r := range res {
		fmt.Println("result", now.Add(-(span.ToDuration() * time.Duration(idx))).Format(dateFormat), r.UpdatedAt().Format(dateFormat))
		assert.Equal(rp.T(), now.Add(-(span.ToDuration() * time.Duration(idx))).Format(dateFormat), r.UpdatedAt().Format(dateFormat))
	}
}

func (rp *dbRepositoryTest) TestHistory02Gap() {
	rp.db.Delete(&models2.Rate{}, "id != 0")

	span := rate.Minute
	totalDatasets := 60
	datasets := make([]*models2.Rate, totalDatasets)
	for i := 0; i < totalDatasets; i++ {
		datasets[i] = &models2.Rate{
			ID:         strconv.Itoa(i + 1),
			Base:       "idr",
			Symbol:     "usd",
			Source:     "bank1",
			SourceType: "enote",
			Sell:       0.1 * float64(i + 1),
			Buy:        0.2 * float64(i + 1),
			UpdatedAt:  time.Now().Add(-(time.Hour - (time.Minute * 5 * time.Duration(i)))),
		}
	}
	rp.db.Create(datasets)
	res, err := rp.repo.History(context.Background(), span, time.Now(), time.Now().Add(-time.Hour))
	assert.Nil(rp.T(), err)
	assert.Len(rp.T(), res, 60)
}

func TestQuery(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(dbRepositoryTest))
}
