#![feature(proc_macro_hygiene, decl_macro)]

// import Rocket
#[macro_use] extern crate rocket;

#[get("/")]
fn index() -> &'static str {
    "Hello, welcome to Assorted Jollof"
}

fn main() {
    rocket::ignite().mount("/", routes![index]).launch();
}
