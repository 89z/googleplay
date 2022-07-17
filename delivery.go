package googleplay

import (
	"errors"
	"github.com/89z/rosso/protobuf"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Delivery struct {
	protobuf.Message
}

func (h Header) Delivery(app string, ver uint64) (*Delivery, error) {
	req, err := http.NewRequest(
		"GET", "https://play-fe.googleapis.com/fdfe/delivery", nil,
	)
	if err != nil {
		return nil, err
	}
	h.Set_Agent(req.Header)
	h.Set_Auth(req.Header) // needed for single APK
	h.Set_Device(req.Header)
	req.URL.RawQuery = url.Values{
		"doc": {app},
		"vc":  {strconv.FormatUint(ver, 10)},
	}.Encode()
	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// ResponseWrapper
	response_wrapper, err := protobuf.Unmarshal(body)
	if err != nil {
		return nil, err
	}
	// .payload.deliveryResponse
	delivery_response := response_wrapper.Get(1).Get(21)
	// .status
	status, err := delivery_response.Get_Varint(1)
	if err != nil {
		return nil, err
	}
	switch status {
	case 2:
		return nil, errors.New("geo-blocking")
	case 3:
		return nil, errors.New("purchase required")
	case 5:
		return nil, errors.New("invalid version")
	}
	var del Delivery
	// .appDeliveryData
	del.Message = delivery_response.Get(2)
	return &del, nil
}

// .downloadUrl
func (d Delivery) Download_URL() (string, error) {
	return d.Get_String(3)
}

type Split_Data struct {
	protobuf.Message
}

// .id
func (s Split_Data) ID() (string, error) {
	return s.Get_String(1)
}

// .downloadUrl
func (s Split_Data) Download_URL() (string, error) {
	return s.Get_String(5)
}

func (d Delivery) Split_Data() []Split_Data {
	var splits []Split_Data
	// .splitDeliveryData
	for _, split := range d.Get_Messages(15) {
		splits = append(splits, Split_Data{split})
	}
	return splits
}

func (d Delivery) Additional_File() []File_Metadata {
	var files []File_Metadata
	// .additionalFile
	for _, file := range d.Get_Messages(4) {
		files = append(files, File_Metadata{file})
	}
	return files
}

type File_Metadata struct {
	protobuf.Message
}

// .fileType
func (f File_Metadata) File_Type() (uint64, error) {
	return f.Get_Varint(1)
}

// .downloadUrl
func (f File_Metadata) Download_URL() (string, error) {
	return f.Get_String(4)
}

type File struct {
	Package_Name string
	Version_Code uint64
}

func (f File) APK(id string) string {
	var buf []byte
	buf = append(buf, f.Package_Name...)
	buf = append(buf, '-')
	if id != "" {
		buf = append(buf, id...)
		buf = append(buf, '-')
	}
	buf = strconv.AppendUint(buf, f.Version_Code, 10)
	buf = append(buf, ".apk"...)
	return string(buf)
}

func (f File) OBB(file_type uint64) string {
	var buf []byte
	if file_type >= 1 {
		buf = append(buf, "patch"...)
	} else {
		buf = append(buf, "main"...)
	}
	buf = append(buf, '.')
	buf = strconv.AppendUint(buf, f.Version_Code, 10)
	buf = append(buf, '.')
	buf = append(buf, f.Package_Name...)
	buf = append(buf, ".obb"...)
	return string(buf)
}
