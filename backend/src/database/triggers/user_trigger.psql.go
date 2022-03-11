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
static char *getarg_text(TriggerData *trigdata, HeapTuple rettuple, int idx) {
	bool isnull;

	HeapTupleHeader  t = PG_GETARG_HEAPTUPLEHEADER(0);

	TupleDesc tupdesc = trigdata->tg_relation->rd_att;
	int32 att = DatumGetInt32(heap_getattr(rettuple, idx, tupdesc, &isnull));
	DatumGetInt32(GetAttributeByName(t, "id", &isnull));
	//text *t = DatumGetTextP(heap_getattr(rettuple, idx, tupdesc, &isnull));
	//if (isnull || !t) {
	//return "";
	//}
	// return VARDATA(t);
	return "hi";
}
static void elog_info(char *s) {
	elog(INFO, "%s", s);
}
*/
import "C"
import (
	"fmt"
	"time"
	"os"
	"unsafe"
	// log "github.com/sirupsen/logrus"
)

//export user_update_trigger
func user_update_trigger(fcInfo *C.FunctionCallInfoBaseData) C.Datum {
	// tail -f /var/log/postgresql/postgresql-14-main.log 
	fmt.Println("user_update_trigger",time.Now())

	trigdata := (*C.TriggerData)(unsafe.Pointer(fcInfo.context))

	var rettuple *C.HeapTupleData
	if C.trigger_fired_by_update(trigdata.tg_event) != 0 {
		rettuple = (*C.HeapTupleData)(trigdata.tg_newtuple)
	} else {
		rettuple = (*C.HeapTupleData)(trigdata.tg_trigtuple)
	}

	_log:=[]byte(fmt.Sprintf("%v",C.GoString(C.getarg_text(trigdata, rettuple, 1))))
	os.WriteFile("/tmp/dat1", _log, 0644)


	// first_param := C.GoString(C.getarg_text(trigdata, rettuple, 1))

	// C.elog_info(C.CString(fmt.Sprintf("got url=%s", first_param)))

	// log.WithFields(log.Fields{
	// 	"first_param":first_param,
	// }).Error(">>>>>>>>>>>>>>>> got first param from trigger")

	return C.pointer_get_datum(rettuple)
}