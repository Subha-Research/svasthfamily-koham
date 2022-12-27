package models_mock

import "github.com/golang/mock/gomock"

type MockACLModel struct {
	ctrl *gomock.Controller
}

func NewMockACLModel(ctrl *gomock.Controller) *MockACLModel {
	mock := &MockACLModel{ctrl: ctrl}
	return mock
}
