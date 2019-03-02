package services

import (
	"fmt"
	"net/http"

	"github.com/toothless-discovery/common"
)

type Discovery int

// Register handle http request
func (t *Discovery) Register(r *http.Request, req *common.DiscoveryRequest, result *common.DiscoveryResponse) error {
	resp := string([]byte(*req))
	fmt.Printf("Register: %s\n", resp)
	resp += ": Ok"
	*result = common.DiscoveryResponse(resp)
	return nil
}
