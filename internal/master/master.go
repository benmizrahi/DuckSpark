package master

import (
	"io"
	"net/http"
	"strconv"
	"sync"

	"github.com/benmizrahi/gobig/internal/bigfream"
	"github.com/benmizrahi/gobig/internal/common"
	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/benmizrahi/gobig/internal/shuffle"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

// singeltone instance of master!
var lock = &sync.Mutex{}

// Singel instance
var masterInstance *Master

type Master struct {
	MasterPath string
	context    *bigfream.Context
	Http       *gin.Engine
	shuffle    shuffle.IShuffler
}

func NewMaster(isLocal bool, host string, port int, minWorkers int) *Master {
	if masterInstance == nil {
		lock.Lock()
		log.Info("gobig Master, Creating new master instance")
		m := &Master{
			MasterPath: host + ":" + strconv.Itoa(port),
			Http:       gin.New(),
		}

		m.Http.Use(ginlogrus.Logger(logrus.New()), gin.Recovery())

		m.Http.POST("/api/register", m.RegisterHandler)
		go m.Http.Run(m.MasterPath)
		log.Info("gobig Master, master is listening on ", m.MasterPath)

		m.context = bigfream.NewContext(isLocal, minWorkers, m.MasterPath)
		m.shuffle = shuffle.NewMasterShuffler(m.Http)

		m.context.InitContext()
		lock.Unlock()

		return m
	}
	return masterInstance
}

func (m *Master) RegisterHandler(c *gin.Context) {
	buf, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}
	req := &protos.RegisterReq{}
	if err := proto.Unmarshal(buf, req); err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}

	m.context.Workers[req.Uuid] = req.Http

	data := &protos.RegisterRes{
		Ok: true,
	}

	c.ProtoBuf(http.StatusOK, data)
}

func (m *Master) Parallelize(data [][]string, option bigfream.BigOptions) *bigfream.Bigfream {
	var bf *bigfream.Bigfream
	if option.Columns != nil {
		bf = bigfream.NewBigfream(m.context, &option.Columns)
	} else {
		//infer types
	}
	partitions, err := m.buildParallelizePartitons(common.ConvertStringSliceToInterfaceSlice(data), &option.Repartiton)
	if err != nil {
		logrus.Error("error building Parallelize partitions")
		m.shutDown()
	}
	bf.AssginPartitons(partitions)
	return bf
}

func (m *Master) Load() *bigfream.Bigfream {
	mf := bigfream.NewBigfream(m.context, nil)
	return mf
}

////////////////////////////////////////////////////////////////////////////////////////////////

func (m *Master) buildParallelizePartitons(data [][]interface{}, requestedNumPartitions *int) ([]*protos.IPartition, error) {

	numPartitions := m.calculatePartitons(data)

	partitions := make([]*protos.IPartition, numPartitions)

	// Shuffle the data
	for index := range data {
		partitionIndex := index % numPartitions
		if partitions[partitionIndex] == nil {
			partitions[partitionIndex] = &protos.IPartition{
				Uuid: uuid.New().String(),
			}
		}
		// if partitions[partitionIndex].Rows == nil {
		// 	partitions[partitionIndex].Rows = make([]*protos.Row, 0)
		// }
		// b, err := common.Serialize[interface{}](row)
		// if err != nil {
		// 	logrus.Error("error Serialize row,", err)
		// }
		// partitions[partitionIndex].Rows = append(partitions[partitionIndex].Rows, &protos.Row{
		// 	CompressRow: b,
		// })
	}

	return partitions, nil
}

func (m *Master) calculatePartitons(data [][]interface{}) int {

	totalProcessingUnits := len(m.context.Workers)

	// Calculate the total size of the data
	totalDataSize := common.CalculateTotalDataSize(data)

	// Calculate the optimal partition size based on available resources and data size
	optimalPartitionSize := totalDataSize / totalProcessingUnits

	// Calculate the number of partitions based on the optimal partition size
	numPartitions := len(data) / optimalPartitionSize

	// Adjust the number of partitions to a minimum value of 1
	if numPartitions <= 0 {
		numPartitions = 2
	}
	return numPartitions
}

func (m *Master) shutDown() {
	log.Fatal("error")
}
