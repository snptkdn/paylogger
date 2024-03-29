use anyhow::anyhow;
use anyhow::Result;
use controllers::purchaselog_controller::PurchaseLogController;
use dotenv::dotenv;
use serenity::async_trait;
use std::env;

use serenity::builder::CreateApplicationCommandOption;
use serenity::model::application::interaction::InteractionResponseType;
use serenity::model::gateway::Ready;
use serenity::model::prelude::application_command::ApplicationCommandOption;
use serenity::model::prelude::interaction::Interaction::ApplicationCommand;
use serenity::model::prelude::*;
use serenity::prelude::*;
use services::category_service;
use shuttle_secrets::SecretStore;

mod constraits;
mod controllers;
mod models;
mod services;

use controllers::category_controller::CategoryController;

struct Bot;

#[async_trait]
impl EventHandler for Bot {
    async fn ready(&self, ctx: Context, ready: Ready) {
        let guild_id = GuildId(1044972423587561504);

        let categories = category_service::CategoryService::get_categories()
            .await
            .unwrap();

        let commands = GuildId::set_application_commands(&guild_id, &ctx.http, |commands| {
            commands
                .create_application_command(|command| {
                    command
                        .name("new_category")
                        .description("add new log category.")
                        .create_option(|option| {
                            option
                                .name("name")
                                .kind(command::CommandOptionType::String)
                                .description("category's name.")
                                .required(true)
                        })
                })
                .create_application_command(|command| {
                    command
                        .name("get_categories")
                        .description("get all categories.")
                })
                .create_application_command(|command| {
                    command
                        .name("new_log")
                        .description("add new log.")
                        .create_option(|option| {
                            option
                                .name("price")
                                .kind(command::CommandOptionType::Integer)
                                .description("used money.")
                                .required(true)
                        })
                        .create_option(|option| {
                            categories.into_iter().fold(
                                option
                                    .name("category")
                                    .kind(command::CommandOptionType::Integer)
                                    .description("log's category.")
                                    .required(true),
                                |option, category| {
                                    option.add_string_choice(category.name, category.id)
                                },
                            )
                        })
                        .create_option(|option| {
                            option
                                .name("date")
                                .kind(command::CommandOptionType::Integer)
                                .description("date when used money. default is today.")
                        })
                })
        })
        .await
        .unwrap();
    }

    async fn interaction_create(
        &self,
        ctx: Context,
        interaction: serenity::model::application::interaction::Interaction,
    ) {
        if let ApplicationCommand(command) = interaction {
            let result = match command.data.name.as_str() {
                "new_category" => {
                    CategoryController::add_category(
                        command.data.options[0].value.clone().unwrap().to_string(),
                    )
                    .await
                }
                "get_categories" => Ok(CategoryController::get_categories()
                    .await
                    .unwrap()
                    .iter()
                    .fold("".to_string(), |acc, category| {
                        (acc + &category.name + "\n").to_string()
                    })
                    .to_string()),
                "new_log" => {
                    PurchaseLogController::add_log_purchase(
                        command.data.options[0] // price
                            .value
                            .clone()
                            .unwrap()
                            .as_i64()
                            .unwrap(),
                        command.data.options[1] // category
                            .value
                            .clone()
                            .unwrap()
                            .as_u64()
                            .unwrap(),
                        match command.data.options.len() > 2 {
                            // date
                            true => Some(
                                command.data.options[2]
                                    .value
                                    .clone()
                                    .unwrap()
                                    .as_str()
                                    .unwrap()
                                    .to_string(),
                            ),
                            false => None,
                        },
                    )
                    .await
                }

                command => unreachable!("Unknown command: {}", command),
            };

            let create_interaction_response =
                command.create_interaction_response(&ctx.http, |response| {
                    response
                        .kind(InteractionResponseType::ChannelMessageWithSource)
                        .interaction_response_data(|message| {
                            message.content(match result {
                                Ok(message) => message,
                                Err(e) => e.to_string(),
                            })
                        })
                });

            if let Err(why) = create_interaction_response.await {
                eprintln!("Cannot respond to slash command: {}", why);
            }
        }
    }
}

#[tokio::main]
async fn main() -> Result<()> {
    dotenv().ok();
    // Get the discord token set in `Secrets.toml`
    let token = if let token = env::var("DISCORD_TOKEN").unwrap() {
        token
    } else {
        return Err(anyhow!("'DISCORD_TOKEN' was not found").into());
    };

    // Set gateway intents, which decides what events the bot will be notified about
    let intents = GatewayIntents::GUILD_MESSAGES | GatewayIntents::MESSAGE_CONTENT;

    let mut client = Client::builder(&token, intents)
        .event_handler(Bot)
        .await
        .expect("Err creating client");

    client.start().await;

    Ok(())
}
