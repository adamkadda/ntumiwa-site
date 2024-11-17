use anyhow::Result;
use reqwest;

#[tokio::test]
async fn quick_dev() -> Result<()> {
    let client = reqwest::Client::new();
    let response = client.get("http://0.0.0.0:8080/")
        .send()
        .await?;

    println!("{}", response.status());
    println!("{}", response.text().await?);

    Ok(())
}