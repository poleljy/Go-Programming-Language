# encoding包
``` go
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// json转xml
type Person struct {
	Id        int    `xml:"id,attr" json:"id"`
	FirstName string `xml:"name>first" json:"first"`
	LastName  string `xml:"name>last" json:"last"`
}

func Json2Xml(src []byte) (string, error) {
	var p Person
	if err := json.Unmarshal(src, &p); err != nil {
		return "", err
	}

	des, err := xml.MarshalIndent(p, "", "\t")
	if err != nil {
		return "", err
	}
	return string(des), nil
}

// xml转json
type DataFormat struct {
	ProductList []struct {
		Name     string `xml:"Name" json:"name"`
		Quantity int    `xml:"Quantity" json:"quantity"`
	} `xml:"Product" json:"products"`
}

func Xml2Json(src []byte) (string, error) {
	var data DataFormat
	if err := xml.Unmarshal(src, &data); err != nil {
		return "", err
	}

	des, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(des), nil
}

// xml文件转json文件
type jsonStaff struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
}

type Staff struct {
	XMLName   xml.Name `xml:"staff"`
	ID        int      `xml:"id" json:"ID"`
	FirstName string   `xml:"firstname" json:"FirstName"`
	LastName  string   `xml:"lastname" json:"LastName"`
	UserName  string   `xml:"username" json:"UserName"`
}

type Company struct {
	XMLName xml.Name `xml:"company"`
	Staffs  []Staff  `xml:"staff" json:"company"`
}

func (s Staff) String() string {
	return fmt.Sprintf("\t ID : %d - FirstName : %s - LastName : %s - UserName : %s \n", s.ID, s.FirstName, s.LastName, s.UserName)
}

func XmlFile2JsonFile(xmlFile, jsonFile string) error {
	file, err := os.Open(xmlFile)
	if err != nil {
		return err
	}
	defer file.Close()

	xmlData, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var c Company
	if err := xml.Unmarshal(xmlData, &c); err != nil {
		return err
	}
	fmt.Println(c.Staffs)

	// convert to JSON
	var staffs []jsonStaff
	for _, value := range c.Staffs {
		staff := jsonStaff{
			ID:        value.ID,
			FirstName: value.FirstName,
			LastName:  value.LastName,
			UserName:  value.UserName,
		}
		staffs = append(staffs, staff)
	}
	jsonData, err := json.Marshal(staffs)
	if err != nil {
		return err
	}

	newFile, err := os.Create(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	defer newFile.Close()

	_, err = newFile.Write(jsonData)
	return err
}

func main() {
	fmt.Println("json ==>  xml")
	j := `{"id": 10, "first": "firstname", "last":"lastname"}`
	x, err := Json2Xml([]byte(j))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("json:\n%s\nxml:\n%s\n", j, x)

	fmt.Println("xml ==>  json")
	xmlData := []byte(`
		<?xml version="1.0" encoding="UTF-8" ?>
		<ProductList>
		    <Product>
		        <Name>ABC123</Name>
		        <Quantity>2</Quantity>
		    </Product>
		    <Product>
		        <Name>ABC123</Name>
		        <Quantity>2</Quantity>
		    </Product>
		</ProductList>`)

	jsonData, err := Xml2Json([]byte(xmlData))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("json:\n%s\nxml:\n%s\n", jsonData, xmlData)

	fmt.Println("xml file ==>  json file")
	err = XmlFile2JsonFile("Employees.xml", "Employees.json")
	if err != nil {
		fmt.Println(err)
	}
}

```