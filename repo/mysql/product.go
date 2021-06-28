package mysql

import (
	"context"
	"encoding/json"
	"factory/exam/infra"
	"factory/exam/repo"
	"fmt"
	"os"

	kafkaLib "github.com/lk153/go-lib/kafka"
	"github.com/lk153/go-lib/kafka/ccloud"
	entities_pb "github.com/lk153/proto-tracking-gen/go/entities"
)

var _ repo.ProductRepoInterface = &ProductMySQLRepo{}

//ProductMySQLRepo ...
type ProductMySQLRepo struct {
	db       *infra.ConnPool
	producer *kafkaLib.KafkaProducer
	topic    *string
}

//NewProductMySQLRepo ...
func NewProductMySQLRepo(
	db *infra.ConnPool,
) *ProductMySQLRepo {
	configPath = ccloud.ParseArgs()
	producerLib := &kafkaLib.KafkaProducer{
		ConfigFile: configPath,
	}
	producerLib.InitConfig()
	err := producerLib.CreateProducerInstance()
	if err != nil {
		fmt.Println("create producer has error")
		os.Exit(1)
	}
	producerLib.CreateTopic(PRODUCT_KAFKA_TOPIC)
	topic := PRODUCT_KAFKA_TOPIC
	return &ProductMySQLRepo{
		db:       db,
		producer: producerLib,
		topic:    &topic,
	}
}

//GetProduct ...
func (p *ProductMySQLRepo) Get(ctx context.Context, limit int, page int, ids []uint64) (productDAO []*repo.ProductModel, err error) {
	tx := p.db.Conn.WithContext(ctx)
	if limit != 0 {
		tx = tx.Limit(limit)
	}

	if page != 0 {
		tx = tx.Offset(page * limit)
	}

	if ids != nil {
		tx = tx.Find(&productDAO, ids)
	} else {
		tx = tx.Find(&productDAO)
	}

	if err = tx.Error; err != nil {
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
	p.producer.ProduceMessage(p.topic, string(raw))

	return productDAO, nil
}

//Update ...
func (p *ProductMySQLRepo) Update(ctx context.Context, product *repo.ProductModel) (err error) {
	err = p.db.Conn.WithContext(ctx).Model(product).Updates(product).Error
	return err
}
