// models/models.go
package models

// CallInfo представляет информацию о вызовах
type CallInfo struct {
	CallerNumber  string
	Dst           string
	CallerName    string
	Duration      string
	CallTime      string
	CallStatus    string
	WaitTime      string
	Callrecording string
	// Добавьте другие поля по необходимости
}
