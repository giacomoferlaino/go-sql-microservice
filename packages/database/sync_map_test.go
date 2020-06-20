package database

import (
	"database/sql"
	"encoding/json"
	"testing"
)

var syncMapMock = NewSyncMapMock()

var dbMap = SyncMap{
	connections: syncMapMock,
}

func TestLoad(t *testing.T) {
	wantOk := false
	gotValue, gotOk := dbMap.Load("connectionName")
	// should return nil as a value if the connection is not found
	if nil != gotValue {
		t.Errorf("dbMap.Load got %v, want %v", gotValue, nil)
	}
	// should return false as a flag is the key is not found
	if wantOk != gotOk {
		t.Errorf("dbMap.Load 'ok' flag is %v, want %v", gotOk, wantOk)
	}
}

func TestMarshalJSON(t *testing.T) {
	syncMapMock.rangeMock = func(f func(key interface{}, value interface{}) bool) {
		f("mockedConnection", &sql.DB{})
	}
	// should return the a byte stream containing the db connections in JSON format
	want, _ := json.Marshal(map[string]sql.DBStats{"mockedConnection": {}})
	got, _ := dbMap.MarshalJSON()
	if string(want) != string(got) {
		t.Errorf("dbMap.MarshalJSON got %v, want %v", got, want)
	}
}
