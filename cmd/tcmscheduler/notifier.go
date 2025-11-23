package main

import (
	"github.com/ekkx/tcm-platform/internal/platform/ulid"
	"github.com/ekkx/tcm-platform/internal/platform/ymd"
)

type NotifyMessage struct {
	JobID         ulid.ULID
	ReservationID ulid.ULID
	MasterUserID  ulid.ULID
	YMD           ymd.YMD
	Error         error
}

// エラーなどを外部サービスなどに通知するインターフェース
type Notifier interface {
	Notify(msg NotifyMessage) error
}

type NotifierImpl struct{}

func (n *NotifierImpl) Notify(msg NotifyMessage) error {
	return nil
}
