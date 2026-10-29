FROM mcr.microsoft.com/devcontainers/go:2-1.27-trixie

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends file

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
RUN go install github.com/bootdotdev/bootdev@latest
