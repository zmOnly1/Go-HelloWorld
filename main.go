package main

import "github.com/tidwall/gjson"

//const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
const my = `
[{
	"Node": {
		"ID": "8d0b2a53-d23b-4927-9d31-89f977b0a271",
		"Node": "s1",
		"Address": "192.168.1.104",
		"Datacenter": "dc1",
		"TaggedAddresses": {
			"lan": "192.168.1.104",
			"wan": "192.168.1.104"
		},
		"Meta": {
			"consul-network-segment": ""
		},
		"CreateIndex": 4,
		"ModifyIndex": 4
	},
	"Service": {
		"ID": "app1",
		"Service": "graceful-test",
		"Tags": ["secure=false"],
		"Address": "192.168.1.104",
		"Meta": null,
		"Port": 35497,
		"EnableTagOverride": false,
		"ProxyDestination": "",
		"Connect": {
			"Native": false,
			"Proxy": null
		},
		"CreateIndex": 429,
		"ModifyIndex": 429
	},
	"Checks": [{
		"Node": "s1",
		"CheckID": "serfHealth",
		"Name": "Serf Health Status",
		"Status": "passing",
		"Notes": "",
		"Output": "Agent alive and reachable",
		"ServiceID": "",
		"ServiceName": "",
		"ServiceTags": [],
		"Definition": {},
		"CreateIndex": 6,
		"ModifyIndex": 6
	}, {
		"Node": "s1",
		"CheckID": "service:app1",
		"Name": "Service 'graceful-test' check",
		"Status": "passing",
		"Notes": "",
		"Output": "HTTP GET http://localhost:35497/health: 200  Output: Up app1",
		"ServiceID": "app1",
		"ServiceName": "graceful-test",
		"ServiceTags": ["secure=false"],
		"Definition": {},
		"CreateIndex": 429,
		"ModifyIndex": 435
	}]
}, {
	"Node": {
		"ID": "8d0b2a53-d23b-4927-9d31-89f977b0a271",
		"Node": "s1",
		"Address": "192.168.1.104",
		"Datacenter": "dc1",
		"TaggedAddresses": {
			"lan": "192.168.1.104",
			"wan": "192.168.1.104"
		},
		"Meta": {
			"consul-network-segment": ""
		},
		"CreateIndex": 4,
		"ModifyIndex": 4
	},
	"Service": {
		"ID": "app2",
		"Service": "graceful-test",
		"Tags": ["secure=false"],
		"Address": "192.168.1.105",
		"Meta": null,
		"Port": 40541,
		"EnableTagOverride": false,
		"ProxyDestination": "",
		"Connect": {
			"Native": false,
			"Proxy": null
		},
		"CreateIndex": 394,
		"ModifyIndex": 419
	},
	"Checks": [{
		"Node": "s1",
		"CheckID": "serfHealth",
		"Name": "Serf Health Status",
		"Status": "passing",
		"Notes": "",
		"Output": "Agent alive and reachable",
		"ServiceID": "",
		"ServiceName": "",
		"ServiceTags": [],
		"Definition": {},
		"CreateIndex": 6,
		"ModifyIndex": 6
	}, {
		"Node": "s1",
		"CheckID": "service:app2",
		"Name": "Service 'graceful-test' check",
		"Status": "passing",
		"Notes": "",
		"Output": "HTTP GET http://localhost:40541/health: 200  Output: Up app2",
		"ServiceID": "app2",
		"ServiceName": "graceful-test",
		"ServiceTags": ["secure=false"],
		"Definition": {},
		"CreateIndex": 394,
		"ModifyIndex": 421
	}]
}]
`
const test_json = `
{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}
`
const test_arr = `[]`

func main() {
	//value := gjson.Get(test_json, "name.last")
	//println(gjson.Get(test_json, "name.last").String())
	//println(gjson.Get(test_json, "age").String())
	//println(gjson.Get(test_json, "children").String())
	//println(gjson.Get(test_json, "children.#").String())
	//println(gjson.Get(test_json, "children.@reverse").String())
	//println(gjson.Get(test_json, `{name.first,age,"the_murphys":friends.#(last="Murphy")#.first}`).String())
	//println(gjson.Get(test_json, `{name.first,age,"the_murphys":name.last}`).String())

	println("count:" + gjson.Get(my, "#").String())
	//println("firstObj:" + gjson.Get(my, "0").String())
	println("obj.Service:" + gjson.Get(my, "0.Service").String())
	println("obj.Service.name:" + gjson.Get(my, "0.Service.Service").String())
	println("obj.Service.Address:" + gjson.Get(my, "0.Service.Address").String())
	println("obj.Service.Port:" + gjson.Get(my, "0.Service.Port").String())
	println("empty.Service:" + gjson.Get(test_arr, "0.Service.Service").String())
	println(gjson.Get(my, `{0.Service.Service}`).String())
	//println(gjson.Get(my, `{"cnt":#}`).String())
	println(gjson.Get(my, `{"name":0.Service.Service,"healthCount":#,"Address1":0.Service.Address,"Port1":0.Service.Port,"Address2":1.Service.Address,"Port2":1.Service.Port}`).String())
}
