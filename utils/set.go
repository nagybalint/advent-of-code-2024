package utils

import "reflect"

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(elem T) bool {
	if _, ok := s[elem]; ok {
		return true
	}
	s[elem] = struct{}{}
	return false
}

func (s Set[T]) AddAll(elem []T) bool {
	changed := false
	for _, e := range elem {
		if _, ok := s[e]; ok {
			continue
		}
		s.Add(e)
		changed = true
	}
	return changed
}

func (s Set[T]) Clear() {
	for e := range s {
		delete(s, e)
	}
}

func (s Set[T]) Contains(elem T) bool {
	_, ok := s[elem]
	return ok
}

func (s Set[T]) ContainsAll(elems []T) bool {
	for _, e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

func (s Set[T]) Equals(other Set[T]) bool {
	return reflect.DeepEqual(s, other)
}

func (s Set[T]) ToSlice() []T {
	slice := make([]T, len(s))
	for e := range s {
		slice = append(slice, e)
	}
	return slice
}
