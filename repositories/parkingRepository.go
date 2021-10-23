package repositories

import (
    "fmt"
    "math/rand"

    "../models"
)


type ParkingRepository struct {
    parkings map[int]models.Parking
    parkingVehicleMapping map[int]*models.Vehicle
}

func GetParkingRepository() ParkingRepository {
    return ParkingRepository{}
}

func (parkingRepository ParkingRepository) CreateNaturalOrderingParking(name string, bikeSlot int, carSlot int) models.Parking {

    bikeParkingStrategy := models.CreateParkingStrategy(models.NaturallyOrdering)
    for i := 0; i < bikeSlot ; i++ {
        bikeParkingStrategy.AddParkingSlot(bikeSlot)
    }
    carParkingStrategy := models.CreateParkingStrategy(models.NaturallyOrdering)
    for i := 0; i < bikeSlot ; i++ {
        carParkingStrategy.AddParkingSlot(bikeSlot)
    }

    parking := models.Parking{
        ID:   rand.Int(),
        Name: name,
        Strategy: map[models.VehicleType]models.ParkingStrategy{
            models.BIKE: bikeParkingStrategy,
            models.CAR: carParkingStrategy,
        },
    }
    parkingRepository.parkings[parking.ID] = parking
    return parking
}

func (parkingRepository ParkingRepository) ParkVehicle(parkingId int, vehicle *models.Vehicle) (parkingSlot *models.ParkingSlot, err error)  {
    if _, exists := parkingRepository.parkings[parkingId]; !exists {
        return nil, fmt.Errorf("parking not found")
    }
    parking := parkingRepository.parkings[parkingId]

    if !parking.CanPark(vehicle) {
        return nil, fmt.Errorf("cannot park vehicle")
    }

    parkingSlot, err = parking.Park(vehicle)
    if err != nil {
        return nil, fmt.Errorf("error in parking vehicle due to - %s", err.Error())
    }

    parkingRepository.parkingVehicleMapping[parkingSlot.ID] = vehicle
    return
}

func (parkingRepository ParkingRepository) GetVehicleParked(parkingSlotID int) (vehicle *models.Vehicle, exists bool) {
    if vehicle, exists = parkingRepository.parkingVehicleMapping[parkingSlotID]; !exists {
        return nil, false
    }
    return
}











