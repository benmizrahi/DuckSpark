package buildins

import (
	"github.com/benmizrahi/gobig/internal/common"
	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/sirupsen/logrus"
)

func Count(uuid string, data []*protos.Data) *protos.TaskResult {
	res := protos.TaskResult{
		Uuid:   uuid,
		Status: true,
		Data:   []*protos.Data{},
	}
	for _, r := range data {
		data, err := common.Deserialize(r.CompressData)
		if err != nil {
			logrus.Error("error deserialize data,", err)
		}

		b, err := common.Serialize([]interface{}{len(data)})
		if err != nil {
			logrus.Error("error Serialize result, ", err)
		}

		res.Data = append(res.Data, &protos.Data{
			DataTypes:    []protos.DataType{protos.DataType_int},
			CompressData: b,
		})
	}
	return &res
}
