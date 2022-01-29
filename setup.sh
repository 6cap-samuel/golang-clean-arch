alias k=kubectl
k apply -f certManager
k apply -f configs
k apply -f secrets
k apply -f services/db
k apply -f services/public
k apply -f microservices
k apply -f networkPolicies