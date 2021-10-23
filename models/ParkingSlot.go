package models

type ParkingSlot struct {
    ID int
}

func NewParkingSlot(ID int) *ParkingSlot {
    return &ParkingSlot{ID: ID}
}




