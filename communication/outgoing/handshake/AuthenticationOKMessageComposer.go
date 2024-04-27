package handshake

import (
	"github.com/alexmgriffiths/habbo-go/utils"
)

type AuthenticationOKMessageComposer struct{}

func NewAuthenticationOKMessageComposer() []byte {
	packet := *utils.NewByteBuf([]byte{})
	packet.WriteShort(2491)
	return packet.Wrap()
}
