package v2

import (
	v1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/genericcondition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []User `json:"items"`
}

type UserSpec struct {
	Email       string            `json:"email"`
	Password    string            `json:"password"`
	AccessCodes []string          `json:"access_codes"`
	Settings    map[string]string `json:"settings"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Provider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProviderSpec `json:"spec"`
}

// +k8s:deepcopy-gen

type ProviderSpec struct {
	// ConfigurationItems is a slice of v1.Setting that allows a provider to define configuration options for its operation.
	// This includes items such as api credentials, instance size or type, image, security groups, etc.
	// Note that this field does not set the _value_ of these items, but rather defines them as variables
	// to be configured on a MachineTemplate, Environment, etc.
	ConfigurationItems []v1.Setting
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Provider `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MachineTemplateSpec `json:"spec"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MachineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MachineTemplate `json:"items"`
}

// +k8s:deepcopy-gen

type MachineTemplateSpec struct {
	// MachineType specifies the type of this machine.
	// For instance, a machine may be assigned to a user, in which case MachineTypeUser is valid.
	// A machine could also be shared amongst users, in which case MachineTypeShared is valid.
	MachineType MachineType `json:"machineType"`

	// ProviderSpecifics allows for configuration of providers within the definition of this machine template.
	// Administrators may specify configuration values here, keyed by the name of the provider.
	// An example value (yaml) may look like:
	// providerSpecifics:
	//   aws-ec2:
	//     image: ami-09812309234
	//   azure-vm:
	//     size: Dsv3
	ProviderSpecifics map[string]map[string]string `json:"providerSpecifics"`

	// DisplayName is used to display the name of this template in a 'pretty' manner in the UI
	// or table output in Kubernetes
	DisplayName string `json:"displayName"`

	// ConnectProtocol specifies the ways this machine can receive connections.
	ConnectProtocols []ConnectProtocol `json:"connectProtocols"`
}

// MachineType specifies the type of machine, e.g. shared or user.
// If there becomes a need for more types of machines, they can be added to this type via
// const declarations.
type MachineType string

const (
	MachineTypeUser   MachineType = "user"
	MachineTypeShared MachineType = "shared"
)

// ConnectProtocol defines how HobbyFarm should attempt connection to a machine.
type ConnectProtocol string

const (
	ConnectProtocolSSH  ConnectProtocol = "ssh"
	ConnectProtocolVNC  ConnectProtocol = "vnc"
	ConnectProtocolRDP  ConnectProtocol = "rdp"
	ConnectProtocolGuac ConnectProtocol = "guac"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EnvironmentSpec   `json:"spec"`
	Status EnvironmentStatus `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Environment `json:"items"`
}

// +k8s:deepcopy-gen

type EnvironmentSpec struct {
	// Capacity defines the capacity of the environment for each template.
	// For example, this field could be (yaml):
	// ```
	// capacity:
	//   template-a: 100
	//   template-b: 200
	//   ubuntu-2204: 150
	// ```
	// This field does not make determinations on the size of the template.
	// It is up to the administrator to calculate the overall capacity of their environment and split it up
	// accordingly to the templates.
	// Further, specifying any capacity value here makes this environment a candidate environment for scheduling.
	// When attempting to create a ScheduledEvent or otherwise schedule machines, a capacity value in this
	// field will mark this environment to HobbyFarm as being able to support that template.
	Capacity map[string]int `json:"capacity"`

	// ConnectEndpoints is a map of connection protocols to endpoints that this environment uses when connecting
	// to machines.
	ConnectEndpoints map[ConnectProtocol]string `json:"connectEndpoints"`

	// Provider is the name of the provider which this environment uses when creating machines.
	Provider string `json:"provider"`

	// ProviderConfiguration defines values for configuring the Provider that apply to all machines being scheduled.
	// This is a good location to put top-level config items such as API credentials (or names of secrets that contain those
	// credentials), or other "overarching" config items that apply to this environment.
	ProviderConfiguration map[string]string `json:"providerConfiguration"`

	// TemplateConfiguration defines values for configuring the Provider when provisioning specific templates.
	// This is a good location for template-specific items such as security groups, cloud init data, etc.
	TemplateConfiguration map[string]map[string]string `json:"templateConfiguration"`

	// DisplayName is the pretty name of the environment used in UIs.
	DisplayName string `json:"displayName"`
}

