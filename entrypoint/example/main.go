package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/RandalTeng/go-oauth2-server/manager"
	"github.com/RandalTeng/go-oauth2-server/models"
	"github.com/RandalTeng/go-oauth2-server/server"
	"github.com/RandalTeng/go-oauth2-server/store"
)

func main() {
	clientStorage := store.NewClientStore()
	_ = clientStorage.Set("abc", &models.Client{
		ID:     "abc",
		Secret: "123",
		Public: false,
	})
	tokenStorage := store.NewMemoryTokenStore(3)
	mngr := manager.NewDefaultManager()
	mngr.MapClientStorage(clientStorage)
	mngr.MapTokenStorage(tokenStorage)
	mngr.SetTokenAdapter(models.NewToken())
	mngr.SetCodeAdapter(models.NewCode())
	// set jwt access token generator
	//key := bytes.NewBufferString("1234567890abcdef").Bytes()
	//mngr.MapAccessGenerate(generator.NewJWTAccessGenerate("oauth2", key, jwt.SigningMethodHS256))
	// set generate refresh token, expires at 7 days after.
	//manager.DefaultClientTokenCfg.IsGenerateRefresh = true
	//manager.DefaultClientTokenCfg.RefreshTokenExp = 7 * 24 * time.Hour

	osrv := server.NewDefaultServer(mngr)
	osrv.ClientInfoHandler = server.ClientBodyHandler

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "hello, go oauth server.")
	})
	http.HandleFunc("/oauth2/code", func(w http.ResponseWriter, r *http.Request) {
		err := osrv.HandleAuthorizeRequest(w, r)
		if err != nil {
			log.Printf("handle authorize request error: %s\n", err)
		} else {
			log.Println("handle authorize request completed.")
		}
	})
	http.HandleFunc("/oauth2/client-credentials", func(w http.ResponseWriter, r *http.Request) {
		err := osrv.HandleTokenRequest(w, r)
		if err != nil {
			log.Printf("handle client credentials request error: %s\n", err)
		} else {
			log.Println("handle client credentials request completed.")
		}
	})
	hsrv := &http.Server{Addr: ":8080"}
	go func(srv *http.Server) {
		log.Println("begin listen and handler connections.")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("run server with error: %s\n", err)
		}
	}(hsrv)

	cs := make(chan os.Signal)
	signal.Notify(cs, syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGSTOP, syscall.SIGKILL,
		syscall.SIGTERM, syscall.SIGINT)
	sig := <-cs
	log.Printf("receive signal: %s, try to stop http server\n", sig)
	_ = hsrv.Close()
	log.Println("server closed.")
}
