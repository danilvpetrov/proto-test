
DOCKER_REPO ?= ping-pong
DOCKER_PLATFORMS += linux/amd64
DOCKER_PLATFORMS += linux/arm64

-include .makefiles/Makefile
-include .makefiles/pkg/protobuf/v2/Makefile
-include .makefiles/pkg/go/v1/Makefile
-include .makefiles/pkg/docker/v1/Makefile

.makefiles/%:
	@curl -sfL https://makefiles.dev/v1 | bash /dev/stdin "$@"

################################################################################

.PHONY: pong
pong: $(GO_DEBUG_DIR)/pong
	$<

.PHONY: ping
ping: $(GO_DEBUG_DIR)/ping
	$<
