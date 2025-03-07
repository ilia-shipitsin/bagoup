// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tagatac/bagoup/opsys/pdfgen (interfaces: PDFGenerator)

// Package mock_pdfgen is a generated GoMock package.
package mock_pdfgen

import (
	reflect "reflect"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	gomock "github.com/golang/mock/gomock"
)

// MockPDFGenerator is a mock of PDFGenerator interface.
type MockPDFGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockPDFGeneratorMockRecorder
}

// MockPDFGeneratorMockRecorder is the mock recorder for MockPDFGenerator.
type MockPDFGeneratorMockRecorder struct {
	mock *MockPDFGenerator
}

// NewMockPDFGenerator creates a new mock instance.
func NewMockPDFGenerator(ctrl *gomock.Controller) *MockPDFGenerator {
	mock := &MockPDFGenerator{ctrl: ctrl}
	mock.recorder = &MockPDFGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPDFGenerator) EXPECT() *MockPDFGeneratorMockRecorder {
	return m.recorder
}

// AddPage mocks base method.
func (m *MockPDFGenerator) AddPage(arg0 wkhtmltopdf.PageProvider) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddPage", arg0)
}

// AddPage indicates an expected call of AddPage.
func (mr *MockPDFGeneratorMockRecorder) AddPage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPage", reflect.TypeOf((*MockPDFGenerator)(nil).AddPage), arg0)
}

// Create mocks base method.
func (m *MockPDFGenerator) Create() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPDFGeneratorMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPDFGenerator)(nil).Create))
}
