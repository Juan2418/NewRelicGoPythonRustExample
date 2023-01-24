#[macro_use]
extern crate rocket;
use rocket::{
    serde::json::Json,
    serde::uuid::Uuid,
    serde::{Deserialize, Serialize},
};

#[derive(Serialize, Deserialize)]
#[serde(crate = "rocket::serde")]
struct GoResponse {
    message: String,
    requestId: Option<String>,
}

#[get("/")]
fn index() -> Json<GoResponse> {
    let request_id = Uuid::new_v4().to_string();
    let example_return = GoResponse {
        message: "Hello, world!".to_string(),
        requestId: Some(request_id),
    };

    Json(example_return)
}

#[get("/trace")]
async fn trace() -> Json<GoResponse> {
    let client = reqwest::Client::new();

    let go_request = client
        .post("http://localhost:80/reference")
        .send()
        .await
        .unwrap()
        .json::<GoResponse>()
        .await
        .unwrap();

    let example_return = GoResponse {
        message: go_request.message,
        requestId: go_request.requestId,
    };

    Json(example_return)
}

#[launch]
#[tokio::main]
async fn rocket() -> _ {
    rocket::build().mount("/", routes![index, trace])
}
