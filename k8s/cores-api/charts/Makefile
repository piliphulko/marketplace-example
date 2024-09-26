REGISTRY = localhost:5000
LIB_VERSION = 0.1.0
SAA_VERSION = 0.1.0

helm-push-cores-api-library:

	helm package ./cores-api-library

	helm push cores-api-library-$(LIB_VERSION).tgz oci://$(REGISTRY)/helm-repo

helm-push-service-acct-auth:

	helm dependency update ./service-acct-auth

	helm package ./service-acct-auth

	helm push service-acct-auth-$(SAA_VERSION).tgz oci://$(REGISTRY)/helm-repo