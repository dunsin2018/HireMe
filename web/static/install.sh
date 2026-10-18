#!/bin/sh
# Install the hireme CLI: download the prebuilt binary for this OS/arch from the
# latest GitHub release and place it on PATH.
#   curl -fsSL https://hiremetoday.ca/install.sh | sh
#
# Served at https://hiremetoday.ca/install.sh as a SvelteKit static asset — it must
# live in web/static/, the only directory adapter-node serves. Kept in sync with
# the canonical script in the hireme-cli repo.
set -eu

REPO="dunsin2018/HireMe-cli"
BIN="hireme"
INSTALL_DIR="${HIREME_INSTALL_DIR:-/usr/local/bin}"

os=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$os" in
  linux | darwin) ;;
  *) echo "hireme: unsupported OS '$os' (use 'go install $REPO/cmd/$BIN@latest')" >&2; exit 1 ;;
esac

arch=$(uname -m)
case "$arch" in
  x86_64 | amd64) arch=amd64 ;;
  arm64 | aarch64) arch=arm64 ;;
  *) echo "hireme: unsupported arch '$arch' (use 'go install $REPO/cmd/$BIN@latest')" >&2; exit 1 ;;
esac

asset="${BIN}_${os}_${arch}"
url="https://github.com/${REPO}/releases/latest/download/${asset}"

tmp=$(mktemp)
trap 'rm -f "$tmp"' EXIT
echo "Downloading ${asset} …"
curl -fsSL "$url" -o "$tmp"
chmod +x "$tmp"

target="${INSTALL_DIR}/${BIN}"
if [ -w "$INSTALL_DIR" ]; then
  mv "$tmp" "$target"
else
  echo "Installing to ${target} (requires sudo) …"
  sudo mv "$tmp" "$target"
fi
trap - EXIT

echo "Installed ${BIN} → ${target}"
echo "Next: ${BIN} auth login --token fhk_…    (create a key at https://hiremetoday.ca → API keys)"
