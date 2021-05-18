package mysql

import (
	"context"
	"factory/exam/infra"
	"factory/exam/repo"
)

var _ repo.ProductRepoInterface = &ProductMySQLRepo{}

//ProductMySQLRepo ...
type ProductMySQLRepo struct {
	db *infra.ConnPool
}

//NewProductMySQLRepo ...
func NewProductMySQLRepo(
	db *infra.ConnPool,
) *ProductMySQLRepo {
	return &ProductMySQLRepo{
		db: db,
	}
}

//GetProduct ...
func (p *ProductMySQLRepo) GetProduct(ctx context.Context, limit int) (productDAO []*repo.ProductModel, err error) {
	if err = p.db.Conn.WithContext(ctx).Limit(limit).Find(&productDAO).Error; err != nil {
		return nil, err
	}

	return productDAO, nil
}

func (p *ProductMySQLRepo) Find(ctx context.Context, id int) (productDAO *repo.ProductModel, err error) {
	if err = p.db.Conn.WithContext(ctx).First(&productDAO, id).Error; err != nil {
		return nil, err
	}

	return productDAO, nil
}

//GetProductBy ...
func (p *ProductMySQLRepo) GetProductBy(ctx context.Context, query string) (productDAO *repo.ProductModel, err error) {
	if err = p.db.Conn.WithContext(ctx).Find(&productDAO).Limit(1).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

//Create ...
func (p *ProductMySQLRepo) Create(ctx context.Context, products []repo.ProductModel) (err error) {
	err = p.db.Conn.WithContext(ctx).Create(products).Error
	return err
}

//Update ...
func (p *ProductMySQLRepo) Update(ctx context.Context, product *repo.ProductModel) (err error) {
	err = p.db.Conn.WithContext(ctx).Model(product).Updates(product).Error
	return err
}
