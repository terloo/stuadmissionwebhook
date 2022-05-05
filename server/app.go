package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var RootCmd = &cobra.Command{
	Use:  "myadmission",
	Long: "启动Application，提供AdmissionControllerWebhook服务",
	Run: func(cmd *cobra.Command, args []string) {
		crtPath := cmd.Flag("tls-crt-path").Value.String()
		keyPath := cmd.Flag("tls-key-path").Value.String()
		if crtPath == "" || keyPath == "" {
			klog.Fatalln("can't load tls config")
		}
		pair, err := tls.LoadX509KeyPair(crtPath, keyPath)
		if err != nil {
			klog.Fatalln("can't load tls config")
		}
		admissionServer := &AdmissionServer{
			Server: &http.Server{
				Addr:      fmt.Sprintf(":%v", cmd.Flag("port").Value.String()),
				TLSConfig: &tls.Config{Certificates: []tls.Certificate{pair}},
			},
		}
		http.DefaultServeMux.Handle("/mutate", admissionServer)
		http.DefaultServeMux.Handle("/validate", admissionServer)
		go func() {
			if err := admissionServer.Server.ListenAndServeTLS("", ""); err != nil {
				klog.Fatalf("Failed to listen and serve webhook server: %v\n", err)
			}
		}()

		klog.Info("Server started")

		// listening OS shutdown singal
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan

		klog.Infof("Got OS shutdown signal, shutting down webhook server gracefully...")
		admissionServer.Server.Shutdown(context.Background())
	},
}

func init() {
	RootCmd.Flags().String("tls-crt-path", "", "TLS证书路径")
	RootCmd.Flags().String("tls-key-path", "", "TLS私钥路径")
	RootCmd.Flags().Int("port", 8443, "端口")
}
