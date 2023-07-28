package lwhelper

import (
	"fmt"
	"log"
	"os"

	"github.com/jaytaylor/go-hostsfile"
)

func FixStaticHosts(domain string) {

	staticHost := map[string]string{} // key=domain, value=ip
	staticHost[domain] = GetOutboundIP().String()

	// ---

	hostInFile := map[string]string{} // key=domain, value=ip
	hostFileLine, err := hostsfile.ParseHosts(hostsfile.ReadHostsFile())
	if err != nil {
		log.Fatal(err)
	}

	for ip, domains := range hostFileLine {
		for _, domain := range domains {
			hostInFile[domain] = ip
		}
	}

	anyChange := false
	for domain, ip := range staticHost {

		if hostInFile[domain] != ip {
			hostInFile[domain] = ip
			anyChange = true
		}

		// out += entry.Key + " \t " + strings.Join(entry.Value, " ") + "\n"
	}

	if anyChange {

		hosts := map[string]string{}
		for domain, ip := range hostInFile {
			hosts[ip] += " " + domain
		}

		out := ""
		for ip, domains := range hosts {
			out += ip + "  " + domains + "\n"
		}

		err = os.WriteFile("/etc/hosts", []byte(out), 0644)
		if err != nil {
			fmt.Println(out)
			log.Fatal(err)
		}

		// windows
		f, err := os.OpenFile("/mnt/host/c/Windows/System32/drivers/etc/hosts", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err == nil {
			f.WriteString(
				GetOutboundIP().String() + "  " + domain + "\n",
			)
		}
		defer f.Close()

	}

}
