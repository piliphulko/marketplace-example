# syntax=docker/dockerfile:1

FROM vault:1.13.3

# RUN apk update && apk add gnupg

COPY config.hcl start_vault.sh /usr/local/bin/
COPY certificate.crt private.key /usr/local/bin/
# unseal_vault.sh base64_gpgPublicUnsealKey.asc

# ENV VAULT_ADDR http://127.0.0.1:8200

# Устанавливаем права на выполнение для скрипта
RUN chmod +x /usr/local/bin/start_vault.sh 
# /usr/local/bin/unseal_vault.sh

# Запускаем скрипт при старте контейнера
ENTRYPOINT ["/bin/sh", "/usr/local/bin/start_vault.sh"]