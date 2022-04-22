
-include .makefiles/Makefile
-include .makefiles/pkg/protobuf/v2/Makefile
-include .makefiles/pkg/go/v1/Makefile

.makefiles/%:
	@curl -sfL https://makefiles.dev/v1 | bash /dev/stdin "$@"

################################################################################

.PHONY: pong
pong: $(GO_DEBUG_DIR)/pong
	$<

.PHONY: ping
ping: $(GO_DEBUG_DIR)/ping
	$<
