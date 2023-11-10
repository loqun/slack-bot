package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}
}

func main() {

	

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6151739495125-6151769946741-BrMrrFNW6DSLYNLIrb3FVQoJ")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A064JKUN7NE-6140144840199-749ba189bf02843a2992229fd4c5c119c3b35b686f52b6e00dde98ca2d4294f3")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples: []string{"my yob is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err!=nil{
				println("error")
			} 
			age := 2023- yob 
			r:= fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},

	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err:= bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}

}