package pkg

import (
	"bufio"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/colinmarc/hdfs/v2"
)

type HDFS struct {
	client *hdfs.Client
}

func NewHDFS(client *hdfs.Client) *HDFS {
	return &HDFS{client: client}
}

func (h *HDFS) PrePare(filename string) ([][]string, error) {
	file, err := h.client.Open(filename)
	if err != nil {
		hlog.Error("open file failed")
		return nil, err
	}
	defer file.Close()

	var res []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	if err = scanner.Err(); err != nil {
		hlog.Error("read file failed")
		return nil, err
	}

	return res, nil
}

func (h *HDFS) GetDir(username string) ([]string, error) {
	path := "/user/" + username
	dir, err := h.client.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, file := range dir {
		result = append(result, file.Name())
	}
	return result, nil
}

func (h *HDFS) CreateFile(filename string, description string) (*hdfs.FileWriter, error) {
	path := filename + description
	file, err := h.client.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
