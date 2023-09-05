package serve_connect

import (
	"fmt"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/otelconnect"
	"github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1/mgmtv1alpha1connect"

	logger_interceptor "github.com/nucleuscloud/neosync/backend/internal/connect/interceptors/logger"
	neosync_k8sclient "github.com/nucleuscloud/neosync/backend/internal/k8s/client"
	neosynclogger "github.com/nucleuscloud/neosync/backend/internal/logger"
	v1alpha1_connectionservice "github.com/nucleuscloud/neosync/backend/services/mgmt/v1alpha1/connection-service"
	v1alpha1_jobservice "github.com/nucleuscloud/neosync/backend/services/mgmt/v1alpha1/job-service"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "connect",
		Short: "serves up connect",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			return serve()
		},
	}
}

func serve() error {
	port := viper.GetInt32("PORT")
	if port == 0 {
		port = 8080
	}
	host := viper.GetString("HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	environment := viper.GetString("NUCLEUS_ENV")

	logger := neosynclogger.New(neosynclogger.ShouldFormatAsJson(), nil).
		With("nucleusEnv", environment)
	loglogger := neosynclogger.NewLogLogger(neosynclogger.ShouldFormatAsJson(), nil)

	mux := http.NewServeMux()

	services := []string{
		mgmtv1alpha1connect.UserAccountServiceName,
		mgmtv1alpha1connect.AuthServiceName,
		mgmtv1alpha1connect.ConnectionServiceName,
		mgmtv1alpha1connect.JobServiceName,
	}

	checker := grpchealth.NewStaticChecker(services...)
	mux.Handle(grpchealth.NewHandler(checker))

	reflector := grpcreflect.NewStaticReflector(services...)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	neosyncK8sClient, err := neosync_k8sclient.New()
	if err != nil {
		return err
	}

	stdInterceptors := connect.WithInterceptors(
		otelconnect.NewInterceptor(),
		logger_interceptor.NewInterceptor(logger),
	)

	api := http.NewServeMux()

	jobConfigNamespace := getJobConfigNamespace()

	connectionService := v1alpha1_connectionservice.New(&v1alpha1_connectionservice.Config{
		JobConfigNamespace: jobConfigNamespace,
	}, neosyncK8sClient)
	api.Handle(
		mgmtv1alpha1connect.NewConnectionServiceHandler(
			connectionService,
			stdInterceptors,
		),
	)

	jobService := v1alpha1_jobservice.New(&v1alpha1_jobservice.Config{
		JobConfigNamespace: jobConfigNamespace,
	}, neosyncK8sClient, connectionService)
	api.Handle(
		mgmtv1alpha1connect.NewJobServiceHandler(
			jobService,
			stdInterceptors,
		),
	)

	mux.Handle("/", api)

	addr := fmt.Sprintf("%s:%d", host, port)

	logger.Info(fmt.Sprintf("listening on %s", addr))
	httpServer := http.Server{
		Addr:              addr,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ErrorLog:          loglogger,
		ReadHeaderTimeout: 10 * time.Second,
	}

	if err = httpServer.ListenAndServe(); err != nil {
		logger.Error(err.Error())
	}
	return nil
}

func getJobConfigNamespace() string {
	jobConfigNamespace := viper.GetString("JOB_CONFIG_NAMESPACE")
	if jobConfigNamespace == "" {
		jobConfigNamespace = "default"
	}
	return jobConfigNamespace
}
