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

func NewXYZone(x uint, y uint) (Zone, error) {
	if x > 23 || y > 5 {
		return nil, fmt.Errorf("%d,%d is an invalid xy coordinate. X must be less than 6, Y must be less than 24. See https://www.daskeyboard.io/q-zone-id-explanation#addressing-a-zone-as-a-2d-coordinate-xy-recommend", x, y)
	}

	return &xyZone{
		xCoord: x,
		yCoord: y,
	}, nil
}

func (xy *xyZone) GetZoneName() string {
	return fmt.Sprintf("%d,%d", xy.xCoord, xy.yCoord)
}

type linearZone struct {
	linearCoordinate uint
}

func NewLinearZone(linearCoordinate uint) (Zone, error) {
	if linearCoordinate < 24 || linearCoordinate > 168 {
		return nil, fmt.Errorf("%d is an invalid linear coordinate. Linear coordinates must be between 24 and 168. See https://www.daskeyboard.io/q-zone-id-explanation#addressing-a-zone-as-a-linear-coordinate", linearCoordinate)
	}

	return &linearZone{
		linearCoordinate: linearCoordinate,
	}, nil
}

func (l *linearZone) GetZoneName() string {
	return fmt.Sprintf("%d", l.linearCoordinate)
}
