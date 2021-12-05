package set

type exists struct{}
type Set struct {
	M map[interface{}]exists
}
func NewSet(items ...interface{}) *Set {
	s := &Set{}
	s.M = make(map[interface{}]exists)
	s.Add(items...)
	return s
}
func (s *Set) Add(items ...interface{}) {
	for _, item := range items {
		s.M[item] = exists{}
	}
}
func (s *Set) Remove(item interface{}) {
	delete(s.M, item)
}
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.M[item]
	return ok
}
func (s *Set) Size() int {
	return len(s.M)
}
type ISet struct {
	M map[int]exists
}
func NewSetInt(items ...int) *ISet {
	s := &ISet{}
	s.M = make(map[int]exists)
	s.Add(items...)
	return s
}
func (s *ISet) Add(items ...int) {
	for _, item := range items {
		s.M[item] = exists{}
	}
}
func (s *ISet) Remove(item int) {
	delete(s.M, item)
}
func (s *ISet) Contains(item int) bool {
	_, ok := s.M[item]
	return ok
}
func (s *ISet) Size() int {
	return len(s.M)
}