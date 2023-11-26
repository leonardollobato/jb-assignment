locals {
  # cluster_name        = var.cluster_name
  vpc_name            = "${var.cluster_name}-${var.environment}-vpc"
  iam_cluster_role    = "${var.cluster_name}-${var.environment}-role"
  iam_alb_role        = "aws-load-balancer-${var.cluster_name}-${var.environment}-role"
  iam_autoscaler_role = "autoscaler-${var.cluster_name}-${var.environment}-role"
}
