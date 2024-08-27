package session

import (
	uuid "github.com/satori/go.uuid"
	"sync"
)

//    MemorySeesionMgr设计：
//    定义MemorySeesionMgr对象（字段：存放所有session的map，读写锁）
//    构造函数
//    Init()
//    CreateSeesion()
//    GetSession()

// 定义对象
type MemorySeesionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

func NewMemorySeesionMgr() SessionMgr {
	sr := &MemorySeesionMgr{sessionMap: make(map[string]Session, 1024)}
	return sr
}

func (m *MemorySeesionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (m *MemorySeesionMgr) CreateSession() (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	// go get github.com/satori/go.uuid
	// 用uuid作为sessionId
	id, err := uuid.NewV4()
	if err != nil {
		return
	}
	// 转string
	sessionId := id.String()
	// 创建个session
	session = NewMemorySession(sessionId)
	return
}

func (m *MemorySeesionMgr) Get(sessionId string) (session Session, err error) {
	return
}
