package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// NOTE: vorrei memorizzare i dati relativi alla session sul client. Quindi SessionManager dovrebbe diventare obsoleto.
// Al momento con login/logout non lo è.

var (
	SessMgr = &SessionManager{
		cookieName:  "LiveBlogCookie",
		sessionsWS:  make(map[string]*SessionCtx),
		maxlifetime: 3600 * 24, // Max session life in seconds
	}
)

type SessionCtx struct {
	ID        string
	Username  string
	CreatedAt time.Time
}

type SessionManager struct {
	cookieName  string
	lock        sync.Mutex
	sessionsWS  map[string]*SessionCtx
	maxlifetime int64
}

func (manager *SessionManager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "default"
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *SessionManager) GetSession(w http.ResponseWriter, r *http.Request) (*SessionCtx, error) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	var sn string
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sn = manager.sessionId()
		log.Printf("Set a new session %s", sn)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sn), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sn, _ = url.QueryUnescape(cookie.Value)
		//fmt.Println("Using cookie value ", cookie.Value)
	}
	var session *SessionCtx
	var ok bool
	if session, ok = manager.sessionsWS[sn]; !ok {
		session = &SessionCtx{
			ID:        sn,
			CreatedAt: time.Now(),
		}
		manager.sessionsWS[sn] = session
		//fmt.Println("Insert the new session data ", sn)
	}
	return session, nil
}

func (manager *SessionManager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	//log.Printf("Session GC call after %d seconds", manager.maxlifetime)
	destrKeys := []string{}
	now := time.Now()
	var max float64
	max = float64(manager.maxlifetime)
	for k, v := range manager.sessionsWS {
		elapsed := now.Sub(v.CreatedAt)
		if elapsed.Seconds() > max {
			destrKeys = append(destrKeys, k)
		}
	}
	for _, k := range destrKeys {
		log.Printf("Delete session %s", k)
		delete(manager.sessionsWS, k)
	}
	seconds, _ := time.ParseDuration(fmt.Sprintf("%ds", manager.maxlifetime))
	time.AfterFunc(time.Duration(seconds), func() {
		manager.GC()
	})
}

func init() {
	go SessMgr.GC()
}
