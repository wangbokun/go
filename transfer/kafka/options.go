package kafka



// Options options
type Options struct {
	Brokers []string
	Group   string
	Topics  []string
	Offset  string
}

// Option Options function
type Option func(*Options)

// NewOptions new Options
func NewOptions(opts ...Option) Options {
	opt := Options{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Brokers set brokers
func Brokers(brokers ...string) Option {
	return func(o *Options) {
		o.Brokers = append(o.Brokers, brokers...)
	}
}

// Group set group
func Group(group string) Option {
	return func(o *Options) {
		o.Group = group
	}
}

// Topics set topics
func Topics(topics ...string) Option {
	return func(o *Options) {
		o.Topics = append(o.Topics, topics...)
	}
}

// Offset set offset
func Offset(offset string) Option {
	return func(o *Options) {
		o.Offset = offset
	}
}