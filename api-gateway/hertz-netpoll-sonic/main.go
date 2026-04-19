package main

import (
	"context"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/netpoll"
)

func main() {
	h := server.Default(
		server.WithHostPorts(":8080"),
		server.WithTransport(netpoll.NewTransporter),
	)
	h.GET("/ping", func(
		ctx context.Context,
		c *app.RequestContext,
	) {
		data, _ := sonic.Marshal(map[string]string{
			"message": "ok",
			"service": "hertz-gateway",
		})
		c.Data(http.StatusOK, "application/json", data)
	})
	h.Spin()
}
