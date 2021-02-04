package crons

import (
	"github.com/juju/errors"
	"sync"

	cron "github.com/robfig/cron/v3"
)

/**
1、管理所有的定时任务,需要记录定时任务的编号和相关信息
2、停止一个定时任务
3、支持添加函数类型和接口类型任务
注意:
task id 是唯一的,实际开发可以使用 uuid
这个封装是并发安全的
*/

// https://www.cnblogs.com/zhangzhiping35/p/12661477.html
// 1、Cron 表达式
//cron表达式代表一个时间的集合，使用6个空格分隔的字段表示：
//字段名	        是否必须	允许的值	       允许的特定字符
//秒(Seconds)	      是	0-59	           * / , -
//分(Minute)	      是	0-59	           * / , -
//时(Hours)	        是	0-23	           * / , -
//日(Day of month)	是	1-31	           * / , - ?
//月(Month)	        是	1-12 或 JAN-DEC	 * / , -
//星期(Day of week)	否	0-6 或 SUM-SAT	 * / , - ?

// 2、Corn表达式说明
//月(Month)和星期(Day of week)字段的值不区分大小写，如：SUN、Sun 和 sun 是一样的。
//星期(Day of week)字段如果没提供，相当于是 *

// 3、cron特定字符说明
//符号	说明
//(*)	表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月
//(/)	表示增长间隔，如第1个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后每隔 15 分钟执行一次（即 3、18、33、48 这些时间点执行），这里也可以表示为：3/15
//(,)	用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行
//(-)	表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）
//(?)	只用于 日(Day of month) 和 星期(Day of week)，表示不指定值，可以用于代替 *

//4、常用cron举例
//每隔5秒执行一次：*/5 * * * * ?
//每隔1分钟执行一次：0 */1 * * * ?
//每天23点执行一次：0 0 23 * * ?
//每天凌晨1点执行一次：0 0 1 * * ?
//每月1号凌晨1点执行一次：0 0 1 1 * ?
//每周一和周三晚上22:30: 00 30 22 * * 1,3
//在26分、29分、33分执行一次：0 26,29,33 * * * ?
//每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
//每年三月的星期四的下午14:10和14:40:  00 10,40 14 ? 3 4

// 定时任务管理器
type CronManager struct {
	inner *cron.Cron
	ids   map[string]cron.EntryID
	mutex sync.Mutex
}

// cron.New() 默认从分开始，cron.WithSeconds() 加上后默认从秒开始
func NewCronManager() *CronManager {
	//更改为一个支持至秒级别的cron
	secondParser := cron.NewParser(cron.Second | cron.Minute |
			cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return &CronManager{
		inner: cron.New(cron.WithParser(secondParser), cron.WithChain()),
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
