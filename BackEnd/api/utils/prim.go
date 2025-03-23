package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func PrimCall() {

	url := "https://prim.iledefrance-mobilites.fr/marketplace/v2/navitia/line_reports/line_reports?depth=1&disable_geojson=true&tags%5B%5D=Actualit%C3%A9"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("apikey", "KEY HERE")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	f, err := os.Create("test.json")

	if err != nil {

		fmt.Println(err)

		return

	}

	l, err := f.WriteString(string(body))

	if err != nil {

		fmt.Println(err)

		f.Close()

		return

	}

	fmt.Println(l, "bytes written successfully")

	err = f.Close()

	if err != nil {

		fmt.Println(err)

		return

	}

}
