package clusters

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/authservice/getway/xds/handler"
)

type Cluster struct {
	Type                          string                         `json:"@type"`
	Name                          string                         `json:"name"`
	ClusterType                   string                         `json:"type"`
	ConnectTimeout                string                         `json:"connect_timeout"`
	LbPolicy                      string                         `json:"lb_policy"`
	DNSLookupFamily               string                         `json:"dns_lookup_family,omitempty"`
	TypedExtensionProtocolOptions *TypedExtensionProtocolOptions `json:"typed_extension_protocol_options,omitempty"`
	LoadAssignment                LoadAssignment                 `json:"load_assignment"`
}

type TypedExtensionProtocolOptions struct {
	HttpProtocolOptions *HttpProtocolOptionsWrapper `json:"envoy.extensions.upstreams.http.v3.HttpProtocolOptions"`
}

type HttpProtocolOptionsWrapper struct {
	Type               string              `json:"@type"`
	ExplicitHttpConfig *ExplicitHttpConfig `json:"explicit_http_config"`
}

type ExplicitHttpConfig struct {
	Http2ProtocolOptions map[string]interface{} `json:"http2_protocol_options"`
}

type LoadAssignment struct {
	ClusterName string     `json:"cluster_name"`
	Endpoints   []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	LbEndpoints []LbEndpoint `json:"lb_endpoints"`
}

type LbEndpoint struct {
	Endpoint EndpointAddress `json:"endpoint"`
}

type EndpointAddress struct {
	Address SocketAddressWrapper `json:"address"`
}

type SocketAddressWrapper struct {
	SocketAddress SocketAddress `json:"socket_address"`
}

type SocketAddress struct {
	Address   string `json:"address"`
	PortValue int    `json:"port_value"`
}

// -----------------  Service -----------------

type Address struct {
	Address string `json"address"`
	Port    int    `json"port"`
}

type Service struct {
	Name      string    `json:"name"`
	Endpoints []Address `json"endpoints"`
}

func deepCopy[T any](src T) (T, error) {
	var dst T
	data, err := json.Marshal(src)
	if err != nil {
		return dst, err
	}
	err = json.Unmarshal(data, &dst)
	return dst, err
}
func Clusters(ctx *gin.Context) {
	cls, err := handler.GetJson[Cluster]("json/cls_static.json")
	if err != nil {
		ctx.JSON(500, gin.H{"error cls static": err.Error()})
		return
	}
	svcArray, err := handler.GetJson[[]Service]("json/cls.json")
	if err != nil {
		ctx.JSON(500, gin.H{"error cls": err.Error()})
		return
	}
	discoveryResponse, err := handler.GetJson[handler.DiscoveryResponse[Cluster]]("json/discovery_response.json")
	if err != nil {
		ctx.JSON(500, gin.H{"error discovery": err.Error()})
		return
	}
	var clsArray []Cluster
	for _, svc := range svcArray {
		clsIns, err := deepCopy(cls)
		if err != nil {
			ctx.JSON(500, gin.H{"error discovery": err.Error()})
			return
		}

		log.Println(clsIns.LoadAssignment.Endpoints[0])

		clsIns.Name = svc.Name
		clsIns.LoadAssignment.ClusterName = svc.Name

		var endpoint Endpoint
		for _, addr := range svc.Endpoints {
			endpoint.LbEndpoints = append(endpoint.LbEndpoints, LbEndpoint{
				Endpoint: EndpointAddress{
					SocketAddressWrapper{
						SocketAddress: SocketAddress{
							Address:   addr.Address,
							PortValue: addr.Port,
						},
					},
				},
			})
		}

		clsIns.LoadAssignment.Endpoints[0] = endpoint

		clsArray = append(clsArray, clsIns)
	}

	discoveryResponse.Resources = clsArray
	ctx.JSON(200, discoveryResponse)
}
