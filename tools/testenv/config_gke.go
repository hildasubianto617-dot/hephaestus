package testenv

import (
	"errors"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// GKEConfig is used to create GKE test environments.
type GKEConfig struct {
	// Region in GCP where cluster will be created.
	Region string
	// ProjectID in GCP where cluster will be created.
	ProjectID string
	// KubernetesVersion supported by GKE.
	KubernetesVersion string
	// AgentServiceAccount is the email of a pre-created service account the controller impersonates
	// through Workload Identity to push to the test Google Artifact Registry.
	AgentServiceAccount string
	// NodeServiceAccount is the email of a pre-created service account the cluster nodes run as.
	NodeServiceAccount string
}

// ResourcePath to Terraform module.
func (c GKEConfig) ResourcePath() string {
	return "resources/terraform/gcp-gke"
}

// Vars derived from struct fields.
func (c GKEConfig) Vars() []byte {
	f := hclwrite.NewEmptyFile()

	root := f.Body()
	root.SetAttributeValue("region", cty.StringVal(c.Region))
	root.SetAttributeValue("project_id", cty.StringVal(c.ProjectID))
	root.SetAttributeValue("agent_service_account", cty.StringVal(c.AgentServiceAccount))
	root.SetAttributeValue("node_service_account", cty.StringVal(c.NodeServiceAccount))

	if strings.TrimSpace(c.KubernetesVersion) != "" {
		root.SetAttributeValue("kubernetes_version", cty.StringVal(c.KubernetesVersion))
	}

	return f.Bytes()
}

// Validate ensures required variables are provided.
func (c GKEConfig) Validate() (err error) {
	if strings.TrimSpace(c.Region) == "" {
		err = multierror.Append(err, errors.New("GCP region cannot be blank"))
	}

	if strings.TrimSpace(c.ProjectID) == "" {
		err = multierror.Append(err, errors.New("GCP project id cannot be blank"))
	}

	if strings.TrimSpace(c.AgentServiceAccount) == "" {
		err = multierror.Append(err, errors.New("agent service account cannot be blank"))
	}

	if strings.TrimSpace(c.NodeServiceAccount) == "" {
		err = multierror.Append(err, errors.New("node service account cannot be blank"))
	}

	return err
}
