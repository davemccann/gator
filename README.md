# Gator

Gator is a lightweight command-line tool designed to allow users to follow and manage blog posts from various sources. The tool stores blog post details by scraping RSS feeds and storing it in a PostgreSQL database allowing for offline use. It supports multiple users, each user maintaining their own or shared set of subscriptions.

The project was created as a simple way to explore the Go programming language providing some simpler use cases ranging from CLI, simple web scraping, database and migration support with basic user management.

## Requirements

The project requires the Go Programming language (1.22 or later) and PostgreSQL.

- [Go](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

Install gator with the command: `go install github.com/davemccann/gator@latest`

## Running

All commands are in the following format: `gator [COMMAND] [ARGUMENTS]`

### Commands

| Command     | Usage                                                            | Description                                                           |
| ----------  | --------                                                         | -------------                                                         |
| `reset`     | `gator reset`                                                    |                                                                       |
| `register`  | `gator register John`                                            | Registers a new user                                                  |
| `login`     | `gator login John`                                               | Changes the active user.                                              |
| `listusers` | `gator listusers`                                                | Lists all users registered.                                           |
| `addfeed`   | `gator addfeed "Hacker News" "https://news.ycombinator.com/rss"` | Adds a new feed and subscribes to current active user                 |
| `listfeeds` | `gator listfeeds`                                                | Lists all feeds that have been added.                                 |
| `follow`    | `gator follow "https://news.ycombinator.com/rss"`                | Subscribes current user to a feed                                     |
| `unfollow`  | `gator unfollow "https://news.ycombinator.com/rss"`              | Unsubscribes current user from a feed                                 |
| `following` | `gator following`                                                | Displays the feeds that the current user is subscribed to.            |
| `agg`       | `gator agg "1m"`                                                 | Scrapes the feeds for new posts at provided interval                  |
| `browse`    | `gator browse 4`  (optional param - default: 2)                  | Displays an optional number of posts from the users subscribed feeds  |