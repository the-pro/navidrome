package persistence

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/deluan/gosonic/engine"
)

var (
	nowPlayingKeyName = []byte("nowplaying")
)

type nowPlayingRepository struct {
	ledisRepository
}

func NewNowPlayingRepository() engine.NowPlayingRepository {
	r := &nowPlayingRepository{}
	r.init("nowplaying", &engine.NowPlayingInfo{})
	return r
}

func (r *nowPlayingRepository) Set(id, username string, playerId int, playerName string) error {
	if id == "" {
		return errors.New("Id is required")
	}
	m := &engine.NowPlayingInfo{TrackId: id, Username: username, Start: time.Now(), PlayerId: playerId, PlayerName: playerName}

	h, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return Db().SetEX(nowPlayingKeyName, int64(engine.NowPlayingExpire.Seconds()), []byte(h))
}

func (r *nowPlayingRepository) Clear(playerId int) (*engine.NowPlayingInfo, error) {
	val, err := Db().GetSet(nowPlayingKeyName, nil)
	if err != nil {
		return nil, err
	}
	info := &engine.NowPlayingInfo{}
	err = json.Unmarshal(val, info)
	if err != nil {
		return nil, nil
	}
	return info, nil
}

func (r *nowPlayingRepository) GetAll() (*[]engine.NowPlayingInfo, error) {
	val, err := Db().Get(nowPlayingKeyName)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return &[]engine.NowPlayingInfo{}, nil
	}
	info := &engine.NowPlayingInfo{}
	err = json.Unmarshal(val, info)
	if err != nil {
		return &[]engine.NowPlayingInfo{}, nil
	}
	return &[]engine.NowPlayingInfo{*info}, nil
}

var _ engine.NowPlayingRepository = (*nowPlayingRepository)(nil)