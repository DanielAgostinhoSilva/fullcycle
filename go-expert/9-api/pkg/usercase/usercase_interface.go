package usercase

type UseCase interface {
	Execute(input interface{}) (*interface{}, error)
}

type NullaryUseCase interface {
	Execute() (interface{}, error)
}

type UnitUseCase interface {
	Execute(input interface{}) error
}
