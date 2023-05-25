FROM mcr.microsoft.com/devcontainers/go:0-1.19-bullseye

# install taskfile
COPY scripts/install-taskfile.sh .
RUN ./install-taskfile.sh
RUN rm ./install-taskfile.sh

# install buffbuild
COPY scripts/install-bufbuild.sh .
RUN ./install-bufbuild.sh
RUN rm ./install-bufbuild.sh

USER vscode

RUN sed -i "s/#alias ll='ls -l'/alias ll='ls -al --color=auto'/g" /home/vscode/.bashrc

