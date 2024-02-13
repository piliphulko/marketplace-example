package gatehttp

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofrs/uuid"
	"github.com/piliphulko/marketplace-example/api/basic"
	pbAA "github.com/piliphulko/marketplace-example/api/service-acct-aut"
	"github.com/piliphulko/marketplace-example/internal/pkg/gatehttp/opt"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func handlerCustomerCreateSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf               = bytes.Buffer{}
		login_customer    = r.FormValue("login_customer")
		password_customer = r.FormValue("password_customer")

		country = r.FormValue("country")
		city    = r.FormValue("city")
	)

	fmt.Println(login_customer, password_customer, country, city)

	_, err := ConnAA.CreateAccount(ctx, pbAA.OneofAccount(basic.CustomerChange{
		CustomerAutNew: &basic.CustomerAut{
			LoginCustomer:    login_customer,
			PasswortCustomer: password_customer,
		},
		CustomerInfo: &basic.CustomerInfo{
			CustomerCountry: country,
			CustomerCity:    city,
		},
	}))

	if err != nil {
		if err = opt.WriteRedirectAnswerInfoErr(&buf, cutMessageFromGrpcAnswer(err, cancelCtxError)); err != nil {
			cancelCtxError(err)
		}
	} else {
		if err := opt.WriteRedirectAnswerInfoOk(&buf, "Сreated"); err != nil {
			cancelCtxError(err)
		}
	}

	ch <- buf.Bytes()
}

func handlerCustomerAuthorizationSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf               = bytes.Buffer{}
		login_customer    = r.FormValue("login_customer")
		password_customer = r.FormValue("password_customer")
	)
	stringJWT, err := ConnAA.AutAccount(ctx,
		pbAA.OneofLoginPass(basic.CustomerAut{
			LoginCustomer:    login_customer,
			PasswortCustomer: password_customer,
		}))
	if err != nil {
		if err = opt.WriteRedirectAnswerInfoErr(&buf, cutMessageFromGrpcAnswer(err, cancelCtxError)); err != nil {
			cancelCtxError(err)
			return
		}
		return
	} else {
		if err := opt.WriteRedirectAnswerCookie(&buf, "JWT", stringJWT.StringJwt); err != nil {
			cancelCtxError(err)
			return
		}
	}
	//optHTTP.RedirectHTTPPathOk = strings.ReplaceAll(optHTTP.RedirectHTTPPathOk, "{login_customer}", login_customer)

	ch <- buf.Bytes()
}

func handlerCustomerCreatePage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf = bytes.Buffer{}
	)

	defer r.Body.Close()
	redirecrAnswer, err := opt.TakeRedirectAnswerFromURL(r)
	if err != nil {
		cancelCtxError(err)
		return
	}

	if err := optHTTP.TakeHTML().Execute(&buf, struct {
		RedirectAnswer interface{}
		Countries      []string
		Cities         []string
	}{
		RedirectAnswer: redirecrAnswer,
		Countries: []string{
			"BELARUS",
			"POLAND",
			"UKRAINE",
		},
		Cities: []string{
			"MINSK",
			"WARSAW",
			"KYIV",
		},
	}); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}

func handlerCustomerHomePage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf = bytes.Buffer{}
		//loginFromUrl = chi.URLParam(r, "login_customer")
	)
	defer r.Body.Close()
	redirecrAnswer, err := opt.TakeRedirectAnswerFromURL(r)
	if err != nil {
		cancelCtxError(err)
		return
	}
	cookieJWT, err := r.Cookie("JWT")
	if err != nil {
		cancelCtxError(err)
		return
	}
	loginFromJwt, err := takeNickname(cookieJWT.Value)
	if err != nil {
		cancelCtxError(err)
		return
	}
	/*
		if loginFromJwt != loginFromUrl {
			cancelCtxError(errors.New(""))
			return
		}
	*/
	md := metadata.Pairs("Authorization", fmt.Sprintf("Bearer %s", cookieJWT.Value))
	mdCtx := metadata.NewOutgoingContext(ctx, md)
	ordersUnconfirmed, err := ConnDC.GetCustomerOrders(mdCtx, &basic.OrderStatus{OrderStatus: basic.OrderStatusEnum_UNCONFIRMED})
	if err != nil {
		cancelCtxError(err)
		return
	}
	ordersConfirmed, err := ConnDC.GetCustomerOrders(mdCtx, &basic.OrderStatus{OrderStatus: basic.OrderStatusEnum_CONFIRNED})
	if err != nil {
		cancelCtxError(err)
		return
	}
	ordersCompleted, err := ConnDC.GetCustomerOrders(mdCtx, &basic.OrderStatus{OrderStatus: basic.OrderStatusEnum_COMPLETED})
	if err != nil {
		cancelCtxError(err)
		return
	}
	fmt.Println("oooookkkkkk")
	walletInfo, err := ConnDC.GetWalletInfo(mdCtx, &emptypb.Empty{})
	if err != nil {
		cancelCtxError(err)
		return
	}
	if err := optHTTP.TakeHTML().Execute(&buf, struct {
		RedirectAnswer         interface{}
		LoginCustomer          string
		WalletMoney            float64
		UnconfirmedOrdersARRAY []*basic.Order
		СonfirmedOrdersARRAY   []*basic.Order
		HistoryOrdersARRAY     []*basic.Order
	}{
		RedirectAnswer:         redirecrAnswer,
		LoginCustomer:          loginFromJwt,
		WalletMoney:            float64(walletInfo.AmountMoney),
		UnconfirmedOrdersARRAY: ordersUnconfirmed.Orders,
		СonfirmedOrdersARRAY:   ordersConfirmed.Orders,
		HistoryOrdersARRAY:     ordersCompleted.Orders,
	}); err != nil {
		if err != nil {
			cancelCtxError(err)
			return
		}
	}
	ch <- buf.Bytes()
}

