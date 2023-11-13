//go:generate oapi-codegen -generate types,server -package api -o api.gen.go openapi.yaml

package api

import "github.com/labstack/echo/v4"

type IBolitAPI struct{}

func (i *IBolitAPI) GetApiV1EnvIdentityRunTestRmq(ctx echo.Context, env string, identity string) error {
	// fill the RedisTestResult with dummy data
	redisTestResult := RedisTestResult{
		RMQConsumers:   &[]int{1}[0],
		RMQFaultQueues: &[]string{"RMQFaultQueues"},
		RMQLongQueues:  &[]string{"RMQLongQueues"},
		RMQQueues:      &[]int{1}[0],
		Name:           "RMQ Test",
		Result:         true,
	}

	// return the RedisTestResult
	return ctx.JSON(200, redisTestResult)
}
