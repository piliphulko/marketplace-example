package basic

import "github.com/piliphulko/marketplace-example/api/apierror"

type Basicer interface {
	PrimaryValidation()
}

func PasswotdRule(pass string) bool {
	if len(pass) < 8 || len(pass) > 64 {
		return false
	} else {
		return true
	}
}

func (in *CustomerAut) PrimaryValidation() error {
	switch {
	case in == nil:
		return apierror.ErrEmpty
	case in.LoginCustomer == "":
		return apierror.ErrEmpty
	case in.PasswordCustomer == "":
		return apierror.ErrEmpty
	case !PasswotdRule(in.PasswordCustomer):
		return apierror.ErrPassLen
	default:
		return nil
	}
}

func (in *CustomerInfo) PrimaryValidation() error {
	switch {
	case in == nil:
		return apierror.ErrEmpty
	case in.CustomerCountry == "":
		return apierror.ErrEmpty
	case in.CustomerCity == "":
		return apierror.ErrEmpty
	default:
		return nil
	}
}

func (in *CustomerNew) PrimaryValidation() error {
	if in == nil {
		return apierror.ErrEmpty
	} else if err := in.CustomerAut.PrimaryValidation(); err != nil {
		return err
	} else if err := in.CustomerInfo.PrimaryValidation(); err != nil {
		return err
	} else {
		return nil
	}
}

func (in *CustomerChange) PrimaryValidation() error {
	// Some fields of this type may not be filled in; additional processing logic is required on the API side
	if in == nil {
		return apierror.ErrEmpty
	} else if err := in.CustomerAutOld.PrimaryValidation(); err != nil {
		return err
	} else {
		return nil
	}
}

func (in *WarehouseAut) PrimaryValidation() error {
	switch {
	case in == nil:
		return apierror.ErrEmpty
	case in.LoginWarehouse == "":
		return apierror.ErrEmpty
	case in.PasswordWarehouse == "":
		return apierror.ErrEmpty
	case !PasswotdRule(in.PasswordWarehouse):
		return apierror.ErrEmpty
	default:
		return nil
	}
}

func (in *WarehouseInfo) PrimaryValidation() error {
	switch {
	case in == nil:
		return apierror.ErrEmpty
	case in.WarehouseName == "":
		return apierror.ErrEmpty
	case in.WarehouseCountry == "":
		return apierror.ErrEmpty
	case in.WarehouseCity == "":
		return apierror.ErrEmpty
	case in.WarehouseCommission == 0:
		return apierror.ErrEmpty
	default:
		return nil
	}
}

func (in *WarehouseNew) PrimaryValidation() error {
	if in == nil {
		return apierror.ErrEmpty
	} else if err := in.WarehouseAut.PrimaryValidation(); err != nil {
		return err
	} else if err := in.WarehouseInfo.PrimaryValidation(); err != nil {
		return err
	} else {
		return nil
	}
}

func (in *WarehouseChange) PrimaryValidation() error {
	// Some fields of this type may not be filled in; additional processing logic is required on the API side
	if in == nil {
		return apierror.ErrEmpty
	} else if err := in.WarehouseAutOld.PrimaryValidation(); err != nil {
		return err
	} else {
		return nil
	}
}

func (in *VendorAut) PrimaryValidation() error {
	switch {
	case in == nil:
		return apierror.ErrEmpty
	case in.LoginVendor == "":
		return apierror.ErrEmpty
	case in.PasswordVendor == "":
		return apierror.ErrEmpty
	case !PasswotdRule(in.PasswordVendor):
		return apierror.ErrPassLen
	default:
		return nil
	}
}

func (in *VendorInfo) PrimaryValidation() error {
	switch {
	case in == nil:
		return apierror.ErrEmpty
	case in.VendorName == "":
		return apierror.ErrEmpty
	default:
		return nil
	}
}

func (in *VendorNew) PrimaryValidation() error {
	if in == nil {
		return apierror.ErrEmpty
	} else if err := in.VendorAut.PrimaryValidation(); err != nil {
		return err
	} else if err := in.VendorInfo.PrimaryValidation(); err != nil {
		return err
	} else {
		return nil
	}
}

func (in *VendorChange) PrimaryValidation() error {
	// Some fields of this type may not be filled in; additional processing logic is required on the API side
	if in == nil {
		return apierror.ErrEmpty
	} else if err := in.VendorAutOld.PrimaryValidation(); err != nil {
		return apierror.ErrEmpty
	} else {
		return nil
	}
}
