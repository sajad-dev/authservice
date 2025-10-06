package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Listeners(ctx *gin.Context) {
	log.Println("Listeners")

	response := map[string]interface{}{
		"version_info": "1",
		"resources": []map[string]interface{}{
			{
				"@type": "type.googleapis.com/envoy.config.listener.v3.Listener",
				"name":  "listener_0",
				"address": map[string]interface{}{
					"socket_address": map[string]interface{}{
						"address":    "0.0.0.0",
						"port_value": 10000,
					},
				},
				"filter_chains": []map[string]interface{}{
					{
						"filters": []map[string]interface{}{
							{
								"name": "envoy.filters.network.http_connection_manager",
								"typed_config": map[string]interface{}{
									"@type":       "type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager",
									"stat_prefix": "ingress_http",
									"route_config": map[string]interface{}{
										"name": "local_route",
										"virtual_hosts": []map[string]interface{}{
											{
												"name":    "local_service",
												"domains": []string{"*"},
												"routes": []map[string]interface{}{
													{
														"match": map[string]interface{}{
															"prefix": "/",
														},
														"route": map[string]interface{}{
															"cluster": "service_cluster",
														},
													},
												},
											},
										},
									},
									"http_filters": []map[string]interface{}{
										{
											"name":         "envoy.filters.http.router",
											"typed_config": map[string]interface{}{"@type": "type.googleapis.com/envoy.extensions.filters.http.router.v3.Router"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"type_url": "type.googleapis.com/envoy.config.listener.v3.Listener",
		"nonce":    "1",
		"control_plane": map[string]interface{}{
			"identifier": "go-xds-server",
		},
	}

	ctx.JSON(200, response)
}
