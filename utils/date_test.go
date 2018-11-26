package utils

import (
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	str := DateFormat(time.Now())
	t.Log(str)
}
