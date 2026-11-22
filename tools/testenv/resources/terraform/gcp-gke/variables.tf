variable "project_id" {
  type        = string
  description = "The project ID in which to host the GKE cluster."
}

variable "region" {
  type        = string
  description = "The region in which the cluster is hosted."
}

variable "agent_service_account" {
  type        = string
  description = "Email of a pre-created service account that the controller impersonates through Workload Identity. It needs roles/artifactregistry.writer in the project, and roles/iam.workloadIdentityUser granted to the controller's Kubernetes service account."
}

variable "node_service_account" {
  type        = string
  description = "Email of a pre-created service account that cluster nodes run as. The identity running Terraform needs roles/iam.serviceAccountUser on it."
}

variable "kubernetes_version" {
  type        = string
  description = "The Kubernetes version of the masters. If set to 'latest' it will pull latest available version in the selected region."
  default     = "latest"
}
