package constant

const (
	ShipDelivery  = 1
	TruckDelivery = 2
	Unknown       = 0
	UnknownText   = "UNKNOWN"
)

var DeliveryType = map[string]uint8{
	"UNKNOWN": Unknown,
	"SHIP":    ShipDelivery,
	"TRUCK":   TruckDelivery,
}

func GetDelivery(key string) uint8 {
	if item, ok := DeliveryType[key]; ok {
		return item
	}

	return DeliveryType[UnknownText]
}
