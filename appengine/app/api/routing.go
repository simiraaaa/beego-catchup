package main

import (
	"net/http"

	"github.com/simiraaaa/beego-catchup/src/config"
	"github.com/simiraaaa/beego-catchup/src/handler"
	"github.com/simiraaaa/beego-catchup/src/lib/log"
	"github.com/simiraaaa/beego-catchup/src/middleware"
	"github.com/go-chi/chi"
)

// Routing ... ルーティング設定
func Routing(r *chi.Mux, d *Dependency) {
	// アクセスコントロール
	r.Use(middleware.AccessControl)

	// ログ
	r.Use(log.Handle)

	// 認証なし(Stagingのみ)
	if config.IsEnvStaging() {
		r.Route("/noauth/v1", func(r chi.Router) {
			r.Use(d.DummyFirebaseAuth.Handle)
			r.Use(d.DummyHTTPHeader.Handle)
			subRouting(r, d)
		})
	}

	// 認証あり
	r.Route("/v1", func(r chi.Router) {
		r.Use(d.FirebaseAuth.Handle)
		r.Use(d.HTTPHeader.Handle)
		subRouting(r, d)
	})

	// Ping
	r.Get("/ping", handler.Ping)

	http.Handle("/", r)
}

func subRouting(r chi.Router, d *Dependency) {
	// API
	r.Get("/sample", d.SampleHandler.Sample)
	r.Get("/testdatastore", d.SampleHandler.TestDataStore)
	r.Get("/testcloudsql", d.SampleHandler.TestCloudSQL)
	r.Get("/testhttp", d.SampleHandler.TestHTTP)

	// API(JSONRPC2)
	r.Route("/rpc", func(r chi.Router) {
		r.Use(d.JSONRPC2.Handle)
		r.Post("/", handler.Empty)
	})
}
