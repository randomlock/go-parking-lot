package controller

import (
    "fmt"

    "../models"
    "../repositories"
)

type ParkingController struct {
}

func (pc ParkingController) CreateNaturalOrderParking(parkingName string, bikeSlot int, carSlot int) models.Parking {
    parkingRepository := repositories.GetParkingRepository()
    return parkingRepository.CreateNaturalOrderingParking(parkingName, bikeSlot, carSlot)
}

func (pc ParkingController) ParkVehicle(parkingId int, vehicle models.Vehicle) (*models.ParkingSlot, error) {
    parkingRepository := repositories.GetParkingRepository()
    receiptRepository := repositories.GetReceiptRepository()

    parkingSpot, err := parkingRepository.ParkVehicle(parkingId, &vehicle)
    if err != nil {
        return nil, err
    }

    receipt, err := receiptRepository.CreateReceipt(vehicle, *parkingSpot)

    if err != nil {
        return nil, err
    }

    fmt.Printf("receipt generated succesfully - %s", receipt.GetID())
    return parkingSpot, nil
}

func (pc ParkingController) UnParkVehicle(vehicle models.Vehicle)  {
    panic("implement me")
}