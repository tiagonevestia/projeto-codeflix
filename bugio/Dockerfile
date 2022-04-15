FROM node:14.15.4-slim

RUN useradd -r -u 1002 node
USER 1002

WORKDIR /home/node/app

CMD [ "sh", "-c", "npm intall && tail -f /dev/null" ]