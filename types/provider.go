package types

type Provider interface {
	IsRequestTrace(l RawLog) bool
	IsResponseTrace(l RawLog) bool
	ParseRequest(l RawLog) (*RequestTrace, error)
	ParseResponse(l RawLog) (*RequestTrace, error)
}
