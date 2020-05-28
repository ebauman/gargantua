package apiserver

import "github.com/hobbyfarm/gargantua/pkg/apiserver"

type PrePostFunc func(ctx apiserver.Context, args ...interface{}) error
type ActionFunc func(ctx apiserver.Context, args ...interface{}) (error, interface{})

type Config struct {
	PreList *PrePostFunc
	List *ActionFunc
	PostList *PrePostFunc

	PreGet *PrePostFunc
	Get *ActionFunc
	PostGet *PrePostFunc

	PreCreate *PrePostFunc
	Create *ActionFunc
	PostCreate *PrePostFunc

	PreUpdate *PrePostFunc
	Update *ActionFunc
	PostUpdate *ActionFunc

	PreDelete *PrePostFunc
	Delete *ActionFunc
	PostDelete *PrePostFunc
}