// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pborzenkov/tg-bot-transmission/pkg/bot (interfaces: Telegram,Transmission)

// Package bot is a generated GoMock package.
package bot

import (
	context "context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	gomock "github.com/golang/mock/gomock"
	transmission "github.com/pborzenkov/go-transmission/transmission"
	url "net/url"
	reflect "reflect"
)

// MockTelegram is a mock of Telegram interface
type MockTelegram struct {
	ctrl     *gomock.Controller
	recorder *MockTelegramMockRecorder
}

// MockTelegramMockRecorder is the mock recorder for MockTelegram
type MockTelegramMockRecorder struct {
	mock *MockTelegram
}

// NewMockTelegram creates a new mock instance
func NewMockTelegram(ctrl *gomock.Controller) *MockTelegram {
	mock := &MockTelegram{ctrl: ctrl}
	mock.recorder = &MockTelegramMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTelegram) EXPECT() *MockTelegramMockRecorder {
	return m.recorder
}

// AnswerCallbackQuery mocks base method
func (m *MockTelegram) AnswerCallbackQuery(arg0 tgbotapi.CallbackConfig) (tgbotapi.APIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnswerCallbackQuery", arg0)
	ret0, _ := ret[0].(tgbotapi.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnswerCallbackQuery indicates an expected call of AnswerCallbackQuery
func (mr *MockTelegramMockRecorder) AnswerCallbackQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnswerCallbackQuery", reflect.TypeOf((*MockTelegram)(nil).AnswerCallbackQuery), arg0)
}

// GetFileDirectURL mocks base method
func (m *MockTelegram) GetFileDirectURL(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileDirectURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileDirectURL indicates an expected call of GetFileDirectURL
func (mr *MockTelegramMockRecorder) GetFileDirectURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileDirectURL", reflect.TypeOf((*MockTelegram)(nil).GetFileDirectURL), arg0)
}

// GetUpdates mocks base method
func (m *MockTelegram) GetUpdates(arg0 tgbotapi.UpdateConfig) ([]tgbotapi.Update, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdates", arg0)
	ret0, _ := ret[0].([]tgbotapi.Update)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpdates indicates an expected call of GetUpdates
func (mr *MockTelegramMockRecorder) GetUpdates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdates", reflect.TypeOf((*MockTelegram)(nil).GetUpdates), arg0)
}

// MakeRequest mocks base method
func (m *MockTelegram) MakeRequest(arg0 string, arg1 url.Values) (tgbotapi.APIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRequest", arg0, arg1)
	ret0, _ := ret[0].(tgbotapi.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeRequest indicates an expected call of MakeRequest
func (mr *MockTelegramMockRecorder) MakeRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRequest", reflect.TypeOf((*MockTelegram)(nil).MakeRequest), arg0, arg1)
}

// Send mocks base method
func (m *MockTelegram) Send(arg0 tgbotapi.Chattable) (tgbotapi.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(tgbotapi.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockTelegramMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTelegram)(nil).Send), arg0)
}

// MockTransmission is a mock of Transmission interface
type MockTransmission struct {
	ctrl     *gomock.Controller
	recorder *MockTransmissionMockRecorder
}

// MockTransmissionMockRecorder is the mock recorder for MockTransmission
type MockTransmissionMockRecorder struct {
	mock *MockTransmission
}

// NewMockTransmission creates a new mock instance
func NewMockTransmission(ctrl *gomock.Controller) *MockTransmission {
	mock := &MockTransmission{ctrl: ctrl}
	mock.recorder = &MockTransmissionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransmission) EXPECT() *MockTransmissionMockRecorder {
	return m.recorder
}

// AddTorrent mocks base method
func (m *MockTransmission) AddTorrent(arg0 context.Context, arg1 *transmission.AddTorrentReq) (*transmission.NewTorrent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTorrent", arg0, arg1)
	ret0, _ := ret[0].(*transmission.NewTorrent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTorrent indicates an expected call of AddTorrent
func (mr *MockTransmissionMockRecorder) AddTorrent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTorrent", reflect.TypeOf((*MockTransmission)(nil).AddTorrent), arg0, arg1)
}

// GetSession mocks base method
func (m *MockTransmission) GetSession(arg0 context.Context, arg1 ...transmission.SessionField) (*transmission.Session, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSession", varargs...)
	ret0, _ := ret[0].(*transmission.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession
func (mr *MockTransmissionMockRecorder) GetSession(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockTransmission)(nil).GetSession), varargs...)
}

// GetSessionStats mocks base method
func (m *MockTransmission) GetSessionStats(arg0 context.Context) (*transmission.SessionStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionStats", arg0)
	ret0, _ := ret[0].(*transmission.SessionStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionStats indicates an expected call of GetSessionStats
func (mr *MockTransmissionMockRecorder) GetSessionStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionStats", reflect.TypeOf((*MockTransmission)(nil).GetSessionStats), arg0)
}

// GetTorrents mocks base method
func (m *MockTransmission) GetTorrents(arg0 context.Context, arg1 transmission.Identifier, arg2 ...transmission.TorrentField) ([]*transmission.Torrent, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTorrents", varargs...)
	ret0, _ := ret[0].([]*transmission.Torrent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTorrents indicates an expected call of GetTorrents
func (mr *MockTransmissionMockRecorder) GetTorrents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTorrents", reflect.TypeOf((*MockTransmission)(nil).GetTorrents), varargs...)
}

// IsPortOpen mocks base method
func (m *MockTransmission) IsPortOpen(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPortOpen", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsPortOpen indicates an expected call of IsPortOpen
func (mr *MockTransmissionMockRecorder) IsPortOpen(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPortOpen", reflect.TypeOf((*MockTransmission)(nil).IsPortOpen), arg0)
}

// RemoveTorrents mocks base method
func (m *MockTransmission) RemoveTorrents(arg0 context.Context, arg1 transmission.Identifier, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTorrents", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTorrents indicates an expected call of RemoveTorrents
func (mr *MockTransmissionMockRecorder) RemoveTorrents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTorrents", reflect.TypeOf((*MockTransmission)(nil).RemoveTorrents), arg0, arg1, arg2)
}

// SetSession mocks base method
func (m *MockTransmission) SetSession(arg0 context.Context, arg1 *transmission.SetSessionReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSession indicates an expected call of SetSession
func (mr *MockTransmissionMockRecorder) SetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSession", reflect.TypeOf((*MockTransmission)(nil).SetSession), arg0, arg1)
}

// StartTorrents mocks base method
func (m *MockTransmission) StartTorrents(arg0 context.Context, arg1 transmission.Identifier) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTorrents", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartTorrents indicates an expected call of StartTorrents
func (mr *MockTransmissionMockRecorder) StartTorrents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTorrents", reflect.TypeOf((*MockTransmission)(nil).StartTorrents), arg0, arg1)
}

// StopTorrents mocks base method
func (m *MockTransmission) StopTorrents(arg0 context.Context, arg1 transmission.Identifier) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopTorrents", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopTorrents indicates an expected call of StopTorrents
func (mr *MockTransmissionMockRecorder) StopTorrents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopTorrents", reflect.TypeOf((*MockTransmission)(nil).StopTorrents), arg0, arg1)
}