func handlerCustomerHomeChangePage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf = bytes.Buffer{}
		//loginFromUrl = chi.URLParam(r, "login_customer")
	)
	defer r.Body.Close()
	redirecrAnswer, err := opt.TakeRedirectAnswerFromURL(r)
	if err != nil {
		cancelCtxError(err)
		return
	}
	cookieJWT, err := r.Cookie("JWT")
	if err != nil {
		cancelCtxError(err)
		return
	}
	loginFromJwt, err := takeNickname(cookieJWT.Value)
	if err != nil {
		cancelCtxError(err)
		return
	}
	/*
		if loginFromJwt != loginFromUrl {
			cancelCtxError(errors.New(""))
			return
		}
	*/

	md := metadata.Pairs("Authorization", fmt.Sprintf("Bearer %s", cookieJWT.Value))
	mdCtx := metadata.NewOutgoingContext(ctx, md)

	customerInfo, err := ConnDC.GetCustomerInfo(mdCtx, &emptypb.Empty{})
	if err != nil {
		cancelCtxError(err)
		return
	}
	countryCity, err := ConnAA.GetCountryCity(ctx, &emptypb.Empty{})
	if err != nil {
		cancelCtxError(err)
		return
	}
	var countrys, citys []string
	for _, v := range countryCity.CountryCityPairs {
		countrys = append(countrys, v.Country)
		citys = append(citys, v.City)
	}
	if err := optHTTP.TakeHTML().Execute(&buf, struct {
		RedirectAnswer  interface{}
		LoginCustomer   string
		CountryCustomer string
		CityCustomer    string
		Countries       []string
		Cities          []string
	}{
		RedirectAnswer:  redirecrAnswer,
		LoginCustomer:   loginFromJwt,
		CountryCustomer: customerInfo.CustomerCountry,
		CityCustomer:    customerInfo.CustomerCity,
		Countries:       countrys,
		Cities:          citys,
	}); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}

func handlerCustomerHomeChangeSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf = bytes.Buffer{}
		//login_customer = chi.URLParam(r, "login_customer")

		login_old    = r.FormValue("login")
		password_old = r.FormValue("password_old")
		login_new    = r.FormValue("login_new")
		password_new = r.FormValue("password_new")
		country      = r.FormValue("country")
		city         = r.FormValue("city")
	)
	/*
		cookieJWT, err := r.Cookie("JWT")
		if err != nil {
			cancelCtxError(err)
			return
		}
		loginFromJwt, err := takeNickname(cookieJWT.Value)
		if err != nil {
			cancelCtxError(err)
			return
		}
		optHTTP.RedirectHTTPPathMistake = strings.ReplaceAll(optHTTP.RedirectHTTPPathMistake, "{login_customer}", loginFromJwt)
		/*
		if loginFromJwt != login_old || loginFromJwt != login_customer {
			cancelCtxError(errors.New(""))
			return
		}
	*/
	/*
		md := metadata.Pairs("Authorization", fmt.Sprintf("Bearer %s", cookieJWT.Value))
		mdCtx := metadata.NewOutgoingContext(ctx, md)
	*/
	_, err := ConnAA.ChangeAccount(ctx, pbAA.OneofAccount(basic.CustomerChange{
		CustomerAutOld: &basic.CustomerAut{
			LoginCustomer:    login_old,
			PasswortCustomer: password_old,
		},
		CustomerAutNew: &basic.CustomerAut{
			LoginCustomer:    login_new,
			PasswortCustomer: password_new,
		},
		CustomerInfo: &basic.CustomerInfo{
			CustomerCountry: country,
			CustomerCity:    city,
		},
	}))
	if err != nil {
		if err = opt.WriteRedirectAnswerInfoErr(&buf, cutMessageFromGrpcAnswer(err, cancelCtxError)); err != nil {
			cancelCtxError(err)
			return
		}
		return
	}
	if err := opt.WriteRedirectAnswerCookie(&buf, "JWT", ""); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}

