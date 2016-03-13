package engine

import (
	"github.com/deluan/gosonic/domain"
)

type Playlists interface {
	GetAll() (*domain.Playlists, error)
	Get(id string) (*PlaylistInfo, error)
}

func NewPlaylists(pr domain.PlaylistRepository, mr domain.MediaFileRepository) Playlists {
	return playlists{pr, mr}
}

type playlists struct {
	plsRepo   domain.PlaylistRepository
	mfileRepo domain.MediaFileRepository
}

func (p playlists) GetAll() (*domain.Playlists, error) {
	return p.plsRepo.GetAll(domain.QueryOptions{})
}

type PlaylistInfo struct {
	Id      string
	Name    string
	Entries []Entry
}

func (p playlists) Get(id string) (*PlaylistInfo, error) {
	pl, err := p.plsRepo.Get(id)
	if err != nil {
		return nil, err
	}

	if pl == nil {
		return nil, ErrDataNotFound
	}

	pinfo := &PlaylistInfo{Id: pl.Id, Name: pl.Name}
	pinfo.Entries = make([]Entry, len(pl.Tracks))

	// TODO Optimize: Get all tracks at once
	for i, mfId := range pl.Tracks {
		mf, err := p.mfileRepo.Get(mfId)
		if err != nil {
			return nil, err
		}
		pinfo.Entries[i] = FromMediaFile(mf)
		pinfo.Entries[i].Track = 0
	}

	return pinfo, nil
}