package gcp

import (
	"encoding/json"

	gcpprovider "github.com/openshift/cluster-api-provider-gcp/pkg/apis/gcpprovider/v1beta1"
)

// Auth is the collection of credentials that will be used by terrform.
type Auth struct {
	ProjectID      string `json:"gcp_project_id,omitempty"`
	ServiceAccount string `json:"gcp_service_account,ompitempty"`
}

type config struct {
	Auth                  `json:",inline"`
	Region                string `json:"gcp_region,omitempty"`
	BootstrapInstanceType string `json:"gcp_bootstrap_instance_type,omitempty"`
	MasterInstanceType    string `json:"gcp_master_instance_type,omitempty"`
	ImageID               string `json:"gcp_image_id,omitempty"`
	VolumeType            string `json:"gcp_master_root_volume_type,omitempty"`
	VolumeSize            int64  `json:"gcp_master_root_volume_size,omitempty"`
}

// TFVars generates gcp-specific Terraform variables launching the cluster.
func TFVars(auth Auth, masterConfigs []*gcpprovider.GCPMachineProviderSpec) ([]byte, error) {
	masterConfig := masterConfigs[0]
	cfg := &config{
		Auth:                  auth,
		Region:                masterConfig.Region,
		BootstrapInstanceType: masterConfig.MachineType,
		MasterInstanceType:    masterConfig.MachineType,
		VolumeType:            masterConfig.Disks[0].Type,
		VolumeSize:            masterConfig.Disks[0].SizeGb,
		ImageID:               masterConfig.Disks[0].Image,
	}

	return json.MarshalIndent(cfg, "", "  ")
}
