package sqlmicroservice

import (
	"testing"
)

var dbMap = DatabaseMap{
	connections: NewSyncMapMock(),
}

func TestLoad(t *testing.T) {
	wantOk := false
	gotValue, gotOk := dbMap.Load("connectionName")
	// should return nil as a value if the connection is not found
	if nil != gotValue {
		t.Errorf("db.Load got %v, want %v", gotValue, nil)
	}
	// should return false as a flag is the key is not found
	if wantOk != gotOk {
		t.Errorf("db.Load 'ok' flag is %v, want %v", gotOk, wantOk)
	}
}
