package mysql

import (
	"context"
	"encoding/json"
	"factory/exam/infra"
	"factory/exam/repo"
	"fmt"

	kafkaLib "github.com/lk153/go-lib/kafka"
	entities_pb "github.com/lk153/proto-tracking-gen/go/entities"
)

var _ repo.ProductRepoInterface = &ProductMySQLRepo{}
var configPath = "librdkafka.config"
var topic = "testing"

//ProductMySQLRepo ...
type ProductMySQLRepo struct {
	db       *infra.ConnPool
	producer *kafkaLib.KafkaProducer
}

//NewProductMySQLRepo ...
func NewProductMySQLRepo(
	db *infra.ConnPool,
) *ProductMySQLRepo {
	producerLib := &kafkaLib.KafkaProducer{
		ConfigFile: &configPath,
	}
	producerLib.InitConfig()
	err := producerLib.CreateProducerInstance()
	if err != nil {
		fmt.Println("create producer has error")
	}
	return &ProductMySQLRepo{
		db:       db,
		producer: producerLib,
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
func (p *ProductMySQLRepo) Create(ctx context.Context, data *entities_pb.ProductInfo) (productDAO *repo.ProductModel, err error) {
	productDAO = &repo.ProductModel{}
	productDAO.ID = uint64(data.Id)
	productDAO.Name = data.Name
	productDAO.Price = data.Price
	productDAO.Status = uint8(data.Status)
	productDAO.Type = data.Type

	result := p.db.Conn.WithContext(ctx).Create(&productDAO)
	if result.Error != nil {
		return nil, result.Error
	}

	raw, err := json.Marshal(productDAO)
	if err != nil {
		fmt.Println("parse data has error")
	}
	p.producer.ProduceMessage(&topic, string(raw))

	return productDAO, nil
}

//Update ...
func (p *ProductMySQLRepo) Update(ctx context.Context, product *repo.ProductModel) (err error) {
	err = p.db.Conn.WithContext(ctx).Model(product).Updates(product).Error
	return err
}