func handlerCustomerHomeWalletPage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf = bytes.Buffer{}
		//loginFromUrl = chi.URLParam(r, "login_customer")
	)
	defer r.Body.Close()
	redirecrAnswer, err := opt.TakeRedirectAnswerFromURL(r)
	if err != nil {
		cancelCtxError(err)
		return
	}
	cookieJWT, err := r.Cookie("JWT")
	if err != nil {
		cancelCtxError(err)
		return
	}
	loginFromJwt, err := takeNickname(cookieJWT.Value)
	if err != nil {
		cancelCtxError(err)
		return
	}
	/*
		if loginFromJwt != loginFromUrl {
			cancelCtxError(errors.New(""))
			return
		}
	*/
	md := metadata.Pairs("Authorization", fmt.Sprintf("Bearer %s", cookieJWT.Value))
	mdCtx := metadata.NewOutgoingContext(ctx, md)

	walletInfo, err := ConnDC.GetWalletInfo(mdCtx, &emptypb.Empty{})
	if err != nil {
		cancelCtxError(err)
		return
	}

	if err := optHTTP.TakeHTML().Execute(&buf, struct {
		RedirectAnswer interface{}
		WalletMoney    float64
		BlockedMoney   float64
		LoginCustomer  string
		GoodsArray     []*basic.Goods
	}{
		RedirectAnswer: redirecrAnswer,
		WalletMoney:    float64(walletInfo.AmountMoney),
		BlockedMoney:   float64(walletInfo.BlockedMoney),
		LoginCustomer:  loginFromJwt,
	}); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}

func handlerMarketplaceCustomerPage(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf = bytes.Buffer{}
	)
	defer r.Body.Close()
	redirecrAnswer, err := opt.TakeRedirectAnswerFromURL(r)
	if err != nil {
		cancelCtxError(err)
		return
	}
	cookieJWT, err := r.Cookie("JWT")
	if err != nil {
		cancelCtxError(err)
		return
	}
	loginFromJwt, err := takeNickname(cookieJWT.Value)
	if err != nil {
		cancelCtxError(err)
		return
	}

	goodsArray, err := ConnOP.GetMarketplace(ctx, &emptypb.Empty{})
	if err != nil {
		cancelCtxError(err)
		return
	}

	md := metadata.Pairs("Authorization", fmt.Sprintf("Bearer %s", cookieJWT.Value))
	mdCtx := metadata.NewOutgoingContext(ctx, md)
	ordersUnconfirmed, err := ConnDC.GetCustomerOrders(mdCtx, &basic.OrderStatus{OrderStatus: basic.OrderStatusEnum_UNCONFIRMED})
	if err != nil {
		cancelCtxError(err)
		return
	}

	var newUuid string
	if ordersUnconfirmed.Orders != nil {
		fmt.Println("newUuid")
		newUuid = ordersUnconfirmed.Orders[0].OrderUuid
	} else {
		newUuidTyper, err := uuid.NewV4()
		if err != nil {
			cancelCtxError(err)
			return
		}
		newUuid = newUuidTyper.String()
	}
	fmt.Println(newUuid)
	if err := optHTTP.TakeHTML().Execute(&buf, struct {
		RedirectAnswer interface{}
		LoginCustomer  string
		NewUuid        string
		GoodsArray     []*basic.Goods
		OrdersArray    []*basic.Order
	}{
		RedirectAnswer: redirecrAnswer,
		LoginCustomer:  loginFromJwt,
		NewUuid:        newUuid,
		GoodsArray:     goodsArray.Goods,
		OrdersArray:    ordersUnconfirmed.Orders,
	}); err != nil {
		cancelCtxError(err)
		return
	}
	ch <- buf.Bytes()
}

func handlerMarketplaceSend(ctx context.Context, cancelCtxError context.CancelCauseFunc, optHTTP *opt.OptionsHTTP, r *http.Request, ch chan []byte) {
	var (
		buf            = bytes.Buffer{}
		name_warehouse = r.FormValue("name_warehouse")
		name_vendor    = r.FormValue("name_vendor")
		name_goods     = r.FormValue("name_goods")
		amount_goods   = r.FormValue("amount_goods")
	)
	amount_goods_int, err := strconv.Atoi(amount_goods)
	if err != nil {
		cancelCtxError(err)
		return
	}
	cookieJWT, err := r.Cookie("JWT")
	if err != nil {
		cancelCtxError(err)
		return
	}
	md := metadata.Pairs("Authorization", fmt.Sprintf("Bearer %s", cookieJWT.Value))
	mdCtx := metadata.NewOutgoingContext(ctx, md)

	uuidNewOrder, err := ConnOP.CreateOrder(mdCtx, &basic.NewOrderARRAY{NewOrder: []*basic.NewOrder{
		{
			NameWarehouse: name_warehouse,
			NameVendor:    name_vendor,
			NameGoods:     name_goods,
			AmountGoods:   uint32(amount_goods_int),
		},
	}})

	if err != nil {
		if err = opt.WriteRedirectAnswerInfoErr(&buf, cutMessageFromGrpcAnswer(err, cancelCtxError)); err != nil {
			cancelCtxError(err)
			return
		}
		return
	}
	if err := opt.WriteRedirectAnswerInfoOk(&buf, fmt.Sprintf("order created: %s", uuidNewOrder)); err != nil {
		cancelCtxError(err)
		return
	}

	ch <- buf.Bytes()
}
