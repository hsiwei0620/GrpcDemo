package main

import (
	"flag"
	"fmt"
	// "github.com/golang/protobuf/proto"
	// "google.golang.org/grpc/encoding/proto"

	// "log"
	// "net"
	"os"
	"strconv"
	// "sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"path/to/project/MqGrpcProject"
)

var (
	port = flag.Uint("port", 8080, "")
	// subno, factory, barcode string
	resveqptid string
)


func callService() {
	logger := zap.L()
	defer logger.Sync()

	logger.Info("calling service")

	conn, err := grpc.Dial("192.1.1.61:8080", grpc.WithInsecure())
	
	if err != nil {
		logger.Error("Connection failed ：%v", zap.Error(err))
	}
	defer conn.Close()

	client := MqGrpcProject.NewMqGrpcsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Millisecond)
	defer cancel()

	serverip := "192.1.1.185"
	actiontype := "Q"
	onlyresvflg := "Y"

	// fmt.Printf("%s\n", ctx)

	rep, err := client.APMEQRSV_Send(ctx, &MqGrpcProject.APMEQRSV_Request{Serverip: serverip, Actiontype: actiontype, Resveqptid: resveqptid, Onlyresvflg: onlyresvflg})
	if err != nil {
		// logger.Error("error happened", zap.Error(err))
		//fmt.Printf(err.Error() + "\n")
		fmt.Printf("%s\n", err)
	} else {

		lcnt, _ := strconv.Atoi(rep.GetLotarycnt())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Printf("lot_cnt : " + rep.GetLotarycnt() + "\n")

		for i := 0; i < lcnt; i++ {
			fmt.Printf("lot_id  : " + rep.Oary[i].GetLotid() + "\n")
		}

		// logger.Info("resv information", zap.Uint32("lot", rep.Oary[0].Lotid))
	}
}



func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)



	for {

		fmt.Printf("Eqpt_ID : ")
		fmt.Scanln(&resveqptid)

		callService()

	}
}
