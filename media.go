package webrtc

import (
	"github.com/fangelod/webrtc/pkg/media"
	"github.com/fangelod/webrtc/pkg/rtcp"
	"github.com/fangelod/webrtc/pkg/rtp"
)

// RTCTrack represents a track that is communicated
type RTCTrack struct {
	ID          string
	PayloadType uint8
	Kind        RTCRtpCodecType
	Label       string
	Ssrc        uint32
	Codec       *RTCRtpCodec
	Packets     <-chan *rtp.Packet
	RTCPPackets <-chan rtcp.Packet
	Samples     chan<- media.RTCSample
	RawRTP      chan<- *rtp.Packet
}
