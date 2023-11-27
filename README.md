# Assignment (Work In Progress)

1. Terraform
2. Applications
3. Docker Images
4. Argo CD Application Deployment
5. Monitoring
6. Interaction with Apps
    5.1. Load Balancer URL
    5.2. Commands 
7. GitHub Actions CI
8. Assets for Testing



## 4 Argo CD Application Deployment



### Retrieve ArgoCD UI Admin password

```
kubectl -n argocd get secret argocd-initial-admin-secret \
          -o jsonpath="{.data.password}" | base64 -d; echo
```

### Retrieve Grafana Dashboard Admin password
```
kubectl get secret \
-n argocd \
kube-prometheus-stack-grafana \
-o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```

## 6.2 Commands

### Add Product

Replace "load-balancer-url"

```
curl <load-balancer-url>/products \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '[{"title": "Jumbo Krulfriet Mild Gekruid 750g","url": "https://jumbo.com/dam-images/fit-in/720x720/Products/10082023_1691675073211_1691675123395_8718452431991_1.png"}]'
```


## 7 Assets for Testing

- Jumbo Krulfriet Mild Gekruid 750g
    - https://jumbo.com/dam-images/fit-in/720x720/Products/10082023_1691675073211_1691675123395_8718452431991_1.png

- Jumbo Sinaasappelsap 1L
    - https://jumbo.com/dam-images/fit-in/720x720/Products/28092023_1695921475729_1695921478695_8718452649280_1.png 

- Broodgeluk - Volkoren Bollen - 6 Stuks
    - https://jumbo.com/dam-images/fit-in/720x720/Products/18082023_1692358384367_1692358390660_8718452713592_1.png

- Bolletje Kruidnoot Letters 200g
    - https://jumbo.com/dam-images/fit-in/720x720/Products/17082023_1692235190233_1692235208193_369994_ZK_08710482533744_C1N1.png

```
curl lb-url/products \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '[{"title": "Jumbo Krulfriet Mild Gekruid 750g","url": "https://jumbo.com/dam-images/fit-in/720x720/Products/10082023_1691675073211_1691675123395_8718452431991_1.png"}, {"title": "Jumbo Sinaasappelsap 1L","url": "https://jumbo.com/dam-images/fit-in/720x720/Products/28092023_1695921475729_1695921478695_8718452649280_1.png "}, {"title": "Broodgeluk - Volkoren Bollen - 6 Stuks","url": "https://jumbo.com/dam-images/fit-in/720x720/Products/18082023_1692358384367_1692358390660_8718452713592_1.png"}]'
```
```
curl lb-url/products \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '[{"title": "Jumbo Krulfriet Mild Gekruid 750g","url": "https://jumbo.com/dam-images/fit-in/720x720/Products/10082023_1691675073211_1691675123395_8718452431991_1.png"}, {"title": "Jumbo Sinaasappelsap 1L","url": "https://jumbo.com/dam-images/fit-in/720x720/Products/28092023_1695921475729_1695921478695_8718452649280_1.png "}, {"title": "Broodgeluk - Volkoren Bollen - 6 Stuks","url": "https://jumbo.com/dam-images/fit-in/720x720/Products/17082023_1692235190233_1692235208193_369994_ZK_08710482533744_C1N1.png"}]'
```