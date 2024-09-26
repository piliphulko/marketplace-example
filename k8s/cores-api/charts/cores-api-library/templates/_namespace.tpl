{{- define "cores-api-library.namespace" -}}
apiVersion: v1
kind: Namespace
metadata:
  name: {{ .Values.NAMESPACE }}
{{- end }}