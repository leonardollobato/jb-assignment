module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "19.19.0"

  cluster_name    = var.cluster_name
  cluster_version = var.cluster_version

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets

  cluster_endpoint_private_access = true
  cluster_endpoint_public_access  = true

  cluster_addons = {
    aws-ebs-csi-driver = {
      most_recent = true
    }
    coredns = {
      most_recent = true
    }
    kube-proxy = {
      most_recent = true
    }
  }

  kms_key_administrators = [
    "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
  ]

  enable_irsa = true

  cluster_enabled_log_types = [
    "scheduler"
  ]

  cloudwatch_log_group_retention_in_days = var.eks_cloudwatch_log_group_retention_in_days

  iam_role_name = local.iam_cluster_role

  eks_managed_node_group_defaults = {
    instance_types = ["t3.medium", "t3a.medium"]
    iam_role_additional_policies = {
      AmazonEBSCSIDriverPolicy           = "arn:aws:iam::aws:policy/service-role/AmazonEBSCSIDriverPolicy"
      AmazonEC2ContainerRegistryReadOnly = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
      AWSLambdaSQSQueueExecutionRole     = "arn:aws:iam::aws:policy/service-role/AWSLambdaSQSQueueExecutionRole"
    }
  }

  eks_managed_node_groups = {

    ondemand_instances = {
      desired_size = 1
      min_size     = 1
      max_size     = 5

      instance_types = ["t3a.medium", "t3.medium"]

      capacity_type = "ON_DEMAND"

      update_config = {
        max_unavailable_percentage = 50
      }

    }

    spot_instances = {
      desired_size = 1
      min_size     = 0
      max_size     = 5

      instance_types = ["t3a.large", "t3.large", "t3.medium", "t3a.medium"]

      capacity_type = "SPOT"

      update_config = {
        max_unavailable_percentage = 50
      }
    }
  }

  manage_aws_auth_configmap = true

  aws_auth_roles = var.aws_auth_roles
  aws_auth_users = var.aws_auth_users
}
