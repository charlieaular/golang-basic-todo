
ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/devcontainers/go:0-${VARIANT}

ARG NODE_VERSION="lts/*"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c ". /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# RUN go mod init 
# RUN go mod tidy

# RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/codegangsta/gin

