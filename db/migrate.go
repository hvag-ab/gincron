package db

import (

	"log"

)

func Migrate(lis []interface{}){
	if len(lis) == 0{
		return 
	}else{
		for _,st := range lis{
			err := DB.Sync2(st)
			if err !=nil{
				log.Fatalf("Fail to sync database: %v\n", err)
			}
		}
	}
	
}