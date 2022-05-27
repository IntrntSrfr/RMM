package discord

import (
	"errors"
	"fmt"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type DiscordService struct {
	sessions []*discordgo.Session
	s        *discordgo.Session
	state    *State
}

var (
	ErrStateNotFound = errors.New("not found in state")
)

type State struct {
	sync.RWMutex
	memberMap map[string][]*GuildMember
}

func NewState() *State {
	return &State{
		memberMap: make(map[string][]*GuildMember),
	}
}

func (s *State) MemberList(guildID string) ([]*GuildMember, error) {
	s.RLock()
	defer s.RUnlock()
	members, ok := s.memberMap[guildID]
	if !ok {
		return nil, ErrStateNotFound
	}
	return members, nil
}

func (s *State) createMemberList(guildID string) {
	s.Lock()
	defer s.Unlock()
	s.memberMap[guildID] = []*GuildMember{}
}

func (s *State) memberAdd(m *GuildMember) error {
	members, ok := s.memberMap[m.GuildID]
	if !ok {
		return ErrStateNotFound
	}
	members = append(members, m)
	return nil
}

func (s *State) MemberAdd(m *GuildMember) error {
	s.Lock()
	defer s.Unlock()
	return s.memberAdd(m)
}

type GuildMember struct {
	GuildID string
	Member  *discordgo.Member
}

const shardCount = 1

func NewDiscordService(token string) (*DiscordService, error) {
	d := &DiscordService{
		state: NewState(),
	}
	d.sessions = make([]*discordgo.Session, shardCount)
	for i := 0; i < shardCount; i++ {
		s, err := discordgo.New("Bot " + token)
		if err != nil {
			return nil, err
		}

		s.ShardCount = shardCount
		s.ShardID = i
		s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers | discordgo.IntentMessageContent)

		s.AddHandler(d.memberAdd)
		s.AddHandler(d.memberRemove)
		s.AddHandler(d.guildCreate)
		s.AddHandler(d.ready)

		d.sessions[i] = s
	}
	d.s = d.sessions[0]
	return d, nil
}

func (d *DiscordService) Open() {
	for _, s := range d.sessions {
		s.Open()
	}
}

func (d *DiscordService) Close() {
	for _, s := range d.sessions {
		s.Close()
	}
}

func (d *DiscordService) RecentMembers(guildID string) ([]*GuildMember, error) {
	return d.state.MemberList(guildID)
}

func (d *DiscordService) ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("logged in as", r.User, r.SessionID)
}

func (d *DiscordService) guildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	fmt.Println("new guild", g.Name)
	d.state.createMemberList(g.ID)
}

func (d *DiscordService) memberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	fmt.Println("new member joined", m.GuildID, m.User.ID)
	_ = d.state.MemberAdd(&GuildMember{
		GuildID: m.GuildID,
		Member:  m.Member,
	})
}

func (d *DiscordService) memberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	fmt.Println("new member left", m.GuildID, m.User.ID)
}
