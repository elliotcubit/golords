
# GolordsBot

A discord bot that has completely random features, but it's mine :)

# Commands

- !8ball - ask 8 ball a question
- !roll {}d{}[+{}d{}]* - roll however many dice
- !poll - start a vote with reactions
- !lookup - dnd 5e srd information lookup, which is a little broken
- !addquote & !getquote - fun quote storage. State is shared between servers, which will change soon
- ``@someone ++`` \ `@someone ++12` \ `@someone --` - give someone karma
- !stacks @someone - check someone's karma
- (Karma is also shared between servers, which will change soon)
- !eqn {LaTeX} - Quick view any latex as a transparent image

At the moment a lot of things are hard-coded, so, for example, bean counting only works in my server.

This will change in the future, so the counting channel can be configured by server admins.

# How to run

You will need:
- A Postgres server
- A discord bot account

Set the env vars:
- `DISCORD_ID`
- `DISCORD_SECRET`
- `DISCORD_TOKEN`
- `YUGABYTE_IP`
- `YUGABYTE_USER`
- `YUGABYTE_PASS`
- `YUGABYTE_DB_NAME`

The database in your Postgres will need the tables:
stacks(
  serverID text,
  userID text,
  amount int
)
beans(
  serverID text,
  userID text,
  amount int
)
quote(
  serverID text,
  userID text,
  quote text,
  timestamp text
)


With these variables set and Postgres actually existing,

`go run .`

To run the bot locally with your own account.
