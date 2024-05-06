package rtc

import (
	"github.com/pion/webrtc/v4"
)

func StartRTCManager() (*RTCManager, error) {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	return &RTCManager{config: config}, nil
}

type RTCManager struct {
	config webrtc.Configuration
}

func (r RTCManager) AddConnection() {
}
