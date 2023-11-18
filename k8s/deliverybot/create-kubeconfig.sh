# Update these to match your environment
SERVICE_ACCOUNT_NAME=deliverybot-service-account
CONTEXT=$(kubectl config current-context)
NAMESPACE=cicd

NEW_CONTEXT=deliverybot
KUBECONFIG_FILE="kubeconfig-sa"


SECRET_NAME=$(kubectl get serviceaccount ${SERVICE_ACCOUNT_NAME} \
  --context ${CONTEXT} \
  --namespace ${NAMESPACE} \
  -o jsonpath='{.secrets[0].name}')
#TOKEN_DATA=$(kubectl get secret ${SECRET_NAME} \
#  --context ${CONTEXT} \
#  --namespace ${NAMESPACE} \
#  -o jsonpath='{.data.token}')

TOKEN_DATA="eyJhbGciOiJSUzI1NiIsImtpZCI6Imd2Qm52dF8zbzl6YXpSczBHY1IyVVdRUVptUnB6ZnA1U21Od3hCa3JtMEEifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJjaWNkIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImRlbGl2ZXJ5Ym90LXNlY3JldCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJkZWxpdmVyeWJvdC1zZXJ2aWNlLWFjY291bnQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiI4ZTY3NGMwZi02ZjFjLTRiNTctOWJhOS01ODQ1NjJlOGMwMzIiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6Y2ljZDpkZWxpdmVyeWJvdC1zZXJ2aWNlLWFjY291bnQifQ.tjRVq0GxaAsFqSr6l-flujkBRKPtZPfgD-FpBnf_MDWi3t_6M9HE_Y-ChaDH64ymQMt8PZ9oY102_HcMm8X3aREG1XukCKHZUJPXL6YEdkIk_xzLzcHHoiEBSsp8HjNUL_i53tMQBPHWRKHNw0ZJLWP2xJdH6TTLiilB2z3VFSUQSvf0MJjD5QbxF_fKC86UzN1HdG-wpAbsiEI7sFH43WWUCcWHjaq0ELGREbmUwi394SoEoWkSgX4LSXDNJqUIR4YxvDps_4QRGOqFOSZsbS3e_7XsU0kOSVIiTYMNrWX3NUD6vvXJHz8RPSvYz4B5bgUWC1qDaHUSkTgYt85xwQ"

TOKEN=$(echo ${TOKEN_DATA} | base64 -d)

# Create dedicated kubeconfig
# Create a full copy
kubectl config view --raw > ${KUBECONFIG_FILE}.full.tmp
# Switch working context to correct context
kubectl --kubeconfig ${KUBECONFIG_FILE}.full.tmp config use-context ${CONTEXT}
# Minify
kubectl --kubeconfig ${KUBECONFIG_FILE}.full.tmp \
  config view --flatten --minify > ${KUBECONFIG_FILE}.tmp
# Rename context
kubectl config --kubeconfig ${KUBECONFIG_FILE}.tmp \
  rename-context ${CONTEXT} ${NEW_CONTEXT}
# Create token user
kubectl config --kubeconfig ${KUBECONFIG_FILE}.tmp \
  set-credentials ${CONTEXT}-${NAMESPACE}-token-user \
  --token ${TOKEN}
# Set context to use token user
kubectl config --kubeconfig ${KUBECONFIG_FILE}.tmp \
  set-context ${NEW_CONTEXT} --user ${CONTEXT}-${NAMESPACE}-token-user
# Set context to correct namespace
kubectl config --kubeconfig ${KUBECONFIG_FILE}.tmp \
  set-context ${NEW_CONTEXT} --namespace ${NAMESPACE}
# Flatten/minify kubeconfig
kubectl config --kubeconfig ${KUBECONFIG_FILE}.tmp \
  view --flatten --minify > ${KUBECONFIG_FILE}
# Remove tmp
rm ${KUBECONFIG_FILE}.full.tmp
rm ${KUBECONFIG_FILE}.tmp
