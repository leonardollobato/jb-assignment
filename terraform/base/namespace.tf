################################################################################
# Namespaces
################################################################################
resource "kubernetes_namespace" "tools_namespace" {
  metadata {
    name = "tools"
  }
}

resource "kubernetes_namespace" "argocd_namespace" {
  metadata {
    name = "argocd"
  }
}
