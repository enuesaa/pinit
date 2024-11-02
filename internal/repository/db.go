package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/guregu/dynamo/v2"
)

type DbRepositoryInterface interface {
	Put(data interface{}) error
	List(partition string, result interface{}) error
	Get(partition string, rangev string, result interface{}) error
	Delete(partition string, rangev string) error
}

type DbRepository struct {}

func (repo *DbRepository) getAwsConfig() (aws.Config, error) {
	ctx := context.Background()

	return config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"))
}

func (repo *DbRepository) getPinitTable() (*dynamo.Table, error) {
	awsconf, err := repo.getAwsConfig()
	if err != nil {
		return nil, err
	}
	db := dynamo.New(awsconf)
	table := db.Table("pinit")

	return &table, nil
}

func (repo *DbRepository) Put(data interface{}) error {
	ctx := context.Background()
	table, err := repo.getPinitTable()
	if err != nil {
		return nil
	}

	return table.Put(data).Run(ctx)
}

func (repo *DbRepository) List(partition string, result interface{}) error {
	ctx := context.Background()
	table, err := repo.getPinitTable()
	if err != nil {
		return nil
	}

	return table.Get("BinderName", partition).Order(dynamo.Descending).Limit(10).All(ctx, result)
}

func (repo *DbRepository) Get(partition string, rangev string, result interface{}) error {
	ctx := context.Background()
	table, err := repo.getPinitTable()
	if err != nil {
		return nil
	}
	query := table.Get("BinderName", partition)
	query.Range("NoteName", dynamo.Equal, rangev)

	return query.One(ctx, result)
}

func (repo *DbRepository) Delete(partition string, rangev string) error {
	ctx := context.Background()
	table, err := repo.getPinitTable()
	if err != nil {
		return nil
	}
	query := table.Delete("BinderName", partition)
	query.Range("NoteName", rangev)

	return query.Run(ctx)
}
