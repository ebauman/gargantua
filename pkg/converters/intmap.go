package converters

import "github.com/hobbyfarm/gargantua/pkg/protobuf"

func Int32MapFromRPC(req map[string]int32) map[string]int {
	out := map[string]int{}

	for k, v := range req {
		out[k] = int(v)
	}

	return out
}

func Int32MapToRPC(req map[string]int) map[string]int32 {
	out := map[string]int32{}

	for k, v := range req {
		out[k] = int32(v)
	}

	return out
}

func IntMapFromRPC(req *protobuf.Int32Map) map[string]int {
	out := map[string]int{}

	for k, v := range req.Map {
		out[k] = int(v)
	}

	return out
}

func IntMapToRPC(req map[string]int) *protobuf.Int32Map {
	out := map[string]int32{}

	for k, v := range req {
		out[k] = int32(v)
	}

	return &protobuf.Int32Map{Map: out}
}
