package main

import (
	"fmt"

	"github.com/KumKeeHyun/toiot/health-check/dataService/memory"
	"github.com/KumKeeHyun/toiot/health-check/usecase/healthCheckUC"
)

/*

sink로 부터 받는 파일
json{
	state : int[] 인덱스가 센서 상태
}


싱크 리스트 받고, 싱크마다 돌며 POST 받고, 받은 정보로  TABLE[SINK][NODE]에 SHORT데이터로 헬스 정상 여부 저장 RED:0,YELLOW:1,GREEN:2
2차원 배열 TABLE[SINK][NODE] 째로 웹으로 SHOOT

*/

func main() {
	sr := memory.NewStatusRepo()

	event := make(chan interface{}, 10)
	_ = healthCheckUC.NewHealthCheckUsecase(sr, event)
	fmt.Scanln()
	/*
		wu := websocketUC.NewWebsocketUsecase(event)

		r := gin.New()

		r.GET("/health-check", func(c *gin.Context) {
			listen := make(chan interface{})
			wu.Register(listen)
			defer wu.Unregister(listen)

			conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
			if err != nil {
				log.Printf("upgrade: %s", err.Error())
			}
			fmt.Println("connect websocket!")

			for data := range listen {
				conn.WriteJSON(data)
			}
			fmt.Println("disconnect websocket!")
		})

		go log.Fatal(r.Run(setting.Healthsetting.Server))

		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
		<-sigterm
	*/
}
