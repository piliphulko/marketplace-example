{{- define "cores-api-library.deployment" -}}
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deployment
  namespace: {{ .Values.NAMESPACE }}
spec:
  replicas: {{ .Values.REPLICAS }}
  selector:
    matchLabels:
      app: separate-copy
  template:
    metadata:
      labels:
        app: separate-copy
    spec:
      containers:
      - name: container
        image: {{ .Values.CONTAINER_IMAGE }}
        ports:
        - containerPort: {{ .Values.PORTS.CONTAINER_PORT }}
          protocol: TCP
        env:
        - name: APP_PORT
          value: "{{ .Values.PORTS.CONTAINER_PORT }}"
        resources:
          requests:
            memory: {{ .Values.LIMITES.CONTAINER.DEFAULT.MEM }}
            cpu: {{ .Values.LIMITES.CONTAINER.DEFAULT.CPU }}
          limits:
            memory: {{ .Values.LIMITES.CONTAINER.LIMITS.MEM }}
            cpu: {{ .Values.LIMITES.CONTAINER.LIMITS.CPU }}
        livenessProbe:
          grpc:
            port: {{ .Values.PORTS.CONTAINER_PORT }}
          initialDelaySeconds: {{ .Values.PROBE.LIVENESS.DELAY_SEC }}
          periodSeconds: {{ .Values.PROBE.LIVENESS.PERIOD_SEC }}
          timeoutSeconds: {{ .Values.PROBE.LIVENESS.TIMEOUT_SEC }}
        readinessProbe:
          grpc:
            port: {{ .Values.PORTS.CONTAINER_PORT }}
          initialDelaySeconds: {{ .Values.PROBE.READINESS.DELAY_SEC }}
          periodSeconds: {{ .Values.PROBE.READINESS.PERIOD_SEC }}
          timeoutSeconds: {{ .Values.PROBE.READINESS.TIMEOUT_SEC }}
{{- end }}