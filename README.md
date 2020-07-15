# PDK
Internet of Things platform for device management, data collection, analytics and visualization, remote control and more.

<img width="780" alt="pdk_arch" src="https://user-images.githubusercontent.com/38535571/85942352-a20ec800-b963-11ea-8d19-60207486f2a5.png">

Table of contents
=================
<!--ts-->
   * [Result](#Result)
   * [How to use](#How-to-use)
   * [Requirement](#Requirement)
   * [Build & Installation](#Build--Installation)
   * [API](#API)
   * [Help](#Help)
<!--te-->

Result
=======

Requirement
=======

Build & Installation
=======

API
=======
* GET /node/regist
  * get all node informations
* POST /node/regist -d {"name":string, "location":string, "sensors":[string...]}
  * register node info
* POST /node/sensor -d {"node_uuid":string, "sensor_uuid":string}
  * connect sensor info to node
* GET /sensor/info 
  * get all sensor informations
* POST /sensor/regist -d {"name":string, "num_of_values":int, "value_names":[string...]}
  * register sensor info


Help
=======
