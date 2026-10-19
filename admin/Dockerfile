FROM node:22.23.3-bookworm@sha256:0e5f906573693feaa1e21057ebdcfdb5bd5021f050b2dc7c9deceb629c7da2a8 AS build

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

FROM node:22.23.3-bookworm-slim@sha256:c3de60bf2f9dd0ac6370e6117950ff62d6e339527e7472301c9c78a017978392

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
