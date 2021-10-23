package parkingLot

import (
    "fmt"

    "../parkingLot/controller"
    "../parkingLot/models"
)

func main() {
    vehicleBuilder := models.GetVehicleBuilder(models.CAR)
    vB1 := vehicleBuilder.SetPlateNumber("11111").SetMake("toyota").SetModel("corolla").Build()
    vB2 := vehicleBuilder.SetPlateNumber("22222").SetMake("maruti").SetModel("suzuki").Build()

    parkingController := controller.ParkingController{}
    parking := parkingController.CreateNaturalOrderParking("my parking", 10, 20)
    slot, err := parkingController.ParkVehicle(parking.ID, *vB1)
    if err != nil {
        fmt.Printf("Error in parking vehicle - %s due to - %s", vB1.PlateNumber, err.Error())
    } else {
        fmt.Printf("vehicle - %s succesfully parked with parking ID - %d", vB1.PlateNumber, slot.ID)
    }
    parkingController.ParkVehicle(parking.ID, *vB2)
    parkingController.UnParkVehicle(*vB1)
}
