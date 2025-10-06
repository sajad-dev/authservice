package clusters_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/authservice/getway/xds/handler/clusters"
)

func TestClusters(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/clusters", clusters.Clusters)

	req := httptest.NewRequest("GET", "/clusters", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)

	// var resp map[string]any
	// err := json.Unmarshal(w.Body.Bytes(), &resp)
	// assert.NoError(t, err)

	t.Log(w.Body.String())
	// 	assert.Equal(t, "1", resp["version_info"])
	// 	assert.Equal(t, "type.googleapis.com/envoy.config.cluster.v3.Cluster", resp["type_url"])

	// 	resources, ok := resp["resources"].([]any)
	// 	assert.True(t, ok)
	// 	assert.Len(t, resources, 1)

	// 	t.Log(resp)

	//	clusterMap, ok := resources[0].(map[string]any)
	//	assert.True(t, ok)
	//
	// assert.Equal(t, "cluster1", clusterMap["name"])
}
