package plugins

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/benmizrahi/duckspark/internal/common"
	"github.com/benmizrahi/duckspark/internal/protos"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FSPlugin struct {
	Path        string
	Format      string
	Parallelism int
}

// Configs implements contract.IPluginContract
func (p FSPlugin) Configs(conf map[string]string) IPluginContract {
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
func (p FSPlugin) Plan(args ...interface{}) common.Maplan {
	plan := "LOAD"
	path := args[len(args)-1].(string)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	distribution := []*protos.Task{}
	for _, file := range files {
		distribution = append(distribution, &protos.Task{
			Uuid:         uuid.New().String(),
			Commands:     []string{"SELECT * FROM '" + path + file.Name() + "'"},
			Plugin:       p.Name(),
			CreationTime: timestamppb.Now(),
		})
	}

	return common.Maplan{
		Plan:  &plan,
		Tasks: distribution,
	}
}

// Name implements plugins.IPluginContract
func (p FSPlugin) Name() string {
	return "fs_analyzer"
}

// Name must be New + struct name
func NewFSPlugin() IPluginContract {
	return FSPlugin{}
}
