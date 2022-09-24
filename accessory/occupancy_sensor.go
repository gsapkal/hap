package accessory

import (
	"github.com/gsapkal/hap/service"
)


type OccupancySensor struct {
	*A
	OccupancySensor *service.OccupancySensor
}

func NewOccupancySensor(info Info) *OccupancySensor {
	a := OccupancySensor{}
	a.A = New(info, TypeOccupancySensor)
	a.A.Id = 1
	a.OccupancySensor = service.NewOccupancySensor()
	a.AddS(a.OccupancySensor.S)

	return &a
}
