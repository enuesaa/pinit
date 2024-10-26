package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/guregu/dynamo/v2"
)

type DbRepositoryInterface interface {
	Put(data interface{}) error
	Get(name string, value interface{}, result interface{}) error
	List(name string, value interface{}, result interface{}) error
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

func (repo *DbRepository) Get(name string, value interface{}, result interface{}) error {
	ctx := context.Background()
	table, err := repo.getPinitTable()
	if err != nil {
		return nil
	}

	return table.Get(name, value).One(ctx, &result)
}

func (repo *DbRepository) List(name string, value interface{}, result interface{}) error {
	ctx := context.Background()
	table, err := repo.getPinitTable()
	if err != nil {
		return nil
	}

	return table.Get(name, value).Limit(10).All(ctx, &result)
}
