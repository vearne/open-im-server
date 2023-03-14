package main

import (
	"OpenIM/internal/api"
	"OpenIM/pkg/common/cmd"
	"OpenIM/pkg/common/config"
	"OpenIM/pkg/common/log"
	"OpenIM/pkg/common/mw"
	"context"
	"errors"
	"fmt"
	"github.com/OpenIMSDK/openKeeper"
	"net"
	"strconv"

	"OpenIM/pkg/common/constant"
)

func main() {
	apiCmd := cmd.NewApiCmd()
	apiCmd.AddPortFlag()
	apiCmd.AddApi(run)
	if err := apiCmd.Execute(); err != nil {
		panic(err.Error())
	}
}

func run(port int) error {
	if port == 0 {
		port = config.Config.Api.GinPort[0]
	}
	zk, err := openKeeper.NewClient(config.Config.Zookeeper.ZkAddr, config.Config.Zookeeper.Schema, 10, config.Config.Zookeeper.UserName, config.Config.Zookeeper.Password)
	if err != nil {
		return err
	}
	log.NewPrivateLog(constant.LogFileName)
	zk.AddOption(mw.GrpcClient())
	router := api.NewGinRouter(zk)
	var address string
	if config.Config.Api.ListenIP != "" {
		address = net.JoinHostPort(config.Config.Api.ListenIP, strconv.Itoa(port))
	} else {
		address = net.JoinHostPort("0.0.0.0", strconv.Itoa(port))
	}
	fmt.Println("start api server, address: ", address, ", OpenIM version: ", config.Version)
	log.ZInfo(context.Background(), "start server success", "address", address, "version", config.Version)
	log.ZDebug(context.Background(), "start server success", "address", address, "version", config.Version)
	log.ZError(context.Background(), "start server success", errors.New("ss"), "address", address)
	log.ZWarn(context.Background(), "start server success", errors.New("ss"), "address", address)
	err = router.Run(address)
	if err != nil {
		log.Error("", "api run failed ", address, err.Error())
		return err
	}
	return nil
}