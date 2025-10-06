package listeners

import (
	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/authservice/getway/xds/handler"
)

type Config struct {
	Listeners []Listener `json:"listeners"`
}

type Listener struct {
	Name         string        `json:"name"`
	Address      Address       `json:"address"`
	FilterChains []FilterChain `json:"filter_chains"`
}

type Address struct {
	SocketAddress SocketAddress `json:"socket_address"`
}

type SocketAddress struct {
	Address   string `json:"address"`
	PortValue int    `json:"port_value"`
}

type FilterChain struct {
	Filters []Filter `json:"filters"`
}

type Filter struct {
	Name        string      `json:"name"`
	TypedConfig TypedConfig `json:"typed_config"`
}

type TypedConfig struct {
	Type        string       `json:"@type"`
	StatPrefix  string       `json:"stat_prefix,omitempty"`
	CodecType   string       `json:"codec_type,omitempty"`
	RouteConfig *RouteConfig `json:"route_config,omitempty"`
	HttpFilters []HttpFilter `json:"http_filters,omitempty"`
}

type RouteConfig struct {
	Name         string        `json:"name"`
	VirtualHosts []VirtualHost `json:"virtual_hosts"`
}

type VirtualHost struct {
	Name    string   `json:"name"`
	Domains []string `json:"domains"`
	Routes  []Route  `json:"routes"`
}

type Route struct {
	Match Match       `json:"match"`
	Route RouteAction `json:"route"`
}

type Match struct {
	Prefix string `json:"prefix"`
}

type RouteAction struct {
	Cluster string `json:"cluster"`
	Timeout string `json:"timeout"`
}

type HttpFilter struct {
	Name        string       `json:"name"`
	TypedConfig FilterConfig `json:"typed_config"`
}

type FilterConfig struct {
	Type            string        `json:"@type"`
	ProtoDescriptor string        `json:"proto_descriptor,omitempty"`
	Services        []string      `json:"services,omitempty"`
	PrintOptions    *PrintOptions `json:"print_options,omitempty"`
	GrpcService     *GrpcService  `json:"grpc_service,omitempty"`
	Timeout         string        `json:"timeout,omitempty"`
}

type PrintOptions struct {
	AddWhitespace              bool `json:"add_whitespace"`
	AlwaysPrintPrimitiveFields bool `json:"always_print_primitive_fields"`
	AlwaysPrintEnumsAsInts     bool `json:"always_print_enums_as_ints"`
	PreserveProtoFieldNames    bool `json:"preserve_proto_field_names"`
}

type GrpcService struct {
	EnvoyGrpc *EnvoyGrpc `json:"envoy_grpc,omitempty"`
	Timeout   string     `json:"timeout,omitempty"`
}

type EnvoyGrpc struct {
	ClusterName string `json:"cluster_name"`
}

type RouteItem struct {
	Prefix  string `json:"prefix"`
	Cluster string `json:"cluster"`
}

type Services struct {
	Pb     string      `json:"pb"`
	Routes []RouteItem `json:"routes"`
}

func Listeners(ctx *gin.Context) {
	lds, err := handler.GetJson[Config]("json/lds_static.json")
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to load LDS: " + err.Error()})
		return
	}
	svc, err := handler.GetJson[Services]("json/services.json")
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to load Services: " + err.Error()})
		return
	}
	discoveryResponse, err := handler.GetJson[handler.DiscoveryResponse[Listener]]("json/discovery_response.json")
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to load Discovery Response: " + err.Error()})
		return
	}

	if len(lds.Listeners) == 0 || len(lds.Listeners[0].FilterChains) == 0 ||
		len(lds.Listeners[0].FilterChains[0].Filters) == 0 {
		ctx.JSON(500, gin.H{"error": "invalid LDS structure"})
		return
	}

	filter := &lds.Listeners[0].FilterChains[0].Filters[0]
	if filter.TypedConfig.RouteConfig == nil ||
		len(filter.TypedConfig.RouteConfig.VirtualHosts) == 0 {
		ctx.JSON(500, gin.H{"error": "missing route_config or virtual_hosts"})
		return
	}

	var routesArr []Route
	for _, r := range svc.Routes {
		prefix := r.Prefix
		if prefix != "" && prefix[0] != '/' {
			prefix = "/" + prefix
		}
		routesArr = append(routesArr, Route{
			Match: Match{Prefix: prefix},
			Route: RouteAction{
				Cluster: r.Cluster,
				Timeout: "60s",
			},
		})
	}

	filter.TypedConfig.RouteConfig.VirtualHosts[0].Routes = routesArr

	for i := range filter.TypedConfig.HttpFilters {
		if filter.TypedConfig.HttpFilters[i].Name == "envoy.filters.http.grpc_json_transcoder" {
			filter.TypedConfig.HttpFilters[i].TypedConfig.ProtoDescriptor = svc.Pb
			filter.TypedConfig.HttpFilters[i].TypedConfig.Services = nil
			for _, r := range svc.Routes {
				filter.TypedConfig.HttpFilters[i].TypedConfig.Services = append(
					filter.TypedConfig.HttpFilters[i].TypedConfig.Services,
					r.Prefix,
				)
			}
			break
		}
	}

	discoveryResponse.Resources = lds.Listeners
	ctx.JSON(200, discoveryResponse)
}
