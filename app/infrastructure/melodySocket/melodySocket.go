package melodySocket

import (
	"go.uber.org/zap"
	"gopkg.in/olahol/melody.v1"
)

var melodySocket *melody.Melody
var err error

func Init() {
	melodySocket = melody.New()
	zap.S().Infow("Melody initialized")
}

func GetMelody() *melody.Melody {
	return melodySocket
}

func CloseMelody() {
	melodySocket.Close()
}
