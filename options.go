package transcoder

// Options ...
type Options interface {
	GetStrArguments() []string
	GetGlobalStrArguments() []string
}
