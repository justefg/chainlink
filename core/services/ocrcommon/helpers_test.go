package ocrcommon

import ocrnetworking "github.com/justefg/libocr/networking"

func (p *SingletonPeerWrapper) PeerConfig() (ocrnetworking.PeerConfig, error) {
	return p.peerConfig()
}
