#![allow(unused)]

use axum::{routing::get_service, Router};
use tower_http::services::ServeDir;

mod web;

#[tokio::main]
async fn main() {
    let app = Router::new()
        .merge(web::routes::router())
        .fallback_service(routes_static());

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8080").await.unwrap();
    println!("->> LISTENING on {:?}", listener.local_addr().unwrap());
    axum::serve(listener, app).await.unwrap();
}

fn routes_static() -> Router {
    Router::new().nest_service("/", get_service(ServeDir::new("./")))
}