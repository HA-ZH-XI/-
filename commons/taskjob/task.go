package taskjob

import (
	"context"
	"github.com/beego/beego/v2/task"
)

func InitTask() {
	/*开始执行一个任务
	参数1：任务的名称
	参数2：cron表达式
	参数3：当前tk1这个定时具体要做什么事情
	*/
	Task1 := task.NewTask("tk1", "0 0 2 * * ？", GenerateWarning)
	task.AddTask("Task1", Task1)
}

func GenerateWarning(ctx context.Context) error {
	//业务
	return nil
}
