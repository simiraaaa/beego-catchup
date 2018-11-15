package service

import (
	"context"

	"github.com/aikizoku/beego/src/model"
)

// Sample ... サービスのインターフェース
type Sample interface {
	Sample(ctx context.Context) (model.Sample, error)
	TestDataStore(ctx context.Context) error
	TestCloudSQL(ctx context.Context) error
	TestHTTP(ctx context.Context) error
}
