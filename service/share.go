package service

import (
	"errors"

	"github.com/mskydream/audio-cloud/repository"
	"github.com/mskydream/audio-cloud/storage"
)

type ShareService struct {
	repo repository.Share
}

func NewShareService(repo repository.Share) *ShareService {
	return &ShareService{repo: repo}
}

func (s *ShareService) ShareAudio(userID, audioId, shareId int) error {
	if userID == shareId {
		return errors.New("can't share own audio to yourself")
	}
	return s.repo.ShareAudio(userID, audioId, shareId)
}

func (s *ShareService) UnshareAudio(userID, audioId, shareId int) error {
	if userID == shareId {
		return errors.New("can't unshare own audio from yourself")
	}
	return s.repo.UnshareAudio(userID, audioId, shareId)
}

func (s *ShareService) GetSharedList(input storage.ShareListParam) (storage.ShareListJson, error) {
	return s.repo.GetSharedList(input)
}
