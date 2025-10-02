FROM envoyproxy/envoy:v1.30-latest

COPY envoy.yaml /etc/envoy/envoy.yaml
# COPY envoyz.yaml /etc/envoy/envoyz.yaml

# CMD ["envoy", "-c", "/etc/envoy/envoyz.yaml"]
CMD ["envoy", "-c", "/etc/envoy/envoy.yaml"]

