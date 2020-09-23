package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
)

func CMSStructFromRPC(req *protobuf.CMSStruct) hfv1.CMSStruct {
	return hfv1.CMSStruct{
		CPU:     int(req.CPU),
		Memory:  int(req.Memory),
		Storage: int(req.Storage),
	}
}

func CMSStructToRPC(req hfv1.CMSStruct) *protobuf.CMSStruct {
	return &protobuf.CMSStruct{
		CPU:     int32(req.CPU),
		Memory:  int32(req.Memory),
		Storage: int32(req.Storage),
	}
}
