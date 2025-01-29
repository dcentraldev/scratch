use std::str::FromStr;

use config::Config as ConfigBuilder;
use futures_util::StreamExt;
use serde::Deserialize;
use solana_client::{nonblocking::pubsub_client::PubsubClient, rpc_config::RpcTransactionConfig};
use solana_sdk::commitment_config::CommitmentConfig;
use solana_pubkey::Pubkey;
use solana_transaction_status::UiTransactionEncoding;

#[derive(Debug, Deserialize)]
struct Config {
    solana: SolanaConfig,
    transaction: TransactionConfig,
}

#[derive(Debug, Deserialize)]
struct SolanaConfig {
    usdc_mint_address: String,
    ws_url: String,
    rpc_url: String,
}

#[derive(Debug, Deserialize)]
struct TransactionConfig {
    min_transfer_amount: f64,
    max_transactions: u32,
}

impl Config {
    fn load() -> Result<Self, config::ConfigError> {
        let settings = ConfigBuilder::builder()
            .add_source(config::File::with_name("config"))
            .build()?;

        settings.try_deserialize()
    }

    fn usdc_mint_address(&self) -> Result<Pubkey, solana_sdk::pubkey::ParsePubkeyError> {
        Pubkey::from_str(&self.solana.usdc_mint_address)
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {

    // Load configuration
    let config = Config::load()?;

    let pubsub_client = PubsubClient::new(&config.solana.ws_url).await?;

    let (mut subscription, _) = pubsub_client.logs_subscribe(
        solana_client::rpc_config::RpcTransactionLogsFilter::Mentions(vec![config.usdc_mint_address()?.to_string()]),
        solana_client::rpc_config::RpcTransactionLogsConfig {
            commitment: Some(solana_sdk::commitment_config::CommitmentConfig {
                commitment: solana_sdk::commitment_config::CommitmentLevel::Confirmed,
            }),
        }
    ).await?;

    let mut counter = 0;
    while let Some(notification) = subscription.next().await {
        // if notification.method == "logsNotification" {
        let signature = notification.value.signature;
        println!("Getting Transaction: {}", signature);
        let rpc_client = solana_client::rpc_client::RpcClient::new(config.solana.rpc_url.to_string());
        let txn_config = RpcTransactionConfig {
            encoding: Some(UiTransactionEncoding::Json),
            commitment: Some(CommitmentConfig::confirmed()),
            max_supported_transaction_version: Some(0),
        };
        
        let txn = rpc_client
        .get_transaction_with_config(&signature.parse().unwrap(), txn_config)
            // .get_transaction(&signature.parse().unwrap(), solana_transaction_status::UiTransactionEncoding::JsonParsed)
            // .await
            .expect("Failed to fetch transaction");
        if let Some(meta) = txn.transaction.meta {
            let (pre_token_balances, post_token_balances) = (meta.pre_token_balances.unwrap(), meta.post_token_balances.unwrap());
            for (pre, post) in pre_token_balances.iter().zip(post_token_balances.iter()) {
                if pre.mint == config.usdc_mint_address()?.to_string() {
                    let amount_transferred = post.ui_token_amount.ui_amount.unwrap_or(0.0) - pre.ui_token_amount.ui_amount.unwrap_or(0.0);
                    // Check if the transfer amount is 50 USDC
                    if amount_transferred.abs() >= config.transaction.min_transfer_amount {
                        println!("{} USDC Transfer Detected in Transaction: {}", config.transaction.min_transfer_amount, signature);
                        println!("From: {}", pre.owner.clone().unwrap());
                        println!("To: {}", post.owner.clone().unwrap());
                        println!("Amount: {} USDC", amount_transferred);
                        counter += 1;
                        if counter > config.transaction.max_transactions {
                            break;
                        }
                    }
                }
            }
        }
    }

    drop(subscription);
    pubsub_client.shutdown().await?;

    // unsubscribe.await;

    Ok(())
}