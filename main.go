package main

import (
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strings"

	"github.com/makehlv/ept/clients"
	"github.com/makehlv/ept/config"
	"github.com/makehlv/ept/repositories"
	"github.com/makehlv/ept/services"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ept <command> [flags]")
		os.Exit(1)
	}

	clients := clients.NewClients()
	logger := slog.New(NewColorHandler(os.Stderr, slog.LevelInfo))
	config := config.NewConfig()
	repos := repositories.NewRepositories(logger, config)
	svc := services.NewServices(clients, logger, config, repos)

	command := os.Args[1]
	switch command {
	case "swg":
		if len(os.Args) < 3 {
			names, err := svc.Swagger.ListServers()
			if err != nil {
				logger.Error("swg failed", "error", err)
				os.Exit(1)
			}
			for _, name := range names {
				fmt.Println(name)
			}
			break
		}
		serverName := os.Args[2]
		genOp := parseFlag(os.Args[3:], "--gen")
		specPath := parseFlag(os.Args[3:], "--spec")
		if genOp != "" && specPath != "" {
			fmt.Println("use either --gen or --spec, not both")
			os.Exit(1)
		}
		if genOp != "" {
			curlCmd, err := svc.Swagger.BuildCurl(serverName, genOp)
			if err != nil {
				logger.Error("swg failed", "error", err)
				os.Exit(1)
			}
			fmt.Println(curlCmd)
		} else if specPath != "" {
			if err := svc.Swagger.SaveServerSpec(serverName, specPath); err != nil {
				logger.Error("swg failed", "error", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("usage: kk swg [<server_name> --gen <operationId> | --spec <absolute_path_to_swagger>]")
			os.Exit(1)
		}
	case "var":
		if len(os.Args) < 3 {
			vars, err := svc.Variable.ListVars()
			if err != nil {
				logger.Error("var failed", "error", err)
				os.Exit(1)
			}
			keys := make([]string, 0, len(vars))
			for k := range vars {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				fmt.Printf("%s=%s\n", k, vars[k])
			}
			break
		}
		if len(os.Args) < 4 {
			fmt.Println("usage: kk var [<key> <value>]")
			os.Exit(1)
		}
		key := os.Args[2]
		value := strings.Join(os.Args[3:], " ")
		if err := svc.Variable.Add(key, value); err != nil {
			logger.Error("var failed", "error", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("unknown command: %s\n", command)
		os.Exit(1)
	}
}

func parseFlag(args []string, flag string) string {
	for i, arg := range args {
		if arg == flag && i+1 < len(args) {
			return args[i+1]
		}
	}
	return ""
}
