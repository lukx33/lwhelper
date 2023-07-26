package lwhelper

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/lukx33/lwhelper/out"
)

type ConfigStore interface {
	out.Info

	String(key string) (func() string, func(value string) out.Info)
	Int(key string) (func() int, func(value int) out.Info)
	Bool(key string) (func() bool, func(value bool) out.Info)
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
	if len(buf) > 0 {
		if store.CatchError(json.Unmarshal(buf, &store.Data)) {
			return store
		}
	}

	return out.SetSuccess(store)
}

// ---

type configStoreS struct {
	out.DontUseMeInfoS
	FilePath string
	Data     map[string]interface{}
}

func (c *configStoreS) save() out.Info {

	buf, err := json.MarshalIndent(c.Data, "", "  ")
	if err != nil {
		return out.NewError(err)
	}

	err = os.WriteFile(c.FilePath, buf, 0644)
	if err != nil {
		return out.NewError(err)
	}

	return out.NewSuccess()
}

// ---
// string

func (c *configStoreS) String(key string) (func() string, func(value string) out.Info) {

	return func() string { // <======== Get
			v, _ := c.Data[key].(string)
			return v

		}, func(value string) out.Info { // <======== Set
			c.Data[key] = value
			return c.save()
		}
}

// ---
// int

func (c *configStoreS) Int(key string) (func() int, func(value int) out.Info) {

	return func() int { // <======== Get
			v, _ := c.Data[key].(float64)
			return int(v)

		}, func(value int) out.Info { // <======== Set
			c.Data[key] = value
			return c.save()
		}
}

// ---
// bool

func (c *configStoreS) Bool(key string) (func() bool, func(value bool) out.Info) {

	return func() bool { // <======== Get
			v, _ := c.Data[key].(bool)
			return v

		}, func(value bool) out.Info { // <======== Set
			c.Data[key] = value
			return c.save()
		}
}

// ---
