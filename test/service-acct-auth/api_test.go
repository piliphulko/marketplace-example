package serviceacctaut_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/piliphulko/marketplace-example/api/basic"
	pbClient "github.com/piliphulko/marketplace-example/api/service-acct-aut"
	"github.com/piliphulko/marketplace-example/internal/pkg/logwriter"
	pb "github.com/piliphulko/marketplace-example/internal/service/service-acct-auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestMain(m *testing.M) {
	pb.InitJWTSecret("testSecret")
	//logerSync := logwriter.InitStdoutLoggerGRPC(&pb.LogGRPC)
	logerSync := logwriter.NewZapLogStdoutGRPC(&pb.LogGRPC, zapcore.ErrorLevel)
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

func TestInterceptor(t *testing.T) {
	conn, closeConn, err := pbClient.ConnToServiceAccountAuthentication(":50050")
	require.Nil(t, err)
	defer closeConn()
	ctx1, cancelCtx1 := context.WithCancel(context.Background())
	cancelCtx1()

	_, err = conn.AutAccount(ctx1, nil)
	st, _ := status.FromError(err)
	assert.Equal(t, st.Code(), codes.Canceled)

	ctx2, cancelCtx2 := context.WithTimeout(context.Background(), time.Microsecond*100)
	defer cancelCtx2()
	time.Sleep(time.Microsecond * 200)

	_, err = conn.UpdateAccount(ctx2, nil)
	st, _ = status.FromError(err)
	assert.Equal(t, st.Code(), codes.DeadlineExceeded)
}

