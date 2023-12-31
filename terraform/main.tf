terraform {

  # For remote state
  # backend "s3" {
  #     bucket = ""
  #     key = ""
  #     region = ""
  # }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.2"
    }
  }
}

provider "aws" {
  region = var.region

  default_tags {
    tags = {
      Environment = "${var.environment}"
      Terraform   = "true"
    }
  }
}

provider "kubernetes" {
  host                   = data.aws_eks_cluster.cluster.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority[0].data)
  token                  = data.aws_eks_cluster_auth.cluster.token
}

provider "helm" {
  kubernetes {
    host                   = data.aws_eks_cluster.cluster.endpoint
    cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority[0].data)
    token                  = data.aws_eks_cluster_auth.cluster.token
  }
}

resource "random_string" "suffix" {
  length  = 3
  special = false
}

