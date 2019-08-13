
package types
import (
	"encoding/csv"
	"os"
	"fmt"
)
// write append csv files
func CsvWriter(file string,yourSliceGoesHere []string) error {
    f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
	    fmt.Println("Error: ", err)
	    return err
    }
    w := csv.NewWriter(f)
    w.Write(yourSliceGoesHere)
	w.Flush()
	defer f.Close()

	return nil
}