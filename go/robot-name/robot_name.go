package robotname

import (
	"time"
	"math/rand"
	"bytes"
)

const testVersion = 1

type Robot struct {
	name string
}

func (robot *Robot) Name() string {
	if robot.name == "" {
		rand.Seed(time.Now().UnixNano())
		var buf bytes.Buffer
		buf.WriteByte(byte(rand.Intn(26) + 'A'))
		buf.WriteByte(byte(rand.Intn(26) + 'A'))
		buf.WriteByte(byte(rand.Intn(10) + '0'))
		buf.WriteByte(byte(rand.Intn(10) + '0'))
		buf.WriteByte(byte(rand.Intn(10) + '0'))
		robot.name = buf.String()
	}
	return robot.name
}


func (robot *Robot) Reset() {
	robot.name = ""
}