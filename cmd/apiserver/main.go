package main

import (
	"fmt"
	"github.com/hobbyfarm/gargantua/v3/pkg/server"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfg = server.Config{
		HTTPPort:       8080,
		HTTPSPort:      8443,
		DSN:            "",
		KubeconfigPath: "",
		KubeContext:    "",
	}
)

func init() {
	rootCmd.Flags().StringVar(&cfg.DSN, "dsn", "", "mysql dsn")
	rootCmd.Flags().IntVar(&cfg.HTTPPort, "http-port", 8080, "http port")
	rootCmd.Flags().IntVar(&cfg.HTTPSPort, "https-port", 8443, "https port")
	rootCmd.Flags().StringVar(&cfg.KubeconfigPath, "kubeconfig", "",
		"path to kubeconfig (leave empty for in-cluster)")
	rootCmd.Flags().StringVar(&cfg.KubeContext, "context", "default", "context in kubeconfig file")
}

var rootCmd = &cobra.Command{
	Use:   "hobbyfarm-apiserver",
	Short: "run apiserver for hobbyfarm",
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	svr, err := server.New(cfg)
	if err != nil {
		return err
	}

	if err := svr.Run(cmd.Context()); err != nil {
		return err
	}

	<-cmd.Context().Done()

	return cmd.Context().Err()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
