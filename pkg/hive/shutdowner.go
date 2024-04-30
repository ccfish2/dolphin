package hive

type Shutdowner interface {
	Shutdown(...ShutdownOption)
}

type ShutdownOption interface {
	apply(*shutdownOptions)
}

type shutdownOptions struct {
	err error
}

type optionFunc func(*shutdownOptions)

func (fn optionFunc) apply(opts *shutdownOptions) { fn(opts) }
