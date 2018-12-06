package main

import (
	"github.com/simiraaaa/beego-catchup/src/handler/api"
	"github.com/simiraaaa/beego-catchup/src/lib/firebaseauth"
	"github.com/simiraaaa/beego-catchup/src/lib/httpheader"
	"github.com/simiraaaa/beego-catchup/src/lib/jsonrpc2"
	"github.com/simiraaaa/beego-catchup/src/repository"
	"github.com/simiraaaa/beego-catchup/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	DummyFirebaseAuth *firebaseauth.Middleware
	FirebaseAuth      *firebaseauth.Middleware
	DummyHTTPHeader   *httpheader.Middleware
	HTTPHeader        *httpheader.Middleware
	JSONRPC2          *jsonrpc2.Middleware
	SampleHandler     *api.SampleHandler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject() {
	// Config
	// dbCfg := config.GetCSQLConfig("sample")

	// Lib
	// dbConn := cloudsql.NewCSQLClient(dbCfg)

	// Repository
	repo := repository.NewSample(nil)

	// Service
	dfaSvc := firebaseauth.NewDummyService()
	faSvc := firebaseauth.NewService()
	dhhSvc := httpheader.NewDummyService()
	hhSvc := httpheader.NewService()
	svc := service.NewSample(repo)

	// Middleware
	d.DummyFirebaseAuth = firebaseauth.NewMiddleware(dfaSvc)
	d.FirebaseAuth = firebaseauth.NewMiddleware(faSvc)
	d.DummyHTTPHeader = httpheader.NewMiddleware(dhhSvc)
	d.HTTPHeader = httpheader.NewMiddleware(hhSvc)

	// JSONRPC2
	d.JSONRPC2 = jsonrpc2.NewMiddleware()
	d.JSONRPC2.Register("sample", api.NewSampleJSONRPC2Handler(svc))

	// Handler
	d.SampleHandler = api.NewSampleHandler(svc)
}
