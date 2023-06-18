package serviceacctaut_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"testing"

	"github.com/piliphulko/marketplace-example/api/basic"
	pbClient "github.com/piliphulko/marketplace-example/api/service-acct-aut"
	"github.com/piliphulko/marketplace-example/internal/pkg/logwriter"
	pb "github.com/piliphulko/marketplace-example/internal/service/service-acct-aut"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	pb.InitJWTSecret("testSecret")
	logerSync := logwriter.InitStdoutLoggerGRPC(&pb.LogGRPC)
	defer logerSync()
	lis, err := net.Listen(
		"tcp",
		":50050",
	)
	if err != nil {
		panic(err)
	}
	var (
		grpcServer = grpc.NewServer(
			grpc.ChainUnaryInterceptor(pb.InterceptotCheckCtx),
		)
		server = pb.StartServer()
	)

	close, err := server.ConnPostrgresql("postgres://postgres:5432@localhost:5432/test_db")
	defer close()
	if err != nil {
		panic(err)
	}

	pb.RegisterServer(grpcServer, server)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Println(err)
		}
	}()
	mRun := m.Run()
	lis.Close()
	conn, _ := server.AcquireConn(context.TODO())
	conn.Exec(context.TODO(),
		`TRUNCATE table_warehouse, table_warehouse_info, table_vendor, table_vendor_info, table_customer, table_customer_info CASCADE;`)
	os.Exit(mRun)
}

func TestCustomer(t *testing.T) {
	conn, closeConn, err := pbClient.ConnToMicroserverAccountAuthentication(":50050")
	require.Nil(t, err)
	defer closeConn()
	var (
		ctx = context.Background()
		wg  = sync.WaitGroup{}
	)
	wg.Add(1)

	go func() {
		defer wg.Done()
		_, err := conn.CreateAccount(ctx, pbClient.OneofAccount(basic.CustomerChange{
			CustomerAutNew: &basic.CustomerAut{
				LoginCustomer:    "newUser",
				PasswortCustomer: "123456as",
			},
			CustomerInfo: &basic.CustomerInfo{
				CustomerCountry: "BELARUS",
				CustomerCiry:    "MINSK",
			},
		}))
		require.Nil(t, err)
		jwtString, err := conn.AutAccount(ctx, pbClient.OneofLoginPass(basic.CustomerAut{
			LoginCustomer:    "newUser",
			PasswortCustomer: "123456as",
		}))
		require.Nil(t, err)
		_, err = conn.CheckJWT(ctx, jwtString)
		require.Nil(t, err)
	}()

	wg.Wait()
}
