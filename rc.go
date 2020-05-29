package main

type Hub struct {
	broadcast chan []byte
	reg       chan *Device
	unreg     chan *Device
	devices   map[*Device]bool
}

type Device struct {
	hub *Hub
	// conn net.Conn
	send chan []byte
	// Device local adress
	localAddr string
}

func (h *Hub) run(msg chan *Message) {
	msgStr := &Message{}
	for {
		select {
		case device := <-h.reg:
			h.devices[device] = true
			devList := []string{}
			for device := range h.devices {
				devList = append(devList, device.localAddr)

			}
			msgStr.DevicesLocAddr = devList
			msg <- msgStr

		case device := <-h.unreg:
			if _, ok := h.devices[device]; ok {
				delete(h.devices, device)
				// close(device.send)
			}
			devList := []string{}
			for device := range h.devices {
				devList = append(devList, device.localAddr)
			}
			msgStr.DevicesLocAddr = devList
			msg <- msgStr
		}
	}
}

func newHub() *Hub {
	return &Hub{
		broadcast: make(chan []byte),
		reg:       make(chan *Device),
		unreg:     make(chan *Device),
		devices:   make(map[*Device]bool),
	}
}
