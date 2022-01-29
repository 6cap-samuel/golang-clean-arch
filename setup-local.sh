alias k=kubectl

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.1.1/deploy/static/provider/cloud/deploy.yaml

k apply -f configs
k apply -f secrets
k apply -f services/db
k apply -f services/public
k apply -f microservices
k apply -f networkPolicies