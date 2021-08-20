# NGINX Ingress Controller for Kubernetes Tutorial: YAML & Helm | Example | Prometheus | Grafana | EKS

[YouTube Tutorial]()

## Prerequisites
- [Kubernetes](https://kubernetes.io/)
- [Helm](https://helm.sh/)

## Deploy Prometheus on Kubernetes Cluster 
- Deploy CRDs
```bash
kubectl apply -f prometheus/0-crd
```
- Deploy Prometheus Operator
```bash
kubectl apply -f prometheus/1-prometheus-operator
```
- Deploy Prometheus
```bash
kubectl apply -f prometheus/2-prometheus
```
- Check Prometheus pods
```bash
kubectl get pods -n monitoring
```

## Deploy Nginx Ingress Controller (YAML & HELM)
- Add Nginx ingress helm repo
```bash
helm repo add ingress-nginx \
  https://kubernetes.github.io/ingress-nginx
```
- Update Helm repo
```bash
helm repo update
```
- Search for Helm
```bash
helm search repo nginx
```

- Create `values.yaml` to override default [parameters](https://github.com/kubernetes/ingress-nginx/blob/main/charts/ingress-nginx/values.yaml)

- Generate YAML from the Helm chart
```bash
helm template my-ing ingress-nginx/ingress-nginx \
  --namespace ingress \
  --version 3.35.0 \
  --values values.yaml \
  --output-dir my-ing
```

- Deploy Nginx ingress with Helm
```bash
helm install my-ing ingress-nginx/ingress-nginx \
  --namespace ingress \
  --version 3.35.0 \
  --values values.yaml \
  --create-namespace
```

- List Helm releases
```bash
helm list -n ingress
```

- Get pods
```bash
kubectl get pods -n ingress
```
- Get services
```bash
kubectl get svc -n ingress
```

## Monitor Nginx Ingress with Prometheus
- Port forward Prometheus to localhost for now
```bash
kubectl get svc -n monitoring
kubectl port-forward svc/prometheus-operated 9090 -n monitoring
```

- Add monitoring label to ingress namespace
```bash
kubectl edit namespace ingress
```
```yaml
monitoring: prometheus
```

## Deploy Grafana on Kubernetes Cluster
- Generate admin password for Grafana
```bash
echo -n "devops123" | base64
```

- Decode
```
echo "devops123" | base64 -d
```

- Deploy grafana
```bash
kubectl apply -f grafana
```

- Get pods
```bash
kubectl get pods -n monitoring
```

## Import Nginx Ingress Controller Grafana Dashboard
- Port forward Grafana to localhost for now
```bash
kubectl get svc -n monitoring
kubectl port-forward svc/grafana 3000 -n monitoring
```

- Login to Grafana, user: `admin`, password: `devops123`

- Add Prometheus datasource
```bash
http://prometheus-operated:9090
```

- Google it, `nginx ingress grafana dashboard`

- Import `9614` dashboard

## Create Ingress for Prometheus
- Get services
```bash
kubectl get svc -n monitoring
```

- Create ingress
```bash
kubectl apply -f example-1/prometheus.yaml
```

- Get ingresses
```bash
kubectl get ing -n monitoring
kubectl get svc -n ingress
```

- Create CNAME record for `prometheus.devopsbyexample.io`

- Go to `http://prometheus.devopsbyexample.io`

## Test Nginx Ingress Admission Webhook

- Add valid ngin directive annotation to Prometheus ingress
```yaml
annotations:
  nginx.ingress.kubernetes.io/configuration-snippet: |
    more_set_headers 'Foo: bar';
```
```bash
kubectl apply -f example-1/prometheus.yaml
```
- Make a mistake in the directive and apply
- Render the full nginx config
```bash
kubectl get pods -n ingress
kubectl exec <pod> -n ingress -- cat /etc/nginx/nginx.conf
```
- Search for `more_set_headers`

## Create Ingress for Grafana
- Get services
```bash
kubectl get svc -n monitoring
```

- Create ingress
```bash
kubectl apply -f example-2/grafana.yaml
```

- Get ingresses
```bash
kubectl get ing -n monitoring
```

- Create CNAME record for `grafana.devopsbyexample.io`

- Go to `http://grafana.devopsbyexample.io`

## Simple Fanout Ingress Example

- Create `app`

- Create `example-3`

- Apply example-3
```bash
kubectl apply -f example-3
```
- Get ingress
```bash
kubectl get ing -n staging
```

- Create CNAME record for `api`

- Test with curl
```bash
curl http://api.devopsbyexample.io/foo/ytafsd
curl http://api.devopsbyexample.io/bar/123
```

## Virtual Hosting Ingress Example

- Create `example-4`

- Apply example-4
```bash
kubectl apply -f example-4
```
- Get ingress
```bash
kubectl get ing -n staging
```

- Create CNAME record for `foo` and `bar`

- Test with curl
```bash
curl http://foo.devopsbyexample.io/hello
curl http://bar.devopsbyexample.io/blog
```

## Nginx Ingress TLS Example
- Install cfssl
```bash
brew install cfssl
```
- Create config `certs/0-config.json`
- Create CA certificate request `certs/1-ca-csr.json`
- Change directory to `certs` and generate CA
```bash
cfssl gencert -initca 1-ca-csr.json | cfssljson -bare ca
openssl x509 -in ca.pem -text -noout
```
- Create certificate request for `foo-api.devopsbyexample.io` domain `certs/2-foo-api-csr.json`
- Generate certificate
```bash
cfssl gencert \
  -ca=ca.pem \
  -ca-key=ca-key.pem \
  -config=0-config.json \
  -profile=demo \
  2-foo-api-csr.json | cfssljson -bare foo-api
```
- Open with OpenSSL
```bash
openssl x509 -in foo-api.pem -text -noout
```
- Create Kubernetes secret `example-5/7-tls-secret.yaml`
```yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: foo-api-devopsbyexample-io-tls
  namespace: staging
type: kubernetes.io/tls
data:
  tls.crt: base64
  tls.key: base64
```

- Encode tls certificate `certs/foo-api.pem` to base64
```bash
echo -n "tls-cert-content" | base64
```
- Encode private key `certs/foo-api.pem` to base64
```bash
echo -n "tls-private-key-content" | base64
```
- Create ingress `example-5/8-tls-ingress.yaml`
- Switch directory and apply
```bash
cd ..
kubectl apply -f example-5
```
- Get ingress
```bash
kubectl get ing -n staging
```
- Create CNAME record for `foo-api`
- Go to `https://foo-api.devopsbyexample.io`
- Add CA to KeyChain

## Nginx Ingress Different Namespaces Example

- Create `example-6`
- Apply
```bash
kubectl apply -f example-6
```
- Use local service
- Get pods in both namespaces
```bash
kubectl get pods -n foo
kubectl get pods -n bar
```
- Get ing
```bash
kubectl get ing -n foo
```
- Create CNAME for `api-ns`
- Test
```bash
curl http://api-ns.devopsbyexample.io/foo/asd
curl http://api-ns.devopsbyexample.io/bar/asd
```
- Create `example-6/6-bar-external.yaml`
- Update `example-6/5-ingress.yaml`
- Apply
```bash
kubectl apply -f example-6
```
- Test
```bash
curl http://api-ns.devopsbyexample.io/bar/asd
```
## Nginx Ingress TCP Example

- Create following files
  - `example-7/0-namespace.yaml`
  - `example-7/1-secrets.yaml`
  - `example-7/2-statefulset.yaml`
  - `example-7/3-service.yaml`

- Apply

```bash
kubectl apply -f example-7
```
```bash
kubectl get pods -n database
```

- Create `example-7/4-configmap.yaml`
```bash
kubectl apply -f example-7/4-configmap.yaml
```
- Add tcp configmap flag
```bash
- --tcp-services-configmap=$(POD_NAMESPACE)/tcp-services
```

```bash
kubectl get svc my-ing-ingress-nginx-controller -n ingress
kubectl get deployment -n ingress
kubectl edit deployment -n ingress my-ing-ingress-nginx-controller
kubectl edit svc my-ing-ingress-nginx-controller -n ingress
```
```yaml
- name: postgres
  port: 5444
  protocol: TCP
```
```bash
kubectl get svc my-ing-ingress-nginx-controller -n ingress
```
- Go to AWS open LB and security group
- Create CNAME for postgres
```bash
kubectl get svc my-ing-ingress-nginx-controller -n ingress
```
```bash
psql --host postgres.devopsbyexample.io \
  --port 5444 \
  --username postgres \
  --password
```
```bash
\l
```

## Monitor Nginx Ingress with Grafana
- Open `http://grafana.devopsbyexample.io`
- Create some traffic
```bash
curl http://api.devopsbyexample.io/bar/123asd
curl http://bar.devopsbyexample.io/bar
curl http://foo-ns.devopsbyexample.io/bar
curl http://api-ns.devopsbyexample.io/foo/asdad
```

Topics
- rewrite-target
- An Ingress allows you to keep the number of load balancers down to a minimum
- [Simple fanout](https://kubernetes.io/docs/concepts/services-networking/ingress/#simple-fanout)
- [Name based virtual hosting](https://kubernetes.io/docs/concepts/services-networking/ingress/#name-based-virtual-hosting)
- [TLS](https://kubernetes.io/docs/concepts/services-networking/ingress/#tls) great segue to cert-manager
- classic loadbalancer vs network vs application
- internal vs external
- http/2 h2 h2c (plain text http)
- distrolless image include ca time zones users
- Starting on September 1st (2020), SSL/TLS certificates cannot be issued for longer than 13 months (397 days).

## Clean Up
- Remove Helm repos
```bash
helm repo remove ingress-nginx
helm repo remove bitnami
```
- remove ca from keychain
```bash
brew remove cfssl
```

## Links
- [K8s Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)
- [NGINX Ingress](https://kubernetes.github.io/ingress-nginx/)
- [NGINX inc version](https://github.com/nginxinc/kubernetes-ingress/)
- [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx)
- [Installation Guide](https://kubernetes.github.io/ingress-nginx/deploy/)
- [Which Ingress Controller Do I Need?](https://docs.nginx.com/nginx-ingress-controller/intro/nginx-ingress-controllers)
- [values.yaml](https://github.com/kubernetes/ingress-nginx/blob/main/charts/ingress-nginx/values.yaml)
- [Configuration options](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#configuration-options)
- [Exposing TCP and UDP services](https://kubernetes.github.io/ingress-nginx/user-guide/exposing-tcp-udp-services/)
- [Admission web hook] (https://kubernetes.github.io/ingress-nginx/how-it-works/#avoiding-outage-from-wrong-configuration)
