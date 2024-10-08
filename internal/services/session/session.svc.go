package session

import (
	"github.com/alexedwards/scs/v2"
)

type SessionService struct {
	SessionMgr *scs.SessionManager
}

var SessionSvc *SessionService

func NewSessionService() *SessionService {
	return &SessionService{}
}

func init() {
	SessionSvc = NewSessionService()
}

func (ss *SessionService) RegisterSessionManager(sessionMgr *scs.SessionManager) {
	ss.SessionMgr = sessionMgr
}