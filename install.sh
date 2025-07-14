#!/bin/bash

set -e

VERSION="v0.1.0"
BINARY_NAME="shtp"
BINARY_URL="https://github.com/00lg/shtp/releases/download/$VERSION/${BINARY_NAME}-linux-amd64"
INSTALL_PATH="/usr/local/bin/$BINARY_NAME"
CONFIG_PATH="$HOME/.shtp"

echo "🔧 Instalando $BINARY_NAME $VERSION..."

if ! command -v curl &> /dev/null; then
  echo "❌ 'curl' no está instalado. Instalalo y reintentá."
  exit 1
fi

if ! command -v docker &> /dev/null; then
  echo "🐳 Docker no encontrado. Instalando..."
  curl -fsSL https://get.docker.com | sh
else
  echo "✅ Docker ya está instalado."
fi

echo "⬇️ Descargando binario de $BINARY_NAME..."
curl -L "$BINARY_URL" -o "$BINARY_NAME"
chmod +x "$BINARY_NAME"
sudo mv "$BINARY_NAME" "$INSTALL_PATH"

mkdir -p "$CONFIG_PATH"

echo ""
echo "✅ Instalación completa."
echo "👉 Ejecutá '$BINARY_NAME --help' para comenzar."
