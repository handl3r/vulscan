package enums

import "errors"

var ErrClientNewTask = errors.New("can not create new task sqlmapapi")

var ErrClientSetOptionGET = errors.New("can not set option for GET")
var ErrClientSetOptionPOST = errors.New("can not set option for POST")
var ErrClientStartScan = errors.New("can not start scan")
var ErrClientEndScan = errors.New("error when scan target")