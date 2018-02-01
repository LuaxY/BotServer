package log

import (
    "io"
    "log"
)

var (
    Debug   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
    Log     *log.Logger
)

func Init(
    debugHandle io.Writer,
    infoHandle io.Writer,
    warningHandle io.Writer,
    errorHandle io.Writer,
    logHandle io.Writer) {

    Debug = log.New(debugHandle,
        "DEBUG: ",
        log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

    Info = log.New(infoHandle,
        " INFO: ",
        log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

    Warning = log.New(warningHandle,
        " WARN: ",
        log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

    Error = log.New(errorHandle,
        "ERROR: ",
        log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

    Log = log.New(logHandle,
        "",
        log.Ldate|log.Ltime|log.Lmicroseconds)
}