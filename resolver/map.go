package resolver

import (
	"iter"
)

type addressMapEntry[T any] struct {
	addr  Address
	value T
}

type AddressMap = AddressMapV2[any]

type AddressMapV2[T any] struct {
	m map[Address]addressMapEntryList[T]
}

func toMapKey(addr *Address) Address { _ = "STUB: not implemented"; return *new(Address) }

type addressMapEntryList[T any] []*addressMapEntry[T]

func NewAddressMap() *AddressMap { _ = "STUB: not implemented"; return nil }

func NewAddressMapV2[T any]() *AddressMapV2[T] { _ = "STUB: not implemented"; return nil }

func (l addressMapEntryList[T]) find(addr Address) int { _ = "STUB: not implemented"; return 0 }

func (a *AddressMapV2[T]) Get(addr Address) (value T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (a *AddressMapV2[T]) Set(addr Address, value T) { _ = "STUB: not implemented"; return }

func (a *AddressMapV2[T]) Delete(addr Address) { _ = "STUB: not implemented"; return }

func (a *AddressMapV2[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (a *AddressMapV2[T]) Keys() []Address { _ = "STUB: not implemented"; return nil }

func (a *AddressMapV2[T]) Values() []T { _ = "STUB: not implemented"; return nil }

func (a *AddressMapV2[T]) All() iter.Seq2[Address, T] { _ = "STUB: not implemented"; return nil }

type endpointMapKey string

type EndpointMap[T any] struct {
	endpoints map[endpointMapKey]endpointData[T]
}

type endpointData[T any] struct {
	decodedKey Endpoint
	value      T
}

func NewEndpointMap[T any]() *EndpointMap[T] { _ = "STUB: not implemented"; return nil }

func encodeEndpoint(e Endpoint) endpointMapKey {
	_ = "STUB: not implemented"
	return *new(endpointMapKey)
}

func (em *EndpointMap[T]) Get(e Endpoint) (value T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (em *EndpointMap[T]) Set(e Endpoint, value T) { _ = "STUB: not implemented"; return }

func (em *EndpointMap[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (em *EndpointMap[T]) Keys() []Endpoint { _ = "STUB: not implemented"; return nil }

func (em *EndpointMap[T]) Values() []T { _ = "STUB: not implemented"; return nil }

func (em *EndpointMap[T]) All() iter.Seq2[Endpoint, T] { _ = "STUB: not implemented"; return nil }

func (em *EndpointMap[T]) Delete(e Endpoint) { _ = "STUB: not implemented"; return }
