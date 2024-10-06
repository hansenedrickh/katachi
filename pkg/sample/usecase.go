package sample

import "context"

type repository interface {
	Insert(ctx context.Context, sample Sample) error
	Fetch(ctx context.Context, id int) (Sample, error)
	Get(ctx context.Context, page, limit int) ([]Sample, error)
	Update(ctx context.Context, sample Sample) error
	Delete(ctx context.Context, id int) error
}

type Usecase struct {
	repository repository
}

func NewUsecase(repo repository) *Usecase {
	return &Usecase{
		repository: repo,
	}
}

func (u *Usecase) Insert(ctx context.Context, name string) error {
	sample := Sample{
		Name: name,
	}

	return u.repository.Insert(ctx, sample)
}

func (u *Usecase) Fetch(ctx context.Context, id int) (Sample, error) {
	return u.repository.Fetch(ctx, id)
}

func (u *Usecase) Get(ctx context.Context, page, limit int) ([]Sample, error) {
	return u.repository.Get(ctx, page, limit)
}

func (u *Usecase) Update(ctx context.Context, id int, name string) error {
	sample, err := u.repository.Fetch(ctx, id)
	if err != nil {
		return err
	}

	sample.Name = name

	return u.repository.Update(ctx, sample)
}

func (u *Usecase) Delete(ctx context.Context, id int) error {
	return u.repository.Delete(ctx, id)
}
