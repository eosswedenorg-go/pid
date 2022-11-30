package pid

import (
	"io/ioutil"
	"os"
	"strconv"
)

func Get() int {
	return os.Getpid()
}

func Save(path string) error {
	err := ioutil.WriteFile(path, []byte(strconv.Itoa(Get())), 0o644)
	if err != nil {
		return err
	}
	return nil
}

func Read(path string) (int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return -1, err
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return -1, err
	}
	return pid, nil
}
