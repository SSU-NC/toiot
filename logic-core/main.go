package main

import (
	"github.com/KumKeeHyun/PDK/logic-core/dataService/memory"
	"github.com/KumKeeHyun/PDK/logic-core/elasticClient"
	"github.com/KumKeeHyun/PDK/logic-core/kafkaConsumer"
	"github.com/KumKeeHyun/PDK/logic-core/logicCore"
	"github.com/KumKeeHyun/PDK/logic-core/rest"
	_ "github.com/KumKeeHyun/PDK/logic-core/setting"
	"github.com/KumKeeHyun/PDK/logic-core/usecase/logicCoreUC"
	"github.com/KumKeeHyun/PDK/logic-core/usecase/metaDataUC"
)

func main() {
	mr := memory.NewMetaRepo()
	ks := kafkaConsumer.NewKafkaConsumer()
	es := elasticClient.NewElasticClient()
	ls := logicCore.NewLogicCore()

	mduc := metaDataUC.NewMetaDataUsecase(mr, ls)
	lcuc := logicCoreUC.NewLogicCoreUsecase(mr, ks, es, ls)

	h := rest.NewHandler(mduc, lcuc)
	rest.RunServer(h)
}
