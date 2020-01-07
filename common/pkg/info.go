/*
 *应用程序本身的一些信息
 */
package pkg

import (
	"time"
)

type Info struct {
	AppName   string    `json:"app_name"`
	Version   string    `json:"version"`
	StartTime time.Time `json:"start_time"`
}
