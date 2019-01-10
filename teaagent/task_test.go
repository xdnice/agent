package teaagent

import (
	"github.com/TeaWeb/code/teaconfigs/agents"
	"sync"
	"testing"
	"time"
)

func TestTask_Run(t *testing.T) {
	config := agents.NewTaskConfig()
	config.Id = "test"
	config.Script = `/usr/bin/env bash\n\necho "Hello"`
	config.Env = []*agents.EnvVariable{
		{
			Name:  "name",
			Value: "Tom",
		},
	}
	//config.Cwd = "/home/www"

	task := NewTask(config)
	proc, stdout, stderr, err := task.Run()
	t.Log("stdout:", stdout)
	t.Log("stderr:", stderr)

	defer t.Log(proc)

	if err != nil {
		t.Fatal("err:" + err.Error())
	}
}

func TestTask_RunConcurrent(t *testing.T) {
	config := agents.NewTaskConfig()
	config.Id = "test"
	config.Script = `/usr/bin/env bash\n\necho "Hello"`
	config.Env = []*agents.EnvVariable{
		{
			Name:  "name",
			Value: "Tom",
		},
	}
	//config.Cwd = "/home/www"

	task := NewTask(config)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i ++ {
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			_, stdout, _, _ := task.Run()
			t.Log("stdout:", stdout)
		}()
	}
	wg.Wait()
}

func TestTask_Schedule(t *testing.T) {
	config := agents.NewTaskConfig()

	schedule := agents.NewScheduleConfig()
	schedule.AddSecondRanges(&agents.ScheduleRangeConfig{
		Value: -1,
		From:  0,
		To:    59,
		Step:  2,
	})
	config.AddSchedule(schedule)
	config.Validate()

	task := NewTask(config)
	task.Schedule()
	time.Sleep(60 * time.Second)
}
