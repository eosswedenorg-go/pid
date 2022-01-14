
package pid

import (
    "os"
    "io/ioutil"
    "strconv"
)

func Get() (int) {
    return os.Getpid()
}

func Save(path string) (error) {

    fd, err := os.Create(path)
    if err != nil {
        return err
    }

    _, err = fd.WriteString(strconv.Itoa(Get()))

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
