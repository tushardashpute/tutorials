# NGINX Ingress Controller for Kubernetes Tutorial: YAML & Helm | Example | Prometheus | Grafana | EKS

[YouTube Tutorial]()

Topics
- rewrite-target
- An Ingress allows you to keep the number of load balancers down to a minimum
- [Simple fanout](https://kubernetes.io/docs/concepts/services-networking/ingress/#simple-fanout)
- [Name based virtual hosting](https://kubernetes.io/docs/concepts/services-networking/ingress/#name-based-virtual-hosting)
- [TLS](https://kubernetes.io/docs/concepts/services-networking/ingress/#tls) great segue to cert-manager



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

## Clean Up
- `helm repo remove nginx-stable`

## Links
- [K8s Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)
- [NGINX Ingress](https://kubernetes.github.io/ingress-nginx/)
- [NGINX inc version](https://github.com/nginxinc/kubernetes-ingress/)
- [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx)
- [Installation Guide](https://kubernetes.github.io/ingress-nginx/deploy/)
- [Which Ingress Controller Do I Need?](https://docs.nginx.com/nginx-ingress-controller/intro/nginx-ingress-controllers)
- [values.yaml](https://github.com/kubernetes/ingress-nginx/blob/main/charts/ingress-nginx/values.yaml)
- [Configuration options](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#configuration-options)
