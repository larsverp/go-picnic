package picnic

import "encoding/json"

type DeliverySlot struct {
	Id                string `json:"slot_id"`
	HubId             string `json:"hub_id"`
	WindowStart       string `json:"window_start"`
	WindowEnd         string `json:"window_end"`
	CutOffTime        string `json:"cut_off_time"`
	IsAvailable       bool   `json:"is_available"`
	Selected          bool   `json:"selected"`
	Reserved          bool   `json:"reserved"`
	MinimumOrderValue int    `json:"minimum_order_value"`
}

type GetDeliverySlotsResponse struct {
	DeliverySlots []DeliverySlot `json:"delivery_slots"`
}

func (cl Client) GetDeliverySlots() ([]DeliverySlot, error) {
	deliveriesRaw, err := cl.get("/cart/delivery_slots")
	if err != nil {
		return nil, err
	}

	return toDeliveries(deliveriesRaw)
}

func (cl Client) GetCurrentDeliveries() ([]DeliverySlot, error) {
	deliveriesRaw, err := cl.get("/cart/delivery_slots/current", []byte(`["CURRENT"]`))
	if err != nil {
		return nil, err
	}
	//Did not test this result yet, since there where no active deliveries at the moment
	return toDeliveries(deliveriesRaw)
}

func toDeliveries(deliveriesRaw []byte) ([]DeliverySlot, error) {
	var response GetDeliverySlotsResponse

	err := json.Unmarshal(deliveriesRaw, &response)
	if err != nil {
		return nil, err
	}
	return response.DeliverySlots, nil
}
