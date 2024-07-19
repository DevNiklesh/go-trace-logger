package logger

type MultiTransporter struct {
	transporters []Transporter
}

func NewMultiTransporter(transporters ...Transporter) *MultiTransporter {
	return &MultiTransporter{transporters: transporters}
}

func (t *MultiTransporter) Transport(level Level, message string) error {
	for _, tt := range t.transporters {
		tt.Transport(level, message)
	}
	return nil
}
