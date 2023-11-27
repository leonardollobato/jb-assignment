################################################################################
# Metrics Server
################################################################################
# https://artifacthub.io/packages/helm/metrics-server/metrics-server
resource "helm_release" "metrics-server" {
  name = "metrics-server"

  repository = "https://kubernetes-sigs.github.io/metrics-server/"
  chart      = "metrics-server"
  namespace  = "kube-system"
  version    = "3.11.0"

  values = [file("${path.module}/values/metrics.values.yaml")]
}

################################################################################
# Cluster AutoScaler
################################################################################
# https://artifacthub.io/packages/helm/cluster-autoscaler/cluster-autoscaler
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
  name               = local.iam_autoscaler_role
  assume_role_policy = data.aws_iam_policy_document.cluster_autoscaler_assume_role_policy.json

  inline_policy {
    name   = "cluster-autoscaler-policy"
    policy = file("${path.module}/iam-policies/cluster-auto-scaler-policy.json")
  }
}

resource "helm_release" "cluster_autoscaler" {
  name = "cluster-autoscaler"

  repository = "https://kubernetes.github.io/autoscaler"
  chart      = "cluster-autoscaler"
  namespace  = "kube-system"
  version    = "9.32.0"

  values = [
    templatefile("${path.module}/values/autoscaler.values.yaml", {
      awsRegion = var.region,
      roleArn   = aws_iam_role.cluster_autoscaler_role.arn
    })
  ]
}


################################################################################
# AWS Load Balancer Controller
################################################################################
# https://artifacthub.io/packages/helm/aws/aws-load-balancer-controller
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

resource "helm_release" "aws_load_balancer_controller" {
  name = "aws-load-balancer-controller"

  repository = "https://aws.github.io/eks-charts"
  chart      = "aws-load-balancer-controller"
  namespace  = "kube-system"
  version    = "1.6.2"

  values = [
    templatefile("${path.module}/values/alb.values.yaml", {
      clusterName = module.eks.cluster_name
    })
  ]
}


# ################################################################################
# # Argo CD
# ################################################################################
resource "kubernetes_namespace" "argocd_namespace" {
  metadata {
    name = "argocd"
  }
}

# # https://artifacthub.io/packages/helm/argo/argo-cd
resource "helm_release" "argocd" {
  name = "argocd"

  repository = "https://argoproj.github.io/argo-helm"
  chart      = "argo-cd"
  namespace  = kubernetes_namespace.argocd_namespace.metadata[0].name
  version    = "5.51.4"

  values = [file("./values/argocd.values.yaml")]

  depends_on = [time_sleep.wait_argocd_termination, time_sleep.wait_argocd_creation]
}

resource "time_sleep" "wait_argocd_termination" {
  destroy_duration = "30s"
}

resource "time_sleep" "wait_argocd_creation" {
  create_duration = "30s"

  depends_on = [helm_release.aws_load_balancer_controller]
}
