package bimap

func NewBiMap[P comparable, Q comparable]() *BiMap[P, Q] {
	return &BiMap[P, Q]{
		forward: make(map[P]Q),
		reverse: make(map[Q]P),
	}
}

func MapToBimap[P comparable, Q comparable](m map[P]Q) *BiMap[P, Q] {
	bm := NewBiMap[P, Q]()

	for p, q := range m {
		bm.Set(p, q)
	}

	return bm
}

type BiMap[P comparable, Q comparable] struct {
	forward map[P]Q
	reverse map[Q]P
}

func (b *BiMap[P, Q]) GetFor(p P) Q {
	return b.forward[p]
}

func (b *BiMap[P, Q]) GetRev(q Q) P {
	return b.reverse[q]
}

func (b *BiMap[P, Q]) DelFor(p P) {
	v := b.forward[p]
	delete(b.reverse, v)
	delete(b.forward, p)
}

func (b *BiMap[P, Q]) DelRev(q Q) {
	v := b.reverse[q]
	delete(b.forward, v)
	delete(b.reverse, q)
}

func (b *BiMap[P, Q]) Set(p P, q Q) {
	if b.ExistsFor(p) {
		b.DelFor(p)
	}

	b.forward[p] = q
	b.reverse[q] = p
}

func (b *BiMap[P, Q]) ExistsFor(p P) bool {
	_, ok := b.forward[p]
	return ok
}

func (b *BiMap[P, Q]) ExistsRev(q Q) bool {
	_, ok := b.reverse[q]
	return ok
}
