#![feature(proc_macro_hygiene)]
#![feature(decl_macro)]
#[macro_use]
extern crate rocket;

mod models;
mod services;
mod controllers;

fn main() {
    rocket::ignite()
        .mount("/", routes![controllers::log_controllers::index])  // ここにルーティングをセットする
        .mount("/log", routes![controllers::log_controllers::post_new_log])  // ここにルーティングをセットする
        .mount("/log", routes![controllers::log_controllers::this_month])  // ここにルーティングをセットする
        .mount("/log", routes![controllers::log_controllers::this_month_per_day])  // ここにルーティングをセットする
        .mount("/log", routes![controllers::log_controllers::this_month_per_category])  // ここにルーティングをセットする
        .mount("/log", routes![controllers::log_controllers::month_per_category])  // ここにルーティングをセットする
        .mount("/category", routes![controllers::category_controllers::get_category_id])  // ここにルーティングをセットする
        .mount("/category", routes![controllers::category_controllers::post_new_category])  // ここにルーティングをセットする
        .launch();
}
