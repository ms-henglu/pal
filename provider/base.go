package provider

import (
	"github.com/ms-henglu/pal/rawlog"
	"github.com/ms-henglu/pal/types"
)

type Provider interface {
	IsRequestTrace(l rawlog.RawLog) bool
	IsResponseTrace(l rawlog.RawLog) bool
	ParseRequest(l rawlog.RawLog) (*types.RequestTrace, error)
	ParseResponse(l rawlog.RawLog) (*types.RequestTrace, error)
}
