package staty

type Store struct {
	state map[string]interface{}
}

func NewStore() *Store {
	return &Store{
		state: map[string]interface{}{},
	}
}

func (s *Store) Set(name string, value interface{}) {
	s.state[name] = value
}

func (s *Store) Get(name string) interface{} {
	return s.state[name]
}
