package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func main() {
	
	//Setting environment
	//These tokens authenticate the bot with Slack.
	os.Setenv("SLACK_BOT_TOKEN","xoxb-7553825111601-7543697569012-yKiAda80nw2a1un1CHpOVtCN")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A07FX4XHQJW-7563968975248-4dda5cf78e69a2560e06b0863c344111dc24b94b9b78036dce24077255e0fab7")

	//creating bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))
	
	//Starts a goroutine to handle command events. 
	go printCommandEvents(bot.CommandEvents())	//go is used to run the function concurrently
	
	bot.Command("My year of birth is <year>",&slacker.CommandDefinition{	//Registers a new command for the bot
		Description: "Year of birth calculator",							//Describes what the command does.
		Examples: []string{"My year of birth is 2001"},
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {		//This function will be called when the command is triggered
			year := r.Param("year")
			yob,err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			age := 2024-yob
			response := fmt.Sprintf("age is %d", age)
			w.Reply(response)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err:= bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}


}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent)  {
	for event:= range analyticsChannel {	// Iterates over the events in the channel.
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}