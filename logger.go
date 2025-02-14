package filelog

type Logger interface {
	Println(v ...any)
}

type Prefix uint

const (
	DEBUG = "DEBUG:"
	INFO  = "INFO:"
	WARN  = "WARN:"
	ERROR = "ERROR:"
)
