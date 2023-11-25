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

  values = [file("${path.module}/values/metrics-server-values.yaml")]
}

################################################################################
# Cluster AutoScaler
################################################################################
# https://artifacthub.io/packages/helm/cluster-autoscaler/cluster-autoscaler
resource "helm_release" "cluster_autoscaler" {
  name = "cluster-autoscaler"

  repository = "https://kubernetes.github.io/autoscaler"
  chart      = "cluster-autoscaler"
  namespace  = "kube-system"
  version    = "9.32.0"

  set {
    name  = "awsRegion"
    value = var.region
  }

  set {
    name  = "serviceAccount.annotations.eks\\.amazonaws\\.com/role-arn"
    value = aws_iam_role.cluster_autoscaler_role.arn
  }

  values = [file("${path.module}/values/autoscaler-values.yaml")]
}


################################################################################
# AWS Load Balancer Controller
################################################################################
# https://artifacthub.io/packages/helm/aws/aws-load-balancer-controller
resource "helm_release" "aws_load_balancer_controller" {
  name = "aws-load-balancer-controller"

  repository = "https://aws.github.io/eks-charts"
  chart      = "aws-load-balancer-controller"
  namespace  = "kube-system"
  version    = "1.6.2"

  values = [
    templatefile("${path.module}/values/alb-controller-values.yaml", {
      clusterName = module.eks.cluster_name
    })
  ]
}

# ################################################################################
# # Argo CD
# ################################################################################

# # https://artifacthub.io/packages/helm/argo/argo-cd
# resource "helm_release" "argocd" {
#   name = "argocd"

#   repository = "https://argoproj.github.io/argo-helm"
#   chart      = "argo-cd"
#   namespace  = kubernetes_namespace.argocd_namespace.metadata[0].name
#   version    = "5.51.4"

#   values = [file("./values/argocd-values.yaml")]

# }
