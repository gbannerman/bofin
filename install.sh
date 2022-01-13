#!/usr/bin/env bash

# TODO
# .env template

latest_version() {
  echo "v1.0.0"
}

install_dir() {
  echo "$HOME"
}

binary_from_github() {
    local BOFIN_VERSION
    BOFIN_VERSION="$(latest_version)"

    local BOFIN_BINARY_URL
    BOFIN_BINARY_URL="https://github.com/gbannerman/bofin/releases/download/${BOFIN_VERSION}/bofin-amd64"

    echo "$BOFIN_BINARY_URL"
}

example_config_from_github() {
    echo "https://raw.githubusercontent.com/gbannerman/bofin/main/.bofin-conf.example.yaml"
}

download() {
    curl -L "$@"
}

check_for_config_file() {
    local CONFIG_FILE
    CONFIG_FILE="$HOME/.bofin-conf.yaml"

    if test -f "$CONFIG_FILE"; then
        echo ".bofin-conf.yaml already exists."
    else 
        download $(example_config_from_github) -o "$HOME/.bofin-conf.yaml"
        echo ".bofin-conf.yaml created."
    fi
}

main() {
    local INSTALL_DIR
    INSTALL_DIR="$(install_dir)"

    download $(binary_from_github) -o "$INSTALL_DIR/bofin"

    chmod a+x "$INSTALL_DIR/bofin"

    sudo mv "$INSTALL_DIR/bofin" /usr/local/bin/bofin

    check_for_config_file

    echo "Bofin installed successfully!"
}

main
