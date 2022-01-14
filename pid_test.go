
package pid

import(
    "os"
    "io/ioutil"
    "strconv"

    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

var osPid = os.Getpid()

func TestGet(t *testing.T) {
    pid := Get()
    assert.Equal(t, pid, osPid)
}

func TestSave(t *testing.T) {

    tmpfile, err := ioutil.TempFile("", "go_test")
    require.NoError(t, err)

    defer os.Remove(tmpfile.Name())

    err = Save(tmpfile.Name())
    require.NoError(t, err)

    data := make([]byte, 32)
    l, err := tmpfile.Read(data)
    require.NoError(t, err)

    pid, err := strconv.Atoi(string(data[:l]))

    require.NoError(t, err)
    assert.Equal(t, pid, osPid)
}


func TestRead(t *testing.T) {

    tmpfile, err := ioutil.TempFile("", "go_test")
    require.NoError(t, err)

    defer os.Remove(tmpfile.Name())

    _, err = tmpfile.WriteString("1234")
    require.NoError(t, err)

    pid, err := Read(tmpfile.Name())

    require.NoError(t, err)
    assert.Equal(t, pid, 1234)
}

func TestReadEmptyFile(t *testing.T) {

    tmpfile, err := ioutil.TempFile("", "go_test")
    if err != nil {
        t.Error(err)
        return
    }
    defer os.Remove(tmpfile.Name())

    pid, err := Read(tmpfile.Name())

    require.Error(t, err)
    assert.Equal(t, pid, -1)
}

func TestReadRandomContent(t *testing.T) {

    tmpfile, err := ioutil.TempFile("", "go_test")
    if err != nil {
        t.Error(err)
        return
    }
    defer os.Remove(tmpfile.Name())

    _, err = tmpfile.WriteString("some_random_data")
    if err != nil {
        t.Error(err)
        return
    }

    pid, err := Read(tmpfile.Name())
    require.Error(t, err)
    assert.Equal(t, pid, -1)
}
