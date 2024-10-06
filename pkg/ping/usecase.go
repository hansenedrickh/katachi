package ping

import "context"

type Usecase struct{}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (u *Usecase) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}
