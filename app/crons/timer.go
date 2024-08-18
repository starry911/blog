package crons

import (
	"blog/app/crons/jobs"
	"github.com/robfig/cron"
)

type Timer struct {
	cron *cron.Cron
	job  jobs.Job
}

func (t *Timer) Start() {
	t.cron.Start()
}

func (t *Timer) Stop() {
	t.cron.Stop()
}

func New() {
	c := cron.New()

	t := &Timer{
		cron: c,
		job:  jobs.Job{},
	}
	if err := t.CronTasks(); err != nil {
		panic("Cron Start Error:" + err.Error())
	}

	t.Start()
}

func (t *Timer) CronTasks() error {
	for _, j := range []struct {
		Name string
		Spec string
		Cmd  func()
	}{
		//{
		//	Name: "测试定时任务执行",
		//	Spec: func() string {
		//		return enum.CronEverySecond
		//	}(),
		//	Cmd: t.job.TestJob,
		//},
	} {
		if err := t.cron.AddFunc(j.Spec, j.Cmd); err != nil {
			return err
		}
	}
	return nil
}
