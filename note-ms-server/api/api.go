package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"e.coding.net/logonod/note-ms-server/app"
	"e.coding.net/logonod/note-ms-server/model"
)

type statusCodeRecorder struct {
	http.ResponseWriter
	http.Hijacker
	StatusCode int
}

func (r *statusCodeRecorder) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func writeError(w http.ResponseWriter, a *app.AppError) error {
	data, err := json.Marshal(a.Status())
	if err != nil {
		return err
	}

	if _, err = w.Write(data); err != nil {
		return err
	}

	return nil
}

type API struct {
	App    *app.App
	Config *Config
}

func New(a *app.App) (api *API, err error) {
	api = &API{App: a}

	api.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	return api, nil
}

func (a *API) Init(r *mux.Router) {
	// 爬虫相关接口
	r.Handle("/spider/collect/get", a.handler(a.SpiderCollectGet)).Methods("GET")
	r.Handle("/spider/collect/create", a.handler(a.SpiderCollectCreate)).Methods("POST")
	r.Handle("/spider/collect/udpate_status", a.handler(a.SpiderCollectUpdateStatus)).Methods("POST")
	r.Handle("/spider/webpage/get", a.handler(a.SpiderWebpageGet)).Methods("GET")
}

func (a *API) handler(f func(*app.Context, http.ResponseWriter, *http.Request) *app.AppError) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 100*1024*1024)

		beginTime := time.Now()

		hijacker, _ := w.(http.Hijacker)
		w = &statusCodeRecorder{
			ResponseWriter: w,
			Hijacker:       hijacker,
		}

		ctx := a.App.NewContext().WithRemoteAddress(a.IPAddressForRequest(r))
		ctx = ctx.WithLogger(ctx.Logger.WithField("request_id", base64.RawURLEncoding.EncodeToString(model.NewId())))

		defer func() {
			if r := recover(); r != nil {
				ctx.Logger.Error(fmt.Errorf("%v: %s", r, debug.Stack()))
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()

		w.Header().Set("Content-Type", "application/json")

		secretKey := r.Header.Get("Secret-Key")
		if secretKey != string(a.App.Config.SecretKey) {
			aErr := app.AppError{nil, "未授权的访问", 40003}
			defer ctx.Logger.Error(fmt.Errorf("errcode: %d, errmsg: %s, err: %v", aErr.Status().Code, aErr.Status().Message, aErr.Error))
			if err := writeError(w, &aErr); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			return
		}

		if aErr := f(ctx, w, r); aErr != nil {
			if aErr.Code == 500 {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			} else {
				if err := writeError(w, aErr); err != nil {
					http.Error(w, "internal server error", http.StatusInternalServerError)
				}
			}

			defer ctx.Logger.Error(fmt.Errorf("errcode: %d, errmsg: %s, err: %v", aErr.Status().Code, aErr.Status().Message, aErr.Error))
		}

		defer func() {
			statusCode := w.(*statusCodeRecorder).StatusCode
			if statusCode == 0 {
				statusCode = 200
			}
			duration := time.Since(beginTime)

			logger := ctx.Logger.WithFields(logrus.Fields{
				"duration":    duration,
				"status_code": statusCode,
				"remote":      ctx.RemoteAddress,
			})
			logger.Info(r.Method + " " + r.URL.RequestURI())
		}()
	})
}

func (a *API) IPAddressForRequest(r *http.Request) string {
	addr := r.RemoteAddr
	if a.Config.ProxyCount > 0 {
		h := r.Header.Get("X-Forwarded-For")
		if h != "" {
			clients := strings.Split(h, ",")
			if a.Config.ProxyCount > len(clients) {
				addr = clients[0]
			} else {
				addr = clients[len(clients)-a.Config.ProxyCount]
			}
		}
	}
	return strings.Split(strings.TrimSpace(addr), ":")[0]
}
