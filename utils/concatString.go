package utils

import (
    "bytes"
)

func Concat(strs ...string) string {
    var buff bytes.Buffer
    
    for _,str := range strs {
        buff.WriteString(str)
    }
    
    return buff.String()
}