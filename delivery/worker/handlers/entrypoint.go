package handlers

import (
	"context"
	"fmt"
	"github.com/mgufrone/forex/internal/domains/rate"
	"log"
)

type entryPoint struct {
	worker  IWorker
	command rate.ICommand
	query   rate.IQuery
}

func NewEntryPoint(worker IWorker, command rate.ICommand, query rate.IQuery) IWorker {
	return &entryPoint{worker: worker, command: command, query: query}
}

func (e *entryPoint) Run(ctx context.Context) (rates []*rate.Rate, err error) {
	rates, err = e.worker.Run(ctx)
	fmt.Println("attempt to inserting rates", len(rates))
	if err == nil {
		for _, c := range rates {
			_ = e.insertOrFail(ctx, c)
		}
	}

	return
}

func (e *entryPoint) insertOrFail(ctx context.Context, in *rate.Rate) error {
	cb := e.query.CriteriaBuilder().
		Where(
			rate.SavedAt(in.UpdatedAt()),
			rate.WhereSource(in.Source()),
			rate.WhereSourceType(in.SourceType()),
			rate.WhereSymbol(in.SourceType()),
		)
	if total, err := e.query.Apply(cb).Count(ctx); err != nil {
		return err
	} else if total > 0 {
		log.Println("record exist. skipping")

		return nil
	}

	return e.command.Create(ctx, in)
}
