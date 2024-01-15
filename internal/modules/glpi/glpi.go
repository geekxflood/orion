// internal/modules/glpi/glpi.go

package module_glpi

import (
	"net/http"
	"sync/atomic"

	"github.com/geekxflood/orion/internal/helpers"
	"github.com/geekxflood/orion/internal/httpclient"
	"github.com/geekxflood/orion/internal/localtypes"
)

// Client is a struct that holds an atomic.Value for the configuration.
type Client struct {
	Conf *atomic.Value
}

// Full doc: https://github.com/ramylson/glpi/blob/master/glpi/apirest.md

// GetGLPI retrieves GLPI data.
func GetGLPI() (string, error) {

	var client Client

	// Load the current configuration
	currentConf := client.Conf.Load().(*localtypes.Config)

	// Get session token
	_, err := GetSessionToken(currentConf)
	if err != nil {
		return "", err
	}

	return "", nil
}

// GetSessionToken retrieves the session token.
func GetSessionToken(currentConf *localtypes.Config) (string, error) {
	// Set the headers for the request
	req, err := http.NewRequest("GET", currentConf.GlpiConfig.Url+"apirest.php/initSession/", nil)
	if err != nil {
		return "", err
	}

	encodedAuth := helpers.EncodeBase64(currentConf.GlpiConfig.Username + ":" + currentConf.GlpiConfig.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("App-Token", currentConf.GlpiConfig.ApiKey)
	req.Header.Set("Authorization", encodedAuth)

	// Make the request
	resp, err := httpclient.NewRequest(req)
	if err != nil {
		return "", err
	}

	return resp, nil
}

/*
Get an item
URL: apirest.php/:itemtype/:id
Description: Return the instance fields of itemtype identified by id.
Method: GET
Parameters: (Headers)
Session-Token: session var provided by initSession endpoint. Mandatory.
App-Token: authorization string provided by the GLPI api configuration. Optional.
Parameters: (query string)
id: unique identifier of the itemtype. Mandatory.
expand_dropdowns (default: false): show dropdown name instead of id. Optional.
get_hateoas (default: true): Show relations of the item in a links attribute. Optional.
get_sha1 (default: false): Get a sha1 signature instead of the full answer. Optional.
with_devices: Only for [Computer, NetworkEquipment, Peripheral, Phone, Printer], retrieve the associated components. Optional.
with_disks: Only for Computer, retrieve the associated file-systems. Optional.
with_softwares: Only for Computer, retrieve the associated software's installations. Optional.
with_connections: Only for Computer, retrieve the associated direct connections (like peripherals and printers) .Optional.
with_networkports: Retrieve all network's connections and advanced network's informations. Optional.
with_infocoms: Retrieve financial and administrative informations. Optional.
with_contracts: Retrieve associated contracts. Optional.
with_documents: Retrieve associated external documents. Optional.
with_tickets: Retrieve associated itil tickets. Optional.
with_problems: Retrieve associated itil problems. Optional.
with_changes: Retrieve associated itil changes. Optional.
with_notes: Retrieve Notes. Optional.
with_logs: Retrieve historical. Optional.
Returns:
200 (OK) with item data (Last-Modified header should contain the date of last modification of the item).
401 (UNAUTHORIZED).
404 (NOT FOUND).
*/
func GetItem() {

}
