package services_mock

import "github.com/golang/mock/gomock"

type MockACLService struct {
	ctrl *gomock.Controller
}

func NewMockACLService(ctrl *gomock.Controller) *MockACLService {
	mock := &MockACLService{ctrl: ctrl}
	return mock
}
