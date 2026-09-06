package attributes

import (
	"iter"
)

type Attributes struct {
	parent     *Attributes
	key, value any
}

func New(key, value any) *Attributes { _ = "STUB: not implemented"; return nil }

func (a *Attributes) WithValue(key, value any) *Attributes { _ = "STUB: not implemented"; return nil }

func (a *Attributes) Value(key any) any { _ = "STUB: not implemented"; return *new(any) }

func (a *Attributes) Equal(o *Attributes) bool { _ = "STUB: not implemented"; return false }

func (a *Attributes) String() string { _ = "STUB: not implemented"; return "" }

func str(x any) (s string) { _ = "STUB: not implemented"; return "" }

func (a *Attributes) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *Attributes) all() iter.Seq2[any, any] { _ = "STUB: not implemented"; return nil }
