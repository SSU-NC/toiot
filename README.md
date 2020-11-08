<p align="center">
  <img src="https://user-images.githubusercontent.com/38535571/92803983-e8b48000-f3f2-11ea-9213-1cdc70fd7df1.png" width="50%">
</p>

# Toiot
Internet of Things platform for device management, data collection, analytics and visualization and more. Toiot provides a powerful enterprise-class platform for those with little web skills, and provides drivers for sensors and hardware boards.

<img width="1856" alt="toiot_arch" src="https://user-images.githubusercontent.com/38535571/98469939-d48bd400-2225-11eb-9dd7-ccbce2ebe750.png">


Table of contents
=================
<!--ts-->
   * [Output](#Output)
   * [Quickstart](#Quickstart)
   * [Installation & Run](#Installation--Run)
   * [Detail](#Detail)
   * [Help](#Help)
   * [License](#License)
<!--te-->

Output
=======
- Dashboard
![dashboard](https://user-images.githubusercontent.com/38535571/92531615-96942300-f269-11ea-83a6-144addd100d4.png)
- Sink and Sensor Mangement
![sink_register](https://user-images.githubusercontent.com/38535571/92531650-ae6ba700-f269-11ea-8cd4-b9ba0e04c24f.png)<br>
![sensor_register](https://user-images.githubusercontent.com/38535571/92531663-b9263c00-f269-11ea-9896-25ba747deb55.png)<br>
![sensor_table](https://user-images.githubusercontent.com/38535571/92531768-e7a41700-f269-11ea-80b7-a0f8c37ccaf2.png)<br>
![register_node_ex](https://user-images.githubusercontent.com/59961690/93308957-e8c2de80-f83d-11ea-8fb7-97688f9be285.png)
- Service
![service](https://user-images.githubusercontent.com/38535571/92531789-f5599c80-f269-11ea-963a-269f53424760.gif)



Quickstart
======= 
This feature is used to `lightly` run this application. All platform elements run on a single server. If you want to run enterprise-class, please refer to [Installation & Run](#Installation--Run) and customize.

#### Installation
```bash
$ git clone https://github.com/SSU-NC/toiot
```

#### Run
```sh
$ docker-compose up
```

#### Stop
```sh
$ docker-compose down
```

#### Port Forwarding
|Host|Container|Service|
|:---:|:---:|:---:|
|3307|3306|mysql|
|8081|8081|application|
|8082|8082|logic|
|8083|8083|health check|
|3000|3000|ui|
|2181|2181|zookeeper|
|9092|9092|kafka|
|9200|9200|elasticsearch|
|5601|5601|kibana|

Installation & Run
=======
This document describes how to personally set up and use this platform. You can launch applications on multiple servers, or you can set various options to make use of a lot of computing resources.  
* [User Interface](./ui/README.md)
* [Backend Server](./application/README.md)


Detail
=======
[WIKI](https://github.com/SSU-NC/toiot/wiki)

Help
=======
yoonje.choi.dev@gmail.com

License
=======
```
Copyright 2020 Keehyun Kum, Haegyeong Im, Jungsu Kim, Sehee Jeong, Yoonje Choi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
