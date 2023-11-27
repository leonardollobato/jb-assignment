# Assignment 

1. Terraform
2. Docker Images
3. Argo CD Application Deployment
4. Interaction with Apps
6. Assets for Testing


## 1 Terraform

Terraform was used to create the underlying Infrastructure like the EKS cluster.

Also using terraform some complementary applications were installed using helm_release, like:
- metrics server
- cluster auto scaler
- aws load balancer controller
- Argocd



## 2 Docker Images

[jb-assignment-crawler](https://hub.docker.com/repository/docker/llleonardo/jb-assignment-crawler)

[jb-assignment-api](https://hub.docker.com/repository/docker/llleonardo/jb-assignment-api)

---

## 3 Argo CD - Applications Deployment

ArgoCD is used to maintain the two sample applications:

### API
A very simple golang api with the following endpoints:

```
/products (GET)
```
List objects in the bucket where the generated images are stored

```
/products (POST)

curl <load-balancer-url>/products \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '[{"title": "Jumbo Krulfriet Mild Gekruid 750g","url": "https://jumbo.com/dam-images/fit-in/720x720/Products/10082023_1691675073211_1691675123395_8718452431991_1.png"}]'

```
Receive the Products and send it to a SQS that is consumed by Crawler app

### Crawler

Consumes from a SQS queue that contains product images and title.

It apply some overlay in the image and take a screenshot.

The screenshot is sent to an S3 bucket.


### Install Applications

```
kubectl apply -f argocd/application/<application-yaml-file>
```



## 4 Interaction with Apps 

### Retrieve ArgoCD UI Admin password

```
kubectl -n argocd get secret argocd-initial-admin-secret \
          -o jsonpath="{.data.password}" | base64 -d; echo
```

## 6 Assets for Testing

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