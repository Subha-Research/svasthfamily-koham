package test_services

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreateAccessRelationship(t *testing.T) {
	mock_ctrl := gomock.NewController(t)
	defer mock_ctrl.Finish()

	// mock_acl_model := models_mock.NewMockACLModel(mock_ctrl)
	// test_acl_service := &services.ACLService{}
	// test_acl_service.CreateSFRelationship()
}

// func TestFoo(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	m := Foo(ctrl)

// 	// Does not make any assertions. Executes the anonymous functions and returns
// 	// its result when Bar is invoked with 99.
// 	m.
// 		EXPECT().
// 		Bar(gomock.Eq(99)).
// 		DoAndReturn(func(_ int) int {
// 			time.Sleep(1 * time.Second)
// 			return 101
// 		}).
// 		AnyTimes()

// 	// Does not make any assertions. Returns 103 when Bar is invoked with 101.
// 	m.
// 		EXPECT().
// 		Bar(gomock.Eq(101)).
// 		Return(103).
// 		AnyTimes()

// 	SUT(m)
// }

// func NewMockFoo(ctrl *gomock.Controller) {
// 	panic("unimplemented")
// }
