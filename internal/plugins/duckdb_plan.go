package plugins

import (
	"database/sql"

	"github.com/benmizrahi/duckspark/internal/common"
	"github.com/sirupsen/logrus"
)

type DuckPlugin struct {
	Path        string
	Format      string
	Parallelism int
}

// Configs implements IPluginContract.
func (p DuckPlugin) Configs(map[string]string) IPluginContract {
	panic("unimplemented")
}

// Plan implements IPluginContract.
func (p DuckPlugin) Plan(args ...interface{}) common.Maplan {
	query := args[0].(string)
	db := args[1].(*sql.DB)

	db.Exec("PRAGMA enable_profiling = 'json';")
	db.Exec("PRAGMA profiling_output='/tmp/test.json';")

	row, err := db.Query("explain "+query, "")
	if err != nil {
		logrus.Fatal(err)
	}

	db.Exec("PRAGMA disable_profiling;")

	var (
		id   string
		name string
	)
	row.Next()
	err = row.Scan(&id, &name)

n	logrus.Info(id, name)

	row.Close()
	return common.Maplan{
		Plan:  nil,
		Tasks: nil,
	}
}

// Name implements plugins.IPluginContract
func (p DuckPlugin) Name() string {
	return "duckdb_plugin"
}

// Name must be New + struct name
func NewDuckPlugin() IPluginContract {
	return DuckPlugin{
		Path:        "",
		Format:      "",
		Parallelism: 0,
	}
}
