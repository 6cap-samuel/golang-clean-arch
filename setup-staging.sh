alias k=kubectl

k apply -f namespaces
k apply -f certmanagers/staging
k apply -f configs/staging
k apply -f secrets/staging
k apply -f deployments/staging
k apply -f networkpolicies/staging
k apply -f services/db/staging
k apply -f services/public/staging