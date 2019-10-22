FROM registry.cn-hangzhou.aliyuncs.com/dmall/alpine-base:v3.10
MAINTAINER xuekui <kui.xue@dmall.com>

COPY bin/manager .
USER nonroot:nonroot

ENTRYPOINT ["/manager"]
