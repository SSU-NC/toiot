package healthCheckUC
import (
	"encoding/json"
	"io"
	"log"
	"net"
	"time"

	"github.com/KumKeeHyun/toiot/health-check/adapter"
	"github.com/KumKeeHyun/toiot/health-check/domain/repository"
	"github.com/KumKeeHyun/toiot/health-check/setting"
)

type healthCheckUsecase struct {
	sr    repository.StatusRepo
	event chan interface{}
}

func NewHealthCheckUsecase(sr repository.StatusRepo, e chan interface{}) *healthCheckUsecase {
	hu := &healthCheckUsecase{
		sr:    sr,
		event: e,
	}
	l, err := net.Listen("tcp", setting.Healthsetting.Listen)
	if nil != err {
		log.Fatalf("fail to bind address to Listen; err: %v", err)
	}
	//defer l.Close()

	go func() {
		for {
			conn, err := l.Accept()
			if nil != err {
				log.Printf("fail to accept; err: %v", err)
				continue
			}
			go hu.healthCheck(conn)
		}
	}()
	/*
		go func() {
			tick := time.Tick(time.Duration(setting.StatusSetting.Tick) * time.Second)
			for {
				select {
				case <-tick:
					hu.healthCheck()
				}
			}
		}()
	*/
	return hu
}

func (hu *healthCheckUsecase) healthCheck(conn net.Conn) {

	// for {
	recvBuf := make([]byte, 4096)
	n, err := conn.Read(recvBuf)	
	if nil != err {
		if io.EOF == err {
			log.Printf("connection is closed from client; %v", conn.RemoteAddr().String())
			return
		}
		log.Printf("fail to receive data; err: %v", err)
		return
	}
	if n > 0 {
		var healthInfo adapter.HealthInfo
		var states adapter.States

		recvBuf = ClearPadding(recvBuf)
		// log.Println("recv Buf2 :", recvBuf)
		json.Unmarshal(recvBuf, &healthInfo)

		states.State = healthInfo
		states.Timestamp = string(time.Now().Unix())
		log.Println("convert to json :", healthInfo)
		//test_start
		tmphealth := hu.sr.UpdateTable(states) // 변화가 생긴 것들만 뭘로 변했는지 알려줌 ex : {1 [{1 1} {2 1} {8 0}]}
		log.Println(tmphealth)

		// hu.event <- tmphealth
		// //test_end

		// //hu.event <- hu.sr.UpdateTable(sinknum, res)

	}
	// }
}

func ClearPadding(buf []byte) []byte {
	var res []byte
	for i := 1; i < 4096; i++ {
		if (buf[i-1] == 125) && (buf[i] == 0) {
			res = buf[:i]
			break
		}
	}
	return res
}
