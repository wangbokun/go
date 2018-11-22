package influxdb

import (
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)


//add string http://10.10.10.10.:8086
func ConnInflux(add, username , password string) client.Client {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     add,
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	return cli
}

//Insert
func WritesPoints(db_name, table_name string,cli client.Client, tags map[string]string, fields map[string]interface{}) error {

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db_name,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	pt, err := client.NewPoint(
		table_name,
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		return err
	}
	bp.AddPoint(pt)

	if err := cli.Write(bp); err != nil {
		return err
	}
	return nil
}


//query
// func QueryDB(MyDB, cli client.Client, cmd string) (res []client.Result, err error) {
// 	q := client.Query{
// 		Command:  cmd,
// 		Database: MyDB,
// 	}
// 	if response, err := cli.Query(q); err == nil {
// 		if response.Error() != nil {
// 			return res, response.Error()
// 		}
// 		res = response.Results
// 	} else {
// 		return res, err
// 	}
// 	return res, nil
// }
