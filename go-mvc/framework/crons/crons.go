package crons

import (
	"github.com/juju/errors"
	"sync"

	cron "github.com/robfig/cron/v3"
)

// 定时任务管理器
type CronManager struct {
	inner *cron.Cron
	ids   map[string]cron.EntryID
	mutex sync.Mutex
}

/**
1、管理所有的定时任务,需要记录定时任务的编号和相关信息
2、停止一个定时任务
3、支持添加函数类型和接口类型任务
注意:
task id 是唯一的,实际开发可以使用 uuid
这个封装是并发安全的
不足:
未支持初始化参数
定时任务的错误采集,如果某个定时任务出错,应该能够获取到错误信息(这里指的是错误不是 panic)
panic 恢复操作可以参考 withChain和cron.Recover
*/
func NewCronManager() *CronManager {
	// cron.New() 默认从分开始，cron.WithSeconds() 加上后默认从秒开始
	return &CronManager{
		inner: cron.New(),
		ids:   make(map[string]cron.EntryID),
	}
}

// IDs ...
func (c *CronManager) IDs() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	validIDs := make([]string, 0, len(c.ids))
	invalidIDs := make([]string, 0)
	for sid, eid := range c.ids {
		if e := c.inner.Entry(eid); e.ID != eid {
			invalidIDs = append(invalidIDs, sid)
			continue
		}
		validIDs = append(validIDs, sid)
	}
	for _, id := range invalidIDs {
		delete(c.ids, id)
	}
	return validIDs
}

// 开始定时任务
func (c *CronManager) Start() {
	c.inner.Start()
}

// 停止定时任务
func (c *CronManager) Stop() {
	c.inner.Stop()
}

/**
 * 删除定时任务
 * @param id 唯一任务id
 */
func (c *CronManager) DelByID(id string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	eid, ok := c.ids[id]
	if !ok {
		return
	}

	c.inner.Remove(eid)
	delete(c.ids, id)
}

/**
 * 添加定时任务
 * @param id 任务唯一id
 * @param spec 配置定时执行时间
 * @param cj 需要执行的任务方法
 * @return error
 */
func (c *CronManager) AddByID(id string, spec string, cj cron.Job) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.ids[id]; ok {
		return errors.Errorf("id为【%s】的定时任务已存在", id)
	}

	eid, err := c.inner.AddJob(spec, cj)
	if err != nil {
		return err
	}
	c.ids[id] = eid
	return nil
}

/**
 * 通过函数添加定时任务
 * @param id 任务唯一id
 * @param spec 配置定时执行时间
 * @param f 需要执行的方法
 * @return error
 */
func (c *CronManager) AddByFunc(id string, spec string, f func()) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.ids[id]; ok {
		return errors.Errorf("id为【%s】的定时任务已存在", id)
	}

	eid, err := c.inner.AddFunc(spec, f)
	if err != nil {
		return err
	}
	c.ids[id] = eid
	return nil
}

/**
 * 判断是否存在定时任务
 * @param jid 唯一任务jid
 * @return bool
 */
func (c *CronManager) IsExists(jid string) bool {
	_, exist := c.ids[jid]
	return exist
}
