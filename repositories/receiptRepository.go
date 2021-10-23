package repositories

import (
    "fmt"

    "../models"
    "github.com/pkg/errors"
)

type ReceiptRepository struct {
    receipts map[string]models.Receipt
    parkingSlotMapping map[int]string
}

func GetReceiptRepository() ReceiptRepository {
    return ReceiptRepository{}
}

func (r ReceiptRepository) showActiveReceipts() (receipts []models.Receipt)  {
    return r.filterReceipts(models.ACTIVE)
}

func (r ReceiptRepository) showInactiveReceipts() (receipts []models.Receipt)  {
    return r.filterReceipts(models.INACTIVE)
}

func (r ReceiptRepository) filterReceipts(status models.Status) (receipts []models.Receipt)  {


    for _, receipt := range r.receipts {
        if receipt.GetStatus() == status {
            receipts = append(receipts, receipt)
        }
    }
    return receipts
}

func (r ReceiptRepository) CreateReceipt(vehicle models.Vehicle, parkingSpot models.ParkingSlot) (receipt models.Receipt, err error)  {
    receipt, err = models.NewReceipt(vehicle, parkingSpot)
    if err != nil {
        return nil, fmt.Errorf("error in generating receipts due to - %s", err.Error())
    }
    r.AddReceipt(receipt)
    r.AddParkingSlotMapping(parkingSpot.ID, receipt.GetID())
    return
}


func (r ReceiptRepository) AddReceipt(receipt models.Receipt) error {
    if _, exists := r.receipts[receipt.GetID()]; exists {
        return errors.New("receipt is already added")
    }
    r.receipts[receipt.GetID()] = receipt
    return nil
}

func (r ReceiptRepository) AddParkingSlotMapping(parkingSlotId int, receiptId string) error {
    if _, exists := r.parkingSlotMapping[parkingSlotId]; exists {
        return errors.New("parking slot is already added")
    }
    r.parkingSlotMapping[parkingSlotId] = receiptId
    return nil
}








