package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	confg "xm/internal/config"

	//middleware "xm/internal/http/middleware"

	mw "xm/internal/http/middlewares"
	lgg "xm/internal/logger"

	"github.com/gorilla/mux"
)

type muxRouter struct {
	lg              lgg.Logger
	signalChan      chan os.Signal
	router          *mux.Router
	httpServer      *http.Server
	ServerIp        string
	ShutDownTimeout time.Duration
	companyH        CompanyHandler
	userH           UserHandler
}

//NewMuxRouter creates a new Mux router for serving http requests
func NewMuxRouter(compH CompanyHandler, usrH UserHandler, conf confg.Configuration, sigs chan os.Signal, l lgg.Logger) {
	r := mux.NewRouter().StrictSlash(true)
	m := &muxRouter{
		lg:              l,
		signalChan:      sigs,
		router:          r,
		httpServer:      &http.Server{},
		ServerIp:        conf.HttpServerConfig.ServerIp,
		ShutDownTimeout: time.Duration(conf.HttpServerConfig.ShutDownTimeout),
		companyH:        compH,
		userH:           usrH,
	}
	// m.router.Use(middleware.JSONMiddleware)
	// m.router.Use(middleware.LoggingMiddleware)
	// m.router.Use(middleware.TimeoutMiddleware)
	//m.router.Use(middlewares.RequestBodyLimiter)

	//m.confi
	go m.Stop()
	m.Start()
}

// Stop() shuts down the server when signal channel receives a signal
func (m *muxRouter) Stop() {
	<-m.signalChan
	//m.lg.Info("stopping http server")
	context, cancel := context.WithTimeout(context.Background(), time.Duration(m.ShutDownTimeout))
	defer cancel()
	err := m.httpServer.Shutdown(context)
	if err != nil {
		m.lg.Fatal("http server shutdown", err)
	}
}

// Start() starts the http server to handle requests
func (m *muxRouter) Start() {
	//m.lg.Info("server started")
	m.httpServer.Addr = m.ServerIp
	//m.httpServer.Addr = ":8081"
	m.httpServer.Handler = m.router
	if err := m.httpServer.ListenAndServe(); err != nil {
		fmt.Println("server started")
		log.Panic(err)
	}

}

func (m *muxRouter) ConfigureHandler() {

	m.router.Methods("GET").Path("/api/v1/" + "companies/{deploymentId}").
		Handler(http.HandlerFunc(m.companyH.CheckHealth))

	m.router.Methods("GET").Path("/api/v1/" + "companies/{deploymentId}").
		Handler(http.HandlerFunc(m.companyH.GetCompany))

	m.router.Methods("POST").Path("/api/v1/" + "companies").
		Handler(http.HandlerFunc(mw.JWTAuth(m.companyH.CreateCompany)))

	m.router.Methods("PATCH").Path("/api/v1/" + "companies/{deploymentId}").
		Handler(http.HandlerFunc(mw.JWTAuth(m.companyH.UpdateCompany)))

	m.router.Methods("DELETE").Path("/api/v1/" + "companies/{deploymentId}").
		Handler(http.HandlerFunc(mw.JWTAuth(m.companyH.DeleteCompany)))

}
