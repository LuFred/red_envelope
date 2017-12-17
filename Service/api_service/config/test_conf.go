package config

const (
	test_config = `{
		"service_name":"api_service",
		"debug":true,
		"tracing_transport_url":"",
		  "listen":"127.0.0.1:4501",
		  "microservice":{      
			"red_envelope_service_host":"127.0.0.1:50020"      
		  },
			"default_token":"oauth f9894bed32164a49a3452a0802c1b11b",
			"redis_host":{
        "addr":"127.0.0.1:6379",
        "pwd":"112233"
    	}		
	  }`
)
