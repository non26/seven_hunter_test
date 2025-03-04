package usecases

import (
	"backend/third_find_all_meat/internal/core/domain"
	"backend/third_find_all_meat/internal/core/port"
	"context"
	"strconv"
)

type meatUsecase struct {
	meatRepository port.IMeatRepository
}

func NewMeatUsecase(meatRepository port.IMeatRepository) port.IMeatUsecase {
	return &meatUsecase{meatRepository: meatRepository}
}

func (u *meatUsecase) GetAll(ctx context.Context) (map[string]int, error) {
	meat := domain.NewMeat()
	for i := 0; i < 100; i++ {
		println(i)
		meatContent, err := u.meatRepository.Get(ctx, strconv.Itoa(i))
		if err != nil {
			return nil, err
		}
		meat.SetContent(meatContent)
		meat.CountingMeat()
	}
	return meat.GetMeats(), nil
}
