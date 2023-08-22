package plugins

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/benmizrahi/godist/internal/common"
	"github.com/benmizrahi/godist/internal/protos"
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
func (FSPlugin) Execute(task *protos.Task) *protos.TaskResult {

	from := task.Instactions[1]
	d, err := os.Open(from)
	if err != nil {
		logrus.Error("unable to read partition file", err)
		return &protos.TaskResult{
			Uuid:    task.Uuid,
			Status:  false,
			EndTime: timestamppb.Now(),
		}
	}
	defer d.Close()

	csvReader := csv.NewReader(d)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	rows := []*protos.DataRow{}
	for _, row := range data {

		d := protos.DataRow{
			Data: []string{},
		}
		for _, column := range row {
			d.Data = append(d.Data, column)
		}
		rows = append(rows, &d)
	}

	return &protos.TaskResult{
		Uuid:    task.Uuid,
		Status:  true,
		Data:    rows,
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
func (p FSPlugin) PlanRead() []*protos.IPartition {
	files, err := ioutil.ReadDir(p.Path)
	if err != nil {
		log.Fatal(err)
	}
	tasks := []*protos.Task{}
	for _, file := range files {
		tasks = append(tasks, &protos.Task{
			Uuid:         uuid.New().String(),
			Instactions:  []string{"read", p.Path + file.Name()},
			Plugin:       p.Name(),
			CreationTime: timestamppb.Now(),
		})
	}

	sliced := common.ChunkSlice(tasks, p.Parallelism)
	distribution := []*protos.IPartition{}

	for _, tasks := range sliced {
		distribution = append(distribution, &protos.IPartition{Tasks: tasks})
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
