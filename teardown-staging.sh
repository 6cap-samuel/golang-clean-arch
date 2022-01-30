alias k=kubectl

k delete -f certmanagers/staging
k delete -f configs/staging
k delete -f secrets/staging
k delete -f deployments/staging
k delete -f networkpolicies/staging
k delete -f services/db/staging
k delete -f services/public/staging