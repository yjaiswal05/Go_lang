# Slack Bot: Age Calculator

This Slack bot calculates your age based on the year of birth provided. The bot is implemented in Go using the Slacker library.

## Prerequisites

- Go installed on your machine
- Slack bot token and app token

## Installation & Execution

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/your-repo.git
   cd your-repo
2. Install the dependcies:
    go get github.com/shomali11/slacker
3. Set your Slack tokens in the environment:
    export SLACK_BOT_TOKEN="your-slack-bot-token"
    export SLACK_APP_TOKEN="your-slack-app-token"
4. Run the bot:
    go run main.go
5. In your Slack workspace, type the command:
    @Age-bot My year of birth is 2001

