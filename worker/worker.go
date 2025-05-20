package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nktauserum/crawler-service/common"
	"github.com/nktauserum/crawler-service/pkg/crawler"
	"github.com/nktauserum/crawler-service/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Println("Worker started")
	worker := NewWorker(5, "localhost:50000")

	worker.StartPolling()
}

type Worker struct {
	nigga_count int
	client      pb.TaskServiceClient
	conn        *grpc.ClientConn
}

func NewWorker(count int, address string) *Worker {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("error establishing connection: " + err.Error())
	}

	client := pb.NewTaskServiceClient(conn)
	return &Worker{nigga_count: count, conn: conn, client: client}
}

func (w *Worker) StartPolling() {
	var wg sync.WaitGroup

	for range w.nigga_count {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				task, err := w.GetTask()
				if err != nil {
					// Задержка перед повторным запросом
					time.Sleep(500 * time.Millisecond)
					continue
				}

				if task == nil {
					continue
				}

				log.Printf("Получена задача: %s", task.URL)

				page, err := ProcessURL(task.URL)
				if err != nil {
					continue
				}

				log.Printf("Задача завершена: %s", task.URL)

				err = w.CompleteTask(task.UUID, page.Content)
				if err != nil {
					continue
				}
			}

		}()
	}
	wg.Wait()
}

func ProcessURL(url string) (common.Page, error) {
	page, err := crawler.GetContent(context.Background(), url)
	if err != nil {
		return common.Page{}, err
	}

	return page, nil
}

func (c *Worker) GetTask() (*common.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	task, err := c.client.GetAvailableTask(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, nil
	}

	if task.Url == "" {
		return nil, fmt.Errorf("no tasks available")
	}

	return &common.Task{
		UUID: task.Uuid,
		URL:  task.Url,
	}, nil
}

func (c *Worker) CompleteTask(uuid, result string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := c.client.CompleteTask(ctx, &pb.TaskResult{
		Uuid:   uuid,
		Result: result,
	})

	return err
}

func (c *Worker) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}
