package api

import (
	"context"
	"net/http"

	"github.com/simiraaaa/beego-catchup/src/handler"
	"github.com/simiraaaa/beego-catchup/src/lib/firebaseauth"
	"github.com/simiraaaa/beego-catchup/src/lib/httpheader"
	"github.com/simiraaaa/beego-catchup/src/lib/log"
	"github.com/simiraaaa/beego-catchup/src/model"
	"github.com/simiraaaa/beego-catchup/src/service"
)

// SampleHandler ... 記事のハンドラ
type SampleHandler struct {
	Svc service.Sample
}

// Sample ... サンプルハンドラ
func (h *SampleHandler) Sample(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// HTTPHeaderの値を取得
	headerParams := httpheader.GetParams(ctx)
	log.Debugf(ctx, "HeaderParams: %v", headerParams)

	// URLParamの値を取得
	urlParam := handler.GetURLParam(r, "sample")
	if urlParam == "" {
		h.handleError(ctx, w, http.StatusBadRequest, "invalid url param is empty")
		return
	}
	log.Debugf(ctx, "URLParam: %s", urlParam)

	// フォームの値を取得
	formParam := handler.GetFormValue(r, "sample")
	if formParam == "" {
		h.handleError(ctx, w, http.StatusBadRequest, "invalid form param is empty")
		return
	}
	log.Debugf(ctx, "FormParams: %s", formParam)

	// FirebaseAuthのユーザーIDを取得
	userID := firebaseauth.GetUserID(ctx)
	log.Debugf(ctx, "UserID: %s", userID)

	// FirebaseAuthのJWTClaimsの値を取得
	claims := firebaseauth.GetClaims(ctx)
	log.Debugf(ctx, "Claims: %v", claims)

	// Serviceを実行する
	sample, err := h.Svc.Sample(ctx)
	if err != nil {
		h.handleError(ctx, w, http.StatusInternalServerError, "h.Service.Sample: "+err.Error())
		return
	}

	handler.RenderJSON(w, http.StatusOK, struct {
		Sample model.Sample `json:"sample"`
		Hoge   string       `json:"hoge,omitempty"`
	}{
		Sample: sample,
		Hoge:   "",
	})
}

// TestDataStore ... DataStoreテスト
func (h *SampleHandler) TestDataStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.Svc.TestDataStore(ctx)
	if err != nil {
		h.handleError(ctx, w, http.StatusInternalServerError, "h.Svc.TestDataStore: "+err.Error())
		return
	}

	handler.RenderSuccess(w)
}

// TestCloudSQL ... CloudSQLテスト
func (h *SampleHandler) TestCloudSQL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.Svc.TestCloudSQL(ctx)
	if err != nil {
		h.handleError(ctx, w, http.StatusInternalServerError, "h.Svc.TestCloudSQL: "+err.Error())
		return
	}

	handler.RenderSuccess(w)
}

// TestHTTP ... HTTPテスト
func (h *SampleHandler) TestHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.Svc.TestHTTP(ctx)
	if err != nil {
		h.handleError(ctx, w, http.StatusInternalServerError, "h.Svc.TestHTTP: "+err.Error())
		return
	}

	handler.RenderSuccess(w)
}

func (h *SampleHandler) handleError(ctx context.Context, w http.ResponseWriter, status int, msg string) {
	log.Errorf(ctx, msg)
	handler.RenderError(w, status, msg)
}

// NewSampleHandler ... SampleHandlerを作成する
func NewSampleHandler(svc service.Sample) *SampleHandler {
	return &SampleHandler{
		Svc: svc,
	}
}
