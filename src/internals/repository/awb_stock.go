package repository

type (
	AWBStockRepository interface {
		// TODO: create the functions needed to implement here
	}

	AWBStockRepositoryImpl struct {
	}
)

func NewAWBStockRepository() AWBStockRepository {
	return &AWBStockRepositoryImpl{}
}
