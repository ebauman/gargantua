package converters

import "github.com/hobbyfarm/gargantua/pkg/protobuf"

func StringMapFromRPC(req *protobuf.StringMap) map[string]string {
	return req.Map
}

func StringMapToRPC(req map[string]string) *protobuf.StringMap {
	return &protobuf.StringMap{Map: req}
}