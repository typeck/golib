package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	rediswatcher "github.com/escaletech/casbin-redis-watcher/watcher"
	gormAdapter "github.com/lun-go/gorm-adapter/v3"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CasbinRule struct {
	PType string `json:"p_type" gorm:"size:100;"`
	V0    string `json:"v0" gorm:"size:100;"`
	V1    string `json:"v1" gorm:"size:100;"`
	V2    string `json:"v2" gorm:"size:100;"`
	V3    string `json:"v3" gorm:"size:100;"`
	V4    string `json:"v4" gorm:"size:100;"`
	V5    string `json:"v5" gorm:"size:100;"`
}

func (CasbinRule) TableName() string {
	return "sys_casbin_rule"
}

// perm.conf
// Initialize the model from a string.
var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

func Setup(db *gorm.DB, prefix string) *casbin.SyncedEnforcer {
	Apter, err := NewAdapter(db, prefix)
	if err != nil {
		panic(err)
	}
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewSyncedEnforcer(m, Apter)
	if err != nil {
		panic(err)
	}

	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}

	//log.SetLogger(&Logger{})
	e.EnableLog(true)
	return e
}

func NewAdapter(db *gorm.DB, prefix string) (*gormAdapter.Adapter, error) {
	db = db.Clauses(clause.OnConflict{UpdateAll: true})
	a, err := gormAdapter.NewAdapterByDBUseTableName(db, prefix, "")
	return a, err
}

func RegisterWatcher(e *casbin.SyncedEnforcer, addr, channelName string) (persist.Watcher, error) {
	w, err := rediswatcher.New(rediswatcher.Options{
		RedisURL: addr,
		Channel:  channelName,
		Cluster:  true,
	})
	if err != nil {
		return nil, errors.WithMessage(err, addr)
	}
	e.SetWatcher(w)
	e.EnableAutoNotifyWatcher(false)
	w.SetUpdateCallback(func(s string) {
		e.LoadPolicy()
		fmt.Println("auto LoadPolicy sands/casbin/*", s)
	})

	return w, nil
}
