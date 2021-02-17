// TCP Client
package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "10.5.110.11:8085")
	if nil != err {
		log.Fatalf("failed to connect to server")
	}

	// some event happens
	conn.Write([]byte(`
	{
		"sid": 1,
		"state": [{
			"nid": 1,
			"state": true
		},{
			"nid": 2,
			"state": false
		},{
			"nid": 4,
			"state": false
		},{
			"nid": 8,
			"state": false
		},{
			"nid": 27,
			"state": false
		}]
	}
	`))
	/*
		{
			"sid": 1,
			"state": [{
				"nid": 1,
				"state": false
			},{
				"nid": 2,
				"state": true
			},{
				"nid": 4,
				"state": false
			},{
				"nid": 8,
				"state": true
			},{
				"nid": 27,
				"state": true
			}]
		}
		{
			"sid": 1,
			"state": [{
				"nid": 1,
				"state": false
			},{
				"nid": 2,
				"state": true
			},{
				"nid": 4,
				"state": false
			},{
				"nid": 8,
				"state": true
			},{
				"nid": 27,
				"state": true
			}]
		}
	*/
	//conn.Write([]byte("005111111111111"))
	// for {
	// 	// heartbeat
	// 	conn.Write([]byte("110210101012"))
	// 	time.Sleep(time.Duration(3) * time.Second)
	// }
}
