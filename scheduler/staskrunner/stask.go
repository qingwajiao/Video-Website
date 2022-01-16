package staskrunner

import (
	"Video-Website/scheduler/dbops"
	"log"
	"sync"
)

var hp = dbops.NewVideoHelper()

func VideoClearDispatcher(dc dataChan) error {
	/**
	去数据库查询要删除的videoID，
	将获取到的videoID 放入到 datachanl中
	*/
	ids, err := hp.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Fatalf("")
		return err
	}
	for _, v := range ids {
		dc <- v
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {
	/**
	从datachanl中获取videoID
	根据获取到的videoID去数据库要删除的video
	*/

	errMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case d := <-dc:
			go func(id interface{}) {
				// 删除video文件  todo 待完善

				// 根据videoID删除数据库中的记录
				err = hp.DelVideoDeletionRecord(id.(string))
				if err != nil {
					errMap.Store(id, err)
				}
			}(d)

		default:
			break forloop
		}
	}

	return nil
}
