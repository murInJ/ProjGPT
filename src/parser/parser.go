package parser

import "sync"

func ParseProj(root string, parseStruct bool, parseFileContent bool) (map[string]string, error) {
	m := make(map[string]string)
	var mu sync.Mutex // 创建互斥锁
	var wg sync.WaitGroup
	errCh := make(chan error, 2)
	defer close(errCh)
	if parseStruct {
		wg.Add(1)
		go func() {
			defer wg.Done()
			str, err := getStructTreeString(root)
			if err != nil {
				errCh <- err
				return
			}
			mu.Lock() // 加锁
			m["Project Struct Tree"] = str
			mu.Unlock() // 解锁
		}()
	}

	if parseFileContent {
		wg.Add(1)
		go func() {
			defer wg.Done()
			str, err := getFileContentString(root)
			if err != nil {
				errCh <- err
				return
			}
			mu.Lock()
			m["Project file content"] = str
			mu.Unlock()
		}()
	}

	wg.Wait()

	select {
	case err := <-errCh:
		return nil, err
	default:
		return m, nil
	}
}
