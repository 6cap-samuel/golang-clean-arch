alias k=kubectl
k delete -f microservices
k delete -f configs
k delete -f secrets
k delete -f services/db
k delete -f services/public
k delete -f networkPolicies
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.1.1/deploy/static/provider/cloud/deploy.yaml
