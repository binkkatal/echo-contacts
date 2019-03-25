[[ ! -d bin ]] && mkdir bin

go build \
    -ldflags "-X main.BuildTime=$(date -u +%Y%m%d.%H%M%S)" \
    -o bin/echo-contacts \
    github.com/binkkatal/echo-contacts