package libvirt

import (
	"fmt"
	"math/rand"

	"github.com/libvirt/libvirt-go-xml"
)

const oui = "05abcd"

// note, source is not initialized
func newDefDisk(i int) libvirtxml.DomainDisk {
	return libvirtxml.DomainDisk{
		Device: "disk",
		Target: &libvirtxml.DomainDiskTarget{
			Bus: "virtio",
			Dev: fmt.Sprintf("vd%s", diskLetterForIndex(i)),
		},
		Driver: &libvirtxml.DomainDiskDriver{
			Name: "qemu",
			Type: "raw",
		},
	}
}

func newDefiSCSIDisk(i int, iscsiDiskPath string) libvirtxml.DomainDisk {
	return libvirtxml.DomainDisk{
		Device: "disk",
		Target: &libvirtxml.DomainDiskTarget{
			Bus: "ide",
			Dev: fmt.Sprintf("hd%s", diskLetterForIndex(i)),
		},
		Driver: &libvirtxml.DomainDiskDriver{
			Name:  "qemu",
			Type:  "raw",
			Cache: "none",
			IO:    "native",
		},
		Source: &libvirtxml.DomainDiskSource{
			Block: &libvirtxml.DomainDiskSourceBlock{
				Dev: iscsiDiskPath,
			},
		},
	}
}

func randomWWN(strlen int) string {
	const chars = "abcdef0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return oui + string(result)
}
