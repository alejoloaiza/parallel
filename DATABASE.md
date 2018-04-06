# IRC WEBSCRAPING BOT - POSTGRESQL DATABASE SETUP

## Steps

1. Install postgreSQL 9.6: Follow the steps from here: https://www.postgresql.org/download/linux/ubuntu/
2. After installation is done and postgreSQL is running, execute the following commands:

```
sudo su - postgres
psql
CREATE SCHEMA "yourschema";
CREATE USER youruser PASSWORD 'yourpassword';
GRANT ALL ON SCHEMA yourschema TO youruser;

```
3. Put the schema, user and password in config.json file
4. Create the table structure using this command:

```

```