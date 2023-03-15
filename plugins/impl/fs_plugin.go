package impl

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/benmizrahi/godist/common"
	"github.com/benmizrahi/godist/plugins/contract"
	"github.com/benmizrahi/godist/protos"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FSPlugin struct {
	Path        string
	Format      string
	Parallelism int
}

// Configs implements contract.IPluginContract
func (p FSPlugin) Configs(conf map[string]string) contract.IPluginContract {
	p.Format = conf["format"]
	p.Path = conf["path"]
	p.Parallelism = 1
	marks, err := strconv.ParseInt(conf["parallelism"], 10, 0)
	if err == nil {
		p.Parallelism = int(marks)
	}
	return p
}

// Name implements plugins.IPluginContract
func (p FSPlugin) Name() string {
	return "fsplugin"
}

// Plan implements plugins.IPluginContract
func (p FSPlugin) PlanRead() []contract.IPartition {
	files, err := ioutil.ReadDir(p.Path)
	if err != nil {
		log.Fatal(err)
	}
	tasks := []*protos.Task{}
	for _, file := range files {
		tasks = append(tasks, &protos.Task{
			Uuid:         uuid.New().String(),
			Instactions:  []string{"read", p.Path + file.Name()},
			CreationTime: timestamppb.Now(),
		})
	}

	sliced := common.ChunkSlice(tasks, p.Parallelism)
	distribution := []contract.IPartition{}

	for _, tasks := range sliced {
		distribution = append(distribution, contract.IPartition{Tasks: tasks})
	}

	return distribution
}

// Name must be New + struct name
func NewFSPlugin() contract.IPluginContract {
	return FSPlugin{}
}
