ARG RUST_VERSION=1.78.0
ARG APP_NAME=ntumiwa-site

# Prepare a base image (chef) with cargo-chef installed

FROM clux/muslrust:stable AS chef
USER root
RUN rustup target add x86_64-unknown-linux-musl
RUN cargo install cargo-chef
WORKDIR /app

################################################################################

# Have the chef prepare a list of the project's dependencies

FROM chef AS planner
COPY . .
RUN cargo chef prepare --recipe-path recipe.json

################################################################################

# Ask the sous-chef to build our dependencies with the chef's recipe.json

FROM chef AS sous-chef
COPY --from=planner /app/recipe.json recipe.json
RUN cargo chef cook --release --target x86_64-unknown-linux-musl --recipe-path recipe.json

################################################################################

# Build the application for the target architecture

FROM chef AS builder
COPY --from=sous-chef app/target /target
COPY . .
RUN cargo build --release --target x86_64-unknown-linux-musl --bin ntumiwa-site

################################################################################

# Prepare the final image

FROM alpine:3.18 AS final
ARG APP_NAME

# Create a non-privileged user that the app will run under
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

# Copy over the binary
COPY --from=builder /app/target/x86_64-unknown-linux-musl/release/${APP_NAME} /usr/local/bin/${APP_NAME}

# Copy over the static directory
COPY ./static ./static

EXPOSE 8080
ENTRYPOINT [ "/usr/local/bin/ntumiwa-site" ]
