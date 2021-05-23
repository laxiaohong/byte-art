package str

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestShiftRotate(t *testing.T) {
	var data = []byte{'a','b','e','e','h'}
	ShiftRotate(data, 2)
	logrus.Info(string(data))
}


func TestShiftStrLeftMove(t *testing.T) {
	var data = []byte{'a','b','e','e','h','*','b'}
	ShiftStrLeftMove(data)
	logrus.Info(string(data))
}