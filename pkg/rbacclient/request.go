package rbacclient

const (
	OperatorAnd = "AND"
	OperatorOr  = "OR"
)

type Request struct {
	operator    string
	permissions []Permission
}

func RbacRequest() *Request {
	return &Request{
		permissions: []Permission{},
	}
}

func (r *Request) GetOperator() string {
	if r.operator == "" {
		return OperatorAnd
	}

	return r.operator
}

func (r *Request) GetPermissions() []Permission {
	return r.permissions
}

func (r *Request) And() *Request {
	r.operator = OperatorAnd
	return r
}

func (r *Request) Or() *Request {
	r.operator = OperatorOr
	return r
}

func (r *Request) HobbyfarmPermission(resource string, verb string) *Request {
	return r.HobbyfarmPermissionWithName(resource, verb, "")
}

func (r *Request) Permission(apigroup string, resource string, verb string) *Request {
	return r.PermissionWithName(apigroup, resource, verb, "")
}

func (r *Request) HobbyfarmPermissionWithName(resource string, verb string, name string) *Request {
	r.permissions = append(r.permissions, HobbyfarmPermission{Resource: resource, Verb: verb, ResourceName: name})

	return r
}

func (r *Request) PermissionWithName(apigroup string, resource string, verb string, name string) *Request {
	r.permissions = append(r.permissions, GenericPermission{APIGroup: apigroup, Resource: resource, Verb: verb, ResourceName: name})

	return r
}

type Permission interface {
	GetAPIGroup() string
	GetResource() string
	GetVerb() string
	GetResourceName() string
}

type GenericPermission struct {
	APIGroup     string
	Resource     string
	Verb         string
	ResourceName string
}

type HobbyfarmPermission struct {
	Resource     string
	Verb         string
	ResourceName string
}

func (hf HobbyfarmPermission) WithResourceName(resourceName string) {
	hf.ResourceName = resourceName
}

func (gp GenericPermission) WithResourceName(resourceName string) {
	gp.ResourceName = resourceName
}

func (g GenericPermission) GetAPIGroup() string {
	return g.APIGroup
}

func (g GenericPermission) GetResource() string {
	return g.Resource
}

func (g GenericPermission) GetVerb() string {
	return g.Verb
}

func (g GenericPermission) GetResourceName() string {
	return g.ResourceName
}

func (h HobbyfarmPermission) GetAPIGroup() string {
	return APIGroup
}

func (h HobbyfarmPermission) GetResource() string {
	return h.Resource
}

func (h HobbyfarmPermission) GetVerb() string {
	return h.Verb
}

func (h HobbyfarmPermission) GetResourceName() string {
	return h.ResourceName
}
