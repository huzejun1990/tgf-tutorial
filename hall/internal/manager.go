// @Author huzejun 2024/2/1 17:47:00
package internal

import (
	"github.com/huzejun1990/tgf/tgf-tutorial/hall/internal/entity"
	"github.com/thkhxm/tgf/db"
	"github.com/thkhxm/tgf/log"
)

type Manager struct {
	accountCache db.IAutoCacheService[string, *entity.UserAccount]
}

/*func (m *Manager) CheckAccount(account string) bool {
	_, ok := m.accountCache.Get(account)
	return ok == nil
}*/

func (m *Manager) GetAccount(account, password string) (*entity.UserAccount, int32) {
	if res, ok := m.accountCache.Get(account); ok == nil {
		if res.Password != password {
			log.DebugTag("account", "password error:%s", account)
			return nil, 404
		}
		log.DebugTag("account", "get account:%s", account)
		return res, 0
	}
	log.DebugTag("account", "account is not exists:%s", account)
	return nil, 0
}

func (m *Manager) CreateAccount(account, password, userId string) *entity.UserAccount {
	res := &entity.UserAccount{
		Account:  account,
		Password: password,
		UserId:   userId,
	}
	m.accountCache.Set(res, account)
	//
	log.DebugTag("account", "create account:%s", account)
	return res
}

func (m *Manager) UpdatePassword(account, password string) {
	if res, ok := m.accountCache.Get(account); ok == nil {
		res.Password = password
		//
		//m.accountCache.Set(res, account)
		//
		m.accountCache.Push(account)
		log.DebugTag("account", "updatae password:%s", account)
	}
}

func (m Manager) InitStruct() {
	m.accountCache = db.NewLongevityAutoCacheManager[string, *entity.UserAccount]("account")
	//m.accountCache = db.NewAutoCacheBuilder[string, *entity.UserAccount]()
}

func NewManager() *Manager {
	m := new(Manager)
	m.InitStruct()
	return m
}
