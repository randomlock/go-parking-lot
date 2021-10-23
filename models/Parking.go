package models

import "fmt"

type Parking struct {
    ID int
    Name string
    Strategy map[VehicleType]ParkingStrategy
}

func (p Parking) Park(vehicle *Vehicle) (*ParkingSlot, error) {
    strategy, err := p.getStrategy(vehicle.Type)
    if err != nil {
        return nil, err
    }
    return strategy.AssignParkingSlot(vehicle), nil
}

func (p Parking) CanPark(vehicle *Vehicle) bool  {
    strategy, err := p.getStrategy(vehicle.Type)
    if err != nil {
        return false
    }
    return strategy.CanPark()
}

func (p Parking) UnPark(parkingSpot *ParkingSlot, vehicle *Vehicle) error {
    strategy, err := p.getStrategy(vehicle.Type)
    if err != nil {
        return err
    }
    strategy.UnAssignParkingSlot(parkingSpot)
    return nil
}


func (p Parking) getStrategy(vehicleType VehicleType) (ParkingStrategy, error) {
    if strategy, exists := p.Strategy[vehicleType]; exists {
        return strategy, nil
    }
    return nil, fmt.Errorf("this parking doesn't support vehicle type - %s", vehicleType)
}