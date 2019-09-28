//+build wireinject

package di

import (
	"tat_gogogo/data/crawler/repository"
	"tat_gogogo/interface/controller"
	"tat_gogogo/usecase"
	"tat_gogogo/usecase/service"

	"github.com/google/wire"
)

var provideCurriculumRepo = wire.NewSet(repository.NewCurriculumRepository)
var provideResultRepo = wire.NewSet(repository.NewResultRepository)
var provideInfoRepo = wire.NewSet(repository.NewInfoRepository)

var provideCurriculumService = wire.NewSet(service.NewCurriculumService)
var provideResultService = wire.NewSet(service.NewResultService)
var provideInfoService = wire.NewSet(service.NewInfoService)

var provideCurriculumController = wire.NewSet(
	provideCurriculumRepo,
	provideResultRepo,
	provideInfoRepo,
	provideResultService,
	usecase.NewResultUseCase,
	controller.NewCurriculumController,
)

var provideCoursesController = wire.NewSet(
	provideCurriculumRepo,
	provideResultRepo,
	provideInfoRepo,
	provideResultService,
	provideCurriculumService,
	usecase.NewResultUseCase,
	usecase.NewCurriculumUseCase,
	controller.NewCoursesController,
)
var provideLoginController = wire.NewSet(
	provideCurriculumRepo,
	provideResultRepo,
	provideInfoRepo,
	provideResultService,
	usecase.NewResultUseCase,
	provideCurriculumService,
	usecase.NewCurriculumUseCase,
	controller.NewLoginController,
)

/*
InjectCourseController is a function which will inject all dependencies
*/
func InjectCourseController() controller.CoursesController {
	wire.Build(provideCoursesController)
	return nil
}

/*
InjectCurriculumController is a function which will inject all dependencies
*/
func InjectCurriculumController() controller.CurriculumController {
	wire.Build(provideCurriculumController)
	return nil
}

/*
InjectLoginController is a function which will inject all dependencies
*/
func InjectLoginController() controller.LoginController {
	wire.Build(provideLoginController)
	return nil
}
