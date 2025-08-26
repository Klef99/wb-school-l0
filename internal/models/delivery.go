package models

import (
	"fmt"
	"regexp"
	"strings"
)

type Delivery struct {
	orderUID string
	name     string
	phone    string
	zip      string
	city     string
	address  string
	region   string
	email    string
}

func (d *Delivery) OrderUID() string {
	return d.orderUID
}

func (d *Delivery) Name() string {
	return d.name
}
func (d *Delivery) Phone() string {
	return d.phone
}
func (d *Delivery) Zip() string {
	return d.zip
}
func (d *Delivery) City() string {
	return d.city
}
func (d *Delivery) Address() string {
	return d.address
}
func (d *Delivery) Region() string {
	return d.region
}
func (d *Delivery) Email() string {
	return d.email
}

func (d *Delivery) SetOrderUID(orderUID string) error {
	if err := validateOrderUID(orderUID); err != nil {
		return err
	}

	d.orderUID = orderUID

	return nil
}

func (d *Delivery) SetName(name string) error {
	if err := validateName(name); err != nil {
		return err
	}

	d.name = name

	return nil
}

func (d *Delivery) SetPhone(phone string) error {
	if err := validatePhone(phone); err != nil {
		return err
	}

	d.phone = phone

	return nil
}

func (d *Delivery) SetZip(zip string) error {
	if err := validateZip(zip); err != nil {
		return err
	}

	d.zip = zip

	return nil
}

func (d *Delivery) SetCity(city string) error {
	if err := validateCity(city); err != nil {
		return err
	}

	d.city = city

	return nil
}

func (d *Delivery) SetAddress(address string) error {
	if err := validateAddress(address); err != nil {
		return err
	}

	d.address = address

	return nil
}

func (d *Delivery) SetRegion(region string) error {
	if err := validateRegion(region); err != nil {
		return err
	}

	d.region = region

	return nil
}

func (d *Delivery) SetEmail(email string) error {
	if err := validateEmail(email); err != nil {
		return err
	}

	d.email = email

	return nil
}

func NewDelivery(orderUID, name, phone, zip, city, address, region, email string) (Delivery, error) {
	d := Delivery{
		orderUID: orderUID,
		name:     name,
		phone:    phone,
		zip:      zip,
		city:     city,
		address:  address,
		region:   region,
		email:    email,
	}

	if err := d.Validate(); err != nil {
		return Delivery{}, err
	}

	return d, nil
}

// Validators

func validateOrderUID(orderUid string) error {
	if orderUid == "" {
		return ValidationError{
			Field:   "orderUID",
			Value:   orderUid,
			Message: "empty orderUID",
		}
	}

	return nil
}

func validateName(name string) error {
	if name == "" {
		return ValidationError{
			Field:   "name",
			Value:   name,
			Message: "empty name",
		}
	}

	return nil
}

func validatePhone(phone string) error {
	re := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	phone = strings.TrimSpace(phone)

	if !re.MatchString(phone) {
		return ValidationError{
			Field:   "phone",
			Value:   phone,
			Message: "phone is not valid e164",
		}
	}

	return nil
}

func validateZip(zip string) error {
	if zip == "" {
		return ValidationError{
			Field:   "zip",
			Value:   zip,
			Message: "empty zip",
		}
	}

	return nil
}

func validateCity(city string) error {
	if city == "" {
		return ValidationError{
			Field:   "city",
			Value:   city,
			Message: "empty city",
		}
	}

	return nil
}

func validateAddress(address string) error {
	if address == "" {
		return ValidationError{
			Field:   "address",
			Value:   address,
			Message: "empty address",
		}
	}

	return nil
}

func validateRegion(region string) error {
	if region == "" {
		return ValidationError{
			Field:   "region",
			Value:   region,
			Message: "empty region",
		}
	}

	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return ValidationError{
			Field:   "email",
			Value:   email,
			Message: "empty email",
		}
	}

	return nil
}

func (d *Delivery) Validate() error {
	var res []error

	if err := validateOrderUID(d.OrderUID()); err != nil {
		res = append(res, err)
	}

	if err := validateName(d.Name()); err != nil {
		res = append(res, err)
	}

	if err := validatePhone(d.Phone()); err != nil {
		res = append(res, err)
	}

	if err := validateZip(d.Zip()); err != nil {
		res = append(res, err)
	}

	if err := validateCity(d.City()); err != nil {
		res = append(res, err)
	}

	if err := validateAddress(d.Address()); err != nil {
		res = append(res, err)
	}

	if err := validateRegion(d.Region()); err != nil {
		res = append(res, err)
	}

	if err := validateEmail(d.Email()); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		msgs := make([]string, 0, len(res))

		for _, err := range res {
			msgs = append(msgs, err.Error())
		}

		return fmt.Errorf("validation failure list: " + strings.Join(msgs, "; "))
	}

	return nil

}
