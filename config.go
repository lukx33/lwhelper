package lwhelper

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/lukx33/lwhelper/out"
)

type ConfigStore interface {
	out.Info

	String(key string, defval string) (func() string, func(value string) out.Info)
	Int(key string, defval int) (func() int, func(value int) out.Info)
	Bool(key string, defval bool) (func() bool, func(value bool) out.Info)
}

// ---

func LoadConfigFile(fp string) ConfigStore {

	if strings.Contains(fp, "~") {
		home, _ := os.UserHomeDir()
		fp = strings.ReplaceAll(fp, "~", home)
	}

	store := &configStoreS{
		FilePath: fp,
		Data:     map[string]interface{}{},
	}

	buf, _ := os.ReadFile(store.FilePath)
	return out.CheckErrorFor(store, json.Unmarshal(buf, &store.Data))
}

// ---

type configStoreS struct {
	out.StructS
	FilePath string
	Data     map[string]interface{}
}

func (c *configStoreS) save() out.Info {

	buf, err := json.MarshalIndent(c.Data, "", "  ")
	if err != nil {
		return out.CheckError(err)
	}

	os.MkdirAll(filepath.Dir(c.FilePath), 0755)

	return out.CheckError(os.WriteFile(c.FilePath, buf, 0644))
}

// ---
// string

func (c *configStoreS) String(key string, defval string) (func() string, func(value string) out.Info) {

	return func() string { // <======== Get
			v, exist := c.Data[key].(string)
			if !exist {
				return defval
			}
			return v

		}, func(value string) out.Info { // <======== Set
			c.Data[key] = value
			return c.save()
		}
}

// ---
// int

func (c *configStoreS) Int(key string, defval int) (func() int, func(value int) out.Info) {

	return func() int { // <======== Get
			v, exist := c.Data[key].(float64)
			if !exist {
				return defval
			}
			return int(v)

		}, func(value int) out.Info { // <======== Set
			c.Data[key] = value
			return c.save()
		}
}

// ---
// bool

func (c *configStoreS) Bool(key string, defval bool) (func() bool, func(value bool) out.Info) {

	return func() bool { // <======== Get
			v, exist := c.Data[key].(bool)
			if !exist {
				return defval
			}
			return v

		}, func(value bool) out.Info { // <======== Set
			c.Data[key] = value
			return c.save()
		}
}

// ---
