variable "gcp_project_id" {
  type        = string
  description = "The target GCP project for the cluster."
}

variable "gcp_service_account" {
  type        = string
  description = "The service account for authenticating with GCP APIs."
}

variable "gcp_region" {
  type        = string
  description = "The target GCP region for the cluster."
}

variable "gcp_extra_labels" {
  type = map(string)

  description = <<EOF
(optional) Extra GCP labels to be applied to created resources.
Example: `{ "key" = "value", "foo" = "bar" }`
EOF

  default = {}
}


variable "gcp_bootstrap_instance_type" {
  type = string
  description = "Instance type for the bootstrap node. Example: `n1-standard-4`"
}

variable "gcp_master_instance_type" {
  type = string
  description = "Instance type for the master node(s). Example: `n1-standard-4`"
}

variable "gcp_image_id" {
  type = string
  description = "Image for all nodes."
}

variable "gcp_master_root_volume_type" {
  type = string
  description = "The type of volume for the root block device of master nodes."
}

variable "gcp_master_root_volume_size" {
  type = string
  description = "The size of the volume in gigabytes for the root block device of master nodes."
}
