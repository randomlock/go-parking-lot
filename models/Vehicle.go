package models

type VehicleBuilder interface {
    Build() *Vehicle
    SetPlateNumber(plateNumber string) VehicleBuilder
    SetMake(make string) VehicleBuilder
    SetModel(model string) VehicleBuilder

}

type Vehicle struct {
    PlateNumber string
    Make  string
    Model string
    Type VehicleType
}

func NewVehicle() *Vehicle {
    return &Vehicle{}
}

func (v *Vehicle) SetPlateNumber(plateNumber string) *Vehicle {
    v.PlateNumber = plateNumber
    return v
}

func (v *Vehicle) SetMake(plateNumber string) *Vehicle  {
    v.Make = plateNumber
    return v
}

func (v *Vehicle) SetModel(plateNumber string) *Vehicle  {
    v.PlateNumber = plateNumber
    return v
}

type Car struct {
    Vehicle
}

func (c *Car) SetPlateNumber(plateNumber string) VehicleBuilder {
    panic("implement me")
}

func (c *Car) SetMake(make string) VehicleBuilder {
    panic("implement me")
}

func (c *Car) SetModel(model string) VehicleBuilder {
    panic("implement me")
}

func (c *Car) Build() *Vehicle {
    return &c.Vehicle
}



type Bike struct {
    Vehicle
}

func (b *Bike) SetPlateNumber(plateNumber string) VehicleBuilder {
    panic("implement me")
}

func (b *Bike) SetMake(make string) VehicleBuilder {
    panic("implement me")
}

func (b *Bike) SetModel(model string) VehicleBuilder {
    panic("implement me")
}

func (b *Bike) Build() *Vehicle {
    return &b.Vehicle
}

func GetVehicleBuilder(vehicleType VehicleType) VehicleBuilder  {

    if vehicleType == CAR {
        return &Car{}
    } else if vehicleType == BIKE {
        return &Bike{}
    } else {
        return nil
    }

}