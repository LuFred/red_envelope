package config

const (
	test_config = `{
		"service_name":"red_envelope_service",
		"debug":true,
		"tracing_transport_url":"",
		"listen":"127.0.0.1:50020",
		 "mysql": {
			"host": "192.168.1.2:3306",
			"username": "root",
			"pwd": "root",
			"db":"red_envelope"
		}
	}`
)

