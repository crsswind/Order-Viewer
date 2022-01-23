#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
    CREATE TABLE customer_companies(	
        company_id int PRIMARY KEY NOT NULL,
        company_name varchar(40) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS customers(
        user_id varchar(40) PRIMARY KEY NOT NULL,
        login varchar(40) NOT NULL,
        "password" varchar(40) NOT NULL,
        "name" varchar(40) NOT NULL,
        company_id int NOT NULL,
        credit_cards text,	
        CONSTRAINT fk_company
        FOREIGN KEY(company_id) 
        REFERENCES customer_companies(company_id)
    );

    CREATE TABLE IF NOT EXISTS orders(
        id int PRIMARY KEY NOT NULL,
        created_at date NOT NULL,
        order_name varchar(40),
        customer_id varchar(40) NOT NULL,	
        CONSTRAINT fk_customers
        FOREIGN KEY(customer_id) 
        REFERENCES customers(user_id)
    );

    CREATE TABLE IF NOT EXISTS order_items(
        id int PRIMARY KEY NOT NULL,
        order_id int NOT NULL,
        price_per_unit decimal(20,6),
        quantity decimal(20,6) NOT NULL,
        product varchar(40),
        CONSTRAINT fk_orders
        FOREIGN KEY(order_id) 
        REFERENCES orders(id)
    );

    CREATE TABLE IF NOT EXISTS deliveries(
        id int PRIMARY KEY NOT NULL,
        order_item_id int NOT NULL,	
        delivered_quantity decimal(20,6) NOT NULL,
        CONSTRAINT fk_deliveries
        FOREIGN KEY(order_item_id) 
        REFERENCES order_items(id)
    );

  COMMIT;
EOSQL