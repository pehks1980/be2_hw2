package main

import (
	"context"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"k8s-go-app/server"
	"k8s-go-app/version"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Config задает параметры конфигурации приложения
type Config struct {
	Port 		string `envconfig:"PORT" default:"8080"`
	StaticsPath string `envconfig:"STATICS_PATH" default:"./static"`
}

func main() {

	config := new(Config)
	err := envconfig.Process("", config)
	if err != nil {
		log.Fatalf("Can't process config: %v", err)
	}
	fs := http.FileServer(http.Dir(config.StaticsPath))

	//http.Handle("/", fs)
	log.Printf("statics server started. html path: %s port: %s\n", config.StaticsPath, config.Port)

	/*
	launchMode := config.LaunchMode(os.Getenv("LAUNCH_MODE"))
	if len(launchMode) == 0 {
		launchMode = config.LocalEnv
	}
	log.Printf("LAUNCH MODE: %v", launchMode)

	cfg, err := config.Load(launchMode, "./config")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CONFIG: %+v", cfg)
*/
	info := server.VersionInfo{
		Version: version.Version,
		Commit:  version.Commit,
		Build:   version.Build,
	}
	log.Printf("info %v", info)
	srv := server.New(info, config.Port, &fs)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := srv.Serve(ctx)
		if err != nil {
			log.Println(fmt.Errorf("serve: %w", err))
			return
		}
	}()

	osSigChan := make(chan os.Signal, 1)
	signal.Notify(osSigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-osSigChan
	log.Println("OS interrupting signal has received")


}

