package dramahttp

import (
	"net"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	dramautils "github.com/zhs007/drama/utils"
	"go.uber.org/zap"
)

// APIHandle - handle
type APIHandle func(ctx *fasthttp.RequestCtx, serv *Serv)

// Serv -
type Serv struct {
	bindAddr    string
	mapAPI      map[string]APIHandle
	isDebugMode bool
	listener    net.Listener
}

// NewServ - new a serv
func NewServ(bindAddr string, isDebugMode bool) *Serv {
	s := &Serv{
		bindAddr:    bindAddr,
		mapAPI:      make(map[string]APIHandle),
		isDebugMode: isDebugMode,
	}

	return s
}

// RegHandle - register a handle
func (s *Serv) RegHandle(name string, handle APIHandle) {
	s.mapAPI[name] = handle
}

// HandleFastHTTP -
func (s *Serv) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	if s.isDebugMode {
		s.outputDebugInfo(ctx)
	}

	h, isok := s.mapAPI[string(ctx.Path())]
	if isok && h != nil {
		h(ctx, s)
	} else {
		s.SetHTTPStatus(ctx, fasthttp.StatusNotFound)
	}
}

// Stop - stop a server
func (s *Serv) Stop() error {
	if s.listener != nil {
		s.listener.Close()

		s.listener = nil
	}

	return nil
}

// Start - start a server
func (s *Serv) Start() error {
	if s.listener != nil {
		s.Stop()
	}

	ln, err := net.Listen("tcp4", s.bindAddr)
	if err != nil {
		dramautils.Error("block7http.Serv.Start:Listen",
			zap.Error(err))

		return err
	}

	s.listener = ln

	return fasthttp.Serve(ln, s.HandleFastHTTP)
}

// SetResponse - set a response
func (s *Serv) SetResponse(ctx *fasthttp.RequestCtx, jsonObj interface{}) {
	// ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	// ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")

	if jsonObj == nil {
		ctx.SetContentType("application/json;charset=UTF-8")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody([]byte(""))

		return
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	b, err := json.Marshal(jsonObj)
	if err != nil {
		dramautils.Warn("block7http.Serv.SetResponse",
			zap.Error(err))

		s.SetHTTPStatus(ctx, fasthttp.StatusInternalServerError)

		return
	}

	ctx.SetContentType("application/json;charset=UTF-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(b)

	dramautils.Debug("block7http.Serv.SetResponse",
		zap.String("RequestURI", string(ctx.RequestURI())),
		zap.String("body", string(b)))
}

// SetStringResponse - set a response with string
func (s *Serv) SetStringResponse(ctx *fasthttp.RequestCtx, str string) {
	// ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	// ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")

	ctx.SetContentType("application/json;charset=UTF-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte(str))

	dramautils.Debug("block7http.Serv.SetStringResponse",
		zap.String("RequestURI", string(ctx.RequestURI())),
		zap.String("body", str))
}

// SetHTTPStatus - set a response with status
func (s *Serv) SetHTTPStatus(ctx *fasthttp.RequestCtx, statusCode int) {
	// ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	// ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")

	ctx.SetStatusCode(statusCode)

	dramautils.Debug("block7http.Serv.SetHTTPStatus",
		zap.String("RequestURI", string(ctx.RequestURI())),
		zap.Int("statusCode", statusCode))
}

func (s *Serv) outputDebugInfo(ctx *fasthttp.RequestCtx) {
	dramautils.Debug("Request infomation",
		zap.String("Method", string(ctx.Method())),
		zap.String("RequestURI", string(ctx.RequestURI())),
		zap.String("Path", string(ctx.Path())),
		zap.String("Host", string(ctx.Host())),
		zap.String("UserAgent", string(ctx.UserAgent())),
		zap.String("RemoteIP", ctx.RemoteIP().String()),
		zap.Uint64("ConnRequestNum", ctx.ConnRequestNum()),
		zap.Time("ConnTime", ctx.ConnTime()),
		zap.Time("Time", ctx.Time()),
	)

	if ctx.QueryArgs() != nil {
		dramautils.Debug("Request infomation QueryArgs",
			zap.String("QueryArgs", ctx.QueryArgs().String()),
		)
	}

	if ctx.PostArgs() != nil {
		dramautils.Debug("Request infomation PostArgs",
			zap.String("PostArgs", ctx.PostArgs().String()),
		)
	}

	if ctx.PostBody() != nil {
		dramautils.Debug("Request infomation PostBody",
			zap.String("PostBody", string(ctx.PostBody())),
		)
	}
}

// ParseBody - parse body
func (s *Serv) ParseBody(ctx *fasthttp.RequestCtx, params interface{}) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	err := json.Unmarshal(ctx.PostBody(), params)
	if err != nil {
		return err
	}

	return nil
}
