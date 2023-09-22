package task

import (
	"AnaOrderService/repository"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
)

func SyncReleasedUnusedStocks(orderRepository *repository.OrderRepository) {
	c := cron.New()
	job := runTask(*orderRepository)
	err := c.AddFunc("*/10 * * * * *", job)
	if err != nil {
		log.Fatal("Error scheduling task:", err)
	}

	c.Start()
	// Keep the program running
	select {}
}

func runTask(orderRepository repository.OrderRepository) func() {
	return func() {
		err := orderRepository.ReleaseUnusedStock()
		if err != nil {
			fmt.Println(fmt.Sprintf("%s%s", "Error Released unused stock in cart", err))
			return
		}

		fmt.Println("Task executed at", time.Now())
	}
}
