package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	models2 "github.com/mgufrone/forex/delivery/grpc/rate/models"
	"github.com/mgufrone/forex/internal/domains/rate"
	"gorm.io/gorm"
)

type dbCommand struct {
	db *gorm.DB
}

func NewCommand(db *gorm.DB) rate.ICommand {
	return &dbCommand{db: db}
}

func (d *dbCommand) Create(ctx context.Context, in *rate.Rate) (err error) {
	id, _ := uuid.NewUUID()
	in.SetID(id.String())
	rt := &models2.Rate{}
	rt.FromDomain(in)
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).Create(rt).Error
	})
}

func (d *dbCommand) Update(ctx context.Context, in *rate.Rate) (err error) {
	rt := &models2.Rate{}
	rt.FromDomain(in)
	if in.ID() == "" {
		return errors.New("record not found")
	}
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", in.ID()).Updates(rt).Error
	})
}

func (d *dbCommand) Delete(ctx context.Context, in *rate.Rate) (err error) {
	rt := &models2.Rate{}
	rt.FromDomain(in)
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Delete(rt).Error
	})
}
