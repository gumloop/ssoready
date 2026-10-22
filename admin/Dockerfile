FROM node:22.23.4-trixie@sha256:77d0c7953e80648137a237496c8edbcb3a14bedc5c5298b93a6d0dfa6c3e1271 AS build

WORKDIR /work

COPY package.json package-lock.json ./
RUN npm ci

COPY .env tailwind.config.js build.mjs tsconfig.json ./
COPY bin ./bin
COPY public ./public
COPY src ./src

RUN npm run build-docker

WORKDIR /server

COPY server/package.json server/package-lock.json ./
RUN npm ci

FROM node:22.23.4-trixie-slim@sha256:8eb619e402eaede1748a3c8173bffab391592812e85136452f781741c7242ddf

RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get install -y --no-install-recommends tini \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /usr/local/lib/node_modules /usr/local/bin/npm /usr/local/bin/npx /usr/local/bin/corepack \
       /usr/local/bin/yarn /usr/local/bin/yarnpkg /opt/yarn-*

WORKDIR /work

COPY --from=build /server/node_modules ./node_modules
COPY --from=build /work/bin/prepare-config-json /work/bin/serve-public ./bin/
COPY --from=build /work/src/config.tsx ./src/
COPY --from=build /work/public ./public
RUN touch public/config.json && chown node:node public/config.json

USER node

EXPOSE 8083
# node does not exit on SIGTERM when it is PID 1, so tini forwards the signal
ENTRYPOINT ["tini", "--", "sh", "-c", "./bin/prepare-config-json > ./public/config.json && exec ./bin/serve-public"]
