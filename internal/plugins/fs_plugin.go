package plugins

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/benmizrahi/gobig/internal/common"
	"github.com/benmizrahi/gobig/internal/domains"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FSPlugin struct {
	Path        string
	Format      string
	Parallelism int
}

// Execute implements contract.IPluginContract (worker job)
func (FSPlugin) Execute(task *domains.Task) *domains.TaskResult {

	from := task.Instruction[1]
	d, err := os.Open(from)
	if err != nil {
		logrus.Error("unable to read partition file", err)
		return &domains.TaskResult{
			Uuid:    task.Uuid,
			Status:  false,
			EndTime: timestamppb.Now(),
		}
	}
	defer d.Close()

	// csvReader := csv.NewReader(d)
	// data, err := csvReader.ReadAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	rows := []*domains.Row{}
	// for _, row := range data {

	// d := protos.Data{
	// 	Data: []string{},
	// }
	// for _, column := range row {
	// 	d.Data = append(d.Data, column)
	// }
	// rows = append(rows, &d)
	// }

	return &domains.TaskResult{
		Uuid:    task.Uuid,
		Status:  true,
		Rows:    rows,
		EndTime: timestamppb.Now(),
	}
}

// Configs implements contract.IPluginContract
func (p FSPlugin) Configs(conf map[string]string) common.IPluginContract {
	p.Format = conf["format"]
	p.Path = conf["path"]
	p.Parallelism = 1
	marks, err := strconv.ParseInt(conf["parallelism"], 10, 0)
	if err == nil {
		p.Parallelism = int(marks)
	}
	return p
}

// Plan implements plugins.IPluginContract
func (p FSPlugin) PlanRead() []*domains.IPartition {
	files, err := ioutil.ReadDir(p.Path)
	if err != nil {
		log.Fatal(err)
	}
	tasks := []*domains.Task{}
	for _, file := range files {
		tasks = append(tasks, &domains.Task{
			Uuid:         uuid.New().String(),
			Instruction:  []string{"read", p.Path + file.Name()},
			Plugin:       p.Name(),
			CreationTime: timestamppb.Now(),
		})
	}

	sliced := common.ChunkSlice(tasks, p.Parallelism)
	distribution := []*domains.IPartition{}

	for _, tasks := range sliced {
		distribution = append(distribution, &domains.IPartition{Tasks: tasks})
	}

	return distribution
}

// Name implements plugins.IPluginContract
func (p FSPlugin) Name() string {
	return "fsplugin"
}

// Name must be New + struct name
func NewFSPlugin() common.IPluginContract {
	return FSPlugin{}
}
