use axum::{response::{Html, IntoResponse}, routing::get, Router};
use askama::Template;

use crate::web::templates;

pub fn router() -> Router {
    Router::new()
        .route("/", get(home))
        .route("/biography", get(biography))
        .route("/performances", get(performances))
        .route("/photos", get(photos))
        .route("/videos", get(videos))
        .route("/contact", get(contact))
}

async fn home() -> impl IntoResponse {
    println!("->> {:<12} - home", "HANDLER");
    let template = templates::HomeTemplate;
    Html(template.render().unwrap())
}

async fn biography() -> impl IntoResponse {
    println!("->> {:<12} - biography", "HANDLER");
    let template = templates::BiographyTemplate;
    Html(template.render().unwrap())
}

async fn performances() -> impl IntoResponse {
    println!("->> {:<12} - performances", "HANDLER");
    let template = templates::PerformancesTemplate;
    Html(template.render().unwrap())
}

async fn photos() -> impl IntoResponse {
    println!("->> {:<12} - photos", "HANDLER");
    let template = templates::PhotosTemplate;
    Html(template.render().unwrap())
}

async fn videos() -> impl IntoResponse {
    println!("->> {:<12} - videos", "HANDLER");
    let template = templates::VideosTemplate;
    Html(template.render().unwrap())
}

async fn contact() -> impl IntoResponse {
    println!("->> {:<12} - contact", "HANDLER");
    let template = templates::ContactTemplate;
    Html(template.render().unwrap())
}