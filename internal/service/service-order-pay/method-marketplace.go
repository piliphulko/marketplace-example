package serviceorderpay

import (
	"context"
	"fmt"

	"github.com/piliphulko/marketplace-example/api/basic"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) GetMarketplace(ctx context.Context, in *emptypb.Empty) (*basic.GoodsArray, error) {
	/*
		jwtString, err := grpctools.TakeJWTfromMetadata(ctx)
		if err != nil {
			return &basic.GoodsArray{}, err
		}
		_, err = grpctools.TakeLoginAndCheckJWT(jwtString)
		if err != nil {
			return &basic.GoodsArray{}, err
		}
	*/
	conn, err := s.AcquireConn(ctx)
	if err != nil {
		return &basic.GoodsArray{}, err
	}
	defer conn.Release()

	const query string = `
	SELECT
		name_warehouse,
		CONCAT(country, ', ', city) AS location,
		name_vendor,
		type_goods,
		name_goods,
		info_goods,
		price_goods,
		amount_goods_available
	FROM view_marketplace`

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return &basic.GoodsArray{}, err
	}
	defer rows.Close()
	var result []*basic.Goods
	for rows.Next() {
		var v basic.Goods
		if err := rows.Scan(&v.NameWarehouse, &v.Location, &v.NameVendor, &v.TypeGoods, &v.NameGoods, &v.InfoGoods, &v.PriceGoods, &v.AmountGoods); err != nil {
			return &basic.GoodsArray{}, err
		}
		result = append(result, &v)
	}
	if err := rows.Err(); err != nil {
		return &basic.GoodsArray{}, err
	}
	fmt.Println(result)

	return &basic.GoodsArray{Goods: result}, status.New(codes.OK, "").Err()
}
