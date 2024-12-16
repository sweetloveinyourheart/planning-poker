package pocker_userserver

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do"
	"github.com/spf13/cobra"

	"github.com/sweetloveinyourheart/planning-pocker/pkg/cmdutil"
	"github.com/sweetloveinyourheart/planning-pocker/pkg/config"
	"github.com/sweetloveinyourheart/planning-pocker/pkg/db"
	"github.com/sweetloveinyourheart/planning-pocker/pkg/grpc"
	"github.com/sweetloveinyourheart/planning-pocker/proto/code/userserver/go/grpcconnect"

	log "github.com/sweetloveinyourheart/planning-pocker/pkg/logger"
	"github.com/sweetloveinyourheart/planning-pocker/services/user"
	"github.com/sweetloveinyourheart/planning-pocker/services/user/actions"
)

const serviceType = "userserver"
const dbTablePrefix = "pocker_userserver"
const defDBName = "pocker_userserver"
const envPrefix = "USERSERVER"

func Command(rootCmd *cobra.Command) *cobra.Command {
	var userServerCommand = &cobra.Command{
		Use:   fmt.Sprintf("%s [flags]", serviceType),
		Short: fmt.Sprintf("Run as %s service", serviceType),
		Run: func(cmd *cobra.Command, args []string) {
			app, err := cmdutil.BoilerplateRun(serviceType)
			if err != nil {
				log.GlobalSugared().Fatal(err)
			}

			app.Migrations(user.FS, dbTablePrefix)

			if err := setupDependencies(); err != nil {
				log.GlobalSugared().Fatal(err)
			}

			user.InitializeRepos(app.Ctx())

			signingKey := config.Instance().GetString("userserver.secrets.token_signing_key")
			actions := actions.NewActions(app.Ctx(), signingKey)

			path, handler := grpcconnect.NewUserServiceHandler(
				actions,
			)
			go grpc.ServeBuf(app.Ctx(), path, handler, config.Instance().GetUint64("clientserver.grpc.port"), serviceType)

			app.Run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Instance().Set("service_prefix", serviceType)

			cmdutil.BoilerplateMetaConfig(serviceType)

			config.RegisterService(cmd, config.Service{
				Command: cmd,
			})
			config.AddDefaultServicePorts(cmd, rootCmd)
			config.AddDefaultDatabase(cmd, defDBName)
			return nil
		},
	}

	// config options
	cmdutil.BoilerplateFlagsCore(userServerCommand, serviceType, envPrefix)
	cmdutil.BoilerplateSecureFlags(userServerCommand, serviceType)
	cmdutil.BoilerplateFlagsDB(userServerCommand, serviceType, envPrefix)

	return userServerCommand
}

func setupDependencies() error {
	dbConn, err := db.NewDbWithWait(config.Instance().GetString("userserver.db.url"), db.DBOptions{
		TimeoutSec:      config.Instance().GetInt("userserver.db.postgres.timeout"),
		MaxOpenConns:    config.Instance().GetInt("userserver.db.postgres.max_open_connections"),
		MaxIdleConns:    config.Instance().GetInt("userserver.db.postgres.max_idle_connections"),
		ConnMaxLifetime: config.Instance().GetInt("userserver.db.postgres.max_lifetime"),
		ConnMaxIdleTime: config.Instance().GetInt("userserver.db.postgres.max_idletime"),
		EnableTracing:   config.Instance().GetBool("userserver.db.tracing"),
	})
	if err != nil {
		return err
	}

	do.Provide(nil, func(i *do.Injector) (*pgxpool.Pool, error) {
		return dbConn, nil
	})

	return nil
}
