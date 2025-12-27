package store

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MemoryStore struct {
	contacts []Contact
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: []Contact{
			{ID: 1, Name: "Sanjeev", Email: "sanjeev@example.com"},
			{ID: 2, Name: "Alex", Email: "alex@example.com"},
		},
	}
}

func (s *MemoryStore) ListContacts() []Contact {
	out := make([]Contact, len(s.contacts))
	copy(out, s.contacts)
	return out
}
