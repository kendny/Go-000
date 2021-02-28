package internal

import (
	"github.com/google/wire"
	"github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/internal/repo"
	"github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/internal/service"
	"github.com/mohuishou/go-training/Week04/blog/08_unit_test/01_example/internal/usecase"
)

// Set Set
var Set = wire.NewSet(
	wire.Struct(new(service.PostService), "*"),
	wire.Struct(new(usecase.PostUsecaseOption), "*"),
	usecase.NewPostUsecase,
	repo.NewPostRepo,
)
