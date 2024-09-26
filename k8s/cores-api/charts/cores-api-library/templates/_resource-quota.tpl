{{- define "cores-api-library.resource-quota" -}}
apiVersion: v1
kind: ResourceQuota
metadata:
  name: resource-quota
  namespace: {{ .Values.NAMESPACE }}
spec:
  hard:
    pods: {{ .Values.LIMITES.SPACE.LIMITS_PODS }}
{{- end }}