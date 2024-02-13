package basiccheck

import "github.com/piliphulko/marketplace-example/api/basic"

type AllBasic interface {
	*basic.NewOrder | *basic.OrderUuid
}

func NillOrNull[T AllBasic](v T) bool {
	switch v := any(v).(type) {
	case *basic.NewOrder:
		if &v.NameWarehouse == nil || v.NameWarehouse == "" ||
			&v.NameVendor == nil || v.NameVendor == "" ||
			&v.NameGoods == nil || v.NameGoods == "" ||
			&v.AmountGoods == nil || v.AmountGoods = 0 {
			return false
		}
		return true
	case *basic.OrderUuid:
		if &v == nil || v.OrderUuid == "" {
			return false
		}
		return true
	}
	return true
}
