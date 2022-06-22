
. "$(git rev-parse --show-toplevel || echo ".")/scripts/common.sh"

cd "$PROJECT_DIR" || exit 1

if ! has go; then
  echo_error "Golang binary not found. Please install golang before continue"
  exit 1
fi

if ! has docker; then
  echo_warn "Docker not found. Please install or start it before running db"
fi

# create bin folder to store downloaded tools and compiled binaries
mkdir -p bin/

if ! has ./bin/migrate; then
  echo_info "Install golang-migrate for database versioning"
  version="v4.15.2"
  if ! has apt-get; then
    platform="darwin"
    curl -LO https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz  | tar xvz 
    mv migrate.$platform-amd64.tar.gz bin
    tar -xf ./bin/migrate.$platform-amd64.tar.gz -C ./bin
    rm -rf ./bin/migrate.$platform-amd64.tar.gz 
  elif has apt; then
    platform="linux"
    curl -o ./bin/migrate.tar.gz -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz  | tar xvz
    mv migrate.$platform-amd64.tar.gz bin
    tar -xf ./bin/migrate.$platform-amd64.tar.gz -C ./bin
    rm -rf ./bin/migrate.$platform-amd64.tar.gz 
  fi

  chmod +x ./bin/migrate
fi
