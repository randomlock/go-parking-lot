package models

import (
    "time"

    "github.com/pkg/errors"
)

type Status string

const (
    ACTIVE Status = "active"
    INACTIVE Status = "inactive"
)

type Receipt interface {
    SetPerHourCharge() float64
    GetID() string
    GetStatus() Status
}


type BaseReceipt struct {
    ID string
    Status Status
    Created time.Time
    PerHourCharge float64
    Number string
    Vehicle Vehicle
    ParkingSlot ParkingSlot
}


func (r BaseReceipt) GetStatus() Status {
    return r.Status
}

func (r BaseReceipt) GetID() string {
    return r.ID
}

func NewReceipt(vehicle Vehicle, parkingSlot ParkingSlot) (Receipt, error) {
    receipt := BaseReceipt{
        ID:          vehicle.PlateNumber + time.Now().String(),
        Status:      ACTIVE,
        Created:     time.Now(),
        Number:      vehicle.PlateNumber,
        Vehicle:   vehicle,
        ParkingSlot: parkingSlot,
    }
    if vehicle.Type == CAR {
        return BikeReceipt{receipt}, nil
    } else if vehicle.Type == BIKE {
        return CarReceipt{receipt}, nil
    } else {
        return nil, errors.New("invalid vehicle type. cannot generate receipt")
    }
}



func (r BaseReceipt) totalCharge() float64 {
    return r.PerHourCharge * ( time.Now().Sub(r.Created).Hours())
}

type BikeReceipt struct {
    BaseReceipt
}

func (b BikeReceipt) SetPerHourCharge() float64 {
    return 10
}


type CarReceipt struct {
    BaseReceipt
}

func (c CarReceipt) SetPerHourCharge() float64 {
    return 20
}


