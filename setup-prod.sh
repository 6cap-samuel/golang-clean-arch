alias k=kubectl

k apply -f namespaces
k apply -f certmanagers/prod
k apply -f configs/prod
k apply -f secrets/prod
k apply -f deployments/prod
k apply -f networkpolicies/prod
k apply -f services/db/prod
k apply -f services/public/prod