package database

// NewSyncMapMock allocates and return a new syncMapMock
func NewSyncMapMock() *SyncMapMock {
	loadMock := func(key interface{}) (interface{}, bool) {
		return nil, false
	}
	storeMock := func(key interface{}, value interface{}) {}
	deleteMock := func(key interface{}) {}
	rangeMock := func(f func(key interface{}, value interface{}) bool) {}
	return &SyncMapMock{
		loadMock:   loadMock,
		storeMock:  storeMock,
		deleteMock: deleteMock,
		rangeMock:  rangeMock,
	}
}

// SyncMapMock is used to mock sync.Map behaviours
type SyncMapMock struct {
	loadMock   func(key interface{}) (interface{}, bool)
	storeMock  func(key interface{}, value interface{})
	deleteMock func(key interface{})
	rangeMock  func(f func(key interface{}, value interface{}) bool)
}

// Delete mocks sync.Map.Delete
func (mock *SyncMapMock) Delete(key interface{}) {
	mock.deleteMock(key)
}

// Load mocks sync.Map.Load
func (mock *SyncMapMock) Load(key interface{}) (interface{}, bool) {
	return mock.loadMock(key)
}

// Store mocks sync.Map.Store
func (mock *SyncMapMock) Store(key interface{}, value interface{}) {
	mock.storeMock(key, value)
}

// Range mocks sync.Map.Range
func (mock *SyncMapMock) Range(f func(key interface{}, value interface{}) bool) {
	mock.rangeMock(f)
}
