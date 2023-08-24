package buildins

import (
	"github.com/benmizrahi/gobig/internal/common"
	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/sirupsen/logrus"
)

func Count(uuid string, data []*protos.Row) *protos.TaskResult {
	res := protos.TaskResult{
		Uuid:   uuid,
		Status: true,
		Rows:   []*protos.Row{},
	}
	for _, r := range data {
		data, err := common.Deserialize(r.CompressRow)
		if err != nil {
			logrus.Error("error deserialize data,", err)
		}

		b, err := common.Serialize([]interface{}{len(data)})
		if err != nil {
			logrus.Error("error Serialize result, ", err)
		}

		res.Rows = append(res.Rows, &protos.Row{
			DataTypes:   []protos.DataType{protos.DataType_int},
			CompressRow: b,
		})
	}
	return &res
}
