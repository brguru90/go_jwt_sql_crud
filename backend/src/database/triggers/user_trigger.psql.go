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
	int64 row_id=DatumGetInt64(heap_getattr(rettuple, idx, tupdesc, &isnull));
	return row_id;
}
*/
import "C"
import (
	"fmt"
	"os"
	"time"
	"unsafe"
)

//export user_update_trigger
func user_update_trigger(fcInfo *C.FunctionCallInfoBaseData) C.Datum {
	// tail -f /var/log/postgresql/postgresql-14-main.log
	fmt.Println("user_update_trigger", time.Now())

	trigdata := (*C.TriggerData)(unsafe.Pointer(fcInfo.context))

	var rettuple *C.HeapTupleData
	if C.trigger_fired_by_update(trigdata.tg_event) != 0 {
		rettuple = (*C.HeapTupleData)(trigdata.tg_newtuple)
	} else {
		rettuple = (*C.HeapTupleData)(trigdata.tg_trigtuple)
	}

	created_or_updated_user_id := fmt.Sprintf("%v", (C.get_row_id_first_col(trigdata, rettuple, 1)))
	os.WriteFile("/tmp/dat1", []byte(created_or_updated_user_id), 0644)

	// log.WithFields(log.Fields{
	// 	"created_or_updated_user_id": created_or_updated_user_id,
	// }).Error(">>>>>>>>>>>>>>>> got data from trigger")

	return C.pointer_get_datum(rettuple)
}
