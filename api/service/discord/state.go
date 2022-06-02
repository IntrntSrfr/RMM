package discord

/*
import (
	"errors"
	"sync"
)

type State struct {
	sync.RWMutex
	guildJoinsMap map[string][]*GuildMember
}

func NewState() *State {
	return &State{
		guildJoinsMap: make(map[string][]*GuildMember),
	}
}

var (
	ErrStateNotFound = errors.New("not found in state")
)

func (s *State) MemberList(guildID string) ([]*GuildMember, error) {
	s.RLock()
	defer s.RUnlock()
	members, ok := s.guildJoinsMap[guildID]
	if !ok {
		return nil, ErrStateNotFound
	}
	return members, nil
}

func (s *State) createMemberList(guildID string) {
	s.Lock()
	defer s.Unlock()
	s.guildJoinsMap[guildID] = []*GuildMember{}
}

func (s *State) MemberAdd(m *GuildMember) error {
	s.Lock()
	defer s.Unlock()
	members, ok := s.guildJoinsMap[m.GuildID]
	if !ok {
		return ErrStateNotFound
	}
	s.guildJoinsMap[m.GuildID] = append(members, m)
	return nil
}
*/
