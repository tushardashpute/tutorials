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

- Create `values.yaml` to override default [parameters]()

- Generate YAML from the Helm chart
```bash
helm template my-ing ingress-nginx/ingress-nginx \
  --namespace ingress \
  --version 3.35.0 \
  --values values.yaml \
  --output-dir nginx
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
kubectl apply -f prometheus/2-prometheus/4-ingress.yaml
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
kubectl apply -f prometheus/2-prometheus/4-ingress.yaml
```
- Make a mistake in the directive and apply

## Create Ingress for Grafana
- Get services
```bash
kubectl get svc -n monitoring
```

- Create ingress
```bash
kubectl apply -f grafana/3-ingress.yaml
```

- Get ingresses
```bash
kubectl get ing -n monitoring
```

- Create CNAME record for `grafana.devopsbyexample.io`

- Go to `http://grafana.devopsbyexample.io`

## Simple Fanout Ingress Example

- Create `app`

- Create `example-1`





















7. Simple fanout example
Name based virtual hosting
8. Fanout in different namespaces example
9. Virual hosting
10. tls
11. tcp
12. Prometheus/Grafana monitoring

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

```
eksctl create cluster -f eks.yaml
```

## YAML
## HELM
- Add Nginx Helm repository
```bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm -n ingress template --version 3.35.0 ingress-nginx ingress-nginx/ingress-nginx -f values.yaml --output-dir ingress
helm -n ingress install --version 3.35.0 ingress-nginx ingress-nginx/ingress-nginx -f values.yaml
```

kubectl get ingressclass

https://github.com/cloudflare/cfssl

https://github.com/kelseyhightower/kubernetes-the-hard-way/blob/master/docs/04-certificate-authority.md

brew install cfssl

## Prometheus
kubectl port-forward svc/prometheus-operated 9090 -n monitoring

curl http://api.devopsbyexample.io/foo
curl http://api.devopsbyexample.io/bar

curl http://foo.devopsbyexample.io/
curl http://bar.devopsbyexample.io/

## Clean Up
- Remove Helm repos
```bash
helm repo remove ingress-nginx
helm repo remove bitnami
```
- remove ca from keychain

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
Admission web hook - https://kubernetes.github.io/ingress-nginx/how-it-works/#avoiding-outage-from-wrong-configuration


Generate the CA configuration file, certificate, and private key:
-config=""

cfssl gencert -initca ca-csr.json | cfssljson -bare ca
openssl x509 -in ca.pem -text -noout

CN: CommonName
OU: OrganizationalUnit
O: Organization
L: Locality
S: StateOrProvinceName
C: CountryName


cfssl gencert \
  -ca=ca.pem \
  -ca-key=ca-key.pem \
  -config=config.json \
  -profile=demo \
  foo-api-csr.json | cfssljson -bare foo-api

openssl x509 -in foo-api.pem -text -noout

Default backend:  default-http-backend:80 (<error: endpoints "default-http-backend" not found>)
Starting on September 1st (2020), SSL/TLS certificates cannot be issued for longer than 13 months (397 days).

GRAFANA 9614
kubectl get ns --show-labels

webhook? works?
http2 works??

curl http://api-ns.devopsbyexample.io/foo/asd

helm repo add bitnami https://charts.bitnami.com/bitnami
helm template -n database --version 10.9.3 postgres bitnami/postgresql --output-dir postgres-helm

psql --host localhost --port 5432 --username postgres --password

- --tcp-services-configmap=$(POD_NAMESPACE)/tcp-services

- name: postgres
  port: 5432
  protocol: TCP

psql --host a5919b05642104f4896d2dd0d08c8407-347fbf2a7b44f1ce.elb.us-east-1.amazonaws.com --port 5432 --username postgres --password