// +k8s:deepcopy-gen

type EnvironmentStatus struct {
	Conditions []genericcondition.GenericCondition `json:"conditions,omitempty"`
}

var (
	ConditionProviderFound condition.Cond = condition.Cond("ProviderFound")
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MachineSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MachineSetSpec   `json:"spec"`
	Status MachineSetStatus `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MachineSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MachineSet `json:"items"`
}

// +k8s:deepcopy-gen

type MachineSetSpec struct {
	AvailabilityStrategy AvailabilityStrategy `json:"availabilityStrategy"`
	ProvisioningStrategy ProvisioningStrategy `json:"provisioningStrategy"`
	MaxProvisioned       int                  `json:"maxProvisioned"`
	MinAvailable         int                  `json:"minAvailable"`
}

type AvailabilityStrategy string

const (
	AvailabilityStrategyAccessCode AvailabilityStrategy = "accesscode"
	AvailabilityStrategyPool       AvailabilityStrategy = "pool"
)

type ProvisioningStrategy string

const (
	ProvisioningStrategyOnDemand  ProvisioningStrategy = "ondemand"
	ProvisioningStrategyAutoScale ProvisioningStrategy = "autoscale"
)

// +k8s:deepcopy-gen

type MachineSetStatus struct {
	Conditions  []genericcondition.GenericCondition `json:"conditions"`
	Provisioned int                                 `json:"provisioned"`
	Available   int                                 `json:"available"`
}

var (
	ConditionMinimumMet condition.Cond = condition.Cond("MinimumMet")
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Machine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MachineSpec   `json:"spec"`
	Status MachineStatus `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Machine `json:"items"`
}

type MachineSpec struct {
	MachineType      MachineType                `json:"machineType"`
	Provider         string                     `json:"provider"`
	Configuration    map[string]string          `json:"configuration"`
	MachineSet       string                     `json:"machineSet"`
	MachineTemplate  string                     `json:"machineTemplate"`
	Environment      string                     `json:"environment"`
	ConnectEndpoints map[ConnectProtocol]string `json:"connectEndpoints"`
}

type MachineLifecycle string

const (
	LifecycleUnknown  MachineLifecycle = "unknown"
	LifecycleReady    MachineLifecycle = "ready"
	LifecycleNotReady MachineLifecycle = "notready"
)

type MachineStatus struct {
	Conditions []genericcondition.GenericCondition `json:"conditions"`
	Lifecycle  MachineLifecycle                    `json:"lifecycle"`
	Properties map[string]string                   `json:"properties"`
}

type MachineClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MachineClaimSpec   `json:"spec"`
	Status MachineClaimStatus `json:"status,omitempty"`
}

type ClaimStrategy string

const (
	ClaimStrategyAnyAvailable ClaimStrategy = "anyavailable"
	ClaimStrategyAccessCode   ClaimStrategy = "accesscode"
)

type MachineClaimSpec struct {
	MachineTemplate string        `json:"machineTemplate"`
	User            string        `json:"user"`
	ClaimStrategy   ClaimStrategy `json:"claimStrategy"`
	AccessCode      string        `json:"accessCode"`
}

type MachineClaimLifecycle string

const (
	MachineClaimLifecycleAvailable MachineClaimLifecycle = "available"
	MachineClaimLifecycleBound     MachineClaimLifecycle = "bound"
	MachineClaimLifecycleFailed    MachineClaimLifecycle = "failed"
)

type MachineClaimStatus struct {
	MachineClaimLifecycle MachineClaimLifecycle `json:"machineClaimLifecycle"`
}
