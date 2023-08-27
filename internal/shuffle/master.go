package shuffle

import (
	"io"
	"log"

	"github.com/benmizrahi/gobig/internal/protos"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

type MasterShuffler struct {
	IShuffler
	//worker to partitions
	state map[string][]string
}

func NewMasterShuffler(http *gin.Engine) IShuffler {
	shuffle := MasterShuffler{}
	http.POST("/api/shuffle/track", shuffle.track)
	return &shuffle
}

func (m *MasterShuffler) Orginize() error {
	return nil
}

func (m *MasterShuffler) track(c *gin.Context) {
	buf, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}
	req := &protos.TrackReq{}
	if err := proto.Unmarshal(buf, req); err != nil {
		log.Fatalln("Failed to parse register request:", err)
	}
	m.state[req.Worker] = req.Partitions
}
