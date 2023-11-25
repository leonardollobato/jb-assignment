data "aws_iam_policy_document" "load_balancer_controller_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]
    effect  = "Allow"

    condition {
      test     = "StringEquals"
      variable = "${replace(module.eks.cluster_oidc_issuer_url, "https://", "")}:sub"
      values   = ["system:serviceaccount:kube-system:aws-load-balancer-controller"]
    }

    condition {
      test     = "StringEquals"
      variable = "${replace(module.eks.cluster_oidc_issuer_url, "https://", "")}:aud"
      values   = ["sts.amazonaws.com"]
    }

    principals {
      identifiers = [module.eks.oidc_provider_arn]
      type        = "Federated"
    }
  }
}

resource "aws_iam_role" "load_balancer_controller_role" {
  name               = local.iam_alb_role
  assume_role_policy = data.aws_iam_policy_document.load_balancer_controller_assume_role_policy.json

  inline_policy {
    name   = "aws-load-balancer-controller-policy"
    policy = file("${path.module}/iam-policies/aws-load-balancer-controller-policy.json")
  }
}

resource "kubernetes_service_account" "alb_service_account" {
  metadata {
    name      = "aws-load-balancer-controller"
    namespace = "kube-system"
    labels = {
      "app.kubernetes.io/name"      = "aws-load-balancer-controller"
      "app.kubernetes.io/component" = "controller"
    }
    annotations = {
      "eks.amazonaws.com/role-arn"               = aws_iam_role.load_balancer_controller_role.arn
      "eks.amazonaws.com/sts-regional-endpoints" = "true"
    }
  }
}

data "aws_iam_policy_document" "cluster_autoscaler_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]
    effect  = "Allow"

    condition {
      test     = "StringEquals"
      variable = "${replace(module.eks.cluster_oidc_issuer_url, "https://", "")}:sub"
      values   = ["system:serviceaccount:kube-system:cluster-autoscaler"]
    }

    principals {
      identifiers = [module.eks.oidc_provider_arn]
      type        = "Federated"
    }
  }
}


resource "aws_iam_role" "cluster_autoscaler_role" {
  name               = "cluster-autoscaler-role"
  assume_role_policy = data.aws_iam_policy_document.cluster_autoscaler_assume_role_policy.json

  inline_policy {
    name   = "cluster-autoscaler-policy"
    policy = file("${path.module}/iam-policies/cluster-auto-scaler-policy.json")
  }
}
