package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

type DiscordService struct {
	sessions []*discordgo.Session
	s        *discordgo.Session
	//state    *State
}

/*
type GuildMember struct {
	GuildID string            `json:"guild_id"`
	Member  *discordgo.Member `json:"member"`
}
*/
const shardCount = 1

func NewDiscordService(token string) (*DiscordService, error) {
	d := &DiscordService{
		//state: NewState(),
	}
	d.sessions = make([]*discordgo.Session, shardCount)
	for i := 0; i < shardCount; i++ {
		s, err := discordgo.New("Bot " + token)
		if err != nil {
			return nil, err
		}

		s.State.TrackVoice = false
		s.State.TrackPresences = false
		s.ShardCount = shardCount
		s.ShardID = i
		s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers | discordgo.IntentMessageContent)

		s.AddHandler(d.memberAdd)
		s.AddHandler(d.memberRemove)
		s.AddHandler(d.guildCreate)
		s.AddHandler(d.ready)
		s.AddHandler(memberChunk)

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

func (d *DiscordService) Guild(guildID string) (*discordgo.Guild, error) {
	for _, s := range d.sessions {
		if g, err := s.State.Guild(guildID); err == nil {
			return g, err
		}
	}
	return nil, discordgo.ErrStateNotFound
}

/*
func (d *DiscordService) GuildMembers(guildID string) ([]*discordgo.Member, error) {
	g, err := d.Guild(guildID)
	if err != nil {
		return nil, err
	}
	d.s.State.Member()
}*/

func (d *DiscordService) GuildJoins(guildID string) ([]*discordgo.Member, error) {
	g, err := d.Guild(guildID)
	if err != nil {
		return nil, err
	}

	fmt.Println(len(g.Members))

	newMems := []*discordgo.Member{}
	for _, m := range g.Members {
		if m.JoinedAt.After(time.Now().Add(-time.Hour * 24)) {
			newMems = append(newMems, m)
		}
	}

	/*
		newMems := []*GuildMember{}
		for _, m := range g.Members {
			if m.JoinedAt.After(time.Now().Add(-1 * time.Hour * 24)) {
				newMems = append(newMems, &GuildMember{guildID, m})
			}
		}
	*/
	fmt.Println(len(newMems))
	return newMems, nil
}

func (d *DiscordService) ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("logged in as", r.User, r.SessionID)
}

func (d *DiscordService) guildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	s.RequestGuildMembers(g.ID, "", 0, "", false)
	fmt.Println("loading: ", g.Guild.Name, g.MemberCount, len(g.Members))
	//d.state.createMemberList(g.ID)
}

func (d *DiscordService) memberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	fmt.Println("new member joined", m.GuildID, m.User.ID)
	/*
		_ = d.state.MemberAdd(&GuildMember{
			GuildID: m.GuildID,
			Member:  m.Member,
		})
	*/
}

func (d *DiscordService) memberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	fmt.Println("new member left", m.GuildID, m.User.ID)
}

func memberChunk(s *discordgo.Session, g *discordgo.GuildMembersChunk) {
	fmt.Println(g.GuildID, len(g.Members))
	if g.ChunkIndex == g.ChunkCount-1 {
		// I don't know if this will work with several shards
		guild, err := s.Guild(g.GuildID)
		if err != nil {
			return
		}
		fmt.Println("finished loading "+guild.Name, len(guild.Members))
	}
}
