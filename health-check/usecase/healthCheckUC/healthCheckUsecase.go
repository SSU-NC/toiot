package healthCheckUC

import (
	"io"
	"log"
	"math"
	"net"
	"time"

	"github.com/KumKeeHyun/toiot/health-check/adapter"
	"github.com/KumKeeHyun/toiot/health-check/domain/repository"
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
	l, err := net.Listen("tcp", ":5032") // 포트정보 setting으로 옮겨야 함
	if nil != err {
		log.Fatalf("fail to bind address to 5032; err: %v", err)
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
	recvBuf := make([]byte, 4096) // receive buffer: 4kB 현재 1바이트씩 받는데 1비트씩으로 수정해야 함
	for {
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
			data := make([]byte, n)
			sinknum := bytetoint(recvBuf[:3]) // 앞 3자리로 싱크 번호 가져옴
			log.Println("sinknum = ", sinknum)
			var states []bool
			states = []bool{}

			for i := 3; i < n; i++ {
				data[i-3] = recvBuf[i]
			}
			for i := 0; i < n-3; i++ {
				log.Println(i, "-Node's status : ", data[i]-'0')
				var status bool // 싱크로 부터 tcp전송받은 센서들의 상태
				if int(data[i]-'0') > 0 {
					status = true
				} else {
					status = false
				}
				states = append(states, status)
				//
			}
			res := makeStates(states)

			//test_start
			tmphealth := hu.sr.UpdateTable(sinknum, res)

			for i := 0; i < len(tmphealth); i++ {
				log.Println("test_Result_", i, "-> Node[", tmphealth.[i].NodeID, "] : ", tmphealth[i].State) // Satates
			}
			hu.event <- tmphealth
			//test_end

			//hu.event <- hu.sr.UpdateTable(sinknum, res)

		}
	}

	/*
		sinks, err := getSinkList()
		if err != nil {
			return
		}

		var wg sync.WaitGroup
		for _, sink := range sinks {
			wg.Add(1)
			go func(s adapter.Sink) {
				res := adapter.States{}
				client := resty.New()
				client.SetTimeout(500 * time.Millisecond)
				resp, _ := client.R().SetResult(&res).Get(fmt.Sprintf("http://%s/health-check", s.Addr))

				if resp.IsSuccess() {
					hu.event <- hu.sr.UpdateTable(s.ID, res)
				}
				wg.Done()
			}(sink)
		}
	*/
}
func bytetoint(data []byte) int {
	res := 0
	n := len(data)
	for i := 0; i < n; i++ {
		res += int(float64(data[i]-'0') * math.Pow(10, float64(n-1-i)))
	}
	return res
}
func makeStates(states []bool) adapter.States {
	res := adapter.States{}
	res.Timestamp = string(time.Now().Unix())
	tmpState := adapter.NodeState{}
	for i := 0; i < len(states); i++ {
		tmpState.NodeID = i
		tmpState.State = states[i]
		res.State = append(res.State, tmpState)
	}
	return res

}
