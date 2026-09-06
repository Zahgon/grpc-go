package stats

type payloadCurveRange struct {
	from, to int32
	weight   float64
}

func newPayloadCurveRange(line []string) (*payloadCurveRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pcr *payloadCurveRange) chooseRandom() int { _ = "STUB: not implemented"; return 0 }

func sha256file(file string) (string, error) { _ = "STUB: not implemented"; return "", nil }

type PayloadCurve struct {
	pcrs []*payloadCurveRange

	Sha256 string
}

func NewPayloadCurve(file string) (*PayloadCurve, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *PayloadCurve) ChooseRandom() int { _ = "STUB: not implemented"; return 0 }

func (pc *PayloadCurve) Hash() string { _ = "STUB: not implemented"; return "" }

func (pc *PayloadCurve) ShortHash() string { _ = "STUB: not implemented"; return "" }
