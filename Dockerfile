FROM scratch

ARG TARGETPLATFORM
COPY artifacts/build/release/$TARGETPLATFORM/* /
