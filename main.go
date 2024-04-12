package main

import (
	"fmt"
	"reflect"
	"time"
)

type Stat struct {
	AcquireCount            int64         `json:"acquireCount"`
	AcquireDuration         time.Duration `json:"acquireDuration"`
	AcquiredConns           int32         `json:"acquiredConns"`
	CanceledAcquireCount    int64         `json:"canceledAcquireCount"`
	ConstructingConns       int32         `json:"constructingConns"`
	EmptyAcquireCount       int64         `json:"emptyAcquireCount"`
	IdleConns               int32         `json:"idleConns"`
	MaxConns                int32         `json:"maxConns"`
	TotalConns              int32         `json:"totalConns"`
	NewConnsCount           int64         `json:"newConnsCount"`
	MaxLifetimeDestroyCount int64         `json:"maxLifetimeDestroyCount"`
	MaxIdleDestroyCount     int64         `json:"maxIdleDestroyCount"`
}

func NewStat() *Stat {
	return &Stat{
		AcquireCount:            0,
		AcquireDuration:         1,
		AcquiredConns:           2,
		CanceledAcquireCount:    3,
		ConstructingConns:       1,
		EmptyAcquireCount:       3,
		IdleConns:               2,
		MaxConns:                3,
		TotalConns:              3,
		NewConnsCount:           1,
		MaxLifetimeDestroyCount: 3,
		MaxIdleDestroyCount:     1,
	}
}

func NewDBInfo() *DBInfo {
	return &DBInfo{
		Primary:   NewStat(),
		Secondary: NewStat(),
	}
}

type DBInfo struct {
	Primary   *Stat `json:"primary"`
	Secondary *Stat `json:"secondary"`
}

type Response struct {
	Database *DBInfo `json:"database"`
}

func main() {
	x := &Response{
		Database: NewDBInfo(),
	}
	v := reflect.Indirect(reflect.ValueOf(x.Database))
	vtyp := v.Type()
	result := ""

	for i := 0; i < v.NumField(); i++ {
		inner := reflect.Indirect(v.Field(i))
		typ := inner.Type()
		for j := 0; j < inner.NumField(); j++ {
			//fmt.Println(fmt.Sprintf(vtyp.Field(i).Name + typ.Field(j).Name + fmt.Sprintf(" %v", inner.Field(j).Interface()) + "\n"))
			result += fmt.Sprintf(vtyp.Field(i).Name + typ.Field(j).Name + fmt.Sprintf(" %v", inner.Field(j).Interface()) + "\n")
		}
	}

	fmt.Println(result)
}
