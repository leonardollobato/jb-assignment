data "aws_availability_zones" "available" {}

data "aws_caller_identity" "current" {}

data "aws_eks_cluster" "cluster" {
  name = module.eks.cluster_name

  depends_on = [module.eks.cluster_name]
}

data "aws_eks_cluster_auth" "cluster" {
  name = module.eks.cluster_name

  depends_on = [module.eks.cluster_name]
}
