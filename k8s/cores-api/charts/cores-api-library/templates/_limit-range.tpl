{{- define "cores-api-library.limit-range" -}}
apiVersion: v1
kind: LimitRange
metadata:
  name: limit-range
  namespace: {{ .Values.NAMESPACE }}
spec:
  limits:
  - default:
      cpu: {{ .Values.LIMITES.SPACE.LIMITS.CPU }}
      memory: {{ .Values.LIMITES.SPACE.LIMITS.MEM }}
    defaultRequest:
      cpu: {{ .Values.LIMITES.SPACE.DEFAULT.CPU }}
      memory: {{ .Values.LIMITES.SPACE.DEFAULT.MEM }}
    type: Container
{{- end }}