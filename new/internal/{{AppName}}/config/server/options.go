package server

type Options struct {
	Addr   string `json:"addr" yaml:"addr"`
	Locale string `json:"locale" yaml:"locale"`
}

func NewServerOptions() *Options {
	return &Options{
		Addr:   "",
		Locale: "zh",
	}
}
