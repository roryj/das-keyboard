package keyboard

import (
	"fmt"
	"strings"
)

// https://www.daskeyboard.io/q-zone-id-explanation/
type Zone interface {
	GetZoneName() string
}

type keyZone struct {
	keyId string
}

func NewKeyZone(key string) Zone {
	keyCode := strings.ToUpper(key)
	if !strings.HasPrefix(keyCode, "KEY_") {
		keyCode = fmt.Sprintf("KEY_%s", key)
	}
	return &keyZone{keyId: keyCode}
}

func (k *keyZone) GetZoneName() string {
	return k.keyId
}

type xyZone struct {
	xCoord uint
	yCoord uint
}

func NewXYZone(x uint, y uint) Zone {
	return &xyZone{
		xCoord: x,
		yCoord: y,
	}
}

func (xy *xyZone) GetZoneName() string {
	return fmt.Sprintf("%d,%d", xy.xCoord, xy.yCoord)
}

type linearZone struct {
	linearCoordinate uint
}

func NewLinearZone(linearCoordinate uint) Zone {
	return &linearZone{
		linearCoordinate: linearCoordinate,
	}
}

func (l *linearZone) GetZoneName() string {
	return fmt.Sprintf("%d", l.linearCoordinate)
}
