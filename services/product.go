package services

import (
	"context"
	"factory/exam/repo"

	"golang.org/x/sync/errgroup"
)

var _ ProductServiceInterface = &ProductService{}

//ProductProvider ...
func ProductProvider(
	productRepo repo.ProductRepoInterface,
) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

//ProductService ...
type ProductService struct {
	productRepo repo.ProductRepoInterface
}

//GetProducts ...
func (ps *ProductService) GetProducts(ctx context.Context, limit uint32) []*repo.ProductModel {
	var products []*repo.ProductModel
	var prodChan = make(chan *repo.ProductModel)
	g, ctx := errgroup.WithContext(ctx)
	for i := uint32(0); i < limit; i++ {
		g.Go(func() error {
			prod, err := ps.productRepo.GetProduct(ctx)
			if err != nil {
				return err
			}

			select {
			case prodChan <- prod:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	}

	go func() {
		g.Wait()
		close(prodChan)
	}()

	for prod := range prodChan {
		products = append(products, prod)
	}

	if err := g.Wait(); err != nil {
		return nil
	}

	return products
}
