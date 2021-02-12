package model

import "testing"

func TestRfData(t *testing.T) {
	datas := make(map[string]Model, 0)
	RfData("user", "username", datas)
}
