#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use]
extern crate rocket;

#[macro_use]
extern crate rocket_contrib;

use rocket_contrib::json::JsonValue;
use serde_derive::{Deserialize, Serialize};

#[catch(404)]
fn not_found() -> JsonValue {
    json!({
        "status": "error",
        "reason": "Resource was not found."
    })
}

#[get("/")]
fn index() -> &'static str {
    "Hello, welcome to Assorted Jollof"
}

#[derive(Serialize, Deserialize)]
struct User {
    id: u8,
    name: String,
    phone: String,
}

#[get("/users")]
fn all_users() -> JsonValue {
    let users: Vec<User> = vec![
        User {
            id: 1,
            name: String::from("Julia"),
            phone: String::from("+23324234432"),
        },
        User {
            id: 2,
            name: String::from("Jennifer"),
            phone: String::from("+233202341231"),
        },
    ];

    json!(users)
}

fn rocket() -> rocket::Rocket {
    rocket::ignite().mount("/", routes![index, all_users])
}

fn main() {
    rocket().launch();
}
