package models

import "sync"

type ParkingStrategy interface {
    AddParkingSlot(slot int)
    RemoveParkingSlot(slot int)
    AssignParkingSlot(vehicle *Vehicle) *ParkingSlot
    UnAssignParkingSlot(parkingSpace *ParkingSlot)
    CanPark() bool
}

func CreateParkingStrategy(parkingStrategyType ParkingStrategyType) ParkingStrategy  {
    if parkingStrategyType == NaturallyOrdering {
        return NaturalOrderingStrategy{}
    }
    return nil
}


type ParkingSpace struct {
    parkingSlot *ParkingSlot
    prev *ParkingSpace
    next *ParkingSpace
}

type NaturalOrderingStrategy struct {
    mutex *sync.Mutex
    head *ParkingSpace
    tail *ParkingSpace
}

func (n NaturalOrderingStrategy) AddParkingSlot(slot int) {
    if n.head == nil {
        n.head = &ParkingSpace{
            parkingSlot: NewParkingSlot(slot),
        }
        n.tail = &ParkingSpace{
            parkingSlot: NewParkingSlot(slot),
        }
    } else {
        n.tail.prev = n.tail
        n.tail.next = &ParkingSpace{
            parkingSlot: NewParkingSlot(slot),
        }
        n.tail = n.tail.next
    }
}

func (n NaturalOrderingStrategy) RemoveParkingSlot(slot int) {
    panic("implement me")
}

func (n NaturalOrderingStrategy) AssignParkingSlot(vehicle *Vehicle) *ParkingSlot {
    n.mutex.Lock()
    defer n.mutex.Unlock()
    assigned := n.head
    n.head = n.head.next
    return assigned.parkingSlot
}

func (n NaturalOrderingStrategy) CanPark() bool {
    return n.head != nil
}

func (n NaturalOrderingStrategy) UnAssignParkingSlot(parkingSpace *ParkingSlot) {
    // TODO need to implement it efficiently
}
