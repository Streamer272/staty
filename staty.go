package staty

type Store struct {
	state map[string]any
}

func NewStore() *Store {
	return &Store{
		state: map[string]any{},
	}
}

func (s *Store) Set(name string, value any) {
	s.state[name] = value
}

func (s *Store) Get(name string) any {
	return s.state[name]
}
