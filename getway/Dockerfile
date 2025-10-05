FROM envoyproxy/envoy:v1.30-latest

RUN mkdir -p /etc/envoy

COPY . .

CMD ["envoy", "-c", "/etc/envoy/envoy.yaml"]

