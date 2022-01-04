
package pid

import (
    "os"
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
