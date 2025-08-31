package models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Delivery struct {
	OrderUID string
	Name     string
	Phone    string
	Zip      string
	City     string
	Address  string
	Region   string
	Email    string
}

func (d Delivery) Validate() error {
	var errs []error

	if d.OrderUID == "" {
		errs = append(errs, ValidationError{columnDeliveryOrderUID, d.OrderUID, "empty orderUID"})
	}

	if d.Name == "" {
		errs = append(errs, ValidationError{columnDeliveryName, d.Name, "empty name"})
	}

	if d.Phone == "" {
		errs = append(errs, ValidationError{columnDeliveryPhone, d.Phone, "empty phone"})
	} else {
		re := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
		if !re.MatchString(strings.TrimSpace(d.Phone)) {
			errs = append(errs, ValidationError{columnDeliveryPhone, d.Phone, "invalid phone (must be E.164)"})
		}
	}

	if d.Zip == "" {
		errs = append(errs, ValidationError{columnDeliveryZIP, d.Zip, "empty zip"})
	}

	if d.City == "" {
		errs = append(errs, ValidationError{columnDeliveryCity, d.City, "empty city"})
	}

	if d.Address == "" {
		errs = append(errs, ValidationError{columnDeliveryAddress, d.Address, "empty address"})
	}

	if d.Region == "" {
		errs = append(errs, ValidationError{columnDeliveryRegion, d.Region, "empty region"})
	}

	if d.Email == "" {
		errs = append(errs, ValidationError{columnDeliveryEmail, d.Email, "empty email"})
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failure: %w", errors.Join(errs...))
	}

	return nil
}

func NewDelivery(orderUID, name, phone, zip, city, address, region, email string) (Delivery, error) {
	d := Delivery{
		OrderUID: orderUID,
		Name:     name,
		Phone:    phone,
		Zip:      zip,
		City:     city,
		Address:  address,
		Region:   region,
		Email:    email,
	}

	if err := d.Validate(); err != nil {
		return Delivery{}, err
	}

	return d, nil
}

func NewDeliveryFromDB(orderUID, name, phone, zip, city, address, region, email string) Delivery {
	return Delivery{
		OrderUID: orderUID,
		Name:     name,
		Phone:    phone,
		Zip:      zip,
		City:     city,
		Address:  address,
		Region:   region,
		Email:    email,
	}
}
