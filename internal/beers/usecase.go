package beers

// BeerUseCase es una interfaz que tienen las firmas de los métodos para esta workshop
type BeerUseCase interface {
	SearchBeers() (*[]Beer, error)
	// AddBeers(Beer) error
}
