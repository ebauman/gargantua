package v2

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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

type Provider struct {
	Spec ProviderSpec
}

type ProviderSpec struct {
	Configuration map[string]any
}

type ScalarType string

const (
	ScalarTypeString  ScalarType = "string"
	ScalarTypeInteger ScalarType = "integer"
	ScalarTypeBoolean ScalarType = "boolean"
	ScalarTypeFloat   ScalarType = "float"
)

type ScalarConfigurationItem struct {
	Key   string
	Type  ScalarType
	Value any
}

type MapConfigurationItem struct {
	Key    string
	Values map[string]ScalarConfigurationItem
}

type ArrayConfigurationItem struct {
	Key    string
	Values []ScalarConfigurationItem
}

/*
configuration:
	itemOne: string
	itemTwo:
		itemTwoSub: string
	itemThree:
	- string
	itemFour: boolean
*/
