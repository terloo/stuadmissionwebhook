FROM alpine:latest

ADD stuadmissionwebhook /stuadmissionwebhook
ENTRYPOINT ["./stuadmissionwebhook"]