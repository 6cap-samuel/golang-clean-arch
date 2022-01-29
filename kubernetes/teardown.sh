alias k=kubectl
k delete -f microservices
k delete -f configs
k delete -f secrets
k delete -f services/db
k delete -f services/public
k delete -f networkPolicies