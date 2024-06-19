# checkov:skip=CKV_DOCKER_7: "Ensure the base image uses a non latest version tag"
# checkov:skip=CKV_DOCKER_2: "Ensure that HEALTHCHECK instructions have been added to container images"
FROM cgr.dev/chainguard/static:latest

USER nobody

COPY go-mail-api /usr/local/bin/go-mail-api

ENTRYPOINT ["go-mail-api"]