use askama::Template;

#[derive(Template)]
#[template(path = "home.html")]
pub(super) struct HomeTemplate;

#[derive(Template)]
#[template(path = "biography.html")]
pub(super) struct BiographyTemplate;

#[derive(Template)]
#[template(path = "performances.html")]
pub(super) struct PerformancesTemplate;

#[derive(Template)]
#[template(path = "photos.html")]
pub(super) struct PhotosTemplate;

#[derive(Template)]
#[template(path = "videos.html")]
pub(super) struct VideosTemplate;

#[derive(Template)]
#[template(path = "contact.html")]
pub(super) struct ContactTemplate;