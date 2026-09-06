package token

type Token struct {
	Secret string `json:"secret"`

	Username string `json:"username"`
}

func (t *Token) Encode() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (t *Token) Decode(s string) error { _ = "STUB: not implemented"; return nil }
