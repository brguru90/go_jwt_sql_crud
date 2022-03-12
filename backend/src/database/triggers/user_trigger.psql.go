package main

/*
#include <stdio.h>
#include "postgres.h"
#include <time.h>
#include "executor/spi.h"
#include "utils/fmgrprotos.h"
#include "commands/trigger.h"
#include "utils/elog.h"
#include "utils/rel.h"
#include "access/htup_details.h"
#include "executor/executor.h"

extern Datum trig_test(PG_FUNCTION_ARGS);

static int trigger_fired_by_update(TriggerEvent tg_event) {
	return (TRIGGER_FIRED_BY_UPDATE(tg_event)) != 0;
}
static Datum pointer_get_datum(HeapTuple t) {
	return PointerGetDatum(t);
}
static int64 get_row_id_first_col(TriggerData *trigdata, HeapTuple rettuple, int idx) {
	bool isnull;
	TupleDesc tupdesc = trigdata->tg_relation->rd_att;
	// Extracting value of first column
	int64 row_id=DatumGetInt64(heap_getattr(rettuple, idx, tupdesc, &isnull));
	return row_id;
}
*/
import "C"
import (
	"fmt"
	"net/http"
	"os"
	"unsafe"

	log "github.com/sirupsen/logrus"
)

//export user_update_trigger
func user_update_trigger(fcInfo *C.FunctionCallInfoBaseData) C.Datum {
	// !Careful, if Anything wrong in this file like syntax error or runtime error, then [database may fail to update data]
	// !Careful, The commented C codes are executable and can be called from Go functions
	// !Warning, IDEs may not able to autocomplete few c imports & IDEs may show error on some line
	// !Careful, The C implementation in comment in go is very sensitive,new lines before user_update_trigger() function may throws errors
	// * all the log statement in this source code will be writtenn to bellow postgres log file
	// tail -f /var/log/postgresql/postgresql-14-main.log

	trigdata := (*C.TriggerData)(unsafe.Pointer(fcInfo.context))

	var rettuple *C.HeapTupleData
	if C.trigger_fired_by_update(trigdata.tg_event) != 0 {
		rettuple = (*C.HeapTupleData)(trigdata.tg_newtuple)
	} else {
		rettuple = (*C.HeapTupleData)(trigdata.tg_trigtuple)
	}

	// calling the C function which read first column data,as specfie [1 bellow]
	created_or_updated_user_id := fmt.Sprintf("%v", (C.get_row_id_first_col(trigdata, rettuple, 1)))
	os.WriteFile("/tmp/dat1", []byte(created_or_updated_user_id), 0644)

	log.WithFields(log.Fields{
		"created_or_updated_user_id": created_or_updated_user_id,
	}).Errorln(">>>>>>>>>>>>>>>> got data from trigger")

	go RediscacheInvalidation(created_or_updated_user_id)

	return C.pointer_get_datum(rettuple)
}

func RediscacheInvalidation(user_id string) {
	req, err := http.NewRequest("GET", "http://localhost:3000/api/del_user_cache/"+user_id, nil)
	req.Header.Set("secret", "1234")
	if err == nil {
		client := &http.Client{}
		_, err := client.Do(req)
		if err == nil {
			log.Infoln("successfully sent request to del_user_cache")
		} else {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("failed to send request to del_user_cache")
		}
	} else {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to create request to del_user_cache")
	}

}
