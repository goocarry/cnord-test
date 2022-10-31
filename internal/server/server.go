package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/goocarry/cnord-test/internal/config"
	"github.com/goocarry/cnord-test/internal/store"
	"github.com/goocarry/cnord-test/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run ...
func Run(cfg *config.Config,log *log.Logger, store *store.Store) {

	gs := grpc.NewServer()
	us := NewUserServer(log, store)

	user.RegisterUserServiceServer(gs, us)

	reflection.Register(gs)

	l, err := net.Listen("tcp", cfg.Listen.Port)
	if err != nil {
		log.Fatalf("error-4a09c85b: unable to listen: %v", err)
		os.Exit(1)
	}
	
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-sigCh
		log.Printf("info-81507edf: got signal %v, attempting graceful shutdown", s)
		us.store.Close()
		gs.GracefulStop()
		wg.Done()
	}()

	err  = us.store.Open()
	if err != nil {
		log.Fatalf("error-45ff4095: cannot open database connection, error: %v", err)
	}

	err = gs.Serve(l)
	if err != nil {
		log.Fatalf("error-405895ef: could not serve: %v", err)
	}
	wg.Wait()
	log.Println("info-a7d54316: clean shutdown")
}