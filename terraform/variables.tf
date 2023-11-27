
variable "cluster_name" {
  type        = string
  description = "AWS EKS cluster name"
  default     = "assignment-cluster"
}

variable "cluster_version" {
  type        = string
  description = "AWS EKS cluster version"
  default     = "1.28"
}

variable "region" {
  type        = string
  description = "AWS regions"
  default     = "us-east-1"
}

variable "environment" {
  type        = string
  description = "Infrastructure environment"
  default     = "dev"
}

variable "eks_cloudwatch_log_group_retention_in_days" {
  type        = number
  description = "Number of days to retain log events. Default retention - 90 days"
  default     = 1
}

variable "aws_auth_roles" {
  description = "Additional IAM roles to add to the aws-auth configmap."
  type = list(object({
    rolearn  = string
    username = string
    groups   = list(string)
  }))

  default = []
}

variable "aws_auth_users" {
  description = "Additional IAM users to add to the aws-auth configmap."

  type = list(object({
    userarn  = string
    username = string
    groups   = list(string)
  }))

  default = []
}
