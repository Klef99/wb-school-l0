package models

import (
	"errors"
	"fmt"
)

type Item struct {
	ChrtID      uint
	TrackNumber string
	Price       uint
	RID         string
	Name        string
	Sale        uint
	Size        string
	TotalPrice  uint
	NMID        uint
	Brand       string
	Status      uint
}

func (i Item) Validate() error {
	var errs []error

	if i.ChrtID == 0 {
		errs = append(errs, ValidationError{columnItemChrtID, fmt.Sprint(i.ChrtID), "empty ChrtID"})
	}

	if i.TrackNumber == "" {
		errs = append(errs, ValidationError{columnItemTrackNumber, i.TrackNumber, "empty TrackNumber"})
	}

	if i.Price == 0 {
		errs = append(errs, ValidationError{columnItemPrice, fmt.Sprint(i.Price), "zero Price"})
	}

	if i.RID == "" {
		errs = append(errs, ValidationError{columnItemRid, i.RID, "empty RID"})
	}

	if i.Name == "" {
		errs = append(errs, ValidationError{columnItemName, i.Name, "empty Name"})
	}

	if i.Sale > 99 {
		errs = append(errs, ValidationError{columnItemSale, i.Sale, "invalid Sale"})
	}

	if i.TotalPrice == 0 {
		errs = append(errs, ValidationError{columnItemTotalPrice, fmt.Sprint(i.TotalPrice), "zero TotalPrice"})
	}

	if i.NMID == 0 {
		errs = append(errs, ValidationError{columnItemNMID, fmt.Sprint(i.NMID), "zero NMID"})
	}

	if i.Brand == "" {
		errs = append(errs, ValidationError{columnItemBrand, i.Brand, "empty Brand"})
	}

	if i.Status == 0 {
		errs = append(errs, ValidationError{columnItemStatus, fmt.Sprint(i.Status), "zero Status"})
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failure: %w", errors.Join(errs...))
	}

	return nil
}

func NewItem(
	chrtID uint, trackNumber string, price uint, rid, name string, sale uint, size string, totalPrice, nmid uint,
	brand string, status uint,
) (Item, error) {
	i := Item{
		ChrtID:      chrtID,
		TrackNumber: trackNumber,
		Price:       price,
		RID:         rid,
		Name:        name,
		Sale:        sale,
		Size:        size,
		TotalPrice:  totalPrice,
		NMID:        nmid,
		Brand:       brand,
		Status:      status,
	}

	if err := i.Validate(); err != nil {
		return Item{}, err
	}

	return i, nil
}

func NewItemFromDB(
	chrtID uint, trackNumber string, price uint, rid, name string, sale uint, size string, totalPrice, nmid uint,
	brand string, status uint,
) Item {
	return Item{
		ChrtID:      chrtID,
		TrackNumber: trackNumber,
		Price:       price,
		RID:         rid,
		Name:        name,
		Sale:        sale,
		Size:        size,
		TotalPrice:  totalPrice,
		NMID:        nmid,
		Brand:       brand,
		Status:      status,
	}
}
