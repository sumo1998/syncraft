package server

type Command struct {
	Version uint16
	ID      string
	Action  string
	Data    map[string]string
}

/*
Commands for performing catalog actions.
Catalog actionss maintainin nodes, endpoints, addr etc.
*/
var AddNode string = "add_node"
var DeleteNode string = "delete_node"
var HealthCheck string = "health_check"
var RegisterService string = "register_service"
var DeregisterService string = "deregister_service"

var CatalogCommands = []string{AddNode, DeleteNode, HealthCheck, RegisterService, DeregisterService}
