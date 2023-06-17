package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/status"
	"github.com/colinmarc/hdfs/v2"
	file "hdfs/kitex_gen/file"
	"sort"
	"sync"
)

// FileServiceImpl implements the last service interface defined in the IDL.
type FileServiceImpl struct {
	MapReduce
	HDFS
}

type KeyValue struct {
	Key   string
	Value string
}

type HDFS interface {
	PrePare(filename string) ([][]string, error)
	GetDir(username string) ([]string, error)
	CreateFile(filename string, description string) (*hdfs.FileWriter, error)
}

type MapReduce interface {
	Map(data []string) []KeyValue
	Reduce(key string, values []string) KeyValue
	MapReduce(data []string, mapper func(string) []KeyValue, reducer func(string, []string) KeyValue)
}

// GetFiles implements the FileServiceImpl interface.
func (s *FileServiceImpl) GetFiles(ctx context.Context, req *file.GetFilesRequest) (resp *file.GetFilesResponse, err error) {
	// TODO: Your code here...
	dir, err := s.HDFS.GetDir(req.Username)
	if err != nil {
		klog.Errorf("get dir error: ", err)
		return nil, status.Err(codes.Internal, "get dir error")
	}
	return &file.GetFilesResponse{Files: dir}, nil
}

// RemoveRepeat implements the FileServiceImpl interface.
func (s *FileServiceImpl) RemoveRepeat(ctx context.Context, req *file.RemoveRepeatRequest) (resp *file.RemoveRepeatResponse, err error) {
	// TODO: Your code here...
	datas, err := s.HDFS.PrePare(req.Filename)
	if err != nil {
		klog.Errorf("hdfs prepare error: ", err)
		return nil, status.Err(codes.Internal, "hdfs prepare error")
	}
	var lock sync.Mutex
	var result map[string]int
	for _, data := range datas {
		go func() {
			values := s.MapReduce.Map(data)
			for _, value := range values {
				lock.Lock()
				if _, ok := result[value.Key]; !ok {
					result[value.Key] = 1
				} else {
					result[value.Key]++
				}
				lock.Unlock()
			}
		}()
	}
	var res []byte
	for key := range result {
		keyByte := []byte(key)
		res = append(res, keyByte...)
	}
	file, err := s.HDFS.CreateFile(req.Filename, "rrt")
	if err != nil {
		return nil, err
	}
	_, err = file.Write(res)
	if err != nil {
		klog.Errorf("write file error: ", err)
		return nil, status.Err(codes.Internal, "write file error")
	}
	return &file.RemoveRepeatResponse{req.Filename + "rrt"}, nil
}

// SortByNum implements the FileServiceImpl interface.
func (s *FileServiceImpl) SortByNum(ctx context.Context, req *file.SortByNumRequest) (resp *file.SortByNumResponse, err error) {
	// TODO: Your code here...
	datas, err := s.HDFS.PrePare(req.Filename)
	if err != nil {
		klog.Errorf("hdfs prepare error: ", err)
		return nil, status.Err(codes.Internal, "hdfs prepare error")
	}
	var intermediate map[string][]string
	var lock sync.Mutex
	var wg sync.WaitGroup

	for _, data := range datas {
		wg.Add(1)
		go func(data []string) {
			defer wg.Done()
			keyValues := s.MapReduce.Map(data)
			lock.Lock()
			for _, keyValue := range keyValues {
				intermediate[keyValue.Key] = append(intermediate[keyValue.Key], keyValue.Value)
			}
			lock.Unlock()

		}(data)
	}
	wg.Wait()

	var results []KeyValue
	for key, value := range intermediate {
		result := s.MapReduce.Reduce(key, value)
		lock.Lock()
		results = append(results, result)
		lock.Unlock()
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Value > results[j].Value
	})

	var res []byte
	for _, keyValue := range results {
		s := fmt.Sprintf("%s %s\n", keyValue.Key, keyValue.Value)
		keyByte := []byte(s)
		res = append(res, keyByte...)
	}
	file, err := s.HDFS.CreateFile(req.Filename, "sbn")
	if err != nil {
		return nil, err
	}
	_, err = file.Write(res)
	if err != nil {
		klog.Errorf("write file error: ", err)
		return nil, status.Err(codes.Internal, "write file error")
	}
	return &file.RemoveRepeatResponse{req.Filename + "sbn"}, nil
}

// SortByTime implements the FileServiceImpl interface.
func (s *FileServiceImpl) SortByTime(ctx context.Context, req *file.SortByNumRequest) (resp *file.SortByTimeResponse, err error) {
	// TODO: Your code here...
	datas, err := s.HDFS.PrePare(req.Filename)
	if err != nil {
		klog.Errorf("hdfs prepare error: ", err)
		return nil, status.Err(codes.Internal, "hdfs prepare error")
	}
	var intermediate map[string][]string
	var lock sync.Mutex
	var wg sync.WaitGroup

	for _, data := range datas {
		wg.Add(1)
		go func(data []string) {
			defer wg.Done()
			keyValues := s.MapReduce.Map(data)
			lock.Lock()
			for _, keyValue := range keyValues {
				intermediate[keyValue.Key] = append(intermediate[keyValue.Key], keyValue.Value)
			}
			lock.Unlock()

		}(data)
	}
	wg.Wait()

	var results []KeyValue
	for key, value := range intermediate {
		result := s.MapReduce.Reduce(key, value)
		lock.Lock()
		results = append(results, result)
		lock.Unlock()
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Value > results[j].Value
	})

	var res []byte
	for _, keyValue := range results {
		s := fmt.Sprintf("%s %s\n", keyValue.Key, keyValue.Value)
		keyByte := []byte(s)
		res = append(res, keyByte...)
	}
	file, err := s.HDFS.CreateFile(req.Filename, "sbt")
	if err != nil {
		return nil, err
	}
	_, err = file.Write(res)
	if err != nil {
		klog.Errorf("write file error: ", err)
		return nil, status.Err(codes.Internal, "write file error")
	}
	return &file.RemoveRepeatResponse{req.Filename + "sbt"}, nil
}