func TestCustomer(t *testing.T) {
	conn, closeConn, err := pbClient.ConnToServiceAccountAuthentication(":50050")
	require.Nil(t, err)
	defer closeConn()
	var (
		ctx = context.Background()
		wg  = sync.WaitGroup{}
	)
	wg.Add(3)

	// NO ERRORS
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
		assert.Nil(t, err)

		jwtString, err := conn.AutAccount(ctx, pbClient.OneofLoginPass(basic.CustomerAut{
			LoginCustomer:    "newUser",
			PasswortCustomer: "123456as",
		}))
		assert.Nil(t, err)

		_, err = conn.CheckJWT(ctx, jwtString)
		assert.Nil(t, err)

		_, err = conn.UpdateAccount(ctx, pbClient.OneofAccount(basic.CustomerChange{
			CustomerAutOld: &basic.CustomerAut{
				LoginCustomer:    "newUser",
				PasswortCustomer: "123456as",
			},
			CustomerAutNew: &basic.CustomerAut{
				LoginCustomer:    "newUser2",
				PasswortCustomer: "123456asNew",
			},
			CustomerInfo: &basic.CustomerInfo{
				CustomerCountry: "BELARUS",
				CustomerCiry:    "",
			},
		}))
		assert.Nil(t, err)

	}()

	// CREATE ACCOUNT ERRORS
	go func() {
		defer wg.Done()
		_, err := conn.CreateAccount(ctx, pbClient.OneofAccount(basic.CustomerChange{
			CustomerAutNew: &basic.CustomerAut{
				LoginCustomer:    "testErr",
				PasswortCustomer: "123456as",
			},
			CustomerInfo: &basic.CustomerInfo{
				CustomerCountry: "BELARUS",
				CustomerCiry:    "MINSK",
			},
		}))
		assert.Nil(t, err)

		_, err = conn.CreateAccount(ctx, pbClient.OneofAccount(basic.CustomerChange{
			CustomerAutNew: &basic.CustomerAut{
				LoginCustomer:    "testErrNew",
				PasswortCustomer: "123456a",
			},
			CustomerInfo: &basic.CustomerInfo{
				CustomerCountry: "BELARUS",
				CustomerCiry:    "MINSK",
			},
		}))
		st, _ := status.FromError(err)
		assert.Equal(t, st.Code(), codes.InvalidArgument)
		assert.Equal(t, st.Message(), pbClient.ErrPassLen.Error())

		_, err = conn.CreateAccount(ctx, pbClient.OneofAccount(basic.CustomerChange{
			CustomerAutNew: &basic.CustomerAut{
				LoginCustomer:    "testErr",
				PasswortCustomer: "123456as",
			},
			CustomerInfo: &basic.CustomerInfo{
				CustomerCountry: "BELARUS",
				CustomerCiry:    "MINSK",
			},
		}))
		st, _ = status.FromError(err)
		assert.Equal(t, st.Code(), codes.AlreadyExists)

		_, err = conn.CreateAccount(ctx, pbClient.OneofAccount(basic.CustomerChange{
			CustomerAutNew: &basic.CustomerAut{
				LoginCustomer:    "testErrNew",
				PasswortCustomer: "123456as",
			},
			CustomerInfo: &basic.CustomerInfo{
				CustomerCountry: "BELARUSs",
				CustomerCiry:    "MINSK",
			},
		}))
		st, _ = status.FromError(err)
		assert.Equal(t, st.Code(), codes.InvalidArgument)
		assert.Equal(t, st.Message(), pbClient.ErrIncorrectCountry.Error())
	}()

	// UPDATE ACCOUNT ERRORS
	go func() {
		defer wg.Done()
		_, err := conn.CreateAccount(context.TODO(), pbClient.OneofAccount(basic.CustomerChange{
			CustomerAutNew: &basic.CustomerAut{
				LoginCustomer:    "testUpdate",
				PasswortCustomer: "123456as",
			},
			CustomerInfo: &basic.CustomerInfo{
				CustomerCountry: "BELARUS",
				CustomerCiry:    "MINSK",
			},
		}))
		assert.Nil(t, err)
		// UPDATE LOGIN
		_, err = conn.UpdateAccount(context.TODO(), pbClient.OneofAccount(basic.CustomerChange{
			CustomerAutOld: &basic.CustomerAut{
				LoginCustomer:    "testUpdate",
				PasswortCustomer: "123456as",
			},
			CustomerAutNew: &basic.CustomerAut{
				LoginCustomer:    "testUpdate2",
				PasswortCustomer: "",
			},
			CustomerInfo: nil,
		}))
		assert.Nil(t, err)
		// FAIL UPDATE PASSWORT
		_, err = conn.UpdateAccount(context.TODO(), pbClient.OneofAccount(basic.CustomerChange{
			CustomerAutOld: &basic.CustomerAut{
				LoginCustomer:    "testUpdate2",
				PasswortCustomer: "123456as",
			},
			CustomerAutNew: &basic.CustomerAut{
				LoginCustomer:    "testUpdate2",
				PasswortCustomer: "123456",
			},
		}))
		st, _ := status.FromError(err)
		assert.Equal(t, st.Code(), codes.InvalidArgument)
		assert.Equal(t, st.Message(), pbClient.ErrPassLen.Error())
		// UPDATE PASSWORT
		_, err = conn.UpdateAccount(context.TODO(), pbClient.OneofAccount(basic.CustomerChange{
			CustomerAutOld: &basic.CustomerAut{
				LoginCustomer:    "testUpdate2",
				PasswortCustomer: "123456as",
			},
			CustomerAutNew: &basic.CustomerAut{
				LoginCustomer:    "testUpdate2",
				PasswortCustomer: "12345678",
			},
		}))
		assert.Nil(t, err)
		/*
			poolPgx, err := pgxpool.New(context.Background(), "postgres://postgres:5432@localhost:5432/test_db")
			require.Nil(t, err)
			pc, err := poolPgx.Exec(context.Background(), `
				SELECT true::bool
				FROM table_customer
				WHERE login_customer = 'testUpdate2' AND passwort_customer = '';`)
			assert.Nil(t, err)
		*/
	}()
	/*
		// AUT
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
			assert.Nil(t, err)

			jwtString, err := conn.AutAccount(ctx, pbClient.OneofLoginPass(basic.CustomerAut{
				LoginCustomer:    "newUser",
				PasswortCustomer: "123456as",
			}))
			assert.Nil(t, err)
		}()
	*/

	wg.Wait()
}
