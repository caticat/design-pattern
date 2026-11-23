FROM --platform=$BUILDPLATFORM python:3.14-slim AS build
WORKDIR /src
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
COPY mkdocs.yml ./
COPY docs/ docs/
COPY codes/ codes/
COPY hooks/ hooks/
RUN mkdocs build --strict --site-dir /site

FROM nginxinc/nginx-unprivileged:stable-alpine
LABEL org.opencontainers.image.source="https://github.com/caticat/design-pattern"
COPY deploy/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /site/ /usr/share/nginx/html/design-pattern/
EXPOSE 8080
