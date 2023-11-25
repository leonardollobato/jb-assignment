output "cluster_endpoint" {
  description = "Endpoint for EKS control plane."
  value       = module.eks.cluster_endpoint
}

output "cluster_security_group_id" {
  description = "Security group ids attached to the cluster control plane."
  value       = module.eks.cluster_security_group_id
}

output "node_groups" {
  description = "Outputs from node groups"
  value       = module.eks.eks_managed_node_groups
}

output "oidc_provider_arn" {
  value = module.eks.oidc_provider_arn
}

output "oidc_issuer_url" {
  value = module.eks.cluster_oidc_issuer_url
}

# Cluster Autoscaler
output "cluster_autoscaler_role_arn" {
  value = aws_iam_role.cluster_autoscaler_role.arn
}
