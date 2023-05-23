package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kvapi "kubevirt.io/api/core/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +genclient:nonNamespaced

const VirtualMachineFinalizer = "finalizers.virtualization.ecpaas.io/virtualmachine"

type DiskVolumeTemplateSourceImage struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type DiskVolumeTemplateSource struct {
	Image DiskVolumeTemplateSourceImage `json:"image,omitempty"`
}

type ResourceRequirements struct {
	// Requests is a description of the initial vmi resources.
	// Valid resource keys are "memory" and "cpu".
	// +optional
	Requests v1.ResourceList `json:"requests,omitempty"`
	// Limits describes the maximum amount of compute resources allowed.
	// Valid resource keys are "memory" and "cpu".
	// +optional
	Limits v1.ResourceList `json:"limits,omitempty"`
}

type DiskVolumeTemplateSpec struct {
	// Resources represents the minimum resources the volume should have.
	Resources ResourceRequirements     `json:"resources,omitempty"`
	Source    DiskVolumeTemplateSource `json:"source,omitempty"`
}

type DiskVolumeTemplateStatus struct {
}

type DiskVolumeTemplate struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// DiskVolumeSpec is the spec for a DiskVolume resource
	Spec   DiskVolumeTemplateSpec   `json:"spec,omitempty"`
	Status DiskVolumeTemplateStatus `json:"status,omitempty"`
}

type Cpu struct {
	Cores int32 `json:"cores,omitempty"`
}

type MacVtap struct {
}

type Interface struct {
	MacVtap MacVtap `json:"macvtap,omitempty"`
	Name    string  `json:"name,omitempty"`
}

type Devices struct {
	Interfaces []Interface `json:"interfaces,omitempty"`
}

type Domain struct {
	Cpu       Cpu                  `json:"cpu,omitempty"`
	Devices   Devices              `json:"devices,omitempty"`
	Resources ResourceRequirements `json:"resources,omitempty"`
}

type Multus struct {
	NetworkName string `json:"networkName,omitempty"`
}

type Network struct {
	Multus Multus `json:"multus,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Hardware struct {
	Domain           Domain         `json:"domain,omitempty"`
	EvictionStrategy string         `json:"evictionStrategy,omitempty"`
	Hostname         string         `json:"hostname,omitempty"`
	Networks         []Network      `json:"networks,omitempty"`
	Volumes          []kvapi.Volume `json:"volumes,omitempty"`
}

// VirtualMachineSpec defines the desired state of VirtualMachine
type VirtualMachineSpec struct {
	// DiskVolumeTemplate is the name of the DiskVolumeTemplate.
	DiskVolumeTemplates []DiskVolumeTemplate `json:"diskVolumeTemplates,omitempty"`
	// DiskVolume is the name of the DiskVolume.
	DiskVolumes []string `json:"diskVolumes,omitempty"`
	// Hardware is the hardware of the VirtualMachine.
	Hardware Hardware `json:"hardware,omitempty"`
}

// +kubebuilder:resource:shortName={ksvm,ksvms}
// +kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualMachine runs a vm at a given name.
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec         `json:"spec,omitempty"`
	Status kvapi.VirtualMachineStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualMachineList contains a list of VirtualMachine
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachine{}, &VirtualMachineList{})
}