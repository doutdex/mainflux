package api

import (
	"github.com/asaskevich/govalidator"
	"github.com/mainflux/mainflux/manager"
)

const maxLimitSize = 100

type apiReq interface {
	validate() error
}

type userReq struct {
	user manager.User
}

func (req userReq) validate() error {
	return req.user.Validate()
}

type identityReq struct {
	key string
}

func (req identityReq) validate() error {
	if req.key == "" {
		return manager.ErrUnauthorizedAccess
	}

	return nil
}

type addClientReq struct {
	key    string
	client manager.Client
}

func (req addClientReq) validate() error {
	if req.key == "" {
		return manager.ErrUnauthorizedAccess
	}

	return req.client.Validate()
}

type updateClientReq struct {
	key    string
	id     string
	client manager.Client
}

func (req updateClientReq) validate() error {
	if req.key == "" {
		return manager.ErrUnauthorizedAccess
	}

	if !govalidator.IsUUID(req.id) {
		return manager.ErrNotFound
	}

	return req.client.Validate()
}

type createChannelReq struct {
	key     string
	channel manager.Channel
}

func (req createChannelReq) validate() error {
	if req.key == "" {
		return manager.ErrUnauthorizedAccess
	}

	return nil
}

type updateChannelReq struct {
	key     string
	id      string
	channel manager.Channel
}

func (req updateChannelReq) validate() error {
	if req.key == "" {
		return manager.ErrUnauthorizedAccess
	}

	if !govalidator.IsUUID(req.id) {
		return manager.ErrNotFound
	}

	return nil
}

type viewResourceReq struct {
	key string
	id  string
}

func (req viewResourceReq) validate() error {
	if req.key == "" {
		return manager.ErrUnauthorizedAccess
	}

	if !govalidator.IsUUID(req.id) {
		return manager.ErrNotFound
	}

	return nil
}

type listResourcesReq struct {
	key    string
	offset int
	limit  int
}

func (req *listResourcesReq) validate() error {
	if req.key == "" {
		return manager.ErrUnauthorizedAccess
	}

	if req.offset >= 0 && req.limit > 0 && req.limit <= maxLimitSize {
		return nil
	}

	return manager.ErrMalformedEntity
}

type connectionReq struct {
	key      string
	chanId   string
	clientId string
}

func (req connectionReq) validate() error {
	if req.key == "" {
		return manager.ErrUnauthorizedAccess
	}

	if !govalidator.IsUUID(req.chanId) || !govalidator.IsUUID(req.clientId) {
		return manager.ErrNotFound
	}

	return nil
}
