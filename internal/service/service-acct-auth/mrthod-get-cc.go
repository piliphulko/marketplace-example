package serviceacctauth

import (
	"context"

	"github.com/piliphulko/marketplace-example/api/basic"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *server) GetCountryCity(ctx context.Context, in *emptypb.Empty) (*basic.CountryCityPairs, error) {
	conn, err := c.AcquireConn(ctx)
	if err != nil {
		return &basic.CountryCityPairs{}, err
	}
	defer conn.Release()
	const query string = `
	SELECT country::varchar, city::varchar
	FROM table_country_city`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return &basic.CountryCityPairs{}, err
	}
	var countryCityPairs []*basic.CountryCityPair
	for rows.Next() {
		var cc *basic.CountryCityPair
		if err := rows.Scan(&cc.Country, &cc.City); err != nil {
			return &basic.CountryCityPairs{}, err
		}
		countryCityPairs = append(countryCityPairs, cc)
	}
	if err := rows.Err(); err != nil {
		return &basic.CountryCityPairs{}, err
	}
	return &basic.CountryCityPairs{CountryCityPairs: countryCityPairs}, status.New(codes.OK, "").Err()
}
