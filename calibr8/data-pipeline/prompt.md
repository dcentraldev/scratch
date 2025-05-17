Create me an apache beam pipeline in golang, which:
1. Connect to Solana using [Helius LaserStream](https://www.helius.dev/docs/laserstream) gRPC
2. Filter all transactions greater than 100, for following tokens: USDC, Sol, JUP, DRIFT
3. Save the token value in a BigQuery table.  Create the table, if one doesn't exist
4. Calculate token pair price (compared to SOL) on following DEX: Orca, Meteora, Raydium
5. Persist the calculated price in a BigQuery Table

We need to build a MVP with proper documentation.  Include a README.md file.  
Make assumptions if required or seek clarification