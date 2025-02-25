package User

import "errors"

type MockService struct {
	Users map[uint64]*User
}

func (m *MockService) Add(user *User) (uint64, error) {
	newID := uint64(len(m.Users) + 1)
	user.ID = newID
	m.Users[newID] = user
	return newID, nil
}

func (m *MockService) Read(id *uint64) (*User, error) {
	if user, ok := m.Users[*id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (m *MockService) Update(user *User, id *uint64) error {
	if _, ok := m.Users[*id]; ok {
		user.ID = *id
		m.Users[*id] = user
		return nil
	} else {
		return errors.New("user not found")
	}
}

func (m *MockService) Delete(userID *uint64) error {
	if _, ok := m.Users[*userID]; ok {
		delete(m.Users, *userID)
		return nil
	} else {
		return errors.New("user not found")
	}
}

func GetMockService() *MockService {
	return &MockService{Users: make(map[uint64]*User)}
}